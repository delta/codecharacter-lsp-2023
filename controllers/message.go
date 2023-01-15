package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/delta/codecharacter-lsp-2023/models"
	"github.com/gorilla/websocket"
)

func HandleMessage(ws *models.WebsocketConnection, messageBytes []byte) error {
	var message map[string]interface{}
	err := json.Unmarshal(messageBytes, &message)
	if err != nil {
		return err
	}
	_, isPresent := message["jsonrpc"]
	fmt.Println("Is JSONRPC? : ", isPresent)
	if isPresent {
		return handleJSONRPCRequest(ws, messageBytes)
	}
	return handleWebSocketRequest(ws, message)
}

func SendMessage(ws *models.WebsocketConnection, message map[string]interface{}) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = ws.Connection.WriteMessage(websocket.TextMessage, messageBytes)
	if err != nil {
		return err
	}
	return nil
}

func handleJSONRPCRequest(ws *models.WebsocketConnection, messageBytes []byte) error {
	fmt.Println("JSONRPC Request : ", string(messageBytes), " with ID : ", ws.ID)
	return nil
}

func handleWebSocketRequest(ws *models.WebsocketConnection, message map[string]interface{}) error {
	fmt.Println("Websocket Request : ", message, " with ID : ", ws.ID)
	if message["operation"] == "fileUpdate" {
		return handleFileUpdate(message, ws)
	}
	return nil
}
