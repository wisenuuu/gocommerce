package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("ini-adalah-jwt-secret-key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			// Periksa apakah token kadaluarsa
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					c.JSON(401, gin.H{"message": "Token has expired"})
					c.Abort()
					return
				}
			}
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Jika token valid, ambil klaimnya
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)

		c.Next()
	}
}
