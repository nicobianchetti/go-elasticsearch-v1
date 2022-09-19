package http

import (
	"fmt"
	"net/http"
)

const (
	TypeNotFound            = "not_found"
	TypeForbidden           = "forbidden"
	TypeBadRequest          = "bad_request"
	TypeInternalServerError = "server_error"
)

type RepositoryError struct {
	Causes  []error
	Message string
	Type    string
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("%s, Cause: %v", e.Message, e.Causes)
}

func NotFound(message string, causes ...error) *RepositoryError {
	return create(TypeNotFound, message, causes...)
}

func InternalServerError(message string, causes ...error) *RepositoryError {
	return create(TypeInternalServerError, message, causes...)
}

func Forbidden(message string, causes ...error) *RepositoryError {
	return create(TypeForbidden, message, causes...)
}

func BadRequest(message string, causes ...error) *RepositoryError {
	return create(TypeBadRequest, message, causes...)
}

func CreateFromStatusCode(message string, statusCode int, err error) *RepositoryError {
	switch statusCode {
	case http.StatusBadRequest:
		return BadRequest(message, err)
	case http.StatusNotFound:
		return NotFound(message, err)
	case http.StatusForbidden:
		return Forbidden(message, err)
	default:
		return InternalServerError(message, err)
	}
}

func create(Type, message string, errors ...error) *RepositoryError {
	return &RepositoryError{
		Type:    Type,
		Message: message,
		Causes:  errors,
	}
}
