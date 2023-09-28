package todoservice

import "github.com/go-chi/chi/v5"

func ApplyRoutes(router chi.Router, controller *TodoController) {
	router.Get("/todos", controller.GetTodos)
	router.Get("/todos/{id}", controller.GetTodo)
	router.Post("/todos", controller.CreateTodo)
}
