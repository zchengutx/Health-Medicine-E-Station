package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MtDoctorApprovalSearch struct {
	request.PageInfo
	DoctorId       *uint        `json:"doctorId" form:"doctorId"`
	DoctorName     *string      `json:"doctorName" form:"doctorName"`
	ApprovalStatus *string      `json:"approvalStatus" form:"approvalStatus"`
	ApproverId     *uint        `json:"approverId" form:"approverId"`
	ApproverName   *string      `json:"approverName" form:"approverName"`
	CreatedAtRange []*time.Time `json:"createdAtRange" form:"createdAtRange"`
}
