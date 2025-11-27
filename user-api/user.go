// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
	"net/http"

	"user-api/internal/config"
	"user-api/internal/handler"
	"user-api/internal/middleware"
	"user-api/internal/svc"
	"user-api/util"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 注册跨域中间件（全局生效）
	server.Use(middleware.CorsMiddleware([]string{"http://localhost:3000"}))

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 统一的错误管理
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *util.Errno:
			return http.StatusOK, util.Error(e)
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
