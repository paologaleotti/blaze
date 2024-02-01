package api

import (
	"blaze/pkg/httpcore"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func ApplyRoutes(router chi.Router, controller *TodoController) {
	staticDir := http.Dir(getStaticDirPath())
	staticHandler := http.FileServer(staticDir)
	router.Handle("/*", httpcore.ServeStaticFiles(
		staticDir,
		staticHandler,
		filepath.Join(getStaticDirPath(), "index.html"),
	))

	router.Get("/api/todos", httpcore.Handle(controller.GetTodos))
	router.Get("/api/todos/{id}", httpcore.Handle(controller.GetTodo))
	router.Post("/api/todos", httpcore.Handle(controller.CreateTodo))
}

func getStaticDirPath() string {
	entrypoint, err := os.Executable()
	if err != nil {
		panic(err)
	}
	binDir := filepath.Dir(entrypoint)
	return filepath.Join(binDir, "static")
}
