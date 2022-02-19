package service

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/rogalni/cng-hello-backend/internal/pkg/auth"
)

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(auth.Secret), nil
	})
}
