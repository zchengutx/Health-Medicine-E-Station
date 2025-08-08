package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos_client/internal/biz"
)

// 库存仓储实现
type drugInventoryRepo struct {
	data *Data
	log  *log.Helper
}

// 创建库存仓储
func NewDrugInventoryRepo(data *Data, logger log.Logger) biz.DrugInventoryRepo {
	return &drugInventoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 获取数据库连接（支持事务）
func (r *drugInventoryRepo) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return r.data.Db
}

// 检查库存
func (r *drugInventoryRepo) CheckInventory(ctx context.Context, drugID int64, quantity int32) (bool, error) {
	var drug biz.MtDrug
	db := r.getDB(ctx)
	result := db.Where("id = ?", drugID).First(&drug)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, fmt.Errorf("药品不存在: drugID=%d", drugID)
		}
		r.log.Errorf("查询药品库存失败: %v", result.Error)
		return false, result.Error
	}

	// 检查库存是否充足
	available := int32(drug.Inventory)
	if available < quantity {
		r.log.Warnf("库存不足: drugID=%d, available=%d, required=%d", drugID, available, quantity)
		return false, nil
	}

	return true, nil
}

// 预留库存
func (r *drugInventoryRepo) ReserveInventory(ctx context.Context, drugID int64, quantity int32) error {
	// 使用乐观锁更新库存
	db := r.getDB(ctx)
	result := db.Model(&biz.MtDrug{}).
		Where("id = ? AND inventory >= ?", drugID, quantity).
		Update("inventory", gorm.Expr("inventory - ?", quantity))

	if result.Error != nil {
		r.log.Errorf("预留库存失败: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("库存不足或药品不存在: drugID=%d, quantity=%d", drugID, quantity)
	}

	r.log.Infof("预留库存成功: drugID=%d, quantity=%d", drugID, quantity)
	return nil
}

// 减少库存（支付成功后调用）
func (r *drugInventoryRepo) ReduceInventory(ctx context.Context, drugID int64, quantity int32) error {
	// 由于在预留时已经减少了库存，这里只需要记录日志
	// 在实际生产环境中，可能需要更复杂的库存管理逻辑
	r.log.Infof("确认减少库存: drugID=%d, quantity=%d", drugID, quantity)
	return nil
}

// 释放预留库存（取消订单时调用）
func (r *drugInventoryRepo) ReleaseReservedInventory(ctx context.Context, drugID int64, quantity int32) error {
	// 恢复库存
	db := r.getDB(ctx)
	result := db.Model(&biz.MtDrug{}).
		Where("id = ?", drugID).
		Update("inventory", gorm.Expr("inventory + ?", quantity))

	if result.Error != nil {
		r.log.Errorf("释放预留库存失败: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("药品不存在: drugID=%d", drugID)
	}

	r.log.Infof("释放预留库存成功: drugID=%d, quantity=%d", drugID, quantity)
	return nil
}