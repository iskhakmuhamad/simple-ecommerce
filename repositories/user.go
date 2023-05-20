package repositories

import (
	"context"

	"github.com/iskhakmuhamad/ecommerce/models"
	"gorm.io/gorm"
)

type repository struct {
	qry *gorm.DB
}

type UserRepository interface {
	InsertUser(ctx context.Context, params *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		qry: db,
	}
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User

	if err := r.qry.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) InsertUser(ctx context.Context, params *models.User) error {
	var user *models.User

	if err := r.qry.Model(&user).Create(params).Error; err != nil {
		return err
	}
	return nil
}
