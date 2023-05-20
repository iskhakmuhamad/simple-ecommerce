package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/ecommerce/shared"
	"github.com/iskhakmuhamad/ecommerce/usecases"
	"github.com/iskhakmuhamad/ecommerce/usecases/payment"
)

type paymentController struct {
	paymentUc usecases.Payment
	tokenUc   usecases.Token
}

type PaymentController interface {
	CreatePayment(ctx *gin.Context)
	GetUserPayments(ctx *gin.Context)
}

func NewPaymentController(paymentUc usecases.Payment, tokenUc usecases.Token) PaymentController {
	return &paymentController{
		paymentUc: paymentUc,
		tokenUc:   tokenUc,
	}
}

func (c *paymentController) CreatePayment(ctx *gin.Context) {
	var (
		request payment.AddPaymentRequest
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

	//create payment
	err = c.paymentUc.CreatePayment(ctx, request)
	if err != nil {
		res := shared.BuildErrorResponse("Failed Checkout!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Success Checkout!", nil)
	ctx.JSON(http.StatusCreated, res)
}

func (c *paymentController) GetUserPayments(ctx *gin.Context) {

	//get user id from token
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.tokenUc.ValidateToken(authHeader)
	if err != nil {
		response := shared.BuildErrorResponse("Malformat Token", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["id"])
	userID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		response := shared.BuildErrorResponse("Malformat Token", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	resp, err := c.paymentUc.GetUserPayments(ctx, userID)
	if err != nil {
		res := shared.BuildErrorResponse("Failed Get User Payments!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Success Get User Payments!", resp)
	ctx.JSON(http.StatusOK, res)
}
