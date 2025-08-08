package medicine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
)

type MtExplainService struct{}

// GetMtExplain 根据ID获取说明书记录
func (mtExplainService *MtExplainService) GetMtExplain(ctx context.Context, ID string) (mtExplain medicine.MtExplain, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", ID).First(&mtExplain).Error
	return
}

// GetMtExplainInfoList 分页获取说明书记录
func (mtExplainService *MtExplainService) GetMtExplainInfoList(ctx context.Context, info medicineReq.MtExplainSearch) (list []medicine.MtExplain, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&medicine.MtExplain{})
	var mtExplains []medicine.MtExplain
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mtExplains).Error
	return mtExplains, total, err
}
