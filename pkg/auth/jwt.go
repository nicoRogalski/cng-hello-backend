package auth

import (
	"errors"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/rs/zerolog/log"
)

const (
	bearerSchema = "Bearer "
)

var Jwks *keyfunc.JWKS

func Setup(oauthJwtCertUrl string) {
	if oauthJwtCertUrl == "" {
		log.Warn().Msg("Server starts without OIDC Endpoint for secret")
		return
	}
	options := keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	// Create the JWKS from the resource at the given URL.
	j, err := keyfunc.Get(oauthJwtCertUrl, options)
	if err != nil {
		log.Warn().Msgf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
	}
	Jwks = j
}

func ExtractJWT(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("no authorization header present")
	}
	return authHeader[len(bearerSchema):], nil
}
