package handlers

import (
	_ "records-app/docs"
	"records-app/internal/service"
)

type Handler struct {
	Services *service.ServiceMethods
}

func GetHandler(service *service.ServiceMethods) *Handler {
	return &Handler{service}
}
