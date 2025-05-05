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
		// User
		protected.GET("/me", func(c *gin.Context) {
			userID := c.MustGet("user")
			c.JSON(200, gin.H{"user_id": userID})
		})

		// Products
		protected.POST("/products", CreateProduct)
		protected.GET("/products", GetAllProducts)
		protected.GET("/products/:id", GetProductByID)
		protected.PUT("/products/:id", UpdateProduct)
		protected.DELETE("/products/:id", DeleteProduct)

		// Orders
		protected.POST("/orders", CreateOrder)

		// Cart
		protected.POST("/cart", AddToCart)
		protected.GET("/cart", ViewCart)
		protected.PUT("/cart/:id", UpdateCartItem)
		protected.DELETE("/cart/:id", RemoveCartItem)

		// Checkout
		protected.POST("/checkout", Checkout)

	}
}
