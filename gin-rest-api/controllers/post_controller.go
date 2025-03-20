package controllers

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func GetPostsWithId(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	database.DB.First(&post, id)
	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newPost)
	c.JSON(http.StatusCreated, newPost)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Post{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
