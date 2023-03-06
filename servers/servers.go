package servers

import (
	"os"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func CreateLSPServer(ws *models.WebsocketConnection) error {
	switch ws.Language {
	case models.Cpp:
		return createCppServer(ws)
	case models.Python:
		return createPythonServer(ws)
	case models.Java:
		return createJavaServer(ws)
	}
	return nil
}

func createPipes(ws *models.WebsocketConnection) error {
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
	return ws.LSPServer.Process.Start()
}
