package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtDiscountRouter struct {}

// InitMtDiscountRouter 初始化 mtDiscount表 路由信息
func (s *MtDiscountRouter) InitMtDiscountRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mtDiscountRouter := Router.Group("mtDiscount").Use(middleware.OperationRecord())
	mtDiscountRouterWithoutRecord := Router.Group("mtDiscount")
	mtDiscountRouterWithoutAuth := PublicRouter.Group("mtDiscount")
	{
		mtDiscountRouter.POST("createMtDiscount", mtDiscountApi.CreateMtDiscount)   // 新建mtDiscount表
		mtDiscountRouter.DELETE("deleteMtDiscount", mtDiscountApi.DeleteMtDiscount) // 删除mtDiscount表
		mtDiscountRouter.DELETE("deleteMtDiscountByIds", mtDiscountApi.DeleteMtDiscountByIds) // 批量删除mtDiscount表
		mtDiscountRouter.PUT("updateMtDiscount", mtDiscountApi.UpdateMtDiscount)    // 更新mtDiscount表
	}
	{
		mtDiscountRouterWithoutRecord.GET("findMtDiscount", mtDiscountApi.FindMtDiscount)        // 根据ID获取mtDiscount表
		mtDiscountRouterWithoutRecord.GET("getMtDiscountList", mtDiscountApi.GetMtDiscountList)  // 获取mtDiscount表列表
	}
	{
	    mtDiscountRouterWithoutAuth.GET("getMtDiscountPublic", mtDiscountApi.GetMtDiscountPublic)  // mtDiscount表开放接口
	}
}
