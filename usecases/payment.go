package usecases

import (
	"context"
	"errors"

	"github.com/iskhakmuhamad/ecommerce/models"
	"github.com/iskhakmuhamad/ecommerce/repositories"
	"github.com/iskhakmuhamad/ecommerce/usecases/payment"
)

type paymentUC struct {
	paymentRepo repositories.PaymentRepository
	productRepo repositories.ProductRepository
	cartRepo    repositories.CartRepository
}

type Payment interface {
	CreatePayment(ctx context.Context, params payment.AddPaymentRequest) error
}

func NewPaymentUC(paymentRepo repositories.PaymentRepository,
	productRepo repositories.ProductRepository,
	cartRepo repositories.CartRepository) Payment {
	return &paymentUC{
		paymentRepo: paymentRepo,
		productRepo: productRepo,
		cartRepo:    cartRepo,
	}
}

func (u *paymentUC) CreatePayment(ctx context.Context, params payment.AddPaymentRequest) error {

	if err := params.Validate(); err != nil {
		return err
	}
	product, err := u.productRepo.GetProductByID(ctx, params.ProductID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("Product doesnt found, please check again or reload the data")
	}
	params.TotalPrice = product.Price * float64(params.Amount)

	err = u.paymentRepo.InsertPayment(ctx, &models.Payment{
		UserID:     params.UserID,
		ProductID:  params.ProductID,
		Amount:     params.Amount,
		TotalPrice: params.TotalPrice,
	})
	if err != nil {
		return err
	}

	err = u.cartRepo.DeleteCartProduct(ctx, params.UserID, params.ProductID)
	if err != nil {
		return err
	}

	return nil
}
