package main

import (
	"ai-web/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	// Routes
	e.GET("/", handlers.IndexHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
