package controllers

import (
	"io/ioutil"
	"path/filepath"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func handleFileUpdate(message map[string]interface{}, ws *models.WebsocketConnection) error {
	filename := "player" + ws.Language.GetExtension()
	return ioutil.WriteFile(ws.WorkspacePath+"/"+filename, []byte(message["code"].(string)), 0644)
}

func getAbsPath(ws *models.WebsocketConnection) error {
	absFolderPath, err := filepath.Abs(ws.WorkspacePath)
	responseBody := make(map[string]interface{})
	if err != nil {
		return SendErrorMessage(ws, err)
	}
	responseBody["status"] = "success"
	responseBody["folderpath"] = absFolderPath
	responseBody["filepath"] = absFolderPath + "/player" + ws.Language.GetExtension()
	return SendMessage(ws, responseBody)
}
