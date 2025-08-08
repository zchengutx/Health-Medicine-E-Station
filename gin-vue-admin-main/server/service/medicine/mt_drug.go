package medicine

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"gorm.io/gorm"
)

type MtDrugService struct{}

// CreateMtDrug 创建mtDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtDrugService *MtDrugService) CreateMtDrug(ctx context.Context, mtDrug *medicine.MtDrug) (err error) {
	err = global.GVA_DB.Create(mtDrug).Error
	return err
}

// DeleteMtDrug 删除mtDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtDrugService *MtDrugService) DeleteMtDrug(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDrug{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&medicine.MtDrug{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMtDrugByIds 批量删除mtDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtDrugService *MtDrugService) DeleteMtDrugByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDrug{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&medicine.MtDrug{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMtDrug 更新mtDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtDrugService *MtDrugService) UpdateMtDrug(ctx context.Context, mtDrug medicine.MtDrug) (err error) {
	err = global.GVA_DB.Model(&medicine.MtDrug{}).Where("id = ?", mtDrug.ID).Updates(&mtDrug).Error
	return err
}

// GetMtDrug 根据ID获取mtDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtDrugService *MtDrugService) GetMtDrug(ctx context.Context, ID string) (mtDrug medicine.MtDrug, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtDrug).Error
	return
}

// GetMtDrugInfoList 分页获取mtDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtDrugService *MtDrugService) GetMtDrugInfoList(ctx context.Context, info medicineReq.MtDrugSearch) (list []medicine.MtDrug, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtDrug{})
	var mtDrugs []medicine.MtDrug
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

	err = db.Find(&mtDrugs).Error
	if err != nil {
		return
	}

	var guide []medicine.MtGuide
	err = global.GVA_DB.Table("mt_guide").Select("id, usage_and_dosage").Find(&guide).Error
	if err != nil {
		return
	}

	guideMap := make(map[int]string, len(guide))
	for _, g := range guide {
		guideMap[int(g.Id)] = g.UsageAndDosage
	}

	var explain []medicine.MtExplain
	err = global.GVA_DB.Table("mt_explain").Select("id, goods_name").Find(&explain).Error
	if err != nil {
		return
	}

	explainMap := make(map[int]string, len(explain))
	for _, e := range explain {
		explainMap[int(e.Id)] = e.GoodsName
	}

	for i, drug := range mtDrugs {
		if drug.Guide != nil && *drug.Guide > 0 {
			mtDrugs[i].Guides = guideMap[*drug.Guide]
		}
		if drug.Explain != nil && *drug.Explain > 0 {
			mtDrugs[i].Explains = explainMap[*drug.Explain]
		}
	}

	return mtDrugs, total, err
}
func (mtDrugService *MtDrugService) GetMtDrugPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
