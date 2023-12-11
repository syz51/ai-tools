package main

import (
	"ai-web/internal/config"
	"ai-web/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	cfg, err := config.LoadConfig("configs")
	if err != nil {
		e.Logger.Fatal("error loading config: ", err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	// Routes
	handler := handlers.NewHandler(cfg)
	e.GET("/", handler.Index)
	e.POST("/translate", handler.Translate)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
