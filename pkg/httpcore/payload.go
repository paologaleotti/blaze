package httpcore

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

var ErrInvalidQuery = errors.New("invalid query parameter type")

var validate = validator.New(validator.WithRequiredStructEnabled())

// DecodeBody decodes the request body into a given type T,
// validates it, and returns the decoded payload and any error encountered during decoding.
// It renders the JSON response to the given http.ResponseWriter and http.Request.
func DecodeBody[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var payload T
	errDecode := render.DecodeJSON(r.Body, &payload)
	errValidate := validate.Struct(payload)
	err := errors.Join(errDecode, errValidate)

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrBadRequest.Msg(err))
		return payload, err
	}
	return payload, nil
}

// GetQueryStr retrieves a string query parameter from the request.
func GetQueryStr(r *http.Request, param string) string {
	return r.URL.Query().Get(param)
}

// GetQueryInt retrieves an integer query parameter from the request.
// It returns the integer value and any error encountered during conversion.
func GetQueryInt(r *http.Request, param string) (int, error) {
	query := r.URL.Query().Get(param)
	v, err := strconv.Atoi(query)
	if err != nil {
		return 0, errors.Join(ErrInvalidQuery, err)
	}
	return v, nil
}
