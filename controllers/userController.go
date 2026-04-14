package controllers

import (
	"net/http"

	"golangbootstrap/models"

	"github.com/gin-gonic/gin"
)

// fake in-memory data
var users = []models.User{
	{ID: 1, Name: "Walter"},
	{ID: 2, Name: "John"},
}

// GET /users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if id == string(rune(user.ID)) {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// POST /users
func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	c.JSON(http.StatusCreated, newUser)
}
