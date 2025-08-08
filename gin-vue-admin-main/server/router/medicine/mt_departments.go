package medicine

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtDepartmentsRouter struct{}

// InitMtDepartmentsRouter 初始化 mtDepartments表 路由信息
func (s *MtDepartmentsRouter) InitMtDepartmentsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtDepartmentsRouter := Router.Group("mtDepartments").Use(middleware.OperationRecord())
	mtDepartmentsRouterWithoutAuth := PublicRouter.Group("mtDepartments")
	var mtDepartmentsApi = v1.ApiGroupApp.MedicineApiGroup.MtDepartmentsApi
	{
		mtDepartmentsRouter.POST("createMtDepartments", mtDepartmentsApi.CreateMtDepartments)             // 新建mtDepartments表
		mtDepartmentsRouter.DELETE("deleteMtDepartments", mtDepartmentsApi.DeleteMtDepartments)           // 删除mtDepartments表
		mtDepartmentsRouter.DELETE("deleteMtDepartmentsByIds", mtDepartmentsApi.DeleteMtDepartmentsByIds) // 批量删除mtDepartments表
		mtDepartmentsRouter.PUT("updateMtDepartments", mtDepartmentsApi.UpdateMtDepartments)              // 更新mtDepartments表
		mtDepartmentsRouter.GET("findMtDepartments", mtDepartmentsApi.FindMtDepartments)                  // 根据ID获取mtDepartments表
		mtDepartmentsRouter.GET("getMtDepartmentsList", mtDepartmentsApi.GetMtDepartmentsList)            // 获取mtDepartments表列表
	}
	{
		mtDepartmentsRouterWithoutAuth.GET("getMtDepartmentsListPublic", mtDepartmentsApi.GetMtDepartmentsListPublic) // 获取科室列表（公开接口）
	}
}
