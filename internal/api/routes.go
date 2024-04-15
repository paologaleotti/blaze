package api

import (
	"blaze/internal/api/handlers"
	"blaze/pkg/httpcore"

	"github.com/go-chi/chi/v5"
)

func applyRoutes(router chi.Router, ctl *handlers.ApiController) {
	router.Get("/todos", httpcore.Handle(ctl.GetTodos))
	router.Get("/todos/{id}", httpcore.Handle(ctl.GetTodo))
	router.Post("/todos", httpcore.Handle(ctl.CreateTodo))
}
