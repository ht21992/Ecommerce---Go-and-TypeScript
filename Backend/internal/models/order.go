// internal/models/order.go
package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint            `json:"user_id"`
	ProductID uint            `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Total     decimal.Decimal `json:"total" gorm:"type:decimal(10,2)"`
}
