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
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, doctor *service.DoctorService, patient *service.PatientService, consultation *service.ConsultationService, logger log.Logger) *http.Server {
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
	doctorv1.RegisterDoctorHTTPServer(srv, doctor)
	patientv1.RegisterPatientHTTPServer(srv, patient)
	consultationv1.RegisterConsultationHTTPServer(srv, consultation)

	// 注册WebSocket路由
	srv.HandleFunc("/ws/consultation", consultation.HandleWebSocket)

	// 打印路由信息
	PrintRoutes(logger)

	return srv
}
