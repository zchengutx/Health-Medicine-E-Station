package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtChatMessageRouter struct {}

// InitMtChatMessageRouter 初始化 mtChatMessage表 路由信息
func (s *MtChatMessageRouter) InitMtChatMessageRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mtChatMessageRouter := Router.Group("mtChatMessage").Use(middleware.OperationRecord())
	mtChatMessageRouterWithoutRecord := Router.Group("mtChatMessage")
	mtChatMessageRouterWithoutAuth := PublicRouter.Group("mtChatMessage")
	{
		mtChatMessageRouter.POST("createMtChatMessage", mtChatMessageApi.CreateMtChatMessage)   // 新建mtChatMessage表
		mtChatMessageRouter.DELETE("deleteMtChatMessage", mtChatMessageApi.DeleteMtChatMessage) // 删除mtChatMessage表
		mtChatMessageRouter.DELETE("deleteMtChatMessageByIds", mtChatMessageApi.DeleteMtChatMessageByIds) // 批量删除mtChatMessage表
		mtChatMessageRouter.PUT("updateMtChatMessage", mtChatMessageApi.UpdateMtChatMessage)    // 更新mtChatMessage表
	}
	{
		mtChatMessageRouterWithoutRecord.GET("findMtChatMessage", mtChatMessageApi.FindMtChatMessage)        // 根据ID获取mtChatMessage表
		mtChatMessageRouterWithoutRecord.GET("getMtChatMessageList", mtChatMessageApi.GetMtChatMessageList)  // 获取mtChatMessage表列表
	}
	{
	    mtChatMessageRouterWithoutAuth.GET("getMtChatMessagePublic", mtChatMessageApi.GetMtChatMessagePublic)  // mtChatMessage表开放接口
	}
}
