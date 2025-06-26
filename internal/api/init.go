package api

import (
	"blaze/db"
	"blaze/internal/api/domain"
	"blaze/internal/api/routers"
	"blaze/internal/api/services"
	"blaze/pkg/httpx"
	"blaze/pkg/storage"
	"blaze/pkg/util"
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func InitService() http.Handler {
	util.InitLogger()
	env := InitEnv()
	router := chi.NewRouter()

	router.Use(cors.New(httpx.DefaultCorsOptions).Handler)
	router.Use(middleware.Timeout(20 * time.Second))
	router.Use(middleware.Recoverer)
	router.Use(httpx.LoggerMiddleware)
	router.NotFound(httpx.NotFoundHandler)

	dbClient, err := sql.Open("sqlite3", env.DatabaseUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to database")
	}

	db.RunMigrations(dbClient)

	todoDb := storage.NewTodoStorage(dbClient)

	todoService := services.NewTodoService(todoDb)

	routers.MountTodoRoutes(router, serviceErrorMap, todoService)

	return router
}

var serviceErrorMap = map[error]httpx.ApiError{
	httpx.ErrInvalidBody:   httpx.ErrBadRequest,
	domain.ErrTodoNotFound: httpx.ErrNotFound,
}
