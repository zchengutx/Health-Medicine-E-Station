package medicine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"gorm.io/gorm"
)

type MtDrugTypeLevelService struct{}

// CreateMtDrugTypeLevel 创建mtDrugTypeLevel表记录
func (mtDrugTypeLevelService *MtDrugTypeLevelService) CreateMtDrugTypeLevel(ctx context.Context, mtDrugTypeLevel *medicine.MtDrugTypeLevel) (err error) {
	err = global.GVA_DB.Create(mtDrugTypeLevel).Error
	return err
}

// DeleteMtDrugTypeLevel 删除mtDrugTypeLevel表记录
func (mtDrugTypeLevelService *MtDrugTypeLevelService) DeleteMtDrugTypeLevel(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDrugTypeLevel{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&medicine.MtDrugTypeLevel{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMtDrugTypeLevelByIds 批量删除mtDrugTypeLevel表记录
func (mtDrugTypeLevelService *MtDrugTypeLevelService) DeleteMtDrugTypeLevelByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDrugTypeLevel{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&medicine.MtDrugTypeLevel{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMtDrugTypeLevel 更新mtDrugTypeLevel表记录
func (mtDrugTypeLevelService *MtDrugTypeLevelService) UpdateMtDrugTypeLevel(ctx context.Context, mtDrugTypeLevel medicine.MtDrugTypeLevel) (err error) {
	err = global.GVA_DB.Model(&medicine.MtDrugTypeLevel{}).Where("id = ?", mtDrugTypeLevel.ID).Updates(&mtDrugTypeLevel).Error
	return err
}

// GetMtDrugTypeLevel 根据ID获取mtDrugTypeLevel表记录
func (mtDrugTypeLevelService *MtDrugTypeLevelService) GetMtDrugTypeLevel(ctx context.Context, ID string) (mtDrugTypeLevel medicine.MtDrugTypeLevel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtDrugTypeLevel).Error
	return
}

// GetMtDrugTypeLevelInfoList 分页获取mtDrugTypeLevel表记录
func (mtDrugTypeLevelService *MtDrugTypeLevelService) GetMtDrugTypeLevelInfoList(ctx context.Context, info medicineReq.MtDrugTypeLevelSearch) (list []medicine.MtDrugTypeLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtDrugTypeLevel{})
	var mtDrugTypeLevels []medicine.MtDrugTypeLevel
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

	err = db.Find(&mtDrugTypeLevels).Error
	return mtDrugTypeLevels, total, err
}

// GetAllMtDrugTypeLevel 获取所有二级分类
func (mtDrugTypeLevelService *MtDrugTypeLevelService) GetAllMtDrugTypeLevel(ctx context.Context) (list []medicine.MtDrugTypeLevel, err error) {
	// 暂时返回空数据，避免查询不存在的表
	// err = global.GVA_DB.Find(&list).Error
	list = []medicine.MtDrugTypeLevel{}
	return
}
