package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/iskhakmuhamad/ecommerce/models"

	"gorm.io/gorm"
)

type productRepository struct {
	qry *gorm.DB
}

type ProductRepository interface {
	GetProducts(ctx context.Context, params *models.Product) ([]models.Product, error)
	GetProductByID(ctx context.Context, id int64) (*models.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		qry: db,
	}
}

func (r *productRepository) GetProducts(ctx context.Context, params *models.Product) ([]models.Product, error) {
	var (
		products []models.Product
	)

	db := r.qry.Model(models.Product{})

	if params.ProductCategory != "" {
		db = db.Where("product_category = ?", params.ProductCategory)
	}
	if params.ProductName != "" {
		name := fmt.Sprintf("%%%s%%", params.ProductName)
		db = db.Where("product_name LIKE ?", name)
	}

	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetProductByID(ctx context.Context, ID int64) (*models.Product, error) {
	var (
		product *models.Product
	)

	if err := r.qry.Model(models.Product{}).Where("id = ?", ID).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return product, nil
}
