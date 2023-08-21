package main

import (
	helloservice "blaze/internal/hello-service"
	"blaze/pkg/httpcore"
	"net/http"
)

func main() {
	router := httpcore.NewRouter()
	controller := helloservice.NewHelloController("hello")

	helloservice.ApplyRoutes(router, controller)
	http.ListenAndServe(":3000", router)
}
