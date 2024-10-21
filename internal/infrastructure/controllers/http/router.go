package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leinodev/go-microservice-template/internal/core"
	"github.com/leinodev/go-microservice-template/internal/core/models"
)

func Register(server *fiber.App, core core.Core[int64, *models.Example]) {
	api := server.Group("/v1")
	{
		exampleHandler := NewExampleHandler(core)
		eHandler := api.Group("/examples")
		eHandler.Get("/:id", exampleHandler.getByID)
	}
}
