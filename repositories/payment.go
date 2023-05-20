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
	GetPaymentByUserID(ctx context.Context, userID int64) ([]models.PaymentDetail, error)
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

func (r *paymentRepository) GetPaymentByUserID(ctx context.Context, userID int64) ([]models.PaymentDetail, error) {
	var payments []models.PaymentDetail

	err := r.qry.Model(&models.Payment{}).Select("*").Joins(" JOIN products ON products.id = payments.product_id JOIN users ON users.id = payments.user_id").Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}
