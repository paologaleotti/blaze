package api

import (
	"blaze/pkg/httpcore"
	"blaze/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func (tc *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) httpcore.Reply[[]*models.Todo] {
	return httpcore.ReplyData(tc.db, http.StatusOK)
}

func (tc *TodoController) GetTodo(w http.ResponseWriter, r *http.Request) httpcore.Reply[models.Todo] {
	id := chi.URLParam(r, "id")

	for _, todo := range tc.db {
		if todo.Id == id {
			return httpcore.ReplyData(*todo, http.StatusOK)
		}
	}

	return httpcore.ReplyErr[models.Todo](httpcore.ErrNotFound)
}

func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) httpcore.Reply[models.Todo] {
	newTodo, err := httpcore.DecodeBody[models.NewTodo](w, r)
	if err != nil {
		return httpcore.ReplyErr[models.Todo](httpcore.ErrBadRequest.Msg(err))
	}

	todo := &models.Todo{
		Id:        uuid.New().String(),
		Title:     newTodo.Title,
		Completed: false,
	}

	tc.db = append(tc.db, todo)

	return httpcore.ReplyData(*todo, http.StatusCreated)
}
