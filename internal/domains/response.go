package domains

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

func ServerResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, MessageResponse{Message: message})
}
