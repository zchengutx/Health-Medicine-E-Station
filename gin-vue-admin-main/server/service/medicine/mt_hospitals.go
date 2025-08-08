package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
)

type MtHospitalsService struct{}

// CreateMtHospitals 创建医院记录
// @param mtHospitals medicine.MtHospitals
// @return err error
func (mtHospitalsService *MtHospitalsService) CreateMtHospitals(mtHospitals medicine.MtHospitals) (err error) {
	err = global.GVA_DB.Create(&mtHospitals).Error
	return err
}

// DeleteMtHospitals 删除医院记录
// @param ID string
// @return err error
func (mtHospitalsService *MtHospitalsService) DeleteMtHospitals(ID string) (err error) {
	err = global.GVA_DB.Delete(&medicine.MtHospitals{}, "id = ?", ID).Error
	return err
}

// DeleteMtHospitalsByIds 批量删除医院记录
// @param IDs []string
// @return err error
func (mtHospitalsService *MtHospitalsService) DeleteMtHospitalsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]medicine.MtHospitals{}, "id in ?", IDs).Error
	return err
}

// UpdateMtHospitals 更新医院记录
// @param mtHospitals medicine.MtHospitals
// @return err error
func (mtHospitalsService *MtHospitalsService) UpdateMtHospitals(mtHospitals medicine.MtHospitals) (err error) {
	err = global.GVA_DB.Save(&mtHospitals).Error
	return err
}

// GetMtHospitals 根据id获取医院记录
// @param ID string
// @return mtHospitals medicine.MtHospitals, err error
func (mtHospitalsService *MtHospitalsService) GetMtHospitals(ID string) (mtHospitals medicine.MtHospitals, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtHospitals).Error
	return
}

// GetMtHospitalsInfoList 分页获取医院记录
// @param info request.PageInfo
// @return list []medicine.MtHospitals, total int64, err error
func (mtHospitalsService *MtHospitalsService) GetMtHospitalsInfoList(info request.PageInfo) (list []medicine.MtHospitals, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtHospitals{})
	var mtHospitalss []medicine.MtHospitals
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&mtHospitalss).Error
	return mtHospitalss, total, err
}

// GetMtHospitalsListPublic 获取医院列表（公开接口）
// @return list []medicine.MtHospitals, err error
func (mtHospitalsService *MtHospitalsService) GetMtHospitalsListPublic() (list []medicine.MtHospitals, err error) {
	err = global.GVA_DB.Where("status = ?", "启用").Find(&list).Error
	return
}
