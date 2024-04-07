package api

import (
	"blaze/pkg/httpcore"

	"github.com/go-chi/chi/v5"
)

func applyRoutes(router chi.Router, controller *ApiController) {
	router.Get("/todos", httpcore.Handle(controller.GetTodos))
	router.Get("/todos/{id}", httpcore.Handle(controller.GetTodo))
	router.Post("/todos", httpcore.Handle(controller.CreateTodo))
}
