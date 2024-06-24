package handlers

import (
	"records-app/internal/service"
	_ "records-app/docs"
)

type Handler struct {
	Services *service.ServiceMethods
}

func NewHandler(service *service.ServiceMethods) *Handler {
	return &Handler{Services: service}
}
