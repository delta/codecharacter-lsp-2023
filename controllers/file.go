package controllers

import (
	"fmt"
	"io/ioutil"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func handleFileUpdate(message map[string]interface{}, ws *models.WebsocketConnection) error {
	fmt.Println("Processing File Update Request")
	var filename string
	switch ws.Language {
	case "cpp":
		filename = "run.cpp"
	case "java":
		filename = "run.java"
	case "python":
		filename = "run.py"
	}
	err := ioutil.WriteFile(ws.WorkspacePath+"/"+filename, []byte(message["code"].(string)), 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
