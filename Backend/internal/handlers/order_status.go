package handlers

import (
	"ecommerce/internal/db"
	"ecommerce/internal/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var validStatuses = map[string]bool{
	"pending":   true,
	"paid":      true,
	"shipped":   true,
	"cancelled": true,
}

func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	status := strings.ToLower(input.Status)
	if !validStatuses[status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	var order models.Order
	if err := db.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Status = status
	db.DB.Save(&order)
	c.JSON(http.StatusOK, order)

}
