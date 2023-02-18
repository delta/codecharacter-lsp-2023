package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func handleJSONRPC(ws *models.WebsocketConnection, requestMessageBytes []byte) error {
	// TODO: Send Content Length
	_, err := ws.LSPServer.Stdin.Write(requestMessageBytes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var responseMessageBytes []byte
	reader := bufio.NewReader(ws.LSPServer.Stdout)
	// TODO: Read Content Length
	responseMessageBytes, _, err = reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("JSONRPC Response : ", string(responseMessageBytes), " with ID : ", ws.ID)
	var responseMessage map[string]interface{}
	err = json.Unmarshal(responseMessageBytes, &responseMessage)
	if err != nil {
		fmt.Println("Error in unmarshalling JSONRPC response : ", err)
		return err
	}
	return SendMessage(ws, responseMessage)
}
