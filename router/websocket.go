package router

import (
	"github.com/delta/codecharacter-lsp-2023/utils"
	"github.com/labstack/echo/v4"
)

func handleWebSocketConnection(c echo.Context) error {
	return utils.InitWebsocket(c)
}
