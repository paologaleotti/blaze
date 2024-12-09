package api

import (
	"blaze/internal/api/domain"
	"blaze/internal/api/routers"
	"blaze/internal/api/services"
	"blaze/pkg/httpx"
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

	router.Use(cors.New(httpx.DefaultCorsOptions).Handler)
	router.Use(middleware.Timeout(20 * time.Second))
	router.Use(middleware.Recoverer)
	router.Use(httpx.LoggerMiddleware)
	router.NotFound(httpx.NotFoundHandler)

	// env := InitEnv() // get typed environment

	// TODO sqlite db?
	db := make([]models.Todo, 0)

	todoService := services.NewTodoService(&db)

	routers.ApplyTodoRoutes(router, serviceErrorMap, todoService)

	return router
}

var serviceErrorMap = map[error]httpx.ApiError{
	httpx.ErrInvalidBody:   httpx.ErrBadRequest,
	domain.ErrTodoNotFound: httpx.ErrNotFound,
}
