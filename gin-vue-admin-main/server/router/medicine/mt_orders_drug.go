package medicine

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtOrdersDrugRouter struct{}

// InitMtOrdersDrugRouter 初始化 mtOrdersDrug表 路由信息
func (s *MtOrdersDrugRouter) InitMtOrdersDrugRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtOrdersDrugRouter := Router.Group("mtOrdersDrug").Use(middleware.OperationRecord())
	mtOrdersDrugRouterWithoutAuth := PublicRouter.Group("mtOrdersDrug")
	var mtOrdersDrugApi = v1.ApiGroupApp.MedicineApiGroup.MtOrdersDrugApi
	{
		mtOrdersDrugRouter.POST("createMtOrdersDrug", mtOrdersDrugApi.CreateMtOrdersDrug)             // 新建mtOrdersDrug表
		mtOrdersDrugRouter.DELETE("deleteMtOrdersDrug", mtOrdersDrugApi.DeleteMtOrdersDrug)           // 删除mtOrdersDrug表
		mtOrdersDrugRouter.DELETE("deleteMtOrdersDrugByIds", mtOrdersDrugApi.DeleteMtOrdersDrugByIds) // 批量删除mtOrdersDrug表
		mtOrdersDrugRouter.PUT("updateMtOrdersDrug", mtOrdersDrugApi.UpdateMtOrdersDrug)              // 更新mtOrdersDrug表
		mtOrdersDrugRouter.GET("findMtOrdersDrug", mtOrdersDrugApi.FindMtOrdersDrug)                  // 根据ID获取mtOrdersDrug表
		mtOrdersDrugRouter.GET("getMtOrdersDrugList", mtOrdersDrugApi.GetMtOrdersDrugList)            // 获取mtOrdersDrug表列表
		mtOrdersDrugRouter.GET("testMtOrdersDrugDetail", mtOrdersDrugApi.TestMtOrdersDrugDetail)      // 测试接口
	}
	{
		mtOrdersDrugRouterWithoutAuth.GET("getMtOrdersDrugPublic", mtOrdersDrugApi.GetMtOrdersDrugPublic) // mtOrdersDrug表开放接口
		mtOrdersDrugRouterWithoutAuth.GET("getMtOrdersDrugDetail", mtOrdersDrugApi.GetMtOrdersDrugDetail) // 获取订单详情信息（公开接口）
	}
}
