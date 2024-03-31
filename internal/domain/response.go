package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Detail struct {
	Detail string `json:"detail"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type GetAllRecordResponse struct {
	Data []Record `json:"data"`
}

func WriteDetailsResponse(ctx *gin.Context, code int, text string) {
	ctx.JSON(code, Detail{Detail: text})
}

func ServerResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, MessageResponse{Message: message})
}
