package medicine

import (
	"github.com/gin-gonic/gin"
)

type MtExplainRouter struct{}

// InitMtExplainRouter 初始化 说明书 路由信息
func (s *MtExplainRouter) InitMtExplainRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtExplainRouterWithoutRecord := Router.Group("mtExplain")
	mtExplainRouterWithoutAuth := PublicRouter.Group("mtExplain")

	{
		mtExplainRouterWithoutRecord.GET("findMtExplain", mtExplainApi.FindMtExplain)       // 根据ID获取说明书
		mtExplainRouterWithoutRecord.GET("getMtExplainList", mtExplainApi.GetMtExplainList) // 获取说明书列表
	}
	{
		mtExplainRouterWithoutAuth.GET("getMtExplainPublic", mtExplainApi.GetMtExplainList) // 获取说明书列表
	}
}
