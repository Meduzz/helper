package herror

import (
	"errors"
	"fmt"
	"net/http"
)

// HttpError definition
type (
	HttpError struct {
		Code    int
		Message string
	}
)

func (e HttpError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

// NewHttpError takes a httpy code and message and returns an error
func NewHttpError(code int, message string) error {
	return HttpError{
		Code:    code,
		Message: message,
	}
}

// ErrorFromCode takes a http.StatusX code and returns an error with message from http.StatusMessage
func ErrorFromCode(code int) error {
	return NewHttpError(code, http.StatusText(code))
}

// CodeFroomError attempts to cast err to a HttpError and dig out the code, or returns 500
func CodeFromError(err error) int {
	herr := &HttpError{}

	// try to cast err to a HttpError and fetch the code
	if errors.As(err, herr) {
		return herr.Code
	}

	// else return 500
	return 500
}

// IsError will check if a code is over 399 and then create an error for that code or return nil.
func IsError(code int) error {
	if code > 399 {
		return ErrorFromCode(code)
	}

	return nil
}

// common error variants
var (
	ErrBadRequest         = ErrorFromCode(http.StatusBadRequest)
	ErrUnauthorized       = ErrorFromCode(http.StatusUnauthorized)
	ErrForbidden          = ErrorFromCode(http.StatusForbidden)
	ErrNotFound           = ErrorFromCode(http.StatusNotFound)
	ErrConflict           = ErrorFromCode(http.StatusConflict)
	ErrInternalError      = ErrorFromCode(http.StatusInternalServerError)
	ErrServiceUnavailable = ErrorFromCode(http.StatusServiceUnavailable)
)
