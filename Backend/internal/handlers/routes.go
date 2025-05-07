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
		admin := protected.Group("/admin", middleware.AdminOnly())

		// User
		protected.GET("/me", func(c *gin.Context) {
			userID := c.MustGet("userID")
			c.JSON(200, gin.H{"userID": userID})
		})

		// Products
		protected.POST("/products", CreateProduct)
		protected.GET("/products", GetAllProducts)
		protected.GET("/products/:id", GetProductByID)
		protected.PUT("/products/:id", UpdateProduct)
		protected.DELETE("/products/:id", DeleteProduct)

		// Orders
		protected.POST("/orders", CreateOrder)
		protected.GET("/orders", GetMyOrders)

		// Cart
		protected.POST("/cart", AddToCart)
		protected.GET("/cart", ViewCart)
		protected.PUT("/cart/:id", UpdateCartItem)
		protected.DELETE("/cart/:id", RemoveCartItem)

		// Checkout
		protected.POST("/checkout", Checkout)

		// Admins
		admin.PUT("/orders/:id/status", UpdateOrderStatus)

	}
}
