package handlers

import (
	"github.com/gin-gonic/gin"
)

type GetDocumentsHandler struct{}

type getDcoumentsUriParams struct {
	IndexID string `uri:"index_id" binding:"required"`
}

type getDocumentsBody struct {
	Data interface{} `json:"data"`
}

func NewGetDocumentsHandler() *GetDocumentsHandler {
	return &GetDocumentsHandler{}
}

func (h *GetDocumentsHandler) Handle(c *gin.Context) {
	// params,err :=
}

// func (h *GetDocumentsHandler) buildParams(c *gin.Context)(create_index ) {
// 	var uriParams getDcoumentsUriParams
// 	if err := c.ShouldBindUri(&uriParams); err != nil {
// 		apierr := api.NewInvalidURLParamsError(err)
// 		c.JSON(apierr.Status(), apierr)
// 		return nil, err
// 	}

// 	var body
// 	if err := c.BindJSON(&body); err != nil {
// 		apierr := api.NewInvalidBodyDataError(err)
// 		c.JSON(apierr.Status(), apierr)
// 		return
// 	}

// 	data, _ := json.Marshal(body.Data)
// 	return
// }
