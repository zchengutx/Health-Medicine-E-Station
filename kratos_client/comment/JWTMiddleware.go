package comment

import (
	"context"
	"net/http"
)

var jwtSecret = []byte("your-secret-key") // 替换为实际密钥

// JWTMiddleware 用于验证JWT的HTTP中间件
func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. 获取Authorization头
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "缺少Authorization头", http.StatusUnauthorized)
				return
			}

			//// 2. 解析Bearer token
			//parts := strings.SplitN(authHeader, " ", 2)
			//if len(parts) != 2 || parts[0] != "Bearer" {
			//	http.Error(w, "Authorization格式错误", http.StatusUnauthorized)
			//	return
			//}

			token, s := GetToken(authHeader)

			if token == nil || s != "" {
				http.Error(w, "无效的Token", http.StatusUnauthorized)
				return
			}

			// 将用户ID存入请求上下文
			ctx := context.WithValue(r.Context(), "user_id", token["user"])
			r = r.WithContext(ctx)

			// 5. 继续处理请求
			next.ServeHTTP(w, r)
		})
	}
}
