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
		auth.GET("/refresh-token", RefreshToken)
	}

	api := router.Group("/api")
	{
		api.GET("/all/record", GetAll)
		api.GET("/record/id", GetById)
		api.POST("/new/record", Create)
		api.PUT("/:id", Update)
		api.DELETE("/:id", Delete)
	}
	return router
}
