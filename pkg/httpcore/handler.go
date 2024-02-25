package httpcore

import (
	"net/http"

	"github.com/go-chi/render"
)

// Reply is a generic response type that can return any Data or an Error.
type Reply[T any] struct {
	Data   *T
	Status int
	Error  *ApiError
}

// ReplyData returns a Reply with data and status.
func ReplyData[T any](data T, status int) Reply[T] {
	return Reply[T]{
		Data:   &data,
		Status: status,
		Error:  nil,
	}
}

// ReplyErr returns a Reply with an error and status.
func ReplyErr[T any](err ApiError) Reply[T] {
	return Reply[T]{
		Data:   nil,
		Status: err.Status,
		Error:  &err,
	}
}

// Handler is a generic handler function that can return any type.
type Handler[T any] func(w http.ResponseWriter, r *http.Request) Reply[T]

// Handle wraps a Handler and takes care of rendering to JSON the response.
func Handle[T any](handler Handler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reply := handler(w, r)
		render.Status(r, reply.Status)

		if reply.Error != nil {
			render.JSON(w, r, reply.Error)
			return
		}
		render.JSON(w, r, reply.Data)
	}
}
