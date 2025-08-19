package server

import (
	cartv1 "kratos_client/api/cart/v1"
	doctorsv1 "kratos_client/api/doctors/v1"
	drug "kratos_client/api/drug/v1"
	estimate "kratos_client/api/estimate/v1"
	userv1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/conf"
	"kratos_client/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, doctors *service.DoctorsService, drugs *service.DrugService, estimates *service.EstimateService, user *service.UserService, cart *service.CartService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(comment.CorsFilter()),
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
	// srv.Handle("/", comment.CorsHandler(srv.Handler))
	// 注册核心服务
	doctorsv1.RegisterDoctorsHTTPServer(srv, doctors)
	drug.RegisterDrugHTTPServer(srv, drugs)
	estimate.RegisterEstimateHTTPServer(srv, estimates)

	// 为用户服务添加JWT中间件
	userv1.RegisterUserHTTPServer(srv, user)
	cartv1.RegisterCartHTTPServer(srv, cart)
	// 其他服务暂时注释，避免编译错误
	// orderv1.RegisterOrderHTTPServer(srv, order)
	// paymentv1.RegisterPaymentHTTPServer(srv, payment)
	// couponv1.RegisterCouponHTTPServer(srv, coupon)
	// prescriptionv1.RegisterPrescriptionHTTPServer(srv, prescription)
	// chatv1.RegisterChatHTTPServer(srv, chat)

	srv.Route("/").POST("/upload", user.Upload, comment.JWTMiddleware())
	srv.Route("/").POST("/GetTargeted", user.GetTargeted, comment.JWTMiddleware())

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

	// WebSocket聊天路由 - 不使用JWT中间件，在WebSocket处理器内部验证token
	// WebSocket功能暂时移除以简化启动

	return srv
}
