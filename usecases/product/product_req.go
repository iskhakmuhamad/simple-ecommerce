package product

import (
	"time"
)

type ProductsRequest struct {
	Name      string    `json:"name" form:"name"`
	Price     float64   `json:"price" form:"price"`
	Qty       int64     `json:"qty" form:"qty"`
	Category  string    `json:"category" form:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *ProductsRequest) Validate() error {

	return nil
}
