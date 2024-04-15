package api

import (
	"blaze/internal/api/handlers"
	"blaze/pkg/httpcore"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func applyRoutes(router chi.Router, ctl *handlers.ApiController) {
	staticDir := http.Dir(getStaticDirPath())
	staticHandler := http.FileServer(staticDir)
	router.Handle("/*", httpcore.ServeStaticFiles(
		staticDir,
		staticHandler,
		filepath.Join(getStaticDirPath(), "index.html"),
	))

	router.Get("/api/todos", httpcore.Handle(ctl.GetTodos))
	router.Get("/api/todos/{id}", httpcore.Handle(ctl.GetTodo))
	router.Post("/api/todos", httpcore.Handle(ctl.CreateTodo))
}

func getStaticDirPath() string {
	entrypoint, err := os.Executable()
	if err != nil {
		panic(err)
	}
	binDir := filepath.Dir(entrypoint)
	return filepath.Join(binDir, "static")
}
