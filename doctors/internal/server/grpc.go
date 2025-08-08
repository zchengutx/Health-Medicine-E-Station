package server

import (
	consultationv1 "doctors/api/consultation/v1"
	doctorv1 "doctors/api/doctor/v1"
	v1 "doctors/api/helloworld/v1"
	patientv1 "doctors/api/patient/v1"
	"doctors/internal/conf"
	"doctors/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, doctor *service.DoctorService, patient *service.PatientService, consultation *service.ConsultationService, logger log.Logger) *grpc.Server {
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
	v1.RegisterGreeterServer(srv, greeter)
	doctorv1.RegisterDoctorServer(srv, doctor)
	patientv1.RegisterPatientServer(srv, patient)
	consultationv1.RegisterConsultationServer(srv, consultation)
	return srv
}
