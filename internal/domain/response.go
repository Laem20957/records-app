package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type GetAllRecordResponse struct {
	Data []Records `json:"data"`
}

func ServerResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Info(message)
	ctx.AbortWithStatusJSON(statusCode, MessageResponse{Message: message})
}
