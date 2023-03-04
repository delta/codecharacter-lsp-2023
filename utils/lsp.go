package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/delta/codecharacter-lsp-2023/models"
	// "github.com/delta/codecharacter-lsp-2023/controllers"
)

func CreateLSPServer(ws *models.WebsocketConnection) error {
	switch ws.Language {
	case models.Cpp:
		return createCppServer(ws)
	}
	return nil
}

func createCppServer(ws *models.WebsocketConnection) error {
	var err error
	filename := "compile_commands.json"
	err2 := ioutil.WriteFile(ws.WorkspacePath+"/"+filename, []byte(createCompileCommands(ws)), 0644)
	if(err2 != nil){
		fmt.Println(err)
	}
	// var initialDir string
	// initialDir, err = os.Getwd()
	// if err != nil {
	// 	fmt.Println("Error getting initial directory:", err)
	// 	os.Exit(1)
	// }

	// Change the working directory to the project directory
	// projectDir := fmt.Sprintf("./%s",ws.WorkspacePath)
	// err = os.Chdir(projectDir)
	// if err != nil {
	// 	fmt.Println("Error changing directory:", err)
	// }

	// Generate compile_commands.json file with CMake
	// cmakeCmd := exec.Command("cmake", "-DCMAKE_EXPORT_COMPILE_COMMANDS=1", ".")
	// cmakeCmd.Stderr = os.Stderr
	// cmakeCmd.Stdout = os.Stdout
	// err = os.Chdir(initialDir)
	// if err != nil {
	// 	fmt.Println("Error generating compile_commands.json:", err)
	// }

	// Change the working directory back to the initial directory
	// if err != nil {
	// 	fmt.Println("Error changing directory:", err)
	// }

	// fmt.Println("Project built successfully!")
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

func createCompileCommands(ws* models.WebsocketConnection) string {
	compilejson := fmt.Sprintf(`[
		{
		  "directory": "%[1]s",
		  "command": "/usr/bin/c++  -I%[1]s/player_code.h  -o CMakeFiles/MyProject.dir/main.cpp.o -c %[1]s/main.cpp",
		  "file": "%[1]s/main.cpp"
		},
		{
		  "directory": "%[1]s",
		  "command": "/usr/bin/c++  -I%[1]s/player_code.h  -o CMakeFiles/MyProject.dir/player_code.cpp.o -c %[1]s/player_code.cpp",
		  "file": "%[1]s/player_code.cpp"
		},
		{
		  "directory": "%[1]s",
		  "command": "/usr/bin/c++  -I%[1]s/player_code.h  -o CMakeFiles/MyProject.dir/player.cpp.o -c %[1]s/player.cpp",
		  "file": "%[1]s/player.cpp"
		}
		]`,ws.WorkspacePath)
	fmt.Println("compile commands is ",compilejson)
	return compilejson
}