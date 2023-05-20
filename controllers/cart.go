package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/ecommerce/shared"
	"github.com/iskhakmuhamad/ecommerce/usecases"
	"github.com/iskhakmuhamad/ecommerce/usecases/cart"
)

type cartController struct {
	cartUc  usecases.Cart
	tokenUc usecases.Token
}

type CartController interface {
	CreateCart(ctx *gin.Context)
}

func NewCartController(cartUc usecases.Cart, tokenUc usecases.Token) CartController {
	return &cartController{
		cartUc:  cartUc,
		tokenUc: tokenUc,
	}
}

func (c *cartController) CreateCart(ctx *gin.Context) {
	var (
		request cart.AddCartRequest
	)

	err := ctx.Bind(&request)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	//get user id from token
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.tokenUc.ValidateToken(authHeader)
	if err != nil {
		response := shared.BuildErrorResponse("Malformat Token", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["id"])
	request.UserID, err = strconv.ParseInt(userID, 10, 64)
	if err != nil {
		response := shared.BuildErrorResponse("Malformat Token", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	//create cart
	err = c.cartUc.CreateCart(ctx, request)
	if err != nil {
		res := shared.BuildErrorResponse("Failed Adding Cart!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Success Adding Cart!", nil)
	ctx.JSON(http.StatusCreated, res)
}
