package servers

import (
	"os/exec"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func createPythonServer(ws *models.WebsocketConnection) error {
	ws.LSPServer.Process = exec.Command("pylsp")
	return createPipes(ws)
}
