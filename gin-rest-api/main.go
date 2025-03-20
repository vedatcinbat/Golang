package main

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"gin-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Product{},
	)

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
