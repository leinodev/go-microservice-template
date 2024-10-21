package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leinodev/go-microservice-template/config"
)

// BuildServer returns instance of http server we are using in all controllers.
func BuildServer(cfg config.Service, middlewares ...any) *fiber.App {
	server := fiber.New(fiber.Config{
		AppName: cfg.Name,
	})

	if len(middlewares) > 0 {
		server.Use(middlewares...)
	}

	return server
}
