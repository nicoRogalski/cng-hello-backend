package auth

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

var cts certs

type certs struct {
	Keys []Key `json:"keys"`
}
type Key struct {
	Kid     string   `json:"kid"`
	Kty     string   `json:"kty"`
	Alg     string   `json:"alg"`
	Use     string   `json:"use"`
	N       string   `json:"n"`
	E       string   `json:"e"`
	X5C     []string `json:"x5c"`
	X5T     string   `json:"x5t"`
	X5TS256 string   `json:"x5t#S256"`
}

func Setup(oauthJwtCertUrl string) {
	if oauthJwtCertUrl == "" {
		log.Warn().Msg("Server without OIDC Endpoint for secret")
		return
	}
	r, err := http.Get(oauthJwtCertUrl)
	if err != nil {
		log.Warn().Msg("Could not fetch JWT Certificate")
	}
	er := json.NewDecoder(r.Body).Decode(&cts)
	if er != nil {
		log.Warn().Msg("Could not decode JWT Certificate")
	}
}

func getRsaKey(kid string) (string, bool) {
	cert, found := getCert(kid)
	if !found {
		return "", false
	}
	return cert.X5C[0], true

}

func getCert(kid string) (*Key, bool) {
	for _, v := range cts.Keys {
		if kid == v.Kid {
			return &v, true
		}
	}
	return nil, false
}
