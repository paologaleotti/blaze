package api

import (
	"blaze/internal/api/routers"
	"blaze/internal/api/services"
	"blaze/pkg/httpcore"
	"blaze/pkg/models"
	"blaze/pkg/util"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitService() http.Handler {
	util.InitLogger()

	router := chi.NewRouter()

	router.Use(cors.New(httpcore.DefaultCorsOptions).Handler)
	router.Use(middleware.Timeout(20 * time.Second))
	router.Use(middleware.Recoverer)
	router.Use(httpcore.LoggerMiddleware)

	// env := InitEnv() // get typed environment

	db := make([]models.Todo, 0)

	todoService := services.NewTodoService(&db)

	routers.ApplyTodoRoutes(router, todoService)

	return router
}
