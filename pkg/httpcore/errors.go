package httpcore

import "net/http"

type ApiError struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e ApiError) Msg(err error) ApiError {
	if err != nil {
		e.Message += ": " + err.Error()
	}
	return e
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
