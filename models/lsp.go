package models

import (
	"io"
	"os/exec"
)

type LSPServer struct {
	Process *exec.Cmd
	Stdin   io.WriteCloser
	Stdout  io.ReadCloser
}
