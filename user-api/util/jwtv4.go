package util

import (
	"github.com/golang-jwt/jwt/v4"
)

func GetJwtToken(secretKey string, iat, seconds int64, payload any) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
