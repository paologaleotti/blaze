package api

import (
	"blaze/pkg/httpcore"
	"blaze/pkg/util"

	"github.com/go-chi/chi/v5"
)

func InitService() chi.Router {
	util.InitLogger()

	router := httpcore.NewRouter()
	// env := api.InitEnv() // get typed environment
	controller := NewTodoController()
	ApplyRoutes(router, controller)

	return router
}
