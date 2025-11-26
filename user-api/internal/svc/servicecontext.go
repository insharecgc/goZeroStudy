// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"user-api/internal/config"
	"user-api/internal/mysql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	MySql sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := mysql.NewMySqlConn(c.MySqlConfig)
	return &ServiceContext{
		Config: c,
		MySql: *mysqlConn,
	}
}
