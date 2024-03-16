package httpcore

import (
	"net/http"

	"github.com/go-chi/render"
)

type Handler func(w http.ResponseWriter, r *http.Request) (any, int)

// Handle wraps a Handler and takes care of rendering to JSON the response
func Handle(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, statusCode := handler(w, r)
		render.Status(r, statusCode)

		if data == nil {
			render.NoContent(w, r)
			return
		}
		render.JSON(w, r, data)
	}
}
