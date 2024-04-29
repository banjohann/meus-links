package api

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const secret = "mydeepdarksecret"

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWTToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
