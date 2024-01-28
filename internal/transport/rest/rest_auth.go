package rest

import (
	"fmt"
	"net/http"
	"strings"

	domain "github.com/Laem20957/records-app/internal/domains"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description Create new account
// @ID Create-account
// @Accept json
// @Produce json
// @Param input body domain.user_attributes true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} domain.ServerResponse
// @Failure 500 {object} domain.ServerResponse
// @Failure default {object} domain.ServerResponse
// @Router /auth/sign-up [post]

func (h *Handler) signUp(ctx *gin.Context) {
	var input domain.Users
	if err := ctx.BindJSON(&input); err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.Services.IServiceAuthorizationMethods.CreateUser(ctx, input)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input domain.SignInInput
	if err := ctx.BindJSON(&input); err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.Services.IServiceAuthorizationMethods.SignIn(ctx, input)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Header("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}

func (h *Handler) refresh_token(ctx *gin.Context) {
	cookie, err := ctx.Cookie("refresh-token")
	if err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token := strings.ReplaceAll(cookie, "'", "")
	refreshToken, err := h.Services.RefreshToken(ctx, token)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Header("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": refreshToken,
	})
}
