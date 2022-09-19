package handlers

import (
	"go-elasticsearch-v1/internal/platform/errors/usecase"
	"go-elasticsearch-v1/internal/platform/web/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func returnError(err error, c *gin.Context) {
	status, redirectURL := getStatusFromError(err)
	if redirectURL != nil {
		c.Header("redirect_url", *redirectURL)
		c.Status(status)
		return
	}

	c.JSON(status, api.NewError(err.Error(), status, err))
}

func getStatusFromError(err error) (int, *string) {
	data, ok := err.(*usecase.UseCaseError)

	if !ok {
		return http.StatusInternalServerError, nil
	}

	switch data.Type {
	case usecase.TypeInvalidParam:
		return http.StatusBadRequest, nil
	case usecase.TypeNotFound:
		return http.StatusNotFound, nil
	case usecase.TypeForbidden:
		return http.StatusForbidden, nil
	case usecase.TypeUnknown:
		return http.StatusInternalServerError, nil
	case usecase.TypeRedirect:
		return http.StatusMovedPermanently, data.URLRedirect
	default:
		return http.StatusInternalServerError, nil
	}
}
