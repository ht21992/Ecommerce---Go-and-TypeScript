// internal/models/product.go

package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price" gorm:"type:decimal(10,2)"`
	Stock       int             `json:"stock"`
	ImageURL    string          `json:"image_url"`
}
