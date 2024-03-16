package httpcore

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(cors.AllowAll().Handler)
	router.Use(middleware.Timeout(20 * time.Second))
	router.Use(middleware.Recoverer)
	router.Use(LoggerMiddleware)

	return router
}
