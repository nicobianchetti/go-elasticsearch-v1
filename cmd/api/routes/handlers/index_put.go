package handlers

import (
	"context"
	"go-elasticsearch-v1/internal/platform/web/api"
	"go-elasticsearch-v1/internal/usecase/create_index"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexWritter interface {
	CreateIndex(ctx context.Context, params create_index.Params) (*create_index.IndexCreateResult, error)
}

type PutIndexHandler struct {
	Service IndexWritter
}

func NewPutIndexHandler(service IndexWritter) *PutIndexHandler {
	return &PutIndexHandler{Service: service}
}

type putIndexUriParams struct {
	IndexID string `uri:"index_id" binding:"required"`
}

type putIndexBody struct {
	Data interface{} `json:"data"`
}

type putIndexResponse struct {
	Acknowledged       bool   `json:"acknowledged"`
	ShardsAcknowledged bool   `json:"shards_acknowledged"`
	Index              string `json:"index"`
}

func (h *PutIndexHandler) Handle(c *gin.Context) {
	params, err := h.buildParams(c)
	if err != nil {
		return
	}

	res, err := h.Service.CreateIndex(c.Request.Context(), params)
	if err != nil {
		returnError(err, c)
		return
	}

	c.JSON(http.StatusCreated, &putIndexResponse{
		Acknowledged:       res.Acknowledged,
		ShardsAcknowledged: res.ShardsAcknowledged,
		Index:              res.Index,
	})
}

func (h *PutIndexHandler) buildParams(c *gin.Context) (create_index.Params, error) {
	var uriParams putIndexUriParams
	if err := c.ShouldBindUri(&uriParams); err != nil {
		apierr := api.NewInvalidURLParamsError(err)
		c.JSON(apierr.Status(), apierr)
		return create_index.Params{}, err
	}

	var body putIndexBody
	if err := c.BindJSON(&body); err != nil {
		apierr := api.NewInvalidBodyDataError(err)
		c.JSON(apierr.Status(), apierr)
		return create_index.Params{}, err
	}

	return create_index.Params{
		IndexID: uriParams.IndexID,
		Body:    body,
	}, nil
}
