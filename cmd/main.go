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
	db          *gorm.DB                       = configs.SetupDatabaseConnection()
	userRepo    repositories.UserRepository    = repositories.NewUserRepository(db)
	productRepo repositories.ProductRepository = repositories.NewProductRepository(db)
	cartRepo    repositories.CartRepository    = repositories.NewCartRepository(db)

	tokenUC   usecases.Token   = usecases.NewTokenUc()
	authUC    usecases.Auth    = usecases.NewAuthUC(userRepo)
	productUC usecases.Product = usecases.NewProductUC(productRepo)
	cartUC    usecases.Cart    = usecases.NewCartUC(cartRepo, productRepo)

	authController    controllers.AuthController    = controllers.NewAuthController(authUC, tokenUC)
	productController controllers.ProductController = controllers.NewProductController(productUC)
	cartController    controllers.CartController    = controllers.NewCartController(cartUC, tokenUC)
)

func main() {

	r := gin.Default()

	apiRoutes := r.Group("api/v1/")
	{
		authRoutes := apiRoutes.Group("auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
			authRoutes.GET("/logout", authController.Logout, middleware.AuthorizeJWT(tokenUC))
		}
		productRoutes := apiRoutes.Group("products")
		{
			productRoutes.GET("/", productController.GetProducts, middleware.AuthorizeJWT(tokenUC))
		}
		cartRoutes := apiRoutes.Group("cart")
		{
			cartRoutes.POST("/", cartController.CreateCart, middleware.AuthorizeJWT(tokenUC))
		}
	}
	r.Run()
}
