package rest

import (
	"fmt"
	"net/http"
	"strings"

	handler "records-app/api/rest/v1/handlers"
	"records-app/internal/adapters/database/schemas"

	"records-app/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary HealthCheck
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /healthcheck [get]
func HealthCheck(ctx *gin.Context) {
	models.ServerResponse(ctx, http.StatusOK, "OK")
}

// @Summary SignUp
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body domain.Users true "account info"
// @Success 200 {integer} integer 200
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /auth/sign-up [post]
func SignUp(ctx *gin.Context) {
	var users schemas.Users
	if err := ctx.BindJSON(&users); err != nil {
		models.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.Handler{}.Services.IServiceAuthorizationMethods.CreateUser(ctx, users)
	if err != nil {
		models.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignIn
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body domain.Users true "credentials"
// @Success 200 {integer} integer 200
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /auth/sign-in [post]
func SignIn(ctx *gin.Context) {
	var input schemas.Users

	if err := ctx.BindJSON(&input); err != nil {
		models.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := handler.Handler{}.Services.IServiceAuthorizationMethods.SignIn(ctx, input)
	if err != nil {
		models.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Header("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}

// @Summary RefreshToken
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {integer} integer 200
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /auth/refresh-token [get]
func RefreshToken(ctx *gin.Context) {
	cookie, err := ctx.Cookie("refresh-token")
	if err != nil {
		models.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token := strings.ReplaceAll(cookie, "'", "")
	refreshToken, err := handler.Handler{}.Services.IServiceAuthorizationMethods.RefreshToken(ctx, token)
	if err != nil {
		models.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Header("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": refreshToken,
	})
}
