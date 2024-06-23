package handlers

import (
	_ "records-app/docs"
	"records-app/internal/service"
)

type Handler struct {
	Services *service.ServiceMethods
}

func NewHandler(service *service.ServiceMethods) *Handler {
	return &Handler{Services: service}
}
