package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/leinodev/go-microservice-template/internal/core/models"
	"net/http"
	"strconv"
)

type (
	ExampleHandler struct {
		core exampleCore
	}
	exampleCore interface {
		Execute(ctx context.Context, id int64) (*models.Example, error)
	}
)

func NewExampleHandler(core exampleCore) *ExampleHandler {
	return &ExampleHandler{core: core}
}

func (h *ExampleHandler) getByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	resp, err := h.core.Execute(c.Context(), int64(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.JSON(resp)
}
