package main

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"gin-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	/* r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/get-hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello-world endpoint",
			"nums":    []int{1, 2, 3, 4, 5},
		})
	}) */

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
