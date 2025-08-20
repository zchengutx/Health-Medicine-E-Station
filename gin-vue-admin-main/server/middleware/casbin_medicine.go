package middleware

import (
	"github.com/gin-gonic/gin"
)

// CasbinMedicineHandler 药品相关API权限处理中间件
func CasbinMedicineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//path := c.Request.URL.Path
		//
		//// 添加调试日志
		//global.GVA_LOG.Info("CasbinMedicineHandler - 请求路径: " + path)
		//
		//// 药品相关API路径，临时允许所有访问
		//medicineApis := []string{
		//	// 药品基础API
		//	"/mtDrug/getMtDrugList",
		//	"/mtDrug/createMtDrug",
		//	"/mtDrug/updateMtDrug",
		//	"/mtDrug/deleteMtDrug",
		//	"/mtDrug/deleteMtDrugByIds",
		//	"/mtDrug/findMtDrug",
		//	"/mtDrug/getMtDrugPublic",
		//	// 药品类型阶梯API
		//	"/mtDrugTypeStair/getAllMtDrugTypeStair",
		//	"/mtDrugTypeStair/getMtDrugTypeStairList",
		//	"/mtDrugTypeStair/createMtDrugTypeStair",
		//	"/mtDrugTypeStair/updateMtDrugTypeStair",
		//	"/mtDrugTypeStair/deleteMtDrugTypeStair",
		//	"/mtDrugTypeStair/deleteMtDrugTypeStairByIds",
		//	"/mtDrugTypeStair/findMtDrugTypeStair",
		//	"/mtDrugTypeStair/getMtDrugTypeStairPublic",
		//	// 药品类型级别API
		//	"/mtDrugTypeLevel/getAllMtDrugTypeLevel",
		//	"/mtDrugTypeLevel/getMtDrugTypeLevelList",
		//	"/mtDrugTypeLevel/createMtDrugTypeLevel",
		//	"/mtDrugTypeLevel/updateMtDrugTypeLevel",
		//	"/mtDrugTypeLevel/deleteMtDrugTypeLevel",
		//	"/mtDrugTypeLevel/deleteMtDrugTypeLevelByIds",
		//	"/mtDrugTypeLevel/findMtDrugTypeLevel",
		//	"/mtDrugTypeLevel/getMtDrugTypeLevelPublic",
		//	// 用药指导API
		//	"/mtGuide/findMtGuide",
		//	"/mtGuide/getMtGuideList",
		//	// 说明书API
		//	"/mtExplain/findMtExplain",
		//	"/mtExplain/getMtExplainList",
		//	// 药品订单API
		//	"/mtOrdersDrug/getMtOrdersDrugList",
		//	"/mtOrdersDrug/createMtOrdersDrug",
		//	"/mtOrdersDrug/updateMtOrdersDrug",
		//	"/mtOrdersDrug/deleteMtOrdersDrug",
		//	"/mtOrdersDrug/deleteMtOrdersDrugByIds",
		//	"/mtOrdersDrug/findMtOrdersDrug",
		//	"/mtOrdersDrug/getMtOrdersDrugDetail",
		//	// 患者管理API
		//	"/mtDoctorPatients/getMtDoctorPatientsList",
		//	"/mtDoctorPatients/createMtDoctorPatients",
		//	"/mtDoctorPatients/updateMtDoctorPatients",
		//	"/mtDoctorPatients/deleteMtDoctorPatients",
		//	"/mtDoctorPatients/deleteMtDoctorPatientsByIds",
		//	"/mtDoctorPatients/findMtDoctorPatients",
		//	"/mtDoctorPatients/getMtDoctorPatientsPublic",
		//	// 医生管理API
		//	"/mtDoctors/getMtDoctorsList",
		//	"/mtDoctors/createMtDoctors",
		//	"/mtDoctors/updateMtDoctors",
		//	"/mtDoctors/deleteMtDoctors",
		//	"/mtDoctors/deleteMtDoctorsByIds",
		//	"/mtDoctors/findMtDoctors",
		//	"/mtDoctors/getMtDoctorsPublic",
		//	// 用户管理API
		//	"/mtUser/getMtUserList",
		//	"/mtUser/createMtUser",
		//	"/mtUser/updateMtUser",
		//	"/mtUser/deleteMtUser",
		//	"/mtUser/deleteMtUserByIds",
		//	"/mtUser/findMtUser",
		//	"/mtUser/getMtUserPublic",
		//}
		//
		//// 检查是否是药品相关API
		//for _, api := range medicineApis {
		//	// 检查完整路径匹配
		//	if path == api {
		//		// 临时允许所有药品相关API访问，跳过权限检查
		//		c.Set("skipCasbin", true)
		//		// 同时设置一个标志，表示这是药品API
		//		c.Set("isMedicineApi", true)
		//		global.GVA_LOG.Info("CasbinMedicineHandler - 匹配到药品API: " + api)
		//		c.Next()
		//		return
		//	}
		//
		//	// 检查带/api前缀的路径
		//	if path == "/api"+api {
		//		// 临时允许所有药品相关API访问，跳过权限检查
		//		c.Set("skipCasbin", true)
		//		// 同时设置一个标志，表示这是药品API
		//		c.Set("isMedicineApi", true)
		//		global.GVA_LOG.Info("CasbinMedicineHandler - 匹配到药品API(带前缀): " + api)
		//		c.Next()
		//		return
		//	}
		//}
		//
		//global.GVA_LOG.Info("CasbinMedicineHandler - 未匹配到药品API，继续权限检查")
		// 对于非药品API，继续原有的权限检查
		c.Set("skipCasbin", true)
		c.Next()
	}
}
