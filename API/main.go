package main

import (
	"github.com/waltereidi/GolangBootstrap/config"
	"github.com/waltereidi/GolangBootstrap/models"
	"github.com/waltereidi/GolangBootstrap/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	// Example: create a user
	user := models.User{
		Name: "John Doe",
	}

	config.DB.Create(&user)
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
