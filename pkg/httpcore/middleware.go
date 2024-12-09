package httpcore

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		log.Info().
			Stringer("url", r.URL).
			Str("method", r.Method).
			Int("status", ww.Status()).
			Dur("duration", time.Since(startTime)).
			Msg("Request completed")
	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, ErrNotFound.With(errors.New("Route not mounted")))
}
