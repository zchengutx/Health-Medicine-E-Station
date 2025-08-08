
package medicine

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
    medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
    "gorm.io/gorm"
)

type MtDiscountService struct {}
// CreateMtDiscount 创建mtDiscount表记录
// Author [yourname](https://github.com/yourname)
func (mtDiscountService *MtDiscountService) CreateMtDiscount(ctx context.Context, mtDiscount *medicine.MtDiscount) (err error) {
	err = global.GVA_DB.Create(mtDiscount).Error
	return err
}

// DeleteMtDiscount 删除mtDiscount表记录
// Author [yourname](https://github.com/yourname)
func (mtDiscountService *MtDiscountService)DeleteMtDiscount(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&medicine.MtDiscount{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&medicine.MtDiscount{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMtDiscountByIds 批量删除mtDiscount表记录
// Author [yourname](https://github.com/yourname)
func (mtDiscountService *MtDiscountService)DeleteMtDiscountByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&medicine.MtDiscount{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&medicine.MtDiscount{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMtDiscount 更新mtDiscount表记录
// Author [yourname](https://github.com/yourname)
func (mtDiscountService *MtDiscountService)UpdateMtDiscount(ctx context.Context, mtDiscount medicine.MtDiscount) (err error) {
	err = global.GVA_DB.Model(&medicine.MtDiscount{}).Where("id = ?",mtDiscount.ID).Updates(&mtDiscount).Error
	return err
}

// GetMtDiscount 根据ID获取mtDiscount表记录
// Author [yourname](https://github.com/yourname)
func (mtDiscountService *MtDiscountService)GetMtDiscount(ctx context.Context, ID string) (mtDiscount medicine.MtDiscount, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtDiscount).Error
	return
}
// GetMtDiscountInfoList 分页获取mtDiscount表记录
// Author [yourname](https://github.com/yourname)
func (mtDiscountService *MtDiscountService)GetMtDiscountInfoList(ctx context.Context, info medicineReq.MtDiscountSearch) (list []medicine.MtDiscount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&medicine.MtDiscount{})
    var mtDiscounts []medicine.MtDiscount
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

	err = db.Find(&mtDiscounts).Error
	return  mtDiscounts, total, err
}
func (mtDiscountService *MtDiscountService)GetMtDiscountPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
