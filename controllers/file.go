package controllers

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func handleFileUpdate(message map[string]interface{}, ws *models.WebsocketConnection) error {
	fmt.Println("Processing File Update Request")
	filename := "player" + ws.Language.GetExtension()
	err := ioutil.WriteFile(ws.WorkspacePath+"/"+filename, []byte(message["code"].(string)), 0644)
	return err
}

func getAbsPath(ws *models.WebsocketConnection) error {
	filename := "player" + ws.Language.GetExtension()
	abs, err := filepath.Abs(ws.WorkspacePath + "/" + filename)
	responseBody := make(map[string]interface{})
	if err != nil {
		return SendErrorMessage(ws, err)
	}
	responseBody["status"] = "success"
	responseBody["message"] = abs
	err = SendMessage(ws, responseBody)
	return err
}
