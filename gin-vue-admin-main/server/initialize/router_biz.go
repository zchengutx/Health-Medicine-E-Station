package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		medicineRouter := router.RouterGroupApp.Medicine
		medicineRouter.InitMtDrugRouter(privateGroup, publicGroup)
		medicineRouter.InitMtDrugTypeStairRouter(privateGroup, publicGroup)
		medicineRouter.InitMtDrugTypeLevelRouter(privateGroup, publicGroup)
		medicineRouter.InitMtGuideRouter(privateGroup, publicGroup)
		medicineRouter.InitMtExplainRouter(privateGroup, publicGroup)
		medicineRouter.InitMtDoctorPatientsRouter(privateGroup, publicGroup)
		medicineRouter.InitMtDoctorsRouter(privateGroup, publicGroup)
		medicineRouter.InitMtOrdersDrugRouter(privateGroup, publicGroup)
		medicineRouter.InitMtUserRouter(privateGroup, publicGroup)
		medicineRouter.InitMtOrdersRouter(privateGroup, publicGroup)
		medicineRouter.InitMtChatMessageRouter(privateGroup, publicGroup)
		medicineRouter.InitMtDiscountRouter(privateGroup, publicGroup)
	}
}
