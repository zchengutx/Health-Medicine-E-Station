
package medicine

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
    medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
    "gorm.io/gorm"
)

type MtOrdersService struct {}
// CreateMtOrders 创建mtOrders表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersService *MtOrdersService) CreateMtOrders(ctx context.Context, mtOrders *medicine.MtOrders) (err error) {
	err = global.GVA_DB.Create(mtOrders).Error
	return err
}

// DeleteMtOrders 删除mtOrders表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersService *MtOrdersService)DeleteMtOrders(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&medicine.MtOrders{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&medicine.MtOrders{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMtOrdersByIds 批量删除mtOrders表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersService *MtOrdersService)DeleteMtOrdersByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&medicine.MtOrders{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&medicine.MtOrders{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMtOrders 更新mtOrders表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersService *MtOrdersService)UpdateMtOrders(ctx context.Context, mtOrders medicine.MtOrders) (err error) {
	err = global.GVA_DB.Model(&medicine.MtOrders{}).Where("id = ?",mtOrders.ID).Updates(&mtOrders).Error
	return err
}

// GetMtOrders 根据ID获取mtOrders表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersService *MtOrdersService)GetMtOrders(ctx context.Context, ID string) (mtOrders medicine.MtOrders, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtOrders).Error
	return
}
// GetMtOrdersInfoList 分页获取mtOrders表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersService *MtOrdersService)GetMtOrdersInfoList(ctx context.Context, info medicineReq.MtOrdersSearch) (list []medicine.MtOrders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&medicine.MtOrders{})
    var mtOrderss []medicine.MtOrders
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&mtOrderss).Error
	return  mtOrderss, total, err
}
func (mtOrdersService *MtOrdersService)GetMtOrdersPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
