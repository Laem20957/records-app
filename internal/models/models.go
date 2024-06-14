package models

import (
	"records-app/internal/adapters/database/schemas"
	"records-app/internal/logger"

	"github.com/gin-gonic/gin"
)

var logs = logger.CreateLogs()

type DataResponse struct {
	ID   int               `json:"id,omitempty"`
	Data []schemas.Records `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func ServerResponse(ctx *gin.Context, statusCode int, message string) {
	logs.Log().Info(message)
	ctx.AbortWithStatusJSON(statusCode, MessageResponse{Message: message})
}
