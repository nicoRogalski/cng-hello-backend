package service

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/rogalni/cng-hello-backend/internal/utils/config"
)

var secret string

func init() {
	js := config.Cfg.JwtSecret
	if js == "" {
		// _, err := http.Get(config.Cfg.JwtCertUrl)
		// if err != nil {
		// 	log.Warn().Msg("Could not fetch JWT Certificate")
		// }

		// parse secert
		// secret = r.Body.Read()
	} else {
		secret = js
	}
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
