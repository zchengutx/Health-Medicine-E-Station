package medicine

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"gorm.io/gorm"
)

type MtChatMessageService struct{}

// CreateMtChatMessage 创建mtChatMessage表记录
// Author [yourname](https://github.com/yourname)
func (mtChatMessageService *MtChatMessageService) CreateMtChatMessage(ctx context.Context, mtChatMessage *medicine.MtChatMessage) (err error) {
	err = global.GVA_DB.Create(mtChatMessage).Error
	return err
}

// DeleteMtChatMessage 删除mtChatMessage表记录
// Author [yourname](https://github.com/yourname)
func (mtChatMessageService *MtChatMessageService) DeleteMtChatMessage(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtChatMessage{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&medicine.MtChatMessage{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMtChatMessageByIds 批量删除mtChatMessage表记录
// Author [yourname](https://github.com/yourname)
func (mtChatMessageService *MtChatMessageService) DeleteMtChatMessageByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtChatMessage{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&medicine.MtChatMessage{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMtChatMessage 更新mtChatMessage表记录
// Author [yourname](https://github.com/yourname)
func (mtChatMessageService *MtChatMessageService) UpdateMtChatMessage(ctx context.Context, mtChatMessage medicine.MtChatMessage) (err error) {
	err = global.GVA_DB.Model(&medicine.MtChatMessage{}).Where("id = ?", mtChatMessage.ID).Updates(&mtChatMessage).Error
	return err
}

// GetMtChatMessage 根据ID获取mtChatMessage表记录
// Author [yourname](https://github.com/yourname)
func (mtChatMessageService *MtChatMessageService) GetMtChatMessage(ctx context.Context, ID string) (mtChatMessage medicine.MtChatMessage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtChatMessage).Error
	return
}

// GetMtChatMessageInfoList 分页获取mtChatMessage表记录
// Author [yourname](https://github.com/yourname)
func (mtChatMessageService *MtChatMessageService) GetMtChatMessageInfoList(ctx context.Context, info medicineReq.MtChatMessageSearch) (list []medicine.MtChatMessage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtChatMessage{})
	var mtChatMessages []medicine.MtChatMessage
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

	err = db.Find(&mtChatMessages).Error

	return mtChatMessages, total, err
}
func (mtChatMessageService *MtChatMessageService) GetMtChatMessagePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
