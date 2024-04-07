package handlers

import (
	_ "github.com/Laem20957/records-app/docs"
	"github.com/Laem20957/records-app/internal/service"
)

type Handler struct {
	Services *service.ServiceMethods
}

func GetHandler(service *service.ServiceMethods) *Handler {
	return &Handler{service}
}
