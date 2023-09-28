package main

import (
	"blaze/internal/todoservice"
	"blaze/pkg/httpcore"
	"log"
	"net/http"
)

func main() {
	router := httpcore.NewRouter()
	controller := todoservice.NewTodoController()

	todoservice.ApplyRoutes(router, controller)

	log.Print("Server started")
	http.ListenAndServe("0.0.0.0:3000", router)
}
