package routers

import (
	"blaze/internal/api/services"
	"blaze/pkg/httpcore"
	"blaze/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// TODO pass context

func ApplyTodoRoutes(router chi.Router, errMap httpcore.ApiErrorMap, todoService *services.TodoService) {
	router.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos := todoService.GetTodos()
		render.JSON(w, r, todos)
	})

	router.Get("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		todo, err := todoService.GetTodoById(id)

		if err != nil {
			httpcore.RenderError(w, r, errMap, err)
			return
		}
		render.JSON(w, r, todo)
	})

	router.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		newTodo, err := httpcore.DecodeBody[models.NewTodo](r)
		if err != nil {
			httpcore.RenderError(w, r, errMap, err)
			return
		}

		todo := todoService.CreateTodo(newTodo)
		render.JSON(w, r, todo)
	})
}
