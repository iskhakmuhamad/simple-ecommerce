package models

import "time"

type User struct {
	ID        int64     `gorm:"autoIncrement;primaryKey;column:id" json:"id"`
	Name      string    `gorm:"type:varchar(255);column:name" json:"name"`
	Email     string    `gorm:"type:varchar(255);column:email" json:"email"`
	Password  string    `gorm:"type:varchar(255);column:password" json:"password"`
	Role      string    `gorm:"type:varchar(255);column:role" json:"role"`
	CreatedAt time.Time `gorm:"type:timestamp;column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at"`
}
