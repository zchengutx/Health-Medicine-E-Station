package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MtDoctorsRouter struct {}

// InitMtDoctorsRouter 初始化 mtDoctors表 路由信息
func (s *MtDoctorsRouter) InitMtDoctorsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mtDoctorsRouter := Router.Group("mtDoctors").Use(middleware.OperationRecord())
	mtDoctorsRouterWithoutRecord := Router.Group("mtDoctors")
	mtDoctorsRouterWithoutAuth := PublicRouter.Group("mtDoctors")
	{
		mtDoctorsRouter.POST("createMtDoctors", mtDoctorsApi.CreateMtDoctors)   // 新建mtDoctors表
		mtDoctorsRouter.DELETE("deleteMtDoctors", mtDoctorsApi.DeleteMtDoctors) // 删除mtDoctors表
		mtDoctorsRouter.DELETE("deleteMtDoctorsByIds", mtDoctorsApi.DeleteMtDoctorsByIds) // 批量删除mtDoctors表
		mtDoctorsRouter.PUT("updateMtDoctors", mtDoctorsApi.UpdateMtDoctors)    // 更新mtDoctors表
	}
	{
		mtDoctorsRouterWithoutRecord.GET("findMtDoctors", mtDoctorsApi.FindMtDoctors)        // 根据ID获取mtDoctors表
		mtDoctorsRouterWithoutRecord.GET("getMtDoctorsList", mtDoctorsApi.GetMtDoctorsList)  // 获取mtDoctors表列表
	}
	{
	    mtDoctorsRouterWithoutAuth.GET("getMtDoctorsPublic", mtDoctorsApi.GetMtDoctorsPublic)  // mtDoctors表开放接口
	}
}
