package domain

import (
	"records-app/internal/logger"

	"github.com/gin-gonic/gin"
)

var logs = logger.CreateLogs()

type GetAllRecordResponse struct {
	Data []Records `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ResultResponse struct {
	Status     string `json:"status"`
	ResponseId int    `json:"responseid,omitempty"`
}

func ServerResponse(ctx *gin.Context, statusCode int, message string) {
	logs.Log().Info(message)
	ctx.AbortWithStatusJSON(statusCode, MessageResponse{Message: message})
}
