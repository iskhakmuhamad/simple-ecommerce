package usecases

import (
	"context"
	"errors"

	"github.com/iskhakmuhamad/ecommerce/models"
	"github.com/iskhakmuhamad/ecommerce/repositories"
	"github.com/iskhakmuhamad/ecommerce/usecases/cart"
)

type cartUC struct {
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
}

type Cart interface {
	CreateCart(ctx context.Context, params cart.AddCartRequest) error
}

func NewCartUC(cartRepo repositories.CartRepository, productRepo repositories.ProductRepository) Cart {
	return &cartUC{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (u *cartUC) CreateCart(ctx context.Context, params cart.AddCartRequest) error {

	if err := params.Validate(); err != nil {
		return err
	}

	req := &models.Cart{
		UserID:    params.UserID,
		ProductID: params.ProductID,
		Amount:    params.Amount,
	}
	product, err := u.productRepo.GetProductByID(ctx, req.UserID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("Product doesnt found, please check again or reload the data")
	}

	err = u.cartRepo.UpSertCart(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
