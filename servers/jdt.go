package servers

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func createJavaServer(ws *models.WebsocketConnection) error {
	workspaceDir, _ := filepath.Abs(ws.WorkspacePath)
	ws.LSPServer.Process = exec.Command("java",
		"-Declipse.application=org.eclipse.jdt.ls.core.id1",
		"-Dosgi.bundles.defaultStartLevel=4",
		"-Dosgi.bundles.defaultStartLevel=4",
		"-Declipse.product=org.eclipse.jdt.ls.core.product",
		"-Dlog.level=ERROR",
		"-noverify",
		"-Xmx100M",
		"-jar",
		"/jdt/plugins/org.eclipse.equinox.launcher_1.6.400.v20210924-0641.jar",
		"-configuration",
		"/jdt/config_linux",
		"-data",
		workspaceDir,
	)
	var err error
	ws.LSPServer.Stdin, err = ws.LSPServer.Process.StdinPipe()
	if err != nil {
		return err
	}
	ws.LSPServer.Stdout, err = ws.LSPServer.Process.StdoutPipe()
	if err != nil {
		return err
	}
	ws.LSPServer.Process.Stderr = os.Stderr
	devnull, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
	ws.LSPServer.Process.Stderr = devnull
	ws.LSPServer.DevNullFd = devnull
	err = ws.LSPServer.Process.Start()
	return err
}
