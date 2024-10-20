package example

import (
	"context"
	"github.com/leinodev/go-microservice-template/internal/core/models"
)

type (
	GetExampleByID struct {
	}
)

func (core *GetExampleByID) Execute(ctx context.Context, id int64) (models.Example, error) {

}
