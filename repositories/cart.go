package repositories

import (
	"context"

	"github.com/iskhakmuhamad/ecommerce/models"
	"gorm.io/gorm"
)

type cartRepository struct {
	qry *gorm.DB
}

type CartRepository interface {
	UpSertCart(ctx context.Context, params *models.Cart) error
	GetCartbyUserIDnProductID(ctx context.Context, params *models.Cart) (*models.Cart, error)
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		qry: db,
	}
}

func (r *cartRepository) UpSertCart(ctx context.Context, params *models.Cart) error {

	data := map[string]interface{}{"product_id": params.ProductID, "user_id": params.UserID, "amount": params.Amount}

	if err := r.qry.Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", params.UserID, params.ProductID).Save(data).Error; err != nil {
		return err
	}
	return nil
}

func (r *cartRepository) GetCartbyUserIDnProductID(ctx context.Context, params *models.Cart) (*models.Cart, error) {
	var (
		cart *models.Cart
	)
	if err := r.qry.Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", params.UserID, params.ProductID).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
