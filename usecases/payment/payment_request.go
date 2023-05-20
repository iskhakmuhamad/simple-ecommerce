package payment

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AddPaymentRequest struct {
	UserID     int64
	ProductID  int64 `json:"product_id" form:"product_id"`
	Amount     int64 `json:"amount" form:"amount"`
	TotalPrice float64
	CreatedAt  *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" form:"updated_at"`
}

func (c *AddPaymentRequest) Validate() error {
	if err := validation.Validate(c.ProductID, validation.Required); err != nil {
		return errors.New("product id must be filled")
	}

	if err := validation.Validate(c.Amount, validation.Required); err != nil {
		return errors.New("amount must be filled")
	}

	return nil
}
