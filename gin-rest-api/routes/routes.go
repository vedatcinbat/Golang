package routes

import (
	"gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// User routes
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserById)
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Post routes
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostsWithId)
	r.POST("/posts", controllers.CreatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// Products routes
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.POST("/products", controllers.CreateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
}
