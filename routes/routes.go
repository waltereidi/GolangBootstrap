package routes

import (
	"github.com/waltereidi/GolangBootstrap/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUserByID)
		api.POST("/users", controllers.CreateUser)
	}
}
