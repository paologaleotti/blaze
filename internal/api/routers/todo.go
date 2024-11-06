package routers

import (
	"blaze/internal/api/services"
	"blaze/pkg/httpcore"
	"blaze/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ApplyTodoRoutes(router chi.Router, todoService *services.TodoService) {
	router.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos := todoService.GetTodos()
		render.JSON(w, r, todos)
	})

	router.Get("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		todo, err := todoService.GetTodoById(id)

		// TODO HANDLE ERRORS
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, render.M{"error": err.Error()})
			return
		}
		render.JSON(w, r, todo)
	})

	router.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		newTodo, err := httpcore.DecodeBody[models.NewTodo](r)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, render.M{"error": err.Error()})
			return
			// TODO HANDLE ERRORS
		}

		todo := todoService.CreateTodo(newTodo)
		render.JSON(w, r, todo)
	})
}
