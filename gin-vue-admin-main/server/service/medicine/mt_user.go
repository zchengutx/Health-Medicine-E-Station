
package medicine

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
    medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
    "gorm.io/gorm"
)

type MtUserService struct {}
// CreateMtUser 创建mtUser表记录
// Author [yourname](https://github.com/yourname)
func (mtUserService *MtUserService) CreateMtUser(ctx context.Context, mtUser *medicine.MtUser) (err error) {
	err = global.GVA_DB.Create(mtUser).Error
	return err
}

// DeleteMtUser 删除mtUser表记录
// Author [yourname](https://github.com/yourname)
func (mtUserService *MtUserService)DeleteMtUser(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&medicine.MtUser{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&medicine.MtUser{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMtUserByIds 批量删除mtUser表记录
// Author [yourname](https://github.com/yourname)
func (mtUserService *MtUserService)DeleteMtUserByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&medicine.MtUser{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&medicine.MtUser{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMtUser 更新mtUser表记录
// Author [yourname](https://github.com/yourname)
func (mtUserService *MtUserService)UpdateMtUser(ctx context.Context, mtUser medicine.MtUser) (err error) {
	err = global.GVA_DB.Model(&medicine.MtUser{}).Where("id = ?",mtUser.ID).Updates(&mtUser).Error
	return err
}

// GetMtUser 根据ID获取mtUser表记录
// Author [yourname](https://github.com/yourname)
func (mtUserService *MtUserService)GetMtUser(ctx context.Context, ID string) (mtUser medicine.MtUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtUser).Error
	return
}
// GetMtUserInfoList 分页获取mtUser表记录
// Author [yourname](https://github.com/yourname)
func (mtUserService *MtUserService)GetMtUserInfoList(ctx context.Context, info medicineReq.MtUserSearch) (list []medicine.MtUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&medicine.MtUser{})
    var mtUsers []medicine.MtUser
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

	err = db.Find(&mtUsers).Error
	return  mtUsers, total, err
}
func (mtUserService *MtUserService)GetMtUserPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
