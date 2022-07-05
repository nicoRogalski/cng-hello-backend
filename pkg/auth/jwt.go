package auth

import (
	"errors"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

const (
	bearerSchema = "Bearer "
)

var Jwks *keyfunc.JWKS

func Setup(oauthJwtCertUrl string) {
	if oauthJwtCertUrl == "" {
		otelzap.L().Warn("Server starts without OIDC Endpoint for secret")
		return
	}
	options := keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			otelzap.L().Warn("There was an error with the jwt.Keyfunc")
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	// Create the JWKS from the resource at the given URL.
	j, err := keyfunc.Get(oauthJwtCertUrl, options)
	if err != nil {
		otelzap.L().Warn("Failed to create JWKS from resource")
	}
	Jwks = j
}

func ExtractJWT(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("no authorization header present")
	}
	return authHeader[len(bearerSchema):], nil
}
