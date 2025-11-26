// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MySqlConfig MySqlConfig `json:"mysqlConfig"`
	JWTConfig   JWTConfig   `json:"jwt"`
}

type MySqlConfig struct {
	Dsn             string `json:"dsn"`
	MaxIdleConns    int    `json:"maxIdleConns"`
	MaxOpenConns    int    `json:"maxOpenConns"`
	ConnectTimeout  int    `json:"connectTimeout"`
}

// JWT 配置
type JWTConfig struct {
	SecretKey      string        `json:"secretKey"`
	ExpirationTime string        `json:"expirationTime"`
}