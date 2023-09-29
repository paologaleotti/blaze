package httpcore

import (
	"blaze/pkg/util"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(LoggerMiddleware(&util.RichLog))
	router.Use(middleware.Recoverer)

	return router
}
