package todoservice

import (
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
		if todo.Id.String() == id {
			render.JSON(w, r, todo)
			return
		}
	}

	render.Status(r, http.StatusNotFound)
	render.PlainText(w, r, "Todo not found")
}

func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.NewTodo
	if err := render.DecodeJSON(r.Body, &newTodo); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, "Failed to decode request body")
		return
	}

	todo := &models.Todo{
		Id:        uuid.New(),
		Title:     newTodo.Title,
		Completed: false,
	}

	tc.db = append(tc.db, todo)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, todo)
}
