package httpcore

import "net/http"

type ApiError struct {
	// Title is a short, machine-readable string that describes the error.
	Title string `json:"title"`
	// Message is a human-readable string that describes the error.
	Message string `json:"message"`
	// Status is the HTTP status code that should be returned with the error.
	Status int `json:"status"`
}

// Msg returns a new ApiError with the given string appended to the message.
func (e ApiError) Msg(msg string) ApiError {
	return ApiError{
		Title:   e.Title,
		Message: e.Message + ": " + msg,
		Status:  e.Status,
	}
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

var (
	ErrBadRequest = ApiError{
		Title:   "BAD_REQUEST",
		Message: "The request is invalid",
		Status:  http.StatusBadRequest,
	}

	ErrUnauthorized = ApiError{
		Title:   "UNAUTHORIZED",
		Message: "You are not authorized",
		Status:  http.StatusUnauthorized,
	}

	ErrForbidden = ApiError{
		Title:   "FORBIDDEN",
		Message: "Access to this resource is forbidden",
		Status:  http.StatusForbidden,
	}

	ErrConflict = ApiError{
		Title:   "CONFLICT",
		Message: "Conflict with resources",
		Status:  http.StatusConflict,
	}

	ErrNotFound = ApiError{
		Title:   "NOT_FOUND",
		Message: "The requested resource was not found",
		Status:  http.StatusNotFound,
	}

	ErrUnkownInternal = ApiError{
		Title:   "UNKOWN_INTERNAL",
		Message: "An internal server error occurred",
		Status:  http.StatusInternalServerError,
	}

	ErrImATeapot = ApiError{
		Title:   "IM_A_TEAPOT",
		Message: "I'm a teapot",
		Status:  http.StatusTeapot,
	}
)
