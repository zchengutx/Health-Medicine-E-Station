package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
)

type MtDepartmentsService struct{}

// CreateMtDepartments 创建科室记录
// @param mtDepartments medicine.MtDepartments
// @return err error
func (mtDepartmentsService *MtDepartmentsService) CreateMtDepartments(mtDepartments medicine.MtDepartments) (err error) {
	err = global.GVA_DB.Create(&mtDepartments).Error
	return err
}

// DeleteMtDepartments 删除科室记录
// @param ID string
// @return err error
func (mtDepartmentsService *MtDepartmentsService) DeleteMtDepartments(ID string) (err error) {
	err = global.GVA_DB.Delete(&medicine.MtDepartments{}, "id = ?", ID).Error
	return err
}

// DeleteMtDepartmentsByIds 批量删除科室记录
// @param IDs []string
// @return err error
func (mtDepartmentsService *MtDepartmentsService) DeleteMtDepartmentsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]medicine.MtDepartments{}, "id in ?", IDs).Error
	return err
}

// UpdateMtDepartments 更新科室记录
// @param mtDepartments medicine.MtDepartments
// @return err error
func (mtDepartmentsService *MtDepartmentsService) UpdateMtDepartments(mtDepartments medicine.MtDepartments) (err error) {
	err = global.GVA_DB.Save(&mtDepartments).Error
	return err
}

// GetMtDepartments 根据id获取科室记录
// @param ID string
// @return mtDepartments medicine.MtDepartments, err error
func (mtDepartmentsService *MtDepartmentsService) GetMtDepartments(ID string) (mtDepartments medicine.MtDepartments, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtDepartments).Error
	return
}

// GetMtDepartmentsInfoList 分页获取科室记录
// @param info request.PageInfo
// @return list []medicine.MtDepartments, total int64, err error
func (mtDepartmentsService *MtDepartmentsService) GetMtDepartmentsInfoList(info request.PageInfo) (list []medicine.MtDepartments, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtDepartments{})
	var mtDepartments []medicine.MtDepartments
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&mtDepartments).Error
	return mtDepartments, total, err
}

// GetMtDepartmentsListPublic 获取科室列表（公开接口）
// @param hospitalId string
// @return list []medicine.MtDepartments, err error
func (mtDepartmentsService *MtDepartmentsService) GetMtDepartmentsListPublic(hospitalId string) (list []medicine.MtDepartments, err error) {
	db := global.GVA_DB.Where("status = ?", "启用")
	if hospitalId != "" {
		db = db.Where("hospital_id = ?", hospitalId)
	}
	err = db.Find(&list).Error
	return
}
