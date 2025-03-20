package controllers

import (
	"net/http"
	"strconv"

	"gin-rest-api/database"
	"gin-rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

func GetUserById(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var users []models.User
	database.DB.Find(&users)

	for _, user := range users {
		if user.ID == uint(id) {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newUser)
	c.JSON(201, newUser)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	database.DB.Delete(&models.User{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
