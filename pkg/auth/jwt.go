package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

const (
	// Key in the Headers hashmap of the token that points to the key ID
	keyIdTokenHeaderKey = "kid"

	// Header and footer to attach to base64-encoded key data that we receive from Auth0
	pubKeyHeader = "-----BEGIN CERTIFICATE-----"
	pubKeyFooter = "-----END CERTIFICATE-----"
	bearerSchema = "Bearer "
	AuthHeader   = "Authorization"
)

func ExtractJWT(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("no authorization header present")
	}
	return authHeader[len(bearerSchema):], nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// IMPORTANT: Validating the algorithm per https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf(
				"expected token algorithm '%v' but got '%v'",
				jwt.SigningMethodRS256.Name,
				token.Header)
		}

		untypedKeyId, found := token.Header[keyIdTokenHeaderKey]
		if !found {
			return nil, fmt.Errorf("no key ID key '%v' found in token header", keyIdTokenHeaderKey)
		}
		keyId, ok := untypedKeyId.(string)
		if !ok {
			return nil, fmt.Errorf("found key ID, but value was not a string")
		}

		key, found := getRsaKey(keyId)
		if !found {
			return nil, fmt.Errorf("no public RSA key found corresponding to key ID from token '%v'", keyId)
		}

		keyStr := pubKeyHeader + "\n" + key + "\n" + pubKeyFooter

		// Since the token is RSA (which we validated at the start of this function), the return type of this function actually has to be rsa.PublicKey!
		pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(keyStr))
		if err != nil {
			return nil, fmt.Errorf("an error occurred parsing the public key base64 for key ID '%v'; this is a code bug", keyId)
		}

		return pubKey, nil
	},
	)
}
