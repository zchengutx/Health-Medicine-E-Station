package medicine

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	MtDrugTypeStairApi
	MtDrugTypeLevelApi
	MtDoctorsApi
	MtHospitalsApi
	MtDepartmentsApi
	MtDoctorApprovalApi
	MtOrdersDrugApi
	MtUserApi
	MtDrugApi
	MtGuideApi
	MtExplainApi
	MtDoctorPatientsApi
	MtOrdersApi
	MtChatMessageApi
	MtDiscountApi
}

var (
	mtDrugTypeStairService  = service.ServiceGroupApp.MedicineServiceGroup.MtDrugTypeStairService
	mtDrugTypeLevelService  = service.ServiceGroupApp.MedicineServiceGroup.MtDrugTypeLevelService
	mtDoctorsService        = service.ServiceGroupApp.MedicineServiceGroup.MtDoctorsService
	mtOrdersDrugService     = service.ServiceGroupApp.MedicineServiceGroup.MtOrdersDrugService
	mtUserService           = service.ServiceGroupApp.MedicineServiceGroup.MtUserService
	mtDrugService           = service.ServiceGroupApp.MedicineServiceGroup.MtDrugService
	mtGuideService          = service.ServiceGroupApp.MedicineServiceGroup.MtGuideService
	mtExplainService        = service.ServiceGroupApp.MedicineServiceGroup.MtExplainService
	mtDoctorPatientsService = service.ServiceGroupApp.MedicineServiceGroup.MtDoctorPatientsService
	mtOrdersService         = service.ServiceGroupApp.MedicineServiceGroup.MtOrdersService
	mtChatMessageService    = service.ServiceGroupApp.MedicineServiceGroup.MtChatMessageService
	mtDiscountService       = service.ServiceGroupApp.MedicineServiceGroup.MtDiscountService
)
