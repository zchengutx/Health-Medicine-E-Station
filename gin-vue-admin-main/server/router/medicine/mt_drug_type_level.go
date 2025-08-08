package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtDrugTypeLevelRouter struct{}

// InitMtDrugTypeLevelRouter 初始化 mtDrugTypeLevel表 路由信息
func (s *MtDrugTypeLevelRouter) InitMtDrugTypeLevelRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtDrugTypeLevelRouter := Router.Group("mtDrugTypeLevel").Use(middleware.OperationRecord())
	mtDrugTypeLevelRouterWithoutRecord := Router.Group("mtDrugTypeLevel")
	mtDrugTypeLevelRouterWithoutAuth := PublicRouter.Group("mtDrugTypeLevel")
	{
		mtDrugTypeLevelRouter.POST("createMtDrugTypeLevel", mtDrugTypeLevelApi.CreateMtDrugTypeLevel)             // 新建mtDrugTypeLevel表
		mtDrugTypeLevelRouter.DELETE("deleteMtDrugTypeLevel", mtDrugTypeLevelApi.DeleteMtDrugTypeLevel)           // 删除mtDrugTypeLevel表
		mtDrugTypeLevelRouter.DELETE("deleteMtDrugTypeLevelByIds", mtDrugTypeLevelApi.DeleteMtDrugTypeLevelByIds) // 批量删除mtDrugTypeLevel表
		mtDrugTypeLevelRouter.PUT("updateMtDrugTypeLevel", mtDrugTypeLevelApi.UpdateMtDrugTypeLevel)              // 更新mtDrugTypeLevel表
	}
	{
		mtDrugTypeLevelRouterWithoutRecord.GET("findMtDrugTypeLevel", mtDrugTypeLevelApi.FindMtDrugTypeLevel)       // 根据ID获取mtDrugTypeLevel表
		mtDrugTypeLevelRouterWithoutRecord.GET("getMtDrugTypeLevelList", mtDrugTypeLevelApi.GetMtDrugTypeLevelList) // 获取mtDrugTypeLevel表列表
		mtDrugTypeLevelRouterWithoutRecord.GET("getAllMtDrugTypeLevel", mtDrugTypeLevelApi.GetAllMtDrugTypeLevel)   // 获取所有二级分类
	}
	{
		mtDrugTypeLevelRouterWithoutAuth.GET("getMtDrugTypeLevelPublic", mtDrugTypeLevelApi.GetAllMtDrugTypeLevel) // mtDrugTypeLevel表开放接口
	}
}
