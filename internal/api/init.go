package api

import (
	"blaze/pkg/httpcore"

	"github.com/go-chi/chi/v5"
)

func InitService() chi.Router {
	router := httpcore.NewRouter()
	//env := api.InitEnv() // get typed environment
	controller := NewTodoController()
	ApplyRoutes(router, controller)

	return router
}
