package models

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebsocketConnection struct {
	ID            uuid.UUID
	Connection    *websocket.Conn
	Language      Language
	WorkspacePath string
	LSPServer     LSPServer
}
