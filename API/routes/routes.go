package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/waltereidi/GolangBootstrap/controllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUserByID)
		api.POST("/users", controllers.CreateUser)
	}
}
