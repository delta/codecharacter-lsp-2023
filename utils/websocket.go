package utils

import (
	"fmt"
	"os"

	"github.com/delta/codecharacter-lsp-2023/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func InitWebsocket(c echo.Context) {
	websocket.Handler(CreateWebsocketConnection).ServeHTTP(c.Response(), c.Request())
}

func drop(ws *websocket.Conn, wsConnectionParams models.WebsocketConnectionParams) error {
	err := os.RemoveAll("workspaces/" + wsConnectionParams.ID.String())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dropping WS Connection : ", wsConnectionParams.ID)
	ws.Close()
	return nil
}

func CreateWebsocketConnection(ws *websocket.Conn) {
	var wsConnectionParams models.WebsocketConnectionParams
	id := uuid.New()
	wsConnectionParams.ID = id
	fmt.Println("WS Connection Created with ID : ", id)
	defer drop(ws, wsConnectionParams)

	err := os.Mkdir("workspaces/"+wsConnectionParams.ID.String(), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	listen(ws, wsConnectionParams)
}

func listen(ws *websocket.Conn, wsConnectionParams models.WebsocketConnectionParams) {
	for {
		fmt.Println("Listening for Messages")
		message := ""
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Println("Error occured : ", err)
			break
		}
		fmt.Println("Websocket Message : ", message)
	}
}
