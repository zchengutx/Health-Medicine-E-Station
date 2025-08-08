package medicine

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtDoctorApprovalRouter struct{}

// InitMtDoctorApprovalRouter 初始化 mtDoctorApproval表 路由信息
func (s *MtDoctorApprovalRouter) InitMtDoctorApprovalRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtDoctorApprovalRouter := Router.Group("mtDoctorApproval").Use(middleware.OperationRecord())
	mtDoctorApprovalRouterWithoutAuth := PublicRouter.Group("mtDoctorApproval")

	var mtDoctorApprovalApi = v1.ApiGroupApp.MedicineApiGroup.MtDoctorApprovalApi
	{
		mtDoctorApprovalRouter.POST("createMtDoctorApproval", mtDoctorApprovalApi.CreateMtDoctorApproval)              // 新建mtDoctorApproval表
		mtDoctorApprovalRouter.DELETE("deleteMtDoctorApproval", mtDoctorApprovalApi.DeleteMtDoctorApproval)            // 删除mtDoctorApproval表
		mtDoctorApprovalRouter.DELETE("deleteMtDoctorApprovalByIds", mtDoctorApprovalApi.DeleteMtDoctorApprovalByIds)  // 批量删除mtDoctorApproval表
		mtDoctorApprovalRouter.PUT("updateMtDoctorApproval", mtDoctorApprovalApi.UpdateMtDoctorApproval)               // 更新mtDoctorApproval表
		mtDoctorApprovalRouter.GET("findMtDoctorApproval", mtDoctorApprovalApi.FindMtDoctorApproval)                   // 根据ID获取mtDoctorApproval表
		mtDoctorApprovalRouter.GET("getMtDoctorApprovalList", mtDoctorApprovalApi.GetMtDoctorApprovalList)             // 获取mtDoctorApproval表列表
		mtDoctorApprovalRouter.GET("getMtDoctorApprovalByDoctorId", mtDoctorApprovalApi.GetMtDoctorApprovalByDoctorId) // 根据医生ID获取审核记录
	}
	{
		mtDoctorApprovalRouterWithoutAuth.GET("getMtDoctorApprovalPublic", mtDoctorApprovalApi.GetMtDoctorApprovalPublic) // 获取mtDoctorApproval表列表
	}
}
