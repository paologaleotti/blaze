package api

import (
	"blaze/internal/api/domain"
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
	router.NotFound(httpcore.NotFoundHandler)

	// env := InitEnv() // get typed environment

	// TODO sqlite db?
	db := make([]models.Todo, 0)

	todoService := services.NewTodoService(&db)

	routers.ApplyTodoRoutes(router, serviceErrorMap, todoService)

	return router
}

var serviceErrorMap = map[error]httpcore.ApiError{
	httpcore.ErrInvalidBody: httpcore.ErrBadRequest,
	domain.ErrTodoNotFound:  httpcore.ErrNotFound,
}
