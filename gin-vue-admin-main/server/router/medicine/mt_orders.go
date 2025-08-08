package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtOrdersRouter struct {}

// InitMtOrdersRouter 初始化 mtOrders表 路由信息
func (s *MtOrdersRouter) InitMtOrdersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mtOrdersRouter := Router.Group("mtOrders").Use(middleware.OperationRecord())
	mtOrdersRouterWithoutRecord := Router.Group("mtOrders")
	mtOrdersRouterWithoutAuth := PublicRouter.Group("mtOrders")
	{
		mtOrdersRouter.POST("createMtOrders", mtOrdersApi.CreateMtOrders)   // 新建mtOrders表
		mtOrdersRouter.DELETE("deleteMtOrders", mtOrdersApi.DeleteMtOrders) // 删除mtOrders表
		mtOrdersRouter.DELETE("deleteMtOrdersByIds", mtOrdersApi.DeleteMtOrdersByIds) // 批量删除mtOrders表
		mtOrdersRouter.PUT("updateMtOrders", mtOrdersApi.UpdateMtOrders)    // 更新mtOrders表
	}
	{
		mtOrdersRouterWithoutRecord.GET("findMtOrders", mtOrdersApi.FindMtOrders)        // 根据ID获取mtOrders表
		mtOrdersRouterWithoutRecord.GET("getMtOrdersList", mtOrdersApi.GetMtOrdersList)  // 获取mtOrders表列表
	}
	{
	    mtOrdersRouterWithoutAuth.GET("getMtOrdersPublic", mtOrdersApi.GetMtOrdersPublic)  // mtOrders表开放接口
	}
}
