package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateToken(accessToken string) (*jwt.Token, error){
	return jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte("secret"), nil
	})
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}