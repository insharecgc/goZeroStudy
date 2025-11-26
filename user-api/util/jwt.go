package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"user-api/internal/config"
)

type JWTUtil struct{
	JWTConfig config.JWTConfig
}

// JWTClaims JWT载荷
type JWTClaims struct {
	UserID   uint64   `json:"user_id"`
	Username string   `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func (j *JWTUtil) GenerateToken(userID uint64, username string) (string, error) {
	// 解析 JWT 有效期（字符串转 time.Duration）
	expiration, err := time.ParseDuration(j.JWTConfig.ExpirationTime)
	if err != nil {
		expiration = time.Hour * 24 // 默认 24 小时
	}
	
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.JWTConfig.SecretKey))
}

// ParseToken 解析JWT token，返回用户信息
func (j *JWTUtil) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.JWTConfig.SecretKey), nil
		},
	)
	retData := make(map[string]interface{})
	if err != nil {
		return retData, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		retData["userId"] = claims.UserID
		retData["username"] = claims.Username
		return retData, nil
	}
	return retData, errors.New("invalid token")
}
