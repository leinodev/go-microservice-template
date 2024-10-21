package example

import (
	"context"
	"github.com/leinodev/go-microservice-template/internal/core"
	"github.com/leinodev/go-microservice-template/internal/core/models"
)

type (
	GetExampleByID struct {
		repo exampleRepo
	}
	exampleRepo interface {
		GetByID(id int64) (*models.Example, error)
	}
)

func NewGetExampleByIDCore(repo exampleRepo) core.Core[int64, *models.Example] {
	return &GetExampleByID{repo: repo}
}

func (core *GetExampleByID) Execute(ctx context.Context, id int64) (*models.Example, error) {
	return core.repo.GetByID(id)
}
