// internal/handlers/checkout.go
package handlers

import (
	"ecommerce/internal/db"
	"ecommerce/internal/kafka"
	"ecommerce/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)


/*


POST /api/checkout


*/

func Checkout(c *gin.Context) {
	userID := c.GetUint("userID")

	// Get cart items
	var cart []models.CartItem
	if err := db.DB.Preload("Product").Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cart fetch failed"})
		return
	}
	if len(cart) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	// Create orders
	var orders []models.Order
	for _, item := range cart {
		// Check stock
		if item.Product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product: " + item.Product.Name})
			return
		}

		// Reduce stock
		item.Product.Stock -= item.Quantity
		db.DB.Save(&item.Product)

		// Create order
		order := models.Order{
			UserID:    userID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Total:     decimal.NewFromInt(int64(item.Quantity)).Mul(item.Product.Price),
			Status:    "pending",
		}

		orders = append(orders, order)
		db.DB.Create(&order)

		// Convert order ID to string
		orderIDStr := strconv.FormatUint(uint64(order.ID), 10)

		// Send order event to Kafka
		kafka.ProduceOrderEvent(orderIDStr)

	}

	// Clear cart
	db.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})

	c.JSON(http.StatusOK, orders)
}
