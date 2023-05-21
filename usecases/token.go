package usecases

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iskhakmuhamad/ecommerce/models"
	"github.com/iskhakmuhamad/ecommerce/usecases/token"
)

type Token interface {
	GenerateToken(ctx context.Context, user *models.User) (*token.ResultResponse, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func NewTokenUc() Token {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	return secretKey
}

func (u *jwtService) GenerateToken(ctx context.Context, user *models.User) (*token.ResultResponse, error) {

	now := time.Now()
	end := now.Add(time.Minute * 15)
	claims := &token.AccessCustomClaim{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: end.Unix(),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := newToken.SignedString([]byte(u.secretKey))
	if err != nil {
		return nil, err
	}

	res := &token.ResultResponse{
		Name:      user.Name,
		Token:     tokenStr,
		ExpiredAt: end.Format(time.RFC3339),
	}
	return res, nil
}

func (u *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	header := strings.Split(token, "Bearer ")
	if len(header) > 1 {
		token = header[1]
	}
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(u.secretKey), nil
	})
}
