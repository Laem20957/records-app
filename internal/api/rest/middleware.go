package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Laem20957/records-app/internal/domain"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "UserContext"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	headerParts := strings.Split(header, " ")
	userId, err := h.Services.IServiceAuthorizationMethods.TokenIsSigned(headerParts[1])

	if header == "" {
		domain.ServerResponse(ctx, http.StatusUnauthorized, "empty auth header")
	} else if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		domain.ServerResponse(ctx, http.StatusUnauthorized, "invalid auth header")
	} else if len(headerParts[1]) == 0 {
		domain.ServerResponse(ctx, http.StatusUnauthorized, "token is empty")
	} else if err != nil {
		domain.ServerResponse(ctx, http.StatusUnauthorized, "invalid token")
	} else {
		ctx.Set(userContext, userId)
	}
}

func getAllContext(ctx *gin.Context) (*gin.Context, error) {
	data, ok := ctx.Get(userContext)
	if !ok {
		return nil, errors.New("context not found")
	} else {
		return data.(*gin.Context), nil
	}
}

func getUserContext(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userContext)
	if !ok {
		return 0, errors.New("context not found")
	} else {
		return id.(int), nil
	}
}
