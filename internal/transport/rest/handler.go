package rest

import (
	"net/http"

	service "github.com/Laem20957/records-app/internal/services"
	"github.com/gin-gonic/gin"
	files_swagger "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"

	_ "github.com/Laem20957/records-app/docs"
)

type Handler struct {
	Services *service.Service
}

// func (h *Handler) GetHandler() *Handler {
// 	service_struct := &service.Service{}
// 	handler_struct := &Handler{Services: service_struct}
// 	return handler_struct
// }

func GetHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.refresh)
	}

	api := router.Group("/api", h.userIdentity)
	{
		note := api.Group("/note")
		{
			note.POST("/", h.create)
			note.GET("/", h.getAll)
			note.GET("/:id", h.getById)
			note.PUT("/:id", h.update)
			note.DELETE("/:id", h.delete)
		}
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/swagger/*any", gin_swagger.WrapHandler(files_swagger.Handler))
	return router
}
