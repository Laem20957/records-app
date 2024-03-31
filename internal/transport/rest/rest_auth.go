package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Laem20957/records-app/internal/domain"
	"github.com/gin-gonic/gin"
)

// auth.go

// @Summary SignUp
// @Tags auth
// @Description Create new account
// @ID Create-account
// @Accept json
// @Produce json
// @Param input body domain.Users true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} domain.MessageResponse
// @Failure 500 {object} domain.MessageResponse
// @Failure default {object} domain.MessageResponse
// @Router /auth/sign-up [post]

func (h *Handler) signUp(ctx *gin.Context) {
	var users domain.Users
	if err := ctx.BindJSON(&users); err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.Services.IServiceAuthorizationMethods.CreateUser(ctx, users)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept json
// @Produce json
// @Param input body domain.SignInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} domain.MessageResponse
// @Failure 500 {object} domain.MessageResponse
// @Failure default {object} domain.MessageResponse
// @Router /auth/sign-in [post]

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

// @Summary Refresh_token
// @Tags auth
// @Description refresh token
// @ID refresh-token
// @Accept json
// @Produce json
// @Success 200 {string} string "token"
// @Failure 400,404 {object} domain.MessageResponse
// @Failure 500 {object} domain.MessageResponse
// @Failure default {object} domain.MessageResponse
// @Router /auth/refresh [get]

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
