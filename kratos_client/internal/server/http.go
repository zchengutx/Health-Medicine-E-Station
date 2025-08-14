package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	doctorsv1 "kratos_client/api/doctors/v1"
	drug "kratos_client/api/drug/v1"
	estimate "kratos_client/api/estimate/v1"
	userv1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/conf"
	"kratos_client/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, doctors *service.DoctorsService, drugs *service.DrugService, estimates *service.EstimateService, user *service.UserService, logger log.Logger) *http.Server {
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
	doctorsv1.RegisterDoctorsHTTPServer(srv, doctors)
	drug.RegisterDrugHTTPServer(srv, drugs)
	estimate.RegisterEstimateHTTPServer(srv, estimates)
	//chatv1.RegisterChatHTTPServer(srv, chat)
	//cartv1.RegisterCartHTTPServer(srv, cart)
	// Kiro修改：注册用户相关的HTTP路由
	userv1.RegisterUserHTTPServer(srv, user)

	// Kiro修改：注册需要JWT认证的路由
	srv.Route("/").POST("/upload", user.Upload, comment.JWTMiddleware())
	srv.Route("/").POST("/GetTargeted", user.GetTargeted, comment.JWTMiddleware())

	// Kiro修改：添加心跳检测相关API
	// 获取用户在线状态
	srv.Route("/").GET("/api/heartbeat/user/{userId}/status", func(ctx http.Context) error {
		userID := ctx.Vars()["userId"][0]
		status := comment.GetUserOnlineStatus(userID)
		return ctx.JSON(200, map[string]interface{}{
			"code":    0,
			"message": "success",
			"data":    status,
		})
	})

	// Kiro修改：获取所有在线用户
	srv.Route("/").GET("/api/heartbeat/online-users", func(ctx http.Context) error {
		onlineUsers := comment.GetAllOnlineUsers()
		return ctx.JSON(200, map[string]interface{}{
			"code":    0,
			"message": "success",
			"data": map[string]interface{}{
				"online_users": onlineUsers,
				"count":        len(onlineUsers),
			},
		})
	})

	// Kiro修改：手动更新用户心跳（用于测试）
	srv.Route("/").POST("/api/heartbeat/update", func(ctx http.Context) error {
		var req struct {
			UserID string `json:"user_id"`
		}
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(400, map[string]interface{}{
				"code":    1,
				"message": "invalid request",
			})
		}
		comment.UpdateUserHeartbeat(req.UserID)
		return ctx.JSON(200, map[string]interface{}{
			"code":    0,
			"message": "heartbeat updated",
		})
	}, comment.JWTMiddleware())

	// WebSocket聊天路由
	srv.Route("/").GET("/ws/chat", func(ctx http.Context) error {
		// 从查询参数或Header获取token进行验证
		comment.HandleWebSocket(ctx.Response(), ctx.Request(), ctx)
		return nil
	}, comment.JWTMiddleware())

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
			"code":    0,
			"message": "success",
			"data":    users,
		})
	})

	return srv
}
