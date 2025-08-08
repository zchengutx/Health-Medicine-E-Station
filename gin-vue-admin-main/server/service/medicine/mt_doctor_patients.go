package medicine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"gorm.io/gorm"
)

type MtDoctorPatientsService struct{}

// CreateMtDoctorPatients 创建mtDoctorPatients表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorPatientsService *MtDoctorPatientsService) CreateMtDoctorPatients(ctx context.Context, mtDoctorPatients *medicine.MtDoctorPatients) (err error) {
	err = global.GVA_DB.Create(mtDoctorPatients).Error
	return err
}

// DeleteMtDoctorPatients 删除mtDoctorPatients表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorPatientsService *MtDoctorPatientsService) DeleteMtDoctorPatients(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDoctorPatients{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&medicine.MtDoctorPatients{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMtDoctorPatientsByIds 批量删除mtDoctorPatients表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorPatientsService *MtDoctorPatientsService) DeleteMtDoctorPatientsByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDoctorPatients{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&medicine.MtDoctorPatients{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMtDoctorPatients 更新mtDoctorPatients表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorPatientsService *MtDoctorPatientsService) UpdateMtDoctorPatients(ctx context.Context, mtDoctorPatients medicine.MtDoctorPatients) (err error) {
	err = global.GVA_DB.Model(&medicine.MtDoctorPatients{}).Where("id = ?", mtDoctorPatients.ID).Updates(&mtDoctorPatients).Error
	return err
}

// GetMtDoctorPatients 根据ID获取mtDoctorPatients表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorPatientsService *MtDoctorPatientsService) GetMtDoctorPatients(ctx context.Context, ID string) (mtDoctorPatients medicine.MtDoctorPatients, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtDoctorPatients).Error
	return
}

// GetMtDoctorPatientsInfoList 分页获取mtDoctorPatients表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorPatientsService *MtDoctorPatientsService) GetMtDoctorPatientsInfoList(ctx context.Context, info medicineReq.MtDoctorPatientsSearch) (list []medicine.MtDoctorPatients, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtDoctorPatients{})
	var mtDoctorPatientss []medicine.MtDoctorPatients
	// 如果有条件搜索 下方会自动创建搜索语句
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

	err = db.Find(&mtDoctorPatientss).Error
	return mtDoctorPatientss, total, err
}
func (mtDoctorPatientsService *MtDoctorPatientsService) GetMtDoctorPatientsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
