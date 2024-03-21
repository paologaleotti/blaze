package api

import (
	"blaze/pkg/httpcore"
	"blaze/pkg/util"
	"net/http"
)

func InitService() http.Handler {
	util.InitLogger()

	router := httpcore.NewRouter()
	// env := api.InitEnv() // get typed environment
	controller := NewTodoController()
	ApplyRoutes(router, controller)

	return router
}
