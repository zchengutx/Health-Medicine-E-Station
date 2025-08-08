// 自动生成模板MtDoctorPatients
package medicine

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mtDoctorPatients表 结构体  MtDoctorPatients
type MtDoctorPatients struct {
	global.GVA_MODEL
	DoctorId         *int       `json:"doctorId" form:"doctorId" gorm:"comment:医生ID;column:doctor_id;size:20;" binding:"required"`                 //医生ID
	PatientId        *int       `json:"patientId" form:"patientId" gorm:"comment:患者ID;column:patient_id;size:20;" binding:"required"`              //患者ID
	RelationshipType *string    `json:"relationshipType" form:"relationshipType" gorm:"comment:关系类型: 普通/关注/VIP;column:relationship_type;size:20;"` //关系类型: 普通/关注/VIP
	Tags             *string    `json:"tags" form:"tags" gorm:"comment:患者标签，逗号分隔;column:tags;size:200;"`                                           //患者标签，逗号分隔
	Notes            *string    `json:"notes" form:"notes" gorm:"comment:备注;column:notes;"`                                                        //备注
	FirstVisitTime   *time.Time `json:"firstVisitTime" form:"firstVisitTime" gorm:"comment:首次就诊时间;column:first_visit_time;" binding:"required"`    //首次就诊时间
	LastVisitTime    *time.Time `json:"lastVisitTime" form:"lastVisitTime" gorm:"comment:最后就诊时间;column:last_visit_time;"`                          //最后就诊时间
	VisitCount       *int       `json:"visitCount" form:"visitCount" gorm:"comment:就诊次数;column:visit_count;size:10;"`                              //就诊次数
	CreatedBy        uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName mtDoctorPatients表 MtDoctorPatients自定义表名 mt_doctor_patients
func (MtDoctorPatients) TableName() string {
	return "mt_doctor_patients"
}
