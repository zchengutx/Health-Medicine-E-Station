package server

import (
	nethttp "net/http"
	
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	cartv1 "kratos_client/api/cart/v1"
	chatv1 "kratos_client/api/chat/v1"
	couponv1 "kratos_client/api/coupon/v1"
	doctorsv1 "kratos_client/api/doctors/v1"
	drug "kratos_client/api/drug/v1"
	estimate "kratos_client/api/estimate/v1"
	orderv1 "kratos_client/api/order/v1"
	prescriptionv1 "kratos_client/api/prescription/v1"
	userv1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/conf"
	"kratos_client/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, doctors *service.DoctorsService, drugs *service.DrugService, estimates *service.EstimateService, chat *service.ChatService, cart *service.CartService, order *service.OrderService, coupon *service.CouponService, prescription *service.PrescriptionService, logger log.Logger) *http.Server {
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
	cartv1.RegisterCartHTTPServer(srv, cart)
	orderv1.RegisterOrderServiceHTTPServer(srv, order)
	couponv1.RegisterCouponServiceHTTPServer(srv, coupon)
	prescriptionv1.RegisterPrescriptionServiceHTTPServer(srv, prescription)
	srv.Route("/").POST("/upload", user.Upload, comment.JWTMiddleware())
	srv.Route("/").POST("/GetTargeted", user.GetTargeted, comment.JWTMiddleware())
	
	// 可选的WebSocket聊天功能（用于实时消息推送）
	srv.Route("/").GET("/ws/chat", func(ctx http.Context) error {
		comment.HandleWebSocket(ctx.Response(), ctx.Request())
		return nil
	})
	
	// 聊天测试页面
	srv.Route("/").GET("/chat", func(ctx http.Context) error {
		nethttp.ServeFile(ctx.Response(), ctx.Request(), "./static/chat.html")
		return nil
	})
	return srv
}
