package httpcore

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func LoggerMiddleware(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)

			log.Info("Request completed",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.Duration("duration", time.Since(startTime)),
			)
		})
	}
}
