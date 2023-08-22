package main

import (
	todoservice "blaze/internal/todo-service"
	"blaze/pkg/httpcore"
	"log"
	"net/http"
)

func main() {
	router := httpcore.NewRouter()
	controller := todoservice.NewTodoController()

	todoservice.ApplyRoutes(router, controller)

	log.Print("Server started")
	http.ListenAndServe(":3000", router)
}
