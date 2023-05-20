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
	GetUserPayments(ctx context.Context, userID int64) ([]models.PaymentDetail, error)
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
	params.TotalPrice = product.ProductPrice * float64(params.Amount)

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

func (u *paymentUC) GetUserPayments(ctx context.Context, userID int64) ([]models.PaymentDetail, error) {
	payments, err := u.paymentRepo.GetPaymentByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
