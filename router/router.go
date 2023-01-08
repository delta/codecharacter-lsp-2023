package router

import "github.com/labstack/echo/v4"

func InitRoutes(server *echo.Echo) {
	server.GET("/", home)
	server.GET("/ws", handleWebSocketConnection)
}
