package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func newFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "Stock Info API",
		ReadTimeout:           30 * time.Second,
		WriteTimeout:          0, // Disabled for SSE streaming
		IdleTimeout:           120 * time.Second,
		DisableKeepalive:      false,
		StreamRequestBody:     true,
		DisableStartupMessage: false,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	return app
}

func startServer(app *fiber.App, port string) {
	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
