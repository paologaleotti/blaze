package httpcore

import (
	"net/http"

	"github.com/go-chi/render"
)

func DecodeBody[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var payload T
	err := render.DecodeJSON(r.Body, &payload)

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrBadRequest)
		return payload, err
	}
	return payload, nil
}
