package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("ini-adalah-jwt-secret-key")

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
