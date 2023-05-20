package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/ecommerce/shared"
	"github.com/iskhakmuhamad/ecommerce/usecases"
)

func AuthorizeJWT(u usecases.Token) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := shared.BuildErrorResponse("Failed", "No token found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := u.ValidateToken(authHeader)
		if token.Valid {
			_ = token.Claims.(jwt.MapClaims)
		} else {
			log.Println(err)
			response := shared.BuildErrorResponse("Token is not valid", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
