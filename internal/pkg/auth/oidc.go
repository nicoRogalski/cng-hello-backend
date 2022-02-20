package auth

import (
	"encoding/json"
	"net/http"

	"github.com/rogalni/cng-hello-backend/config"
	"github.com/rs/zerolog/log"
)

var cts Certs

type Certs struct {
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

func Setup() {
	u := config.App.JwtCertUrl
	if u == "" {
		log.Warn().Msg("Server without OIDC Endpoint for secret")
		return
	}
	r, err := http.Get(u)
	if err != nil {
		log.Warn().Msg("Could not fetch JWT Certificate")
	}
	er := json.NewDecoder(r.Body).Decode(&cts)
	if er != nil {
		log.Warn().Msg("Could not decode JWT Certificate")
	}
}

func GetCert(kid string) (*Key, bool) {
	for _, v := range cts.Keys {
		if kid == v.Kid {
			return &v, true
		}
	}
	return nil, false
}
