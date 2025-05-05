package handlers

import (
	"ecommerce/internal/db"
	"ecommerce/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*

API Usage (example)

POST /api/cart

{ "product_id": 1, "quantity": 2 }

GET /api/cart → Get all items for the user.
PUT /api/cart/:id → Update item quantity.
DELETE /api/cart/:id → Remove from cart.


*/

func AddToCart(c *gin.Context) {
	var input struct {
		ProductID uint `json:"product_id"` // ID of the product to add to cart
		Quantity  int  `json:"quantity"`   // Quantity of the product
	}

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	userId := c.GetUint("userID")
	// fmt.Printf(userId)
	var item models.CartItem
	// check if item already exisits in cart
	err := db.DB.Where("user_id = ? And product_id = ?", userId, input.ProductID).First(&item).Error
	if err == nil {
		// if item exisits , update quantitiy
		item.Quantity += input.Quantity
		db.DB.Save(&item)
		c.JSON(http.StatusOK, item)
		return

	}

	// create a new cart item

	newItem := models.CartItem{
		UserID:    userId,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}

	db.DB.Create(&newItem)
	c.JSON(http.StatusCreated, newItem)

}

func ViewCart(c *gin.Context) {
	userId := c.GetUint("userID")
	var cart []models.CartItem
	db.DB.Preload("Product").Where("user_id = ?", userId).Find(&cart)
	c.JSON(http.StatusOK, cart)

}

func UpdateCartItem(c *gin.Context) {
	id := c.Param("id")
	var item models.CartItem
	if err := db.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Item not found"})
		return

	}

	userID := c.GetUint("userID")

	if item.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	item.Quantity = input.Quantity
	db.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

func RemoveCartItem(c *gin.Context) {
	id := c.Param("id")
	var item models.CartItem
	if err := db.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userID := c.GetUint("userID")
	if item.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	db.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item removed"})
}
