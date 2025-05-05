// internal/models/order.go
package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// type Order struct {
// 	gorm.Model
// 	UserID    uint            `json:"user_id"`
// 	ProductID uint            `json:"product_id"`
// 	Quantity  int             `json:"quantity"`
// 	Total     decimal.Decimal `json:"total" gorm:"type:decimal(10,2)"`
// }



type Order struct {
    gorm.Model
    UserID    uint    `json:"user_id"`
    User      User    `json:"user" gorm:"foreignKey:UserID"`
    ProductID uint    `json:"product_id"`
    Product   Product `json:"product" gorm:"foreignKey:ProductID"`
    Quantity  int     `json:"quantity"`
    Total     decimal.Decimal `json:"total" gorm:"type:decimal(10,2)"`
    Status    string  `json:"status"` // e.g., "pending", "paid", "shipped"
}
