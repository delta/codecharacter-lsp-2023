package utils

import (
	"fmt"
	"os"

	"github.com/delta/codecharacter-lsp-2023/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

func InitWebsocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	CreateWebsocketConnection(ws)
	return nil
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
		_, messageBytes, err := ws.ReadMessage()
		message := string(messageBytes[:])
		if err != nil {
			fmt.Println("Error occured : ", err)
			break
		}
		fmt.Println("Websocket Message : ", message, " with ID : ", wsConnectionParams.ID)
	}
}
