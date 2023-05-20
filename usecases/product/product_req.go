package product

import (
	"time"
)

type ProductsRequest struct {
	SearchName      string    `json:"search_name" form:"search_name"`
	ProductPrice    float64   `json:"product_price" form:"product_price"`
	ProductQty      int64     `json:"product_qty" form:"product_qty"`
	ProductCategory string    `json:"product_category" form:"product_category"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (c *ProductsRequest) Validate() error {

	return nil
}
