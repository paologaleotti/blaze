package httpx

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type ApiError struct {
	// Title is a short, machine-readable string that describes the error.
	Title string `json:"title"`
	// Message is a human-readable string that describes the error.
	Message string `json:"message"`
	// Status is the HTTP status code that should be returned with the error.
	Status int `json:"status"`
}

// With returns a new ApiError with the given error appended to the message.
func (e ApiError) With(err error) ApiError {
	if err == nil {
		return e
	}
	return ApiError{
		Title:   e.Title,
		Message: e.Message + ": " + err.Error(),
		Status:  e.Status,
	}
}

// ApiErrorMap is a map of standard errors to their corresponding ApiError.
type ApiErrorMap map[error]ApiError

// RenderError renders the error wrapped in a ApiError to the client, according to the given error mapping.
// If the error is not found in the map, a generic 500 Internal Server Error is returned.
//
// In a handler, remember to return right after calling this function to prevent further processing.
func RenderError(w http.ResponseWriter, r *http.Request, errMap ApiErrorMap, err error) {
	apiErr := ApiError{}
	for k, v := range errMap {
		if errors.Is(err, k) {
			apiErr = v
			break
		} else {
			apiErr = ErrUnkownInternal
		}
	}

	render.Status(r, apiErr.Status)
	render.JSON(w, r, apiErr.With(err))
}

var (
	ErrBadRequest = ApiError{
		Title:   "BadRequest",
		Message: "The request is invalid",
		Status:  http.StatusBadRequest,
	}

	ErrUnauthorized = ApiError{
		Title:   "Unauthorized",
		Message: "You are not authorized",
		Status:  http.StatusUnauthorized,
	}

	ErrForbidden = ApiError{
		Title:   "Forbidden",
		Message: "Access to this resource is forbidden",
		Status:  http.StatusForbidden,
	}

	ErrConflict = ApiError{
		Title:   "Conflict",
		Message: "Conflict with resources",
		Status:  http.StatusConflict,
	}

	ErrNotFound = ApiError{
		Title:   "NotFound",
		Message: "The requested resource was not found",
		Status:  http.StatusNotFound,
	}

	ErrUnkownInternal = ApiError{
		Title:   "UnknownInternal",
		Message: "An internal server error occurred",
		Status:  http.StatusInternalServerError,
	}

	ErrImATeapot = ApiError{
		Title:   "ImATeapot",
		Message: "I'm a teapot",
		Status:  http.StatusTeapot,
	}
)
