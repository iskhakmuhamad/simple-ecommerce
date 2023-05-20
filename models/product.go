package models

import "time"

type Product struct {
	ID        int64     `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	Name      string    `gorm:"type:varchar(255);column:name" json:"name"`
	Price     float64   `gorm:"type:float;column:price" json:"price"`
	Qty       int64     `gorm:"type:integer;column:qty" json:"qty"`
	Category  string    `gorm:"type:varchar(255);column:category" json:"category"`
	CreatedAt time.Time `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
}
