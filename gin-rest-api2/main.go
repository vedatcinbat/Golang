package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})

	})

	// Static route
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Path parameter: route with dynamic user id
	router.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	// Query string: route to search with a query parameter
	router.GET("search", func(c *gin.Context) {
		query := c.Query("q")
		c.JSON(http.StatusOK, gin.H{
			"search_query": query,
		})
	})

	// Route grouing: grouping API endpoints under '/api' path
	api := router.Group("/api")
	{
		api.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		api.POST("/create", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{"status": "created"})
		})
	}

	router.Run(":8080")
}
