package http

import (
	"github.com/agrism/bookstore_oauth-api/src/domain/access_token"
	"github.com/agrism/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler accessTokenHandler) GetById(ctx *gin.Context) {
	accessTokenId := strings.TrimSpace(ctx.Param("access_token_id"))

	accessToken, err := handler.service.GetById(accessTokenId)

	if err != nil {
		ctx.JSON(err.Status, err)
	}

	ctx.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken

	if err := c.ShouldBindJSON(&at); err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, at)
}
