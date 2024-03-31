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
	Data []Record `json:"data"`
}

func ServerResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, MessageResponse{Message: message})
}
