package main

import (
	"blaze/internal/todoservice"
	"blaze/pkg/httpcore"
	"blaze/pkg/util"
	"net/http"
)

func main() {
	util.InitLogger()
	router := httpcore.NewRouter()
	//env := todoservice.InitEnv() // get typed environment
	controller := todoservice.NewTodoController()

	todoservice.ApplyRoutes(router, controller)

	util.Log.Info("Server started at http://localhost:3000")
	http.ListenAndServe("0.0.0.0:3000", router)
}
