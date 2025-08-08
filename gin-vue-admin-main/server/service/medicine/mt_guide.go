package medicine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
)

type MtGuideService struct{}

// GetMtGuide 根据ID获取用药指导记录
func (mtGuideService *MtGuideService) GetMtGuide(ctx context.Context, ID string) (mtGuide medicine.MtGuide, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", ID).First(&mtGuide).Error
	return
}

// GetMtGuideInfoList 分页获取用药指导记录
func (mtGuideService *MtGuideService) GetMtGuideInfoList(ctx context.Context, info medicineReq.MtGuideSearch) (list []medicine.MtGuide, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&medicine.MtGuide{})
	var mtGuides []medicine.MtGuide
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mtGuides).Error
	return mtGuides, total, err
}
