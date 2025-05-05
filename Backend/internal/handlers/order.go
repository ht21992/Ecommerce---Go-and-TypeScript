package handlers

import (
	"ecommerce/internal/db"
	"ecommerce/internal/kafka"
	"ecommerce/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order data"})
		return
	}

	// Save order to DB
	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Convert order ID to string
	orderIDStr := strconv.FormatUint(uint64(order.ID), 10)

	// Send order event to Kafka
	kafka.ProduceOrderEvent(orderIDStr)

	// Respond back to client
	c.JSON(http.StatusOK, order)
}
