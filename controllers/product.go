package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/ecommerce/shared"
	"github.com/iskhakmuhamad/ecommerce/usecases"
	"github.com/iskhakmuhamad/ecommerce/usecases/product"
)

type productController struct {
	productUc usecases.Product
	tokenUc   usecases.Token
}

type ProductController interface {
	GetProducts(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
}

func NewProductController(productUc usecases.Product, tokenUc usecases.Token) ProductController {
	return &productController{
		productUc: productUc,
		tokenUc:   tokenUc,
	}
}

func (c *productController) GetProducts(ctx *gin.Context) {
	var (
		request product.ProductsRequest
	)

	err := ctx.Bind(&request)

	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data, err := c.productUc.GetProducts(ctx, request)
	if err != nil {
		res := shared.BuildErrorResponse("Failed Get Products!", err.Error())
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := shared.BuildResponse("Success Get Products!", data)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var (
		request product.AddProductRequest
	)

	if err := ctx.Bind(&request); err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	//get role from token
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.tokenUc.ValidateToken(authHeader)
	if err != nil {
		response := shared.BuildErrorResponse("Malformat Token", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userRole := fmt.Sprintf("%v", claims["role"])
	if strings.ToLower(userRole) == "customer" {
		response := shared.BuildErrorResponse("Doesnt have permission", "User with role customer dont have permission to adding product to store")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err = c.productUc.CreateProduct(ctx, request)
	if err != nil {
		res := shared.BuildErrorResponse("Failed Adding New Product!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Success Adding New Product!", nil)
	ctx.JSON(http.StatusCreated, res)
}
