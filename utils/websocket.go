package utils

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/delta/codecharacter-lsp-2023/controllers"
	"github.com/delta/codecharacter-lsp-2023/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

func InitWebsocket(c echo.Context) error {
	var ws models.WebsocketConnection
	id := c.QueryParam("id")
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
	fmt.Println("WS Connection Created with ID : ", id, " and Language : ", language)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	wsConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error Upgrading to Websocket Connection")
	}
	ws.Connection = wsConn
	err = createWorkspace(&ws)
	if err != nil {
		return c.String(http.StatusBadGateway, "Something went wrong, contact the event administrator.")
	}
	return nil
}

func dropConnection(ws *models.WebsocketConnection) {
	err := os.RemoveAll(ws.WorkspacePath)
	if err != nil {
		fmt.Println(err)
	}
	if ws.LSPServer.Process != nil {
		err = ws.LSPServer.Process.Process.Signal(os.Interrupt)
		if err != nil {
			fmt.Println(err)
		}
		// Reads process exit state to remove the <defunct> process from the system process table
		err = ws.LSPServer.Process.Wait()
		if err != nil {
			fmt.Println(err)
		}
	}
	ws.Connection.Close()
	fmt.Println("Dropped WS Connection : ", ws.ID)
}

func createWorkspace(ws *models.WebsocketConnection) error {
	defer dropConnection(ws)

	ws.WorkspacePath = "workspaces/" + ws.ID.String()
	err := os.Mkdir(ws.WorkspacePath, os.ModePerm)
	if err != nil {
		// fmt.Println("error is hereeeee")
		// fmt.Println(os.Getwd())
		fmt.Println(err)
		return err
	}
	headerFiles, err := os.ReadDir("player_code/" + ws.Language.GetLanguage() + "/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, headerFile := range headerFiles {
		absFilePath, _ := filepath.Abs("player_code/" + ws.Language.GetLanguage() + "/" + headerFile.Name())
		_ = os.Symlink(absFilePath, ws.WorkspacePath+"/"+headerFile.Name())
	}
	err = CreateLSPServer(ws)
	if err != nil {
		fmt.Println(err)
		return err
	}
	go controllers.SendMessageFunc(ws)
	err = listen(ws)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func listen(ws *models.WebsocketConnection) error {
	for {
		fmt.Println("Listening for Messages")
		_, messageBytes, err := ws.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println("WS Connection ", ws.ID, " closing with error : ", err)
				return err
			}
			return nil
		}
		err = controllers.HandleMessage(ws, messageBytes)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}
