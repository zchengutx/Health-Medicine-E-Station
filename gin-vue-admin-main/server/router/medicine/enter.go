package medicine

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	MtDrugTypeStairRouter
	MtDrugTypeLevelRouter
	MtDoctorsRouter
	MtHospitalsRouter
	MtDepartmentsRouter
	MtDoctorApprovalRouter
	MtOrdersDrugRouter
	MtUserRouter
	MtDrugRouter
	MtGuideRouter
	MtExplainRouter
	MtDoctorPatientsRouter
	MtOrdersRouter
	MtDiscountRouter
}

var (
	mtDrugTypeStairApi  = api.ApiGroupApp.MedicineApiGroup.MtDrugTypeStairApi
	mtDrugTypeLevelApi  = api.ApiGroupApp.MedicineApiGroup.MtDrugTypeLevelApi
	mtDoctorsApi        = api.ApiGroupApp.MedicineApiGroup.MtDoctorsApi
	mtHospitalsApi      = api.ApiGroupApp.MedicineApiGroup.MtHospitalsApi
	mtDepartmentsApi    = api.ApiGroupApp.MedicineApiGroup.MtDepartmentsApi
	mtDoctorApprovalApi = api.ApiGroupApp.MedicineApiGroup.MtDoctorApprovalApi
	mtOrdersDrugApi     = api.ApiGroupApp.MedicineApiGroup.MtOrdersDrugApi
	mtUserApi           = api.ApiGroupApp.MedicineApiGroup.MtUserApi
	mtDrugApi           = api.ApiGroupApp.MedicineApiGroup.MtDrugApi
	mtGuideApi          = api.ApiGroupApp.MedicineApiGroup.MtGuideApi
	mtExplainApi        = api.ApiGroupApp.MedicineApiGroup.MtExplainApi
	mtDoctorPatientsApi = api.ApiGroupApp.MedicineApiGroup.MtDoctorPatientsApi
	mtOrdersApi         = api.ApiGroupApp.MedicineApiGroup.MtOrdersApi
	mtDiscountApi       = api.ApiGroupApp.MedicineApiGroup.MtDiscountApi
)
