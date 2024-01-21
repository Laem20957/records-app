package rest

import (
	"fmt"
	"net/http"
	"strings"

	domain "github.com/Laem20957/records-app/internal/domains"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags Auth
// @Description Create new account
// @ID Create-account
// @Accept json
// @Produce son
// @Param Input body domain.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /auth/sign-up [post]

func (h *Handler) signUp(c *gin.Context) {
	var input domain.Users
	if err := c.BindJSON(&input); err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.Services.ServiceAuthorizationMethods.CreateUsers(c, input)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input domain.SignInInput
	if err := c.BindJSON(&input); err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.Services.ServiceAuthorizationMethods.SignIn(c, input)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}

func (h *Handler) refresh_token(c *gin.Context) {
	cookie, err := c.Cookie("refresh-token")
	if err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token := strings.ReplaceAll(cookie, "'", "")
	accessToken, refreshToken, err := h.Services.RefreshTokens(c, token)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}
