package repositories

import (
	"context"
	"errors"

	"github.com/iskhakmuhamad/ecommerce/models"
	"gorm.io/gorm"
)

type cartRepository struct {
	qry *gorm.DB
}

type CartRepository interface {
	UpSertCart(ctx context.Context, params *models.Cart) error
	GetCartbyUserIDnProductID(ctx context.Context, userID int64, productID int64) (*models.Cart, error)
	GetCartProductsByUserID(ctx context.Context, userID int64) ([]models.CartProducts, error)
	DeleteCartProduct(ctx context.Context, userID int64, productID int64) error
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

func (r *cartRepository) GetCartbyUserIDnProductID(ctx context.Context, userID int64, productID int64) (*models.Cart, error) {
	var (
		cart *models.Cart
	)
	if err := r.qry.Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return cart, nil
}

func (r *cartRepository) GetCartProductsByUserID(ctx context.Context, userID int64) ([]models.CartProducts, error) {
	var (
		carts = []models.CartProducts{}
	)
	err := r.qry.Model(&models.Cart{}).Select("*").Joins(" JOIN products ON products.id = carts.product_id").Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *cartRepository) DeleteCartProduct(ctx context.Context, userID int64, productID int64) error {

	if err := r.qry.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Cart{}).Error; err != nil {
		return err
	}
	return nil
}
