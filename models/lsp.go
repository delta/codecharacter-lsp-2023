package models

import (
	"io"
	"os"
	"os/exec"
)

type LSPServer struct {
	Process   *exec.Cmd
	Stdin     io.WriteCloser
	Stdout    io.ReadCloser
	DevNullFd *os.File
}
