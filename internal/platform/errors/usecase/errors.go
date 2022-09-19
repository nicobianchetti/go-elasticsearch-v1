package usecase

import (
	"fmt"

	"go-elasticsearch-v1/internal/platform/errors/repository/http"
)

const (
	TypeNotFound     = "not_found"
	TypeUnknown      = "unknown"
	TypeForbidden    = "forbidden"
	TypeInvalidParam = "invalid_param"
	TypeRedirect     = "redirect"
)

type UseCaseError struct {
	Causes      []error
	Message     string
	Type        string
	URLRedirect *string
}

func (e *UseCaseError) Error() string {
	if e.Causes == nil {
		return e.Message
	}
	return fmt.Sprintf("%s, Cause: %v", e.Message, e.Causes)
}

func NotFound(message string, causes ...error) *UseCaseError {
	return Create(TypeNotFound, message, causes...)
}

func Unknown(message string, causes ...error) *UseCaseError {
	return Create(TypeUnknown, message, causes...)
}

func Forbidden(message string, causes ...error) *UseCaseError {
	return Create(TypeForbidden, message, causes...)
}

func InvalidParam(message string, causes ...error) *UseCaseError {
	return Create(TypeInvalidParam, message, causes...)
}

func Redirect(message string, urlRedirect string, causes ...error) *UseCaseError {
	err := Create(TypeRedirect, message, causes...)
	err.URLRedirect = &urlRedirect
	return err
}

func Create(Type, message string, errors ...error) *UseCaseError {
	return &UseCaseError{
		Type:    Type,
		Message: message,
		Causes:  errors,
	}
}

func CreateFromError(message string, err error) *UseCaseError {
	data, ok := err.(*http.RepositoryError)
	if !ok {
		return Unknown(message, err)
	}

	switch data.Type {
	case http.TypeBadRequest:
		return InvalidParam(message, err)
	case http.TypeNotFound:
		return NotFound(message, err)
	case http.TypeForbidden:
		return Forbidden(message, err)
	case http.TypeInternalServerError:
		return Unknown(message, err)
	default:
		return Unknown(message, err)
	}
}
