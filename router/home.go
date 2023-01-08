package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.String(http.StatusAccepted, "Hello! Welcome to the CodeCharacter LSP Project")
}
