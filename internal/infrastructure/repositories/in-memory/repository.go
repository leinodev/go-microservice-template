package in_memory

import (
	"errors"
	"fmt"
	"github.com/leinodev/go-microservice-template/internal/core/models"
	"time"
)

type ExampleRepository struct {
	data map[int64]models.Example
}

func NewExampleRepository() *ExampleRepository {
	return &ExampleRepository{
		data: map[int64]models.Example{
			1: models.Example{
				ID:        1,
				Name:      "first",
				CreatedAt: time.Now(),
			},
			2: models.Example{
				ID:        2,
				Name:      "second",
				CreatedAt: time.Now(),
			},
		},
	}
}

func (repo *ExampleRepository) GetByID(id int64) (*models.Example, error) {
	m, ok := repo.data[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("there is not example with id %d", id))
	}
	return &m, nil
}
