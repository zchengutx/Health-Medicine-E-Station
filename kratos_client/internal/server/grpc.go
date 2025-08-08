package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	cartv1 "kratos_client/api/cart/v1"
	couponv1 "kratos_client/api/coupon/v1"
	doctorsv1 "kratos_client/api/doctors/v1"
	drug "kratos_client/api/drug/v1"
	estimate "kratos_client/api/estimate/v1"
	orderv1 "kratos_client/api/order/v1"
	prescriptionv1 "kratos_client/api/prescription/v1"
	"kratos_client/internal/conf"
	"kratos_client/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, doctors *service.DoctorsService, drugs *service.DrugService, estimates *service.EstimateService, cart *service.CartService, order *service.OrderService, coupon *service.CouponService, prescription *service.PrescriptionService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	doctorsv1.RegisterDoctorsServer(srv, doctors)
	drug.RegisterDrugServer(srv, drugs)
	estimate.RegisterEstimateServer(srv, estimates)
	cartv1.RegisterCartServer(srv, cart)
	orderv1.RegisterOrderServiceServer(srv, order)
	couponv1.RegisterCouponServiceServer(srv, coupon)
	prescriptionv1.RegisterPrescriptionServiceServer(srv, prescription)

	return srv
}
