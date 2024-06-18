package handlers

import (
	"errors"
	"net/http"
	"strings"

	"records-app/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	allContext          = "ContextAll"
	userContext         = "ContextUser"
)

func UserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	headerParts := strings.Split(header, " ")
	userId, err := Handler{}.Services.IServiceAuthorizationMethods.TokenIsSigned(headerParts[1])

	if header == "" {
		models.ServerResponse(ctx, http.StatusUnauthorized, "empty auth header")
	} else if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		models.ServerResponse(ctx, http.StatusUnauthorized, "invalid auth header")
	} else if len(headerParts[1]) == 0 {
		models.ServerResponse(ctx, http.StatusUnauthorized, "token is empty")
	} else if err != nil {
		models.ServerResponse(ctx, http.StatusUnauthorized, "invalid token")
	} else {
		ctx.Set(userContext, userId)
	}
}

func GetAllContext(ctx *gin.Context) (*gin.Context, error) {
	data, ok := ctx.Get(allContext)
	if !ok {
		return nil, errors.New("context not found")
	} else if data == nil {
		return nil, errors.New("runtime error:" +
			"invalid memory address or nil pointer dereference")
	} else {
		return data.(*gin.Context), nil
	}
}

func GetUserContext(ctx *gin.Context) (*gin.Context, error) {
	id, ok := ctx.Get(userContext)
	if !ok {
		return nil, errors.New("context not found")
	} else {
		return id.(*gin.Context), nil
	}
}
