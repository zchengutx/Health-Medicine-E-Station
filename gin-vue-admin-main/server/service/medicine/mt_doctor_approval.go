package medicine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
)

type MtDoctorApprovalService struct{}

// CreateMtDoctorApproval 创建mtDoctorApproval表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorApprovalService *MtDoctorApprovalService) CreateMtDoctorApproval(ctx context.Context, mtDoctorApproval medicine.MtDoctorApproval) (err error) {
	err = global.GVA_DB.Create(&mtDoctorApproval).Error
	return err
}

// DeleteMtDoctorApproval 删除mtDoctorApproval表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorApprovalService *MtDoctorApprovalService) DeleteMtDoctorApproval(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&medicine.MtDoctorApproval{}, "id = ?", ID).Error
	return err
}

// DeleteMtDoctorApprovalByIds 批量删除mtDoctorApproval表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorApprovalService *MtDoctorApprovalService) DeleteMtDoctorApprovalByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]medicine.MtDoctorApproval{}, "id in ?", IDs).Error
	return err
}

// UpdateMtDoctorApproval 更新mtDoctorApproval表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorApprovalService *MtDoctorApprovalService) UpdateMtDoctorApproval(ctx context.Context, mtDoctorApproval medicine.MtDoctorApproval) (err error) {
	err = global.GVA_DB.Model(&medicine.MtDoctorApproval{}).Where("id = ?", mtDoctorApproval.ID).Updates(&mtDoctorApproval).Error
	return err
}

// GetMtDoctorApproval 根据ID获取mtDoctorApproval表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorApprovalService *MtDoctorApprovalService) GetMtDoctorApproval(ctx context.Context, ID string) (mtDoctorApproval medicine.MtDoctorApproval, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtDoctorApproval).Error
	return
}

// GetMtDoctorApprovalInfoList 分页获取mtDoctorApproval表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorApprovalService *MtDoctorApprovalService) GetMtDoctorApprovalInfoList(ctx context.Context, info medicineReq.MtDoctorApprovalSearch) (list []medicine.MtDoctorApproval, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtDoctorApproval{})
	var mtDoctorApprovals []medicine.MtDoctorApproval
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.DoctorId != nil {
		db = db.Where("doctor_id = ?", *info.DoctorId)
	}
	if info.DoctorName != nil {
		db = db.Where("doctor_name LIKE ?", "%"+*info.DoctorName+"%")
	}
	if info.ApprovalStatus != nil {
		db = db.Where("approval_status = ?", *info.ApprovalStatus)
	}
	if info.ApproverId != nil {
		db = db.Where("approver_id = ?", *info.ApproverId)
	}
	if info.ApproverName != nil {
		db = db.Where("approver_name LIKE ?", "%"+*info.ApproverName+"%")
	}
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("created_at DESC").Find(&mtDoctorApprovals).Error
	return mtDoctorApprovals, total, err
}

// GetMtDoctorApprovalByDoctorId 根据医生ID获取审核记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorApprovalService *MtDoctorApprovalService) GetMtDoctorApprovalByDoctorId(ctx context.Context, doctorId uint) (mtDoctorApproval medicine.MtDoctorApproval, err error) {
	err = global.GVA_DB.Where("doctor_id = ?", doctorId).Order("created_at DESC").First(&mtDoctorApproval).Error
	return
}

// GetMtDoctorApprovalPublic 不需要鉴权的mtDoctorApproval表接口
// Author [yourname](https://github.com/yourname)
func (mtDoctorApprovalService *MtDoctorApprovalService) GetMtDoctorApprovalPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
