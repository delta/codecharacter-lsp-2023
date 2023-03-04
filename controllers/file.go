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
	// v:= "Content-Length: 143"
	// u := fmt.Sprintf(`{"jsonrpc": "2.0","id": 1,"method": "initialize","params":{"rootUri":"file:///app/workspaces/%s/player.cpp"}}`,ws.ID)
	// res := v + "\n\n" + u
	// _, err := ws.LSPServer.Stdin.Write([]byte(res))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	err2 := ioutil.WriteFile(ws.WorkspacePath+"/"+filename, []byte(message["code"].(string)), 0644)
	return err2
}

func getAbsPath(ws *models.WebsocketConnection) error {
	// filename := "player" + ws.Language.GetExtension()
	abs, err := filepath.Abs(ws.WorkspacePath)
	responseBody := make(map[string]interface{})
	if err != nil {
		return SendErrorMessage(ws, err)
	}
	responseBody["status"] = "success"
	responseBody["message"] = abs
	err = SendMessage(ws, responseBody)
	return err
}
