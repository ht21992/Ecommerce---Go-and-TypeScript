package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID    uint    `json:"user_id"`                             // Reference to the user
	ProductID uint    `json:"product_id"`                          // Reference to the product
	Quantity  int     `json:"quantity"`                            // Quantity of the product in the cart
	Product   Product `json:"product" gorm:"foreignKey:ProductID"` // Embedded product details
}
