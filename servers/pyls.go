package servers

import (
	"os"
	"os/exec"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func createPythonServer(ws *models.WebsocketConnection) error {
	ws.LSPServer.Process = exec.Command("pyls", "-v")
	var err error
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
