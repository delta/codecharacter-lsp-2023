package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitLogger(server *echo.Echo) {
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} | ${method} ${uri} \t | ${latency_human}\n",
		Output: server.Logger.Output(),
	}))
}
