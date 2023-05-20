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
	GetUserCartProducts(ctx context.Context, userID int64) ([]models.CartProducts, error)
	DeleteCartProduct(ctx context.Context, params cart.DeleteCartRequest) error
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
	product, err := u.productRepo.GetProductByID(ctx, req.ProductID)
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

func (u *cartUC) GetUserCartProducts(ctx context.Context, userID int64) ([]models.CartProducts, error) {
	cartProducts, err := u.cartRepo.GetCartProductsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	for index, products := range cartProducts {
		cartProducts[index].TotalPrice = products.ProductPrice * float64(products.Amount)
	}
	return cartProducts, nil
}

func (u *cartUC) DeleteCartProduct(ctx context.Context, params cart.DeleteCartRequest) error {
	if err := params.Validate(); err != nil {
		return err
	}
	product, err := u.cartRepo.GetCartbyUserIDnProductID(ctx, params.UserID, params.ProductID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("Product doesnt found in your cart, please check again or reload the data")
	}
	err = u.cartRepo.DeleteCartProduct(ctx, params.UserID, params.ProductID)
	if err != nil {
		return err
	}
	return nil
}
