package mysql

import (
	"context"
	"time"
	"user-api/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func NewMySqlConn(c config.MySqlConfig) *sqlx.SqlConn {
	mysqlConn := sqlx.NewMysql(c.Dsn)

	db, err := mysqlConn.RawDB()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(c.ConnectTimeout))
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)
	
	return &mysqlConn
}

		