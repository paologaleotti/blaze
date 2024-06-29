package httpcore

import (
	"net/http"

	"github.com/go-chi/cors"
)

var DefaultCorsOptions cors.Options = cors.Options{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{
		http.MethodHead,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	},
	AllowedHeaders:   []string{"*"},
	AllowCredentials: true,
	MaxAge:           7200,
}
