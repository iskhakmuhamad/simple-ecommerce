package usecases

import (
	"context"

	"github.com/iskhakmuhamad/ecommerce/models"
	"github.com/iskhakmuhamad/ecommerce/repositories"
	"github.com/iskhakmuhamad/ecommerce/usecases/product"
)

type productUC struct {
	repo repositories.ProductRepository
}

type Product interface {
	GetProducts(ctx context.Context, params product.ProductsRequest) ([]models.Product, error)
	CreateProduct(ctx context.Context, params product.AddProductRequest) error
}

func NewProductUC(r repositories.ProductRepository) Product {
	return &productUC{
		repo: r,
	}
}

func (u *productUC) GetProducts(ctx context.Context, params product.ProductsRequest) ([]models.Product, error) {

	if err := params.Validate(); err != nil {
		return nil, err
	}

	products, err := u.repo.GetProducts(ctx, &models.Product{
		ProductName:     params.SearchName,
		ProductCategory: params.ProductCategory,
	})

	if err != nil {
		return nil, err
	}

	return products, nil
}
func (u *productUC) CreateProduct(ctx context.Context, params product.AddProductRequest) error {

	if err := params.Validate(); err != nil {
		return err
	}

	err := u.repo.InsertProduct(ctx, &models.Product{
		ProductName:     params.ProductName,
		ProductPrice:    params.ProductPrice,
		ProductQty:      params.ProductQty,
		ProductCategory: params.ProductCategory,
	})

	if err != nil {
		return err
	}

	return nil
}
