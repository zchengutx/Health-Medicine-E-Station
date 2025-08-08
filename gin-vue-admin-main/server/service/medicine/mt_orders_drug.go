package medicine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"gorm.io/gorm"
)

type MtOrdersDrugService struct{}

// CreateMtOrdersDrug 创建mtOrdersDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersDrugService *MtOrdersDrugService) CreateMtOrdersDrug(ctx context.Context, mtOrdersDrug *medicine.MtOrdersDrug) (err error) {
	err = global.GVA_DB.Create(mtOrdersDrug).Error
	return err
}

// DeleteMtOrdersDrug 删除mtOrdersDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersDrugService *MtOrdersDrugService) DeleteMtOrdersDrug(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtOrdersDrug{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&medicine.MtOrdersDrug{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMtOrdersDrugByIds 批量删除mtOrdersDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersDrugService *MtOrdersDrugService) DeleteMtOrdersDrugByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&medicine.MtOrdersDrug{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&medicine.MtOrdersDrug{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMtOrdersDrug 更新mtOrdersDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersDrugService *MtOrdersDrugService) UpdateMtOrdersDrug(ctx context.Context, mtOrdersDrug medicine.MtOrdersDrug) (err error) {
	err = global.GVA_DB.Model(&medicine.MtOrdersDrug{}).Where("id = ?", mtOrdersDrug.ID).Updates(&mtOrdersDrug).Error
	return err
}

// GetMtOrdersDrug 根据ID获取mtOrdersDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersDrugService *MtOrdersDrugService) GetMtOrdersDrug(ctx context.Context, ID string) (mtOrdersDrug medicine.MtOrdersDrug, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mtOrdersDrug).Error
	if err != nil {
		return
	}

	// 获取订单信息
	if mtOrdersDrug.OrderId != nil {
		var order medicine.MtOrders
		err = global.GVA_DB.Where("id = ?", *mtOrdersDrug.OrderId).First(&order).Error
		if err == nil && order.OrderNo != nil {
			mtOrdersDrug.Order = *order.OrderNo
		}
	}

	// 获取药品信息
	if mtOrdersDrug.DrugId != nil {
		var drug medicine.MtDrug
		err = global.GVA_DB.Where("id = ?", *mtOrdersDrug.DrugId).First(&drug).Error
		if err == nil && drug.DrugName != nil {
			mtOrdersDrug.Drug = *drug.DrugName
		}
	}

	// 获取用户信息
	if mtOrdersDrug.UserId != nil {
		var user medicine.MtUser
		err = global.GVA_DB.Where("id = ?", *mtOrdersDrug.UserId).First(&user).Error
		if err == nil {
			mtOrdersDrug.User = *user.NickName
		}
	}

	return
}

// GetMtOrdersDrugInfoList 分页获取mtOrdersDrug表记录
// Author [yourname](https://github.com/yourname)
func (mtOrdersDrugService *MtOrdersDrugService) GetMtOrdersDrugInfoList(ctx context.Context, info medicineReq.MtOrdersDrugSearch) (list []medicine.MtOrdersDrug, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&medicine.MtOrdersDrug{})
	var mtOrdersDrugs []medicine.MtOrdersDrug
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
	err = db.Find(&mtOrdersDrugs).Error
	if err != nil {
		return
	}

	var order []medicine.MtOrders
	err = global.GVA_DB.Table("mt_orders").Select("id,order_no").Find(&order).Error
	if err != nil {
		return
	}
	orderMap := make(map[uint64]string, len(order))
	for _, o := range order {
		if o.OrderNo != nil {
			orderMap[uint64(o.ID)] = *o.OrderNo
		}
	}

	var drug []medicine.MtDrug
	err = global.GVA_DB.Table("mt_drug").Select("id,drug_name").Find(&drug).Error
	if err != nil {
		return
	}
	drugMap := make(map[uint64]string, len(drug))
	for _, d := range drug {
		drugMap[uint64(d.ID)] = *d.DrugName
	}

	var user []medicine.MtUser
	err = global.GVA_DB.Table("mt_user").Select("id,nick_name").Find(&user).Error
	if err != nil {
		return
	}
	userMap := make(map[uint64]string, len(user))
	for _, u := range user {
		userMap[uint64(u.ID)] = *u.NickName
	}

	for i, d := range mtOrdersDrugs {
		if d.OrderId != nil && *d.OrderId > 0 {
			mtOrdersDrugs[i].Order = orderMap[uint64(*d.OrderId)]
		}
		if d.DrugId != nil && *d.DrugId > 0 {
			mtOrdersDrugs[i].Drug = drugMap[uint64(*d.DrugId)]
		}
		if d.UserId != nil && *d.UserId > 0 {
			mtOrdersDrugs[i].User = userMap[uint64(*d.UserId)]
		}
	}
	//fmt.Println(orderMap, drugMap, userMap)

	return mtOrdersDrugs, total, err
}
func (mtOrdersDrugService *MtOrdersDrugService) GetMtOrdersDrugPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetMtOrdersDrugDetail 获取订单详情信息
// Author [yourname](https://github.com/yourname)
func (mtOrdersDrugService *MtOrdersDrugService) GetMtOrdersDrugDetail(ctx context.Context, ID string) (detail map[string]interface{}, err error) {
	var mtOrdersDrug medicine.MtOrdersDrug
	err = global.GVA_DB.Where("id = ?", ID).First(&mtOrdersDrug).Error
	if err != nil {
		return
	}

	detail = make(map[string]interface{})

	// 获取订单信息
	if mtOrdersDrug.OrderId != nil {
		var order medicine.MtOrders
		err = global.GVA_DB.Where("id = ?", *mtOrdersDrug.OrderId).First(&order).Error
		if err == nil {
			detail["orderInfo"] = map[string]interface{}{
				"orderNo":     order.OrderNo,
				"totalAmount": order.TotalAmount,
				"status":      order.Status,
				"payTime":     order.CreatedAt,
			}
		}
	}

	// 获取药品信息
	if mtOrdersDrug.DrugId != nil {
		var drug medicine.MtDrug
		err = global.GVA_DB.Where("id = ?", *mtOrdersDrug.DrugId).First(&drug).Error
		if err == nil {
			detail["drugInfo"] = map[string]interface{}{
				"drugName":      drug.DrugName,
				"specification": drug.Specification,
				"price":         drug.Price,
				"quantity":      mtOrdersDrug.Quantity,
			}
		}
	}

	// 获取用户信息
	if mtOrdersDrug.UserId != nil {
		var user medicine.MtUser
		err = global.GVA_DB.Where("id = ?", *mtOrdersDrug.UserId).First(&user).Error
		if err == nil {
			detail["userInfo"] = map[string]interface{}{
				"nickName": user.NickName,
				"mobile":   user.Mobile,
			}
		}
	}

	// 获取收货地址信息
	if mtOrdersDrug.OrderId != nil {
		var order medicine.MtOrders
		err = global.GVA_DB.Where("id = ?", *mtOrdersDrug.OrderId).First(&order).Error
		if err == nil {
			detail["addressInfo"] = map[string]interface{}{
				"addressDetail": order.AddressDetail,
			}
		}
	}

	return detail, nil
}
