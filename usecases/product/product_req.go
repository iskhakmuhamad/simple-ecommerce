package product

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
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

type AddProductRequest struct {
	ProductName     string    `json:"product_name" form:"product_name"`
	ProductPrice    float64   `json:"product_price" form:"product_price"`
	ProductQty      int64     `json:"product_qty" form:"product_qty"`
	ProductCategory string    `json:"product_category" form:"product_category"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (c *AddProductRequest) Validate() error {
	if err := validation.Validate(c.ProductName, validation.Required); err != nil {
		return errors.New("product name must be filled")
	}
	if err := validation.Validate(c.ProductPrice, validation.Required); err != nil {
		return errors.New("product price must be filled")
	}
	if err := validation.Validate(c.ProductQty, validation.Required); err != nil {
		return errors.New("product qty must be filled")
	}
	if err := validation.Validate(c.ProductCategory, validation.Required); err != nil {
		return errors.New("product category must be filled")
	}

	return nil
}
