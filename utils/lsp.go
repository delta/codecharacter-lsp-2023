package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func CreateLSPServer(ws *models.WebsocketConnection) error {
	switch ws.Language {
	case models.Cpp:
		return createCppServer(ws)
	}
	return nil
}

func createCppServer(ws *models.WebsocketConnection) error {
	ws.LSPServer.Process = exec.Command("ccls", `--init={
		"index":{
		  "onChange":true,
		  "trackDependency":0,
		  "threads":2,
		  "comments":0
		},
		"cache":{
		  "retainInMemory":0,
		  "directory":"./`+ws.WorkspacePath+`"
		},
		"diagnostics":{
		  "onSave":1500
		}
	  }`)
	var err error
	ws.LSPServer.Process.Stderr = os.Stderr
	ws.LSPServer.Stdin, err = ws.LSPServer.Process.StdinPipe()
	if err != nil {
		fmt.Println(err)
		return err
	}
	ws.LSPServer.Stdout, err = ws.LSPServer.Process.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ws.LSPServer.Process.Start()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
