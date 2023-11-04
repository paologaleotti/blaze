package todoservice

import (
	"blaze/pkg/httpcore"
	"blaze/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type TodoController struct {
	db []*models.Todo
}

func NewTodoController() *TodoController {
	return &TodoController{
		db: make([]*models.Todo, 0),
	}
}

func (tc *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, tc.db)
}

func (tc *TodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, todo := range tc.db {
		if todo.Id == id {
			render.JSON(w, r, todo)
			return
		}
	}

	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, httpcore.ErrNotFound)
}

func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	newTodo, err := httpcore.DecodeBody[models.NewTodo](w, r)
	if err != nil {
		return
	}

	todo := &models.Todo{
		Id:        uuid.New().String(),
		Title:     newTodo.Title,
		Completed: false,
	}

	tc.db = append(tc.db, todo)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, todo)
}
