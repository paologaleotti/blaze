package httpcore

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		startTime := time.Now()
		next.ServeHTTP(w, r)

		log.Info().
			Str("message", "Request completed").
			Stringer("url", r.URL).
			Str("method", r.Method).
			Int("status", ww.Status()).
			Dur("duration", time.Since(startTime)).
			Msg("")
	})
}
