package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	userv1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/conf"
	"kratos_client/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		// Kiro修改：使用HTTP过滤器处理跨域，移除中间件方式
		http.Filter(comment.CorsFilter()), // Kiro修改：添加跨域过滤器
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	// Kiro修改：移除错误的跨域处理方式
	// srv.Handle("/", comment.CorsHandler(srv.Handler))

	// Kiro修改：注册用户相关的HTTP路由
	userv1.RegisterUserHTTPServer(srv, user)

	// Kiro修改：注册需要JWT认证的路由
	srv.Route("/").POST("/upload", user.Upload, comment.JWTMiddleware())
	srv.Route("/").POST("/GetTargeted", user.GetTargeted, comment.JWTMiddleware())

	return srv
}
