package utils

import "github.com/labstack/echo/v4"

func SendResponse(c echo.Context, code int, message interface{}) error {
	return c.JSON(code, map[string]interface{}{
		"message": message,
	})
}
