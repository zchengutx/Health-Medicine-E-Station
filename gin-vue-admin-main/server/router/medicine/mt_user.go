package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtUserRouter struct {}

// InitMtUserRouter 初始化 mtUser表 路由信息
func (s *MtUserRouter) InitMtUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mtUserRouter := Router.Group("mtUser").Use(middleware.OperationRecord())
	mtUserRouterWithoutRecord := Router.Group("mtUser")
	mtUserRouterWithoutAuth := PublicRouter.Group("mtUser")
	{
		mtUserRouter.POST("createMtUser", mtUserApi.CreateMtUser)   // 新建mtUser表
		mtUserRouter.DELETE("deleteMtUser", mtUserApi.DeleteMtUser) // 删除mtUser表
		mtUserRouter.DELETE("deleteMtUserByIds", mtUserApi.DeleteMtUserByIds) // 批量删除mtUser表
		mtUserRouter.PUT("updateMtUser", mtUserApi.UpdateMtUser)    // 更新mtUser表
	}
	{
		mtUserRouterWithoutRecord.GET("findMtUser", mtUserApi.FindMtUser)        // 根据ID获取mtUser表
		mtUserRouterWithoutRecord.GET("getMtUserList", mtUserApi.GetMtUserList)  // 获取mtUser表列表
	}
	{
	    mtUserRouterWithoutAuth.GET("getMtUserPublic", mtUserApi.GetMtUserPublic)  // mtUser表开放接口
	}
}
