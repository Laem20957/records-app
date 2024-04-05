package rest

import (
	"github.com/gin-gonic/gin"
	filesSwagger "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(filesSwagger.Handler))

	healthcheck := router.Group("/")
	{
		healthcheck.GET("healthcheck", h.healthcheck)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.refresh_token)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/", h.create)
		api.GET("/allrecords", h.getAll)
		api.GET("/:id", h.getById)
		api.PUT("/:id", h.update)
		api.DELETE("/:id", h.delete)
	}
	return router
}
