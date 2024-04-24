package rest

import (
	"github.com/gin-gonic/gin"
	filesSwagger "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(filesSwagger.Handler))

	healthcheck := router.Group("/")
	{
		healthcheck.GET("healthcheck", HealthCheck)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
		auth.GET("/refresh", RefreshToken)
	}

	api := router.Group("/api")
	{
		api.POST("/", Create)
		api.GET("/allrecords", GetAll)
		api.GET("/:id", GetById)
		api.PUT("/:id", Update)
		api.DELETE("/:id", Delete)
	}
	return router
}
