package main

import (
	"github.com/delta/codecharacter-lsp-2023/config"
	"github.com/delta/codecharacter-lsp-2023/router"
	"github.com/delta/codecharacter-lsp-2023/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	server := echo.New()
	utils.InitLogger(server)
	config.InitConfig()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.FrontendURL},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	router.InitRoutes(server)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
