package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/ecommerce/configs"
	"github.com/iskhakmuhamad/ecommerce/controllers"
	"github.com/iskhakmuhamad/ecommerce/middleware"
	"github.com/iskhakmuhamad/ecommerce/repositories"
	"github.com/iskhakmuhamad/ecommerce/usecases"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB                    = configs.SetupDatabaseConnection()
	repo    repositories.UserRepository = repositories.NewUserRepository(db)
	tokenUC usecases.Token              = usecases.NewTokenUc()
	authUC  usecases.Auth               = usecases.NewAuthUC(repo)
	ac      controllers.AuthController  = controllers.NewAuthController(authUC, tokenUC)
)

func main() {

	r := gin.Default()

	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/register", ac.Register)
		authRoutes.POST("/login", ac.Login)
		authRoutes.POST("/logout", ac.Logout, middleware.AuthorizeJWT(tokenUC))
	}
	r.Run()
}
