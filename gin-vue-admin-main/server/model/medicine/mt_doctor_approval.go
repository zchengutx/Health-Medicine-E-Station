// 自动生成模板MtDoctorApproval
package medicine

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mtDoctorApproval表 结构体  MtDoctorApproval
type MtDoctorApproval struct {
	global.GVA_MODEL
	DoctorId       *uint      `json:"doctorId" form:"doctorId" gorm:"comment:医生ID;column:doctor_id;" binding:"required"`                                            //医生ID
	DoctorName     *string    `json:"doctorName" form:"doctorName" gorm:"comment:医生姓名;column:doctor_name;size:50;"`                                                 //医生姓名
	ApprovalStatus *string    `json:"approvalStatus" form:"approvalStatus" gorm:"comment:审核状态：0-未通过，1-通过，2-待审核;column:approval_status;size:20;" binding:"required"` //审核状态
	ApprovalTime   *time.Time `json:"approvalTime" form:"approvalTime" gorm:"comment:审核时间;column:approval_time;"`                                                   //审核时间
	ApproverId     *uint      `json:"approverId" form:"approverId" gorm:"comment:审核人ID;column:approver_id;"`                                                        //审核人ID
	ApproverName   *string    `json:"approverName" form:"approverName" gorm:"comment:审核人姓名;column:approver_name;size:50;"`                                          //审核人姓名
	ApprovalReason *string    `json:"approvalReason" form:"approvalReason" gorm:"comment:审核理由;column:approval_reason;size:500;"`                                    //审核理由
	RejectReason   *string    `json:"rejectReason" form:"rejectReason" gorm:"comment:拒绝理由;column:reject_reason;size:500;"`                                          //拒绝理由
	SubmitTime     *time.Time `json:"submitTime" form:"submitTime" gorm:"comment:提交审核时间;column:submit_time;"`                                                       //提交审核时间
	CreatedBy      uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName mtDoctorApproval表 MtDoctorApproval自定义表名 mt_doctor_approval
func (MtDoctorApproval) TableName() string {
	return "mt_doctor_approval"
}
