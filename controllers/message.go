package controllers

import (
	"encoding/json"

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
	return ws.Connection.WriteMessage(websocket.TextMessage, messageBytes)
}

func SendErrorMessage(ws *models.WebsocketConnection, message error) error {
	responseBody := make(map[string]interface{})
	responseBody["status"] = "error"
	responseBody["message"] = message.Error()
	return SendMessage(ws, responseBody)
}

func handleJSONRPCRequest(ws *models.WebsocketConnection, messageBytes []byte) error {
	return handleJSONRPC(ws, messageBytes)
}

func handleWebSocketRequest(ws *models.WebsocketConnection, message map[string]interface{}) error {
	switch message["operation"] {
	case "fileUpdate":
		return handleFileUpdate(message, ws)
	case "getAbsPath":
		return getAbsPath(ws)
	}
	return nil
}
