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
	return err
}

func SendErrorMessage(ws *models.WebsocketConnection, message error) error {
	responseBody := make(map[string]interface{})
	responseBody["status"] = "error"
	responseBody["message"] = message.Error()
	err := SendMessage(ws, responseBody)
	return err
}

func handleJSONRPCRequest(ws *models.WebsocketConnection, messageBytes []byte) error {
	fmt.Println("JSONRPC Request : ", string(messageBytes), " with ID : ", ws.ID)
	return handleJSONRPC(ws, messageBytes)
}

func handleWebSocketRequest(ws *models.WebsocketConnection, message map[string]interface{}) error {
	fmt.Println("Websocket Request : ", message, " with ID : ", ws.ID)
	switch message["operation"] {
	case "fileUpdate":
		return handleFileUpdate(message, ws)
	case "getAbsPath":
		return getAbsPath(ws)
	}
	return nil
}
