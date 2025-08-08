// 自动生成模板MtDoctors
package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mtDoctors表 结构体  MtDoctors
type MtDoctors struct {
	global.GVA_MODEL
	DoctorCode   *string `json:"doctorCode" form:"doctorCode" gorm:"comment:医生编码;column:doctor_code;size:32;" binding:"required"`                         //医生编码
	Name         *string `json:"name" form:"name" gorm:"comment:医生姓名;column:name;size:50;" binding:"required"`                                            //医生姓名
	Gender       *string `json:"gender" form:"gender" gorm:"comment:性别：1-男，2-女;column:gender;size:20;" binding:"required"`                                //性别：1-男，2-女
	DepartmentId *int    `json:"departmentId" form:"departmentId" gorm:"comment:科室ID;column:department_id;" binding:"required"`                           //科室ID
	HospitalId   *int    `json:"hospitalId" form:"hospitalId" gorm:"comment:医院ID;column:hospital_id;" binding:"required"`                                 //医院ID
	Title        *string `json:"title" form:"title" gorm:"comment:职称;column:title;size:50;" binding:"required"`                                           //职称
	Status       *string `json:"status" form:"status" gorm:"comment:审核状态：0-未通过，1-已通过，2-未审核;column:status;size:20;" binding:"required"`                    //审核状态：0-未通过，1-已通过，2-未审核
	ServiceAudit *string `json:"serviceAudit" form:"serviceAudit" gorm:"comment:服务审核：0-未审核，1-已审核，2-待审核;column:service_audit;size:20;" binding:"required"` //服务审核：0-未审核，1-已审核，2-待审核
	CreatedBy    uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint    `gorm:"column:deleted_by;comment:删除者"`

	Department string `gorm:"-" json:"department"` //gorm:"-" 表示不映射数据表字段
	Hospital   string `gorm:"-" json:"hospital"`
}

// TableName mtDoctors表 MtDoctors自定义表名 mt_doctors
func (MtDoctors) TableName() string {
	return "mt_doctors"
}
