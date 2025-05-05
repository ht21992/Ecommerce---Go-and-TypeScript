package handlers

import (
	"ecommerce/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/signup", Signup)
		api.POST("/login", Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		protected.GET("/me", func(c *gin.Context) {
			userID := c.MustGet("user")
			c.JSON(200, gin.H{"user_id": userID})
		})
		protected.POST("/products", CreateProduct)
		protected.GET("/products", GetAllProducts)
		protected.GET("/products/:id", GetProductByID)
		protected.PUT("/products/:id", UpdateProduct)
		protected.DELETE("/products/:id", DeleteProduct)
		protected.POST("/orders", CreateOrder)
	}
}
