package rest

import (
	"errors"
	"net/http"
	"strings"

	domain "github.com/Laem20957/records-app/internal/domains"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "userId"
)

func (hs *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		domain.NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		domain.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		domain.NewErrorResponse(ctx, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := hs.Services.ServiceAuthorizationMethods.ParseToken(headerParts[1])
	if err != nil {
		domain.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid token")
		return
	}

	ctx.Set(userContext, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userContext)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
