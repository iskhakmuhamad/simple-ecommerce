package models

import "time"

type Product struct {
	ID              int64     `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	ProductName     string    `gorm:"type:varchar(255);column:product_name" json:"product_name"`
	ProductPrice    float64   `gorm:"type:float;column:product_price" json:"product_price"`
	ProductQty      int64     `gorm:"type:integer;column:product_qty" json:"product_qty"`
	ProductCategory string    `gorm:"type:varchar(255);column:product_category" json:"product_category"`
	CreatedAt       time.Time `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
}
