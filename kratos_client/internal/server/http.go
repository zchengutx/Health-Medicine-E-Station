package server

import (
	nethttp "net/http"
	
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	chatv1 "kratos_client/api/chat/v1"
	doctorsv1 "kratos_client/api/doctors/v1"
	drug "kratos_client/api/drug/v1"
	estimate "kratos_client/api/estimate/v1"
	userv1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/conf"
	"kratos_client/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, doctors *service.DoctorsService, drugs *service.DrugService, estimates *service.EstimateService, chat *service.ChatService, logger log.Logger) *http.Server {
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
	userv1.RegisterUserHTTPServer(srv, user)
	doctorsv1.RegisterDoctorsHTTPServer(srv, doctors)
	drug.RegisterDrugHTTPServer(srv, drugs)
	estimate.RegisterEstimateHTTPServer(srv, estimates)
	chatv1.RegisterChatHTTPServer(srv, chat)
	srv.Route("/").POST("/upload", user.Upload, comment.JWTMiddleware())
	srv.Route("/").POST("/GetTargeted", user.GetTargeted, comment.JWTMiddleware())
	
	// WebSocket聊天路由
	srv.Route("/").GET("/ws/chat", func(ctx http.Context) error {
		comment.HandleWebSocket(ctx.Response(), ctx.Request())
		return nil
	})
	
	// 简化的WebSocket测试路由（用于调试）
	srv.Route("/").GET("/ws/test", func(ctx http.Context) error {
		comment.HandleTestWebSocket(ctx.Response(), ctx.Request())
		return nil
	})
	
	// 获取房间用户列表的API
	srv.Route("/").GET("/api/chat/room/{roomId}/users", func(ctx http.Context) error {
		roomID := ctx.Vars()["roomId"][0] // 取第一个元素
		users := comment.GetRoomUsers(roomID)
		return ctx.JSON(200, map[string]interface{}{
			"code": 0,
			"message": "success",
			"data": users,
		})
	})
	
	// 静态文件服务（用于测试页面）
	srv.Route("/").GET("/chat", func(ctx http.Context) error {
		nethttp.ServeFile(ctx.Response(), ctx.Request(), "./static/chat.html")
		return nil
	})
	return srv
}
