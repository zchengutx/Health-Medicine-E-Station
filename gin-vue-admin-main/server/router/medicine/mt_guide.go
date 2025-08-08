package medicine

import (
	"github.com/gin-gonic/gin"
)

type MtGuideRouter struct{}

// InitMtGuideRouter 初始化 用药指导 路由信息
func (s *MtGuideRouter) InitMtGuideRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mtGuideRouterWithoutRecord := Router.Group("mtGuide")
	mtGuideRouterWithoutAuth := PublicRouter.Group("mtGuide")

	{
		mtGuideRouterWithoutRecord.GET("findMtGuide", mtGuideApi.FindMtGuide)       // 根据ID获取用药指导
		mtGuideRouterWithoutRecord.GET("getMtGuideList", mtGuideApi.GetMtGuideList) // 获取用药指导列表
	}
	{
		mtGuideRouterWithoutAuth.GET("getMtGuidePublic", mtGuideApi.GetMtGuideList) // 获取用药指导列表
	}
}
