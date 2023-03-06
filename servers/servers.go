package servers

import (
	"github.com/delta/codecharacter-lsp-2023/models"
)

func CreateLSPServer(ws *models.WebsocketConnection) error {
	switch ws.Language {
	case models.Cpp:
		return createCppServer(ws)
		// case models.Python:
		// 	return createPythonServer(ws)
		// case models.Java:
		// 	return createPythonServer(ws)
	}
	return nil
}
