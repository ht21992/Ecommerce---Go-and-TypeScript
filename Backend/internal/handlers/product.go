// internal/handlers/product.go
package handlers

import (
	"ecommerce/internal/db"
	"ecommerce/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}
	if err := db.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a product"})
		return
	}

	c.JSON(http.StatusOK, product)

}

func GetAllProducts(c *gin.Context) {

	var products []models.Product
	db.DB.Find(&products)
	c.JSON(http.StatusOK, products)

}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)

}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.ImageURL = input.ImageURL

	db.DB.Save(&product)
	c.JSON(http.StatusOK, product)

}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
