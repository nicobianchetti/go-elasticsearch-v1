package http

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_returnOK(t *testing.T) {
	//Setup
	err := create(TypeInternalServerError, "Could not get pack", errors.New("Timeout 300ms"))

	//Exercise
	result := err.Error()

	//Verify
	assert.Equal(t, "Could not get pack, Cause: [Timeout 300ms]", result)
}

func Test_NotFound_returnOK(t *testing.T) {
	//Exercise
	result := NotFound("Message", nil)

	//Verify
	assert.Equal(t, &RepositoryError{
		Message: "Message",
		Type:    TypeNotFound,
		Causes:  []error{nil},
	}, result)
}

func Test_Unknown_returnOK(t *testing.T) {
	//Exercise
	result := InternalServerError("Message", nil)

	//Verify
	assert.Equal(t, &RepositoryError{
		Message: "Message",
		Type:    TypeInternalServerError,
		Causes:  []error{nil},
	}, result)
}

func Test_Forbidden_returnOK(t *testing.T) {
	//Exercise
	result := Forbidden("Message", nil)

	//Verify
	assert.Equal(t, &RepositoryError{
		Message: "Message",
		Type:    TypeForbidden,
		Causes:  []error{nil},
	}, result)
}

func Test_BadRequest_returnOK(t *testing.T) {
	//Exercise
	result := BadRequest("Message", nil)

	//Verify
	assert.Equal(t, &RepositoryError{
		Message: "Message",
		Type:    TypeBadRequest,
		Causes:  []error{nil},
	}, result)
}

func Test_CreateFromStatusCode_withStatusBadRequest_returnOK(t *testing.T) {
	//Setup
	err := errors.New("Error")
	status := http.StatusBadRequest
	message := "Message"

	//Exercise
	result := CreateFromStatusCode(message, status, err)

	//Verify
	assert.Equal(t, &RepositoryError{
		Message: "Message",
		Type:    TypeBadRequest,
		Causes:  []error{err},
	}, result)
}

func Test_CreateFromStatusCode_withStatusNotFound_returnOK(t *testing.T) {
	//Setup
	err := errors.New("Error")
	status := http.StatusNotFound
	message := "Message"

	//Exercise
	result := CreateFromStatusCode(message, status, err)

	//Verify
	assert.Equal(t, &RepositoryError{
		Message: "Message",
		Type:    TypeNotFound,
		Causes:  []error{err},
	}, result)
}

func Test_CreateFromStatusCode_withStatusForbidden_returnOK(t *testing.T) {
	//Setup
	err := errors.New("Error")
	status := http.StatusForbidden
	message := "Message"

	//Exercise
	result := CreateFromStatusCode(message, status, err)

	//Verify
	assert.Equal(t, &RepositoryError{
		Message: "Message",
		Type:    TypeForbidden,
		Causes:  []error{err},
	}, result)
}

func Test_CreateFromStatusCode_withUnexpectedStatus_returnOK(t *testing.T) {
	//Setup
	err := errors.New("Error")
	status := http.StatusInternalServerError
	message := "Message"

	//Exercise
	result := CreateFromStatusCode(message, status, err)

	//Verify
	assert.Equal(t, &RepositoryError{
		Message: "Message",
		Type:    TypeInternalServerError,
		Causes:  []error{err},
	}, result)
}
