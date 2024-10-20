package http

import (
	"context"
	"encoding/json"
	"github.com/leinodev/go-microservice-template/internal/core/models"
	"net/http"
	"strconv"
)

type (
	ExampleHandler struct {
		core ExampleCore
	}
	ExampleCore interface {
		Execute(ctx context.Context, id int64) (models.Example, error)
	}
)

func (h *ExampleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id must be an integer"))
		return
	}
	m, err := h.core.Execute(r.Context(), int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	resp, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}
