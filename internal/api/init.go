package api

import (
	"blaze/internal/api/handlers"
	"blaze/pkg/httpcore"
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

	router.Use(cors.AllowAll().Handler)
	router.Use(middleware.Timeout(20 * time.Second))
	router.Use(middleware.Recoverer)
	router.Use(httpcore.LoggerMiddleware)

	// env := api.InitEnv() // get typed environment

	controller := handlers.NewApiController()
	applyRoutes(router, controller)

	return router
}
