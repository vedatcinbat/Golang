package handlers

import (
	"net/http"
	"strconv"

	"gin-rest-api/models"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@gmail.com"},
	{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "janedoe@gmail.com"},
}

// GET /users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func GetUserById(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// POST /users
func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
