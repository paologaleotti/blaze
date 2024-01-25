package httpcore

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

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

// GetQueryInt64 retrieves a 64-bit integer query parameter from the request.
// It returns the integer value and any error encountered during conversion.
func GetQueryInt64(r *http.Request, param string) (int64, error) {
	query := r.URL.Query().Get(param)
	if query == "" {
		return 0, nil
	}

	v, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		return 0, errors.Join(ErrInvalidQuery, err)
	}
	return v, nil
}

// GetQueryInt64Opt retrieves a 64-bit integer query parameter from the request.
// It returns the integer value and any error encountered during conversion.
func GetQueryInt64Opt(r *http.Request, param string) (*int64, error) {
	query := r.URL.Query().Get(param)
	if query == "" {
		return nil, nil
	}

	v, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		return nil, errors.Join(ErrInvalidQuery, err)
	}
	return &v, nil
}

// GetQueryArrayStr retrieves an array of string from a comma-separated query parameter.
func GetQueryArrayStr(r *http.Request, param string) []string {
	query := r.URL.Query().Get(param)
	if query == "" {
		return nil
	}
	return strings.Split(query, ",")
}
