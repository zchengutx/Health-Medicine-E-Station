package comment

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
)

// Kiro修改：创建HTTP过滤器形式的跨域处理
// CorsFilter 返回一个HTTP过滤器用于处理跨域
func CorsFilter() kratoshttp.FilterFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Kiro修改：设置允许的源域名列表
			allowedOrigins := []string{
				"http://localhost:3000",
				"http://127.0.0.1:3000",
				"http://localhost:8080",
				"http://127.0.0.1:8080",
				"http://localhost:5173", // Vite默认端口
				"http://127.0.0.1:5173",
			}

			// Kiro修改：处理Origin头
			origin := r.Header.Get("Origin")
			if origin != "" {
				// Kiro修改：检查是否在允许列表中
				for _, allowedOrigin := range allowedOrigins {
					if allowedOrigin == origin {
						w.Header().Set("Access-Control-Allow-Origin", origin)
						break
					}
				}
			} else {
				// Kiro修改：如果没有Origin头，设置为*（适用于某些工具如Postman）
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}

			// Kiro修改：设置允许的HTTP方法
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD")
			
			// Kiro修改：设置允许的请求头
			w.Header().Set("Access-Control-Allow-Headers", 
				"Content-Type, Authorization, X-Requested-With, Accept, Origin, Cache-Control, X-File-Name, X-File-Size, X-File-Type")
			
			// Kiro修改：允许携带凭证（cookies等）
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			
			// Kiro修改：设置预检请求缓存时间（24小时）
			w.Header().Set("Access-Control-Max-Age", "86400")
			
			// Kiro修改：设置暴露的响应头
			w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")

			// Kiro修改：处理OPTIONS预检请求
			if strings.ToUpper(r.Method) == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent) // 204 No Content
				return
			}

			// Kiro修改：继续处理请求
			next.ServeHTTP(w, r)
		})
	}
}

// Kiro修改：保留中间件版本作为备用（简化版本）
// CorsMiddleware 返回一个简化的Kratos CORS中间件
func CorsMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// Kiro修改：直接处理请求，跨域头由过滤器处理
			return handler(ctx, req)
		}
	}
}