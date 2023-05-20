package cart

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AddCartRequest struct {
	UserID    int64
	ProductID int64 `json:"product_id" form:"product_id"`
	Amount    int64 `json:"amount" form:"amount"`
}

func (c *AddCartRequest) Validate() error {

	if err := validation.Validate(c.ProductID, validation.Required); err != nil {
		return errors.New("product_id must be filled")
	}

	return nil
}
