package usecases

import (
	"context"
	"errors"

	"github.com/iskhakmuhamad/ecommerce/models"
	"github.com/iskhakmuhamad/ecommerce/repositories"
	"github.com/iskhakmuhamad/ecommerce/shared"
	"github.com/iskhakmuhamad/ecommerce/usecases/auth"
)

type authUC struct {
	repo repositories.UserRepository
}

type Auth interface {
	Register(ctx context.Context, params auth.RegisterRequest) error
	Login(ctx context.Context, params auth.LoginRequest) (*models.User, error)
	Logout(ctx context.Context, email string) error
}

func NewAuthUC(r repositories.UserRepository) Auth {
	return &authUC{
		repo: r,
	}
}

func (u *authUC) Register(ctx context.Context, params auth.RegisterRequest) error {
	var (
		encryptedPassword string
		err               error
		user              *models.User
	)

	if err = params.Validate(); err != nil {
		return err
	}

	user, _ = u.repo.GetUserByEmail(ctx, params.Email)

	if user != nil {
		return errors.New("email is used")
	}

	encryptedPassword, err = shared.EncryptPassword(params.Password)
	if err != nil {
		return err
	}

	req := &models.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: encryptedPassword,
		Role:     params.Role,
	}

	err = u.repo.InsertUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *authUC) Login(ctx context.Context, params auth.LoginRequest) (*models.User, error) {

	var (
		err  error
		user *models.User
	)

	if err = params.Validate(); err != nil {
		return nil, err
	}

	user, err = u.repo.GetUserByEmail(ctx, params.Email)
	if err != nil || user == nil {
		return nil, errors.New("email not found")
	}

	err = shared.CheckPassword(params.Password, user.Password)
	if err != nil {
		return nil, errors.New("wrong password")
	}

	user, _ = u.repo.GetUserByEmail(ctx, params.Email)

	return user, nil
}

func (u *authUC) Logout(ctx context.Context, email string) error {

	// var (
	// 	err  error
	// 	user *models.User
	// )

	// user, _ = u.repo.GetUserByEmail(ctx, email)

	return nil
}
