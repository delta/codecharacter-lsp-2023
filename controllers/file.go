package controllers

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func handleFileUpdate(message map[string]interface{}, ws *models.WebsocketConnection) error {
	fmt.Println("Processing File Update Request")
	filename := "player" + ws.Language.Extension
	err := ioutil.WriteFile(ws.WorkspacePath+"/"+filename, []byte(message["code"].(string)), 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getAbsPath(message map[string]interface{}, ws *models.WebsocketConnection) error {
	filename := "player" + ws.Language.Extension
	abs, err := filepath.Abs(ws.WorkspacePath + "/" + filename)
	responseBody := make(map[string]interface{})
	if err != nil {
		fmt.Println(err)
		responseBody["status"] = "error"
		responseBody["message"] = err.Error()
		SendMessage(ws, responseBody)
		return err
	}
	responseBody["status"] = "success"
	responseBody["message"] = abs
	SendMessage(ws, responseBody)
	return nil
}
