package usecase

import (
	"errors"
	"testing"

	"go-elasticsearch-v1/internal/platform/errors/repository/http"

	"github.com/stretchr/testify/assert"
)

func Test_Error_returnOK(t *testing.T) {
	//Setup
	err := CreateFromError("Could not get pack",
		&http.RepositoryError{Message: "PackRepository: could not get pack", Causes: []error{errors.New("Timeout 300ms")}})

	//Exercise
	result := err.Error()

	//Verify
	assert.Equal(t, "Could not get pack, Cause: [PackRepository: could not get pack, Cause: [Timeout 300ms]]", result)
}

func Test_Error_withCauseNil_returnOK(t *testing.T) {
	//Setup
	err := CreateFromError("Could not get pack", nil)

	//Exercise
	result := err.Error()

	//Verify
	assert.Equal(t, "Could not get pack, Cause: [<nil>]", result)
}

func Test_NotFound_returnOK(t *testing.T) {
	//Exercise
	result := NotFound("Message", nil)

	//Verify
	assert.Equal(t, &UseCaseError{
		Message: "Message",
		Type:    TypeNotFound,
		Causes:  []error{nil},
	}, result)
}

func Test_Unknown_returnOK(t *testing.T) {
	//Exercise
	result := Unknown("Message", nil)

	//Verify
	assert.Equal(t, &UseCaseError{
		Message: "Message",
		Type:    TypeUnknown,
		Causes:  []error{nil},
	}, result)
}

func Test_Forbidden_returnOK(t *testing.T) {
	//Exercise
	result := Forbidden("Message", nil)

	//Verify
	assert.Equal(t, &UseCaseError{
		Message: "Message",
		Type:    TypeForbidden,
		Causes:  []error{nil},
	}, result)
}

func Test_InvalidParam_returnOK(t *testing.T) {
	//Exercise
	result := InvalidParam("Message", nil)

	//Verify
	assert.Equal(t, &UseCaseError{
		Message: "Message",
		Type:    TypeInvalidParam,
		Causes:  []error{nil},
	}, result)
}

func Test_CreateFromError_withBadRequest_returnOK(t *testing.T) {
	//Exercise
	result := CreateFromError("Message", &http.RepositoryError{Type: http.TypeBadRequest})

	//Verify
	assert.Equal(t, &UseCaseError{
		Message: "Message",
		Type:    TypeInvalidParam,
		Causes:  []error{&http.RepositoryError{Type: http.TypeBadRequest}},
	}, result)
}

func Test_CreateFromError_withForbidden_returnOK(t *testing.T) {
	//Exercise
	result := CreateFromError("Message", &http.RepositoryError{Type: http.TypeForbidden})

	//Verify
	assert.Equal(t, &UseCaseError{
		Message: "Message",
		Type:    TypeForbidden,
		Causes:  []error{&http.RepositoryError{Type: http.TypeForbidden}},
	}, result)
}

func Test_CreateFromError_withNotFound_returnOK(t *testing.T) {
	//Exercise
	result := CreateFromError("Message", &http.RepositoryError{Type: http.TypeNotFound})

	//Verify
	assert.Equal(t, &UseCaseError{
		Message: "Message",
		Type:    TypeNotFound,
		Causes:  []error{&http.RepositoryError{Type: http.TypeNotFound}},
	}, result)
}

func Test_CreateFromError_withInternalServerError_returnOK(t *testing.T) {
	//Exercise
	result := CreateFromError("Message", &http.RepositoryError{Type: http.TypeInternalServerError})

	//Verify
	assert.Equal(t, &UseCaseError{
		Message: "Message",
		Type:    TypeUnknown,
		Causes:  []error{&http.RepositoryError{Type: http.TypeInternalServerError}},
	}, result)
}
