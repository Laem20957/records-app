package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Laem20957/records-app/internal/service"
	filesSwagger "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Laem20957/records-app/docs"
)

type Handler struct {
	Services *service.ServiceMethods
}

func GetHandler(service *service.ServiceMethods) *Handler {
	return &Handler{service}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
		auth.GET("/refresh", handler.refresh_token)
	}

	api := router.Group("/api", handler.userIdentity)
	{
		record := api.Group("/record")
		{
			record.POST("/", handler.create)
			record.GET("/", handler.getAll)
			record.GET("/:id", handler.getById)
			record.PUT("/:id", handler.update)
			record.DELETE("/:id", handler.delete)
		}
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(filesSwagger.Handler))
	return router
}
