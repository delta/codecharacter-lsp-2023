package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func CreateWebsocketConnection(c echo.Context) error {
	fmt.Println("WS Connection Created")
	websocket.Handler(listen).ServeHTTP(c.Response(), c.Request())
	return nil
}

func drop(ws *websocket.Conn) error {
	fmt.Println("Dropping WS Connection")
	ws.Close()
	return nil
}

func listen(ws *websocket.Conn) {
	defer drop(ws)
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
