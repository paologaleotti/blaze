package handlers

import (
	"blaze/pkg/httpcore"
	"blaze/pkg/models"
	"net/http"

	"github.com/google/uuid"
)

func (tc *ApiController) GetTodos(w http.ResponseWriter, r *http.Request) (any, int) {
	return tc.db, http.StatusOK
}

func (tc *ApiController) GetTodo(w http.ResponseWriter, r *http.Request) (any, int) {
	id := r.PathValue("id")

	for _, todo := range tc.db {
		if todo.Id == id {
			return todo, http.StatusOK
		}
	}

	return httpcore.ErrNotFound.Msg("todo not found"), http.StatusNotFound
}

func (tc *ApiController) CreateTodo(w http.ResponseWriter, r *http.Request) (any, int) {
	newTodo, err := httpcore.DecodeBody[models.NewTodo](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	todo := &models.Todo{
		Id:        uuid.New().String(),
		Title:     newTodo.Title,
		Completed: false,
	}

	tc.db = append(tc.db, todo)

	return todo, http.StatusCreated
}
