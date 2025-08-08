package medicine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"gorm.io/gorm"
)

type MtDrugTypeStairService struct{}

// CreateMtDrugTypeStair 创建mtDrugTypeStair表记录
func (mtDrugTypeStairService *MtDrugTypeStairService) CreateMtDrugTypeStair(ctx context.Context, mtDrugTypeStair *medicine.MtDrugTypeStair) (err error) {
	err = global.GVA_DB.Create(mtDrugTypeStair).Error
	return err
}

// DeleteMtDrugTypeStair 删除mtDrugTypeStair表记录
func (mtDrugTypeStairService *MtDrugTypeStairService) DeleteMtDrugTypeStair(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDrugTypeStair{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&medicine.MtDrugTypeStair{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMtDrugTypeStairByIds 批量删除mtDrugTypeStair表记录
func (mtDrugTypeStairService *MtDrugTypeStairService) DeleteMtDrugTypeStairByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtDrugTypeStair{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&medicine.MtDrugTypeStair{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMtDrugTypeStair 更新mtDrugTypeStair表记录
func (mtDrugTypeStairService *MtDrugTypeStairService) UpdateMtDrugTypeStair(ctx context.Context, mtDrugTypeStair medicine.MtDrugTypeStair) (err error) {
	err = global.GVA_DB.Model(&medicine.MtDrugTypeStair{}).Where("id = ?", mtDrugTypeStair.ID).Updates(&mtDrugTypeStair).Error
	return err
}

// GetMtDrugTypeStair 根据ID获取mtDrugTypeStair表记录
func (mtDrugTypeStairService *MtDrugTypeStairService) GetMtDrugTypeStair(ctx context.Context, ID string) (mtDrugTypeStair medicine.MtDrugTypeStair, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtDrugTypeStair).Error
	return
}

// GetMtDrugTypeStairInfoList 分页获取mtDrugTypeStair表记录
func (mtDrugTypeStairService *MtDrugTypeStairService) GetMtDrugTypeStairInfoList(ctx context.Context, info medicineReq.MtDrugTypeStairSearch) (list []medicine.MtDrugTypeStair, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtDrugTypeStair{})
	var mtDrugTypeStairs []medicine.MtDrugTypeStair
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

	err = db.Find(&mtDrugTypeStairs).Error
	return mtDrugTypeStairs, total, err
}

// GetAllMtDrugTypeStair 获取所有一级分类
func (mtDrugTypeStairService *MtDrugTypeStairService) GetAllMtDrugTypeStair(ctx context.Context) (list []medicine.MtDrugTypeStair, err error) {
	// 暂时返回空数据，避免查询不存在的表
	// err = global.GVA_DB.Find(&list).Error
	list = []medicine.MtDrugTypeStair{}
	return
}
