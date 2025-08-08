package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtDrugTypeStairRouter struct{}

// InitMtDrugTypeStairRouter 初始化 mtDrugTypeStair表 路由信息
func (s *MtDrugTypeStairRouter) InitMtDrugTypeStairRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtDrugTypeStairRouter := Router.Group("mtDrugTypeStair").Use(middleware.OperationRecord())
	mtDrugTypeStairRouterWithoutRecord := Router.Group("mtDrugTypeStair")
	mtDrugTypeStairRouterWithoutAuth := PublicRouter.Group("mtDrugTypeStair")
	{
		mtDrugTypeStairRouter.POST("createMtDrugTypeStair", mtDrugTypeStairApi.CreateMtDrugTypeStair)             // 新建mtDrugTypeStair表
		mtDrugTypeStairRouter.DELETE("deleteMtDrugTypeStair", mtDrugTypeStairApi.DeleteMtDrugTypeStair)           // 删除mtDrugTypeStair表
		mtDrugTypeStairRouter.DELETE("deleteMtDrugTypeStairByIds", mtDrugTypeStairApi.DeleteMtDrugTypeStairByIds) // 批量删除mtDrugTypeStair表
		mtDrugTypeStairRouter.PUT("updateMtDrugTypeStair", mtDrugTypeStairApi.UpdateMtDrugTypeStair)              // 更新mtDrugTypeStair表
	}
	{
		mtDrugTypeStairRouterWithoutRecord.GET("findMtDrugTypeStair", mtDrugTypeStairApi.FindMtDrugTypeStair)       // 根据ID获取mtDrugTypeStair表
		mtDrugTypeStairRouterWithoutRecord.GET("getMtDrugTypeStairList", mtDrugTypeStairApi.GetMtDrugTypeStairList) // 获取mtDrugTypeStair表列表
		mtDrugTypeStairRouterWithoutRecord.GET("getAllMtDrugTypeStair", mtDrugTypeStairApi.GetAllMtDrugTypeStair)   // 获取所有一级分类
	}
	{
		mtDrugTypeStairRouterWithoutAuth.GET("getMtDrugTypeStairPublic", mtDrugTypeStairApi.GetAllMtDrugTypeStair) // mtDrugTypeStair表开放接口
	}
}
