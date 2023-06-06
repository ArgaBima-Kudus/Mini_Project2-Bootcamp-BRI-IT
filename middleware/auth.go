package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type response struct {
	Success string `json:"success"`
	Message string `json:"message"`
}

func Auth(c *gin.Context) {

	// Token yang diterima
	authToken := c.GetHeader("Authorization")
	removeBearer := strings.Split(authToken, " ")
	// Verifikasi token dengan kunci rahasia
	token, err := jwt.Parse(removeBearer[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		c.JSON(401, response{Success: "false", Message: "token q tidak valid"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("role", claims["sub"])
	} else {
		c.JSON(401, response{Success: "false", Message: "token 2 tidak valid"})
		c.Abort()
		return
	}
	c.Next()
}
