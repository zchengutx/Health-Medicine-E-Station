package medicine

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"gorm.io/gorm"
)

type MtDoctorsService struct{}

// CreateMtDoctors 创建mtDoctors表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorsService *MtDoctorsService) CreateMtDoctors(ctx context.Context, mtDoctors *medicine.MtDoctors) (err error) {
	err = global.GVA_DB.Create(mtDoctors).Error
	return err
}

// DeleteMtDoctors 删除mtDoctors表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorsService *MtDoctorsService) DeleteMtDoctors(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDoctors{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&medicine.MtDoctors{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMtDoctorsByIds 批量删除mtDoctors表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorsService *MtDoctorsService) DeleteMtDoctorsByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDoctors{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&medicine.MtDoctors{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMtDoctors 更新mtDoctors表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorsService *MtDoctorsService) UpdateMtDoctors(ctx context.Context, mtDoctors medicine.MtDoctors) (err error) {
	err = global.GVA_DB.Model(&medicine.MtDoctors{}).Where("id = ?", mtDoctors.ID).Updates(&mtDoctors).Error
	return err
}

// GetMtDoctors 根据ID获取mtDoctors表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorsService *MtDoctorsService) GetMtDoctors(ctx context.Context, ID string) (mtDoctors medicine.MtDoctors, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtDoctors).Error
	return
}

// GetMtDoctorsInfoList 分页获取mtDoctors表记录
// Author [yourname](https://github.com/yourname)
func (mtDoctorsService *MtDoctorsService) GetMtDoctorsInfoList(ctx context.Context, info medicineReq.MtDoctorsSearch) (list []medicine.MtDoctors, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtDoctors{})
	var mtDoctorss []medicine.MtDoctors
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
	err = db.Find(&mtDoctorss).Error
	if err != nil {
		return // 处理查询医生的错误
	}

	var hospitals []medicine.MtHospitals
	err = global.GVA_DB.Table("mt_hospitals").Select("id, name").Find(&hospitals).Error
	if err != nil {
		return
	}

	hospitalMap := make(map[int]string, len(hospitals))
	for _, hosp := range hospitals {
		hospitalMap[int(hosp.ID)] = hosp.Name
	}

	var departments []medicine.MtDepartments
	err = global.GVA_DB.Table("mt_departments").Select("id, name").Find(&departments).Error
	if err != nil {
		return
	}

	departmentMap := make(map[int]string, len(departments))
	for _, dept := range departments {
		departmentMap[int(dept.ID)] = dept.Name
	}

	for i, doctor := range mtDoctorss {
		if doctor.HospitalId != nil && *doctor.HospitalId > 0 {
			mtDoctorss[i].Hospital = hospitalMap[*doctor.HospitalId]
		}
		if doctor.DepartmentId != nil && *doctor.DepartmentId > 0 {
			mtDoctorss[i].Department = departmentMap[*doctor.DepartmentId]
		}
	}

	return mtDoctorss, total, err

}
func (mtDoctorsService *MtDoctorsService) GetMtDoctorsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
