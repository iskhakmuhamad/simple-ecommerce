package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/ecommerce/shared"
	"github.com/iskhakmuhamad/ecommerce/usecases"
	"github.com/iskhakmuhamad/ecommerce/usecases/product"
)

type productController struct {
	productUc usecases.Product
}

type ProductController interface {
	GetProducts(ctx *gin.Context)
}

func NewProductController(productUc usecases.Product) ProductController {
	return &productController{
		productUc: productUc,
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
