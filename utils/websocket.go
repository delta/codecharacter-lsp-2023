package utils

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/delta/codecharacter-lsp-2023/config"
	"github.com/delta/codecharacter-lsp-2023/controllers"
	"github.com/delta/codecharacter-lsp-2023/models"
	"github.com/delta/codecharacter-lsp-2023/servers"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	EnableCompression: true,
}

func CheckOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	return origin == config.FrontendURL
}

func InitWebsocket(c echo.Context) error {
	var ws models.WebsocketConnection
	ws.ID = uuid.New()
	language := c.Param("language")
	if language != "cpp" && language != "java" && language != "python" {
		return c.String(http.StatusBadRequest, "Invalid Language")
	}
	switch language {
	case "cpp":
		ws.Language = models.Cpp
	case "java":
		ws.Language = models.Java
	case "python":
		ws.Language = models.Python
	}
	c.Echo().Logger.Info("WS Connection Created with ID : ", ws.ID, " and Language : ", language)
	upgrader.CheckOrigin = CheckOrigin
	wsConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.String(http.StatusBadRequest, "Error Upgrading to Websocket Connection")
	}
	wsConn.EnableWriteCompression(true)
	ws.Connection = wsConn
	err = createWorkspace(&ws, c)
	if err != nil {
		return c.String(http.StatusBadGateway, "Something went wrong, contact the event administrator.")
	}
	return nil
}

func dropConnection(ws *models.WebsocketConnection, c echo.Context) {
	err := ws.LSPServer.Process.Process.Signal(os.Interrupt)
	if err != nil {
		c.Echo().Logger.Error(err)
	}
	_ = ws.LSPServer.DevNullFd.Close()
	// Reads process exit state to remove the <defunct> process from the system process table
	_ = ws.LSPServer.Process.Wait()
	err = os.RemoveAll(ws.WorkspacePath)
	if err != nil {
		c.Echo().Logger.Error(err)
	}
	ws.Connection.Close()
	c.Echo().Logger.Info("WS Connection ", ws.ID, " closed")
}

func createWorkspace(ws *models.WebsocketConnection, c echo.Context) error {
	defer dropConnection(ws, c)

	ws.WorkspacePath = "workspaces/" + ws.ID.String()
	err := os.Mkdir(ws.WorkspacePath, os.ModePerm)
	if err != nil {
		c.Echo().Logger.Error(err)
		return err
	}
	headerFiles, err := os.ReadDir("player_code/" + ws.Language.GetLanguage() + "/")
	if err != nil {
		c.Echo().Logger.Error(err)
		return err
	}
	for _, headerFile := range headerFiles {
		absFilePath, _ := filepath.Abs("player_code/" + ws.Language.GetLanguage() + "/" + headerFile.Name())
		_ = os.Symlink(absFilePath, ws.WorkspacePath+"/"+headerFile.Name())
	}
	err = servers.CreateLSPServer(ws)
	if err != nil {
		c.Echo().Logger.Error(err)
		return err
	}
	// Start an async goroutine to listen for messages from the LSP server
	go controllers.StreamReader(ws)
	err = listen(ws, c)
	if err != nil {
		c.Echo().Logger.Error(err)
		return err
	}
	return nil
}

func listen(ws *models.WebsocketConnection, c echo.Context) error {
	for {
		_, messageBytes, err := ws.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Echo().Logger.Info("WS Connection ", ws.ID, " closing with error : ", err)
				return err
			}
			return nil
		}
		err = controllers.HandleMessage(ws, messageBytes)
		if err != nil {
			c.Echo().Logger.Error(err)
			return err
		}
	}
}
