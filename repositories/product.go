package repositories

import (
	"context"
	"fmt"

	"github.com/iskhakmuhamad/ecommerce/models"
	"github.com/iskhakmuhamad/ecommerce/usecases/product"

	"gorm.io/gorm"
)

type productRepository struct {
	qry *gorm.DB
}

type ProductRepository interface {
	GetProducts(ctx context.Context, params *product.ProductsRequest) ([]models.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		qry: db,
	}
}

func (r *productRepository) GetProducts(ctx context.Context, params *product.ProductsRequest) ([]models.Product, error) {
	var (
		products []models.Product
	)

	db := r.qry.Model(models.Product{})

	if params.Category != "" {
		db = db.Where("category = ?", params.Category)
	}
	if params.Name != "" {
		name := fmt.Sprintf("%%%s%%", params.Name)
		db = db.Where("name LIKE ?", name)
	}
	fmt.Printf("params: %v\n", params)

	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
