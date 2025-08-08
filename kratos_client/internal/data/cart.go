package data

import (
	"context"
	"encoding/json"

	"kratos_client/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartRepo 创建购物车仓库
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 创建购物车项目
func (r *cartRepo) CreateCart(ctx context.Context, userID, drugID, number int64) error {
	// 获取药品信息
	drugInfo, err := r.GetDrugInfo(ctx, drugID)
	if err != nil {
		return err
	}

	// 生成购物车项目
	cartItem := &biz.CartItem{
		ID:            drugID, // 使用drugID作为购物车项目ID
		UserID:        userID,
		DrugID:        drugID,
		Number:        number,
		DrugName:      drugInfo.DrugName,
		Specification: drugInfo.Specification,
		Price:         drugInfo.Price,
		Inventory:     drugInfo.Inventory,
		ExhibitionURL: drugInfo.ExhibitionURL,
	}

	// 序列化为JSON
	cartItemJSON, err := json.Marshal(cartItem)
	if err != nil {
		return err
	}

	// 存储到Redis哈希
	cartKey := biz.GenerateCartKey(userID)
	cartField := biz.GenerateCartField(drugID)

	// 检查是否已存在，如果存在则累加数量
	existingData, err := r.data.RDb.HGet(ctx, cartKey, cartField).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	if err != redis.Nil {
		// 已存在，累加数量
		var existingItem biz.CartItem
		if err := json.Unmarshal([]byte(existingData), &existingItem); err != nil {
			return err
		}

		newNumber := existingItem.Number + number
		// 检查库存
		if newNumber > drugInfo.Inventory {
			return biz.ErrInsufficientInventory
		}

		cartItem.Number = newNumber
		cartItemJSON, err = json.Marshal(cartItem)
		if err != nil {
			return err
		}
	}

	return r.data.RDb.HSet(ctx, cartKey, cartField, cartItemJSON).Err()
}

// 更新购物车项目
func (r *cartRepo) UpdateCart(ctx context.Context, userID, drugID, number int64) error {
	cartKey := biz.GenerateCartKey(userID)
	cartField := biz.GenerateCartField(drugID)

	// 获取现有数据
	existingData, err := r.data.RDb.HGet(ctx, cartKey, cartField).Result()
	if err != nil {
		if err == redis.Nil {
			return biz.ErrCartItemNotFound
		}
		return err
	}

	var cartItem biz.CartItem
	if err := json.Unmarshal([]byte(existingData), &cartItem); err != nil {
		return err
	}

	// 更新数量
	cartItem.Number = number

	// 重新序列化
	cartItemJSON, err := json.Marshal(cartItem)
	if err != nil {
		return err
	}

	return r.data.RDb.HSet(ctx, cartKey, cartField, cartItemJSON).Err()
}

// 删除购物车项目
func (r *cartRepo) DeleteCart(ctx context.Context, userID int64, drugIDs []int64) error {
	cartKey := biz.GenerateCartKey(userID)

	// 构建要删除的字段列表
	fields := make([]string, len(drugIDs))
	for i, drugID := range drugIDs {
		fields[i] = biz.GenerateCartField(drugID)
	}
	return r.data.RDb.HDel(ctx, cartKey, fields...).Err()

}

// 获取购物车列表
func (r *cartRepo) ListCart(ctx context.Context, userID int64) ([]*biz.CartItem, error) {
	cartKey := biz.GenerateCartKey(userID)

	// 获取所有购物车项目
	cartData, err := r.data.RDb.HGetAll(ctx, cartKey).Result()
	if err != nil {
		return nil, err
	}

	var cartItems []*biz.CartItem
	for _, itemJSON := range cartData {
		var item biz.CartItem
		if err := json.Unmarshal([]byte(itemJSON), &item); err != nil {
			r.log.Errorf("Failed to unmarshal cart item: %v", err)
			continue
		}

		// 更新药品信息（确保数据一致性）
		drugInfo, err := r.GetDrugInfo(ctx, item.DrugID)
		if err != nil {
			r.log.Errorf("Failed to get drug info for drugID %d: %v", item.DrugID, err)
			continue
		}

		item.DrugName = drugInfo.DrugName
		item.Specification = drugInfo.Specification
		item.Price = drugInfo.Price
		item.Inventory = drugInfo.Inventory
		item.ExhibitionURL = drugInfo.ExhibitionURL

		cartItems = append(cartItems, &item)
	}

	return cartItems, nil
}

// 检查药品库存
func (r *cartRepo) CheckDrugInventory(ctx context.Context, drugID int64) (int64, error) {
	var inventory int64
	err := r.data.Db.WithContext(ctx).
		Table("mt_drug").
		Select("inventory").
		Where("id = ?", drugID).
		Scan(&inventory).Error

	return inventory, err
}

// 获取药品信息
func (r *cartRepo) GetDrugInfo(ctx context.Context, drugID int64) (*biz.CartItem, error) {
	var drug struct {
		ID            int64   `gorm:"column:id"`
		DrugName      string  `gorm:"column:drug_name"`
		Specification string  `gorm:"column:specification"`
		Price         float64 `gorm:"column:price"`
		Inventory     int64   `gorm:"column:inventory"`
		ExhibitionURL string  `gorm:"column:exhibition_url"`
	}

	err := r.data.Db.WithContext(ctx).
		Table("mt_drug").
		Where("id = ?", drugID).
		First(&drug).Error

	if err != nil {
		return nil, err
	}

	return &biz.CartItem{
		ID:            drug.ID,
		DrugID:        drug.ID,
		DrugName:      drug.DrugName,
		Specification: drug.Specification,
		Price:         drug.Price,
		Inventory:     drug.Inventory,
		ExhibitionURL: drug.ExhibitionURL,
	}, nil
}
