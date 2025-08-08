package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtDrugRouter struct {}

// InitMtDrugRouter 初始化 mtDrug表 路由信息
func (s *MtDrugRouter) InitMtDrugRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mtDrugRouter := Router.Group("mtDrug").Use(middleware.OperationRecord())
	mtDrugRouterWithoutRecord := Router.Group("mtDrug")
	mtDrugRouterWithoutAuth := PublicRouter.Group("mtDrug")
	{
		mtDrugRouter.POST("createMtDrug", mtDrugApi.CreateMtDrug)   // 新建mtDrug表
		mtDrugRouter.DELETE("deleteMtDrug", mtDrugApi.DeleteMtDrug) // 删除mtDrug表
		mtDrugRouter.DELETE("deleteMtDrugByIds", mtDrugApi.DeleteMtDrugByIds) // 批量删除mtDrug表
		mtDrugRouter.PUT("updateMtDrug", mtDrugApi.UpdateMtDrug)    // 更新mtDrug表
	}
	{
		mtDrugRouterWithoutRecord.GET("findMtDrug", mtDrugApi.FindMtDrug)        // 根据ID获取mtDrug表
		mtDrugRouterWithoutRecord.GET("getMtDrugList", mtDrugApi.GetMtDrugList)  // 获取mtDrug表列表
	}
	{
	    mtDrugRouterWithoutAuth.GET("getMtDrugPublic", mtDrugApi.GetMtDrugPublic)  // mtDrug表开放接口
	}
}
