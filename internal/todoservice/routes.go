package todoservice

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func ApplyRoutes(router chi.Router, controller *TodoController) {
	staticFs := http.FileServer(http.Dir(getStaticDirPath()))
	router.Handle("/*", http.StripPrefix("/", staticFs))

	router.Get("/api/todos", controller.GetTodos)
	router.Get("/api/todos/{id}", controller.GetTodo)
	router.Post("/api/todos", controller.CreateTodo)
}

func getStaticDirPath() string {
	entrypoint, err := os.Executable()
	if err != nil {
		panic(err)
	}
	binDir := filepath.Dir(entrypoint)
	return filepath.Join(binDir, "static")
}
