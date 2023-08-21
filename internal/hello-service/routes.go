package helloservice

import "github.com/go-chi/chi/v5"

func ApplyRoutes(router chi.Router, controller *HelloController) {
	router.Get("/hello", controller.GetHello)
}
