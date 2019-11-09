package jwt

import (
	"time"

	jose "github.com/dvsekhvalnov/jose2go"
)

//GenerateToken generates and returns a JWT token
func GenerateToken(payload string, key []byte) (string, error) {
	headers := map[string]interface{}{
		"typ":        "JWT",
		"alg":        "HS256",
		"expiration": int64(time.Now().Unix()) + 3600,
	}
	token, err := jose.Sign(payload, jose.HS256, key, jose.Headers(headers))
	if err != nil {
		return "", err
	}
	return token, nil
}

//DecodeToken decodes and validates a token by the key
func DecodeToken(token string, key []byte) (string, interface{}) {
	payload, headers, err := jose.Decode(token, key)
	if err != nil {
		return "", false
	}
	return payload, headers
}
