package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "kratos_client/api/helloworld/v1"
	userv1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/conf"
	"kratos_client/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
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
	v1.RegisterGreeterHTTPServer(srv, greeter)
	userv1.RegisterUserHTTPServer(srv, user)
	srv.Route("/").POST("/upload", user.Upload, comment.JWTMiddleware())
	return srv
}
