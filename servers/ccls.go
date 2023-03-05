package servers

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func createCppServer(ws *models.WebsocketConnection) error {
	filename := "compile_commands.json"
	err := ioutil.WriteFile(ws.WorkspacePath+"/"+filename, []byte(createCompileCommands(ws)), 0644)
	if err != nil {
		return err
	}
	ws.LSPServer.Process = exec.Command("ccls", `--init={
		"index":{
		  "onChange":true,
		  "trackDependency":0,
		  "threads":2,
		  "comments":0
		},
		"cache":{
		  "retainInMemory":1
		},
		"diagnostics":{
		  "onSave":1500
		}
	  }`)
	ws.LSPServer.Stdin, err = ws.LSPServer.Process.StdinPipe()
	if err != nil {
		return err
	}
	ws.LSPServer.Stdout, err = ws.LSPServer.Process.StdoutPipe()
	if err != nil {
		return err
	}
	devnull, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
	ws.LSPServer.Process.Stderr = devnull
	ws.LSPServer.DevNullFd = devnull
	err = ws.LSPServer.Process.Start()
	return err
}

func createCompileCommands(ws *models.WebsocketConnection) string {
	return fmt.Sprintf(`[
		{
		  "directory": "%[1]s",
		  "command": "/usr/bin/c++ -std=c++17 -I%[1]s/player_code.h  -o CMakeFiles/MyProject.dir/main.cpp.o -c %[1]s/main.cpp",
		  "file": "%[1]s/main.cpp"
		},
		{
		  "directory": "%[1]s",
		  "command": "/usr/bin/c++ -std=c++17 -I%[1]s/player_code.h  -o CMakeFiles/MyProject.dir/player_code.cpp.o -c %[1]s/player_code.cpp",
		  "file": "%[1]s/player_code.cpp"
		},
		{
		  "directory": "%[1]s",
		  "command": "/usr/bin/c++ -std=c++17 -I%[1]s/player_code.h  -o CMakeFiles/MyProject.dir/player.cpp.o -c %[1]s/player.cpp",
		  "file": "%[1]s/player.cpp"
		}
		]`, ws.WorkspacePath)
}
