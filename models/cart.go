package models

import "time"

type Cart struct {
	UserID    int64      `gorm:"type:integer;primaryKey;column:user_id" json:"user_id"`
	ProductID int64      `gorm:"type:integer;primaryKey;column:product_id" json:"product_id"`
	Amount    int64      `gorm:"type:integer;column:amount" json:"amount"`
	CreatedAt *time.Time `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
}

type CartProducts struct {
	TotalPrice float64    `json:"total_price"`
	Amount     int64      `json:"amount"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Product    `json:"product_detail"`
}
