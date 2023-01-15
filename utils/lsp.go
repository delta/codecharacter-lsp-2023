package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func CreateLSPServer(wsConnectionParams *models.WebsocketConnection) error {
	switch wsConnectionParams.Language {
	case "cpp":
		return createCppServer(wsConnectionParams)
	}
	return nil
}

func createCppServer(wsConnectionParams *models.WebsocketConnection) error {
	wsConnectionParams.LSPServer = exec.Command("ccls", `--init={
		"index":{
		  "onChange":true,
		  "trackDependency":0,
		  "threads":2,
		  "comments":0
		},
		"cache":{
		  "retainInMemory":0,
		  "directory":"./`+wsConnectionParams.WorkspacePath+`"
		},
		"diagnostics":{
		  "onSave":1500
		}
	  }`)
	wsConnectionParams.LSPServer.Stderr = os.Stderr
	err := wsConnectionParams.LSPServer.Start()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
