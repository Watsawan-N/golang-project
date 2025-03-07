package errs

import (
	"errors"
	"golang-project/pkg/logs"
	"net/http"

	"github.com/labstack/gommon/log"
)

type Error struct {
	Err        error
	StatusCode int
}

func (e Error) Error() string {
	return e.Err.Error()
}

func NewInternalServerError(message string) *Error {
	logs.Error(message)
	return &Error{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(message),
	}
}

func NewNotImplementedError(message string) *Error {
	logs.Error(message)
	return &Error{
		StatusCode: http.StatusNotImplemented,
		Err:        errors.New(message),
	}
}

func NewNotFoundError(message string) *Error {
	logs.Error(message)
	return &Error{
		StatusCode: http.StatusNotFound,
		Err:        errors.New(message),
	}
}

func NewUnauthorizedError(message string) *Error {
	logs.Error(message)
	return &Error{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(message),
	}
}

func NewBadRequestError(message string) *Error {
	logs.Error(message)
	return &Error{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(message),
	}
}

func NewUnprocessableEntityError(message string) *Error {
	logs.Error(message)
	return &Error{
		StatusCode: http.StatusUnprocessableEntity,
		Err:        errors.New(message),
	}
}

func NewConflictError(message string) *Error {
	log.Error(message)
	return &Error{
		StatusCode: http.StatusConflict,
		Err:        errors.New(message),
	}
}
