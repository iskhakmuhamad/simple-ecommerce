package repositories

import (
	"context"

	"github.com/iskhakmuhamad/ecommerce/models"
	"gorm.io/gorm"
)

type paymentRepository struct {
	qry *gorm.DB
}

type PaymentRepository interface {
	InsertPayment(ctx context.Context, params *models.Payment) error
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		qry: db,
	}
}

func (r *paymentRepository) InsertPayment(ctx context.Context, params *models.Payment) error {
	var payment *models.Payment

	if err := r.qry.Model(&payment).Create(params).Error; err != nil {
		return err
	}
	return nil
}
