package servers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func createCppServer(ws *models.WebsocketConnection) error {
	filename := "compile_commands.json"
	err := os.WriteFile(ws.WorkspacePath+"/"+filename, []byte(createCompileCommands(ws)), 0644)
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
	return createPipes(ws)
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
