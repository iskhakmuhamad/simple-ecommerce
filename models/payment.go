package models

import "time"

type Payment struct {
	ID         int64      `gorm:"type:integer;primaryKey;column:id" json:"id"`
	UserID     int64      `gorm:"type:integer;column:user_id" json:"user_id"`
	ProductID  int64      `gorm:"type:integer;column:product_id" json:"product_id"`
	Amount     int64      `gorm:"type:integer;column:amount" json:"amount"`
	TotalPrice float64    `gorm:"type:float;column:total_price" json:"total_price"`
	CreatedAt  *time.Time `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
}
