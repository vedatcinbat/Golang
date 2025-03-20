package controllers

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	database.DB.First(&product, id)
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Product{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
