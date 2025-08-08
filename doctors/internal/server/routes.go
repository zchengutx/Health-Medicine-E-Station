package server

import (
	"github.com/go-kratos/kratos/v2/log"
)

// RouteInfo 路由信息结构
type RouteInfo struct {
	Method string
	Path   string
	Name   string
	Desc   string
}

// GetDoctorRoutes 获取医生模块的所有路由信息
func GetDoctorRoutes() []RouteInfo {
	return []RouteInfo{
		{
			Method: "POST",
			Path:   "/api/v1/doctor/SendSms",
			Name:   "SendSms",
			Desc:   "发送短信验证码",
		},
		{
			Method: "POST",
			Path:   "/api/v1/doctor/RegisterDoctor",
			Name:   "RegisterDoctor",
			Desc:   "医生注册",
		},
		{
			Method: "POST",
			Path:   "/api/v1/doctor/LoginDoctor",
			Name:   "LoginDoctor",
			Desc:   "医生登录",
		},
		{
			Method: "POST",
			Path:   "/api/v1/doctor/Authentication",
			Name:   "Authentication",
			Desc:   "医生认证",
		},
		{
			Method: "POST",
			Path:   "/api/v1/doctor/GetDoctorProfile",
			Name:   "GetDoctorProfile",
			Desc:   "获取医生个人信息",
		},
		{
			Method: "POST",
			Path:   "/api/v1/doctor/UpdateDoctorProfile",
			Name:   "UpdateDoctorProfile",
			Desc:   "更新医生个人信息",
		},
		{
			Method: "POST",
			Path:   "/api/v1/doctor/ChangePassword",
			Name:   "ChangePassword",
			Desc:   "修改密码",
		},
	}
}

// GetPatientRoutes 获取患者模块的所有路由信息
func GetPatientRoutes() []RouteInfo {
	return []RouteInfo{
		{
			Method: "POST",
			Path:   "/api/v1/patient/CreatePatient",
			Name:   "CreatePatient",
			Desc:   "患者信息录入",
		},
		{
			Method: "POST",
			Path:   "/api/v1/patient/GetPatientProfile",
			Name:   "GetPatientProfile",
			Desc:   "获取患者档案",
		},
		{
			Method: "POST",
			Path:   "/api/v1/patient/UpdatePatientProfile",
			Name:   "UpdatePatientProfile",
			Desc:   "更新患者档案",
		},
		{
			Method: "POST",
			Path:   "/api/v1/patient/GetPatientList",
			Name:   "GetPatientList",
			Desc:   "获取患者列表",
		},
		{
			Method: "POST",
			Path:   "/api/v1/patient/GetPatientsByCategory",
			Name:   "GetPatientsByCategory",
			Desc:   "按分类获取患者",
		},
		{
			Method: "POST",
			Path:   "/api/v1/patient/UpdatePatientCategory",
			Name:   "UpdatePatientCategory",
			Desc:   "更新患者分类",
		},
	}
}

// GetConsultationRoutes 获取问诊模块的所有路由信息
func GetConsultationRoutes() []RouteInfo {
	return []RouteInfo{
		{
			Method: "POST",
			Path:   "/api/v1/consultation/StartConsultation",
			Name:   "StartConsultation",
			Desc:   "开始问诊",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/GetConsultationDetail",
			Name:   "GetConsultationDetail",
			Desc:   "获取问诊详情",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/UpdateConsultationStatus",
			Name:   "UpdateConsultationStatus",
			Desc:   "更新问诊状态",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/EndConsultation",
			Name:   "EndConsultation",
			Desc:   "结束问诊",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/GetConsultationHistory",
			Name:   "GetConsultationHistory",
			Desc:   "获取问诊历史",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/SendMessage",
			Name:   "SendMessage",
			Desc:   "发送消息",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/GetMessages",
			Name:   "GetMessages",
			Desc:   "获取消息列表",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/MarkMessageRead",
			Name:   "MarkMessageRead",
			Desc:   "标记消息已读",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/AddConsultationRecord",
			Name:   "AddConsultationRecord",
			Desc:   "添加问诊记录",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/GetConsultationRecords",
			Name:   "GetConsultationRecords",
			Desc:   "获取问诊记录",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/GetConsultationReport",
			Name:   "GetConsultationReport",
			Desc:   "获取问诊报告",
		},
		{
			Method: "POST",
			Path:   "/api/v1/consultation/GetConsultationsByType",
			Name:   "GetConsultationsByType",
			Desc:   "按类型获取问诊",
		},
		{
			Method: "WS",
			Path:   "/ws/WebSocketConsultation",
			Name:   "WebSocketConsultation",
			Desc:   "问诊WebSocket连接",
		},
	}
}

// GetGreeterRoutes 获取问候模块的路由信息
func GetGreeterRoutes() []RouteInfo {
	return []RouteInfo{
		{
			Method: "GET",
			Path:   "/helloworld/{name}",
			Name:   "SayHello",
			Desc:   "问候接口",
		},
	}
}

// PrintRoutes 打印所有路由信息
func PrintRoutes(logger log.Logger) {
	helper := log.NewHelper(logger)

	helper.Info("===========================================")
	helper.Info("           API 接口列表")
	helper.Info("===========================================")

	// 打印医生模块接口
	helper.Info("📱 医生模块接口:")
	doctorRoutes := GetDoctorRoutes()
	for _, route := range doctorRoutes {
		helper.Infof("   %s %s - %s (%s)",
			route.Method,
			route.Path,
			route.Desc,
			route.Name)
	}

	helper.Info("")

	// 打印患者模块接口
	helper.Info("🏥 患者模块接口:")
	patientRoutes := GetPatientRoutes()
	for _, route := range patientRoutes {
		helper.Infof("   %s %s - %s (%s)",
			route.Method,
			route.Path,
			route.Desc,
			route.Name)
	}

	helper.Info("")

	// 打印问诊模块接口
	helper.Info("💬 问诊服务模块接口:")
	consultationRoutes := GetConsultationRoutes()
	for _, route := range consultationRoutes {
		helper.Infof("   %s %s - %s (%s)",
			route.Method,
			route.Path,
			route.Desc,
			route.Name)
	}

	helper.Info("")

	// 打印问候模块接口
	helper.Info("👋 问候模块接口:")
	greeterRoutes := GetGreeterRoutes()
	for _, route := range greeterRoutes {
		helper.Infof("   %s %s - %s (%s)",
			route.Method,
			route.Path,
			route.Desc,
			route.Name)
	}

	helper.Info("")
	helper.Info("===========================================")
	helper.Info("🌐 服务地址:")
	helper.Info("   HTTP: http://localhost:8000")
	helper.Info("   gRPC: localhost:9000")
	helper.Info("===========================================")
}
