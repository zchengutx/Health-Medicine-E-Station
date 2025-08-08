package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtDoctorPatientsRouter struct{}

// InitMtDoctorPatientsRouter 初始化 mtDoctorPatients表 路由信息
func (s *MtDoctorPatientsRouter) InitMtDoctorPatientsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtDoctorPatientsRouter := Router.Group("mtDoctorPatients").Use(middleware.OperationRecord())
	mtDoctorPatientsRouterWithoutRecord := Router.Group("mtDoctorPatients")
	mtDoctorPatientsRouterWithoutAuth := PublicRouter.Group("mtDoctorPatients")
	{
		mtDoctorPatientsRouter.POST("createMtDoctorPatients", mtDoctorPatientsApi.CreateMtDoctorPatients)             // 新建mtDoctorPatients表
		mtDoctorPatientsRouter.DELETE("deleteMtDoctorPatients", mtDoctorPatientsApi.DeleteMtDoctorPatients)           // 删除mtDoctorPatients表
		mtDoctorPatientsRouter.DELETE("deleteMtDoctorPatientsByIds", mtDoctorPatientsApi.DeleteMtDoctorPatientsByIds) // 批量删除mtDoctorPatients表
		mtDoctorPatientsRouter.PUT("updateMtDoctorPatients", mtDoctorPatientsApi.UpdateMtDoctorPatients)              // 更新mtDoctorPatients表
	}
	{
		mtDoctorPatientsRouterWithoutRecord.GET("findMtDoctorPatients", mtDoctorPatientsApi.FindMtDoctorPatients)       // 根据ID获取mtDoctorPatients表
		mtDoctorPatientsRouterWithoutRecord.GET("getMtDoctorPatientsList", mtDoctorPatientsApi.GetMtDoctorPatientsList) // 获取mtDoctorPatients表列表
	}
	{
		mtDoctorPatientsRouterWithoutAuth.GET("getMtDoctorPatientsPublic", mtDoctorPatientsApi.GetMtDoctorPatientsPublic) // mtDoctorPatients表开放接口
	}
}
