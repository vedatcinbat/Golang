package routes

import (
	"gin-rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserById)
	r.POST("/users", handlers.CreateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
}
