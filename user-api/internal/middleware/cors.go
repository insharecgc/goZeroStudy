// middleware/cors.go
package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

// CorsMiddleware 自定义跨域中间件
func CorsMiddleware(allowedOrigins []string) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 1. 获取请求来源（Origin头）
			origin := r.Header.Get("Origin")

			// 2. 判断是否允许该来源（支持通配符或动态逻辑）
			allowed := false
			if len(allowedOrigins) == 0 {
				allowed = true // 允许所有来源（不推荐生产环境）
			} else {
				for _, o := range allowedOrigins {
					if o == "*" || o == origin {
						allowed = true
						break
					}
				}
			}

			// 3. 设置跨域响应头
			if allowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Custom-Header")
				w.Header().Set("Access-Control-Allow-Credentials", "true") // 允许携带Cookie
				w.Header().Set("Access-Control-Max-Age", "86400")          // 预检缓存时间
			}

			// 4. 处理OPTIONS预检请求（直接返回204）
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// 5. 继续执行后续处理
			next(w, r)
		}
	}
}
