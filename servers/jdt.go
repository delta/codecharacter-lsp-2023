package servers

import (
	"os/exec"
	"path/filepath"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func createJavaServer(ws *models.WebsocketConnection) error {
	workspaceDir, err := filepath.Abs(ws.WorkspacePath)
	if err != nil {
		return err
	}
	ws.LSPServer.Process = exec.Command("java",
		"-Declipse.application=org.eclipse.jdt.ls.core.id1",
		"-Dosgi.bundles.defaultStartLevel=4",
		"-Declipse.product=org.eclipse.jdt.ls.core.product",
		"-Dlog.level=ERROR",
		"-noverify",
		"-Xmx512M",
		"-jar",
		"/jdt/plugins/org.eclipse.equinox.launcher_1.6.400.v20210924-0641.jar",
		"-configuration",
		"/jdt/config_linux",
		"-data",
		workspaceDir,
	)
	return createPipes(ws)
}
