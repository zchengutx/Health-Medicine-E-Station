package medicine

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtHospitalsRouter struct{}

// InitMtHospitalsRouter 初始化 mtHospitals表 路由信息
func (s *MtHospitalsRouter) InitMtHospitalsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtHospitalsRouter := Router.Group("mtHospitals").Use(middleware.OperationRecord())
	mtHospitalsRouterWithoutAuth := PublicRouter.Group("mtHospitals")
	var mtHospitalsApi = v1.ApiGroupApp.MedicineApiGroup.MtHospitalsApi
	{
		mtHospitalsRouter.POST("createMtHospitals", mtHospitalsApi.CreateMtHospitals)             // 新建mtHospitals表
		mtHospitalsRouter.DELETE("deleteMtHospitals", mtHospitalsApi.DeleteMtHospitals)           // 删除mtHospitals表
		mtHospitalsRouter.DELETE("deleteMtHospitalsByIds", mtHospitalsApi.DeleteMtHospitalsByIds) // 批量删除mtHospitals表
		mtHospitalsRouter.PUT("updateMtHospitals", mtHospitalsApi.UpdateMtHospitals)              // 更新mtHospitals表
		mtHospitalsRouter.GET("findMtHospitals", mtHospitalsApi.FindMtHospitals)                  // 根据ID获取mtHospitals表
		mtHospitalsRouter.GET("getMtHospitalsList", mtHospitalsApi.GetMtHospitalsList)            // 获取mtHospitals表列表
	}
	{
		mtHospitalsRouterWithoutAuth.GET("getMtHospitalsListPublic", mtHospitalsApi.GetMtHospitalsListPublic) // 获取医院列表（公开接口）
	}
}
