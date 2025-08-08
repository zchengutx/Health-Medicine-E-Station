package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"kratos_client/internal/biz"
)

// 订单数据模型 - 对应 mt_orders 表
type MtOrder struct {
	ID            uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderNo       string          `gorm:"column:order_no;size:50;not null" json:"order_no"`
	UserID        int64           `gorm:"column:user_id" json:"user_id"`
	UserName      string          `gorm:"column:user_name;size:50" json:"user_name"`
	UserPhone     string          `gorm:"column:user_phone;size:20" json:"user_phone"`
	DoctorID      int64           `gorm:"column:doctor_id" json:"doctor_id"`
	DoctorName    string          `gorm:"column:doctor_name;size:50" json:"doctor_name"`
	AddressID     int64           `gorm:"column:address_id" json:"address_id"`
	AddressDetail string          `gorm:"column:address_detail;size:200" json:"address_detail"`
	TotalAmount   decimal.Decimal `gorm:"column:total_amount;type:decimal(10,2);not null" json:"total_amount"`
	PayType       string          `gorm:"column:pay_type;size:10" json:"pay_type"`
	Status         string          `gorm:"column:status;size:20" json:"status"`
	PayTime        *time.Time      `gorm:"column:pay_time;type:datetime(3)" json:"pay_time"`
	DrugTime       *time.Time      `gorm:"column:drug_time;type:datetime(3)" json:"drug_time"`
	SendTime       *time.Time      `gorm:"column:send_time;type:datetime(3)" json:"send_time"`
	FinishTime     *time.Time      `gorm:"column:finish_time;type:datetime(3)" json:"finish_time"`
	CancelTime     *time.Time      `gorm:"column:cancel_time;type:datetime(3)" json:"cancel_time"`
	UserCouponID   *int64          `gorm:"column:user_coupon_id" json:"user_coupon_id"`
	DiscountAmount decimal.Decimal `gorm:"column:discount_amount;type:decimal(10,2);default:0" json:"discount_amount"`
	OriginalAmount decimal.Decimal `gorm:"column:original_amount;type:decimal(10,2);default:0" json:"original_amount"`
	Remark         string          `gorm:"column:remark;size:500" json:"remark"`
}

// 表名
func (MtOrder) TableName() string {
	return "mt_orders"
}

// 订单项数据模型
type MtOrderItem struct {
	ID        int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int64           `gorm:"index;not null" json:"order_id"`
	DrugID    int64           `gorm:"index;not null" json:"drug_id"`
	DrugName  string          `gorm:"size:100;not null" json:"drug_name"`
	DrugSpec  string          `gorm:"size:50" json:"drug_spec"`
	Quantity  int32           `gorm:"not null" json:"quantity"`
	Price     decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"price"`
	Subtotal  decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"subtotal"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
}

// 表名
func (MtOrderItem) TableName() string {
	return "mt_order_items"
}

// 订单仓储实现
type orderRepo struct {
	data *Data
	log  *log.Helper
}

// 创建订单仓储
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 转换业务模型到数据模型
func (r *orderRepo) toBizOrder(do *MtOrder) *biz.MtOrder {
	return &biz.MtOrder{
		ID:             do.ID,
		OrderNo:        do.OrderNo,
		UserID:         do.UserID,
		UserName:       do.UserName,
		UserPhone:      do.UserPhone,
		DoctorID:       do.DoctorID,
		DoctorName:     do.DoctorName,
		AddressID:      do.AddressID,
		AddressDetail:  do.AddressDetail,
		TotalAmount:    do.TotalAmount,
		PayType:        do.PayType,
		Status:         do.Status,
		PayTime:        do.PayTime,
		DrugTime:       do.DrugTime,
		SendTime:       do.SendTime,
		FinishTime:     do.FinishTime,
		CancelTime:     do.CancelTime,
		UserCouponID:   do.UserCouponID,
		DiscountAmount: do.DiscountAmount,
		OriginalAmount: do.OriginalAmount,
		Remark:         do.Remark,
	}
}

// 转换业务模型到数据模型
func (r *orderRepo) toDataOrder(bo *biz.MtOrder) *MtOrder {
	return &MtOrder{
		ID:             bo.ID,
		OrderNo:        bo.OrderNo,
		UserID:         bo.UserID,
		UserName:       bo.UserName,
		UserPhone:      bo.UserPhone,
		DoctorID:       bo.DoctorID,
		DoctorName:     bo.DoctorName,
		AddressID:      bo.AddressID,
		AddressDetail:  bo.AddressDetail,
		TotalAmount:    bo.TotalAmount,
		PayType:        bo.PayType,
		Status:         bo.Status,
		PayTime:        bo.PayTime,
		DrugTime:       bo.DrugTime,
		SendTime:       bo.SendTime,
		FinishTime:     bo.FinishTime,
		CancelTime:     bo.CancelTime,
		UserCouponID:   bo.UserCouponID,
		DiscountAmount: bo.DiscountAmount,
		OriginalAmount: bo.OriginalAmount,
		Remark:         bo.Remark,
	}
}

// 转换订单项业务模型到数据模型
func (r *orderRepo) toBizOrderItem(doi *MtOrderItem) *biz.MtOrderItem {
	return &biz.MtOrderItem{
		ID:        doi.ID,
		OrderID:   doi.OrderID,
		DrugID:    doi.DrugID,
		DrugName:  doi.DrugName,
		DrugSpec:  doi.DrugSpec,
		Quantity:  doi.Quantity,
		Price:     doi.Price,
		Subtotal:  doi.Subtotal,
		CreatedAt: doi.CreatedAt,
	}
}

// 转换订单项业务模型到数据模型
func (r *orderRepo) toDataOrderItem(boi *biz.MtOrderItem) *MtOrderItem {
	return &MtOrderItem{
		ID:        boi.ID,
		OrderID:   boi.OrderID,
		DrugID:    boi.DrugID,
		DrugName:  boi.DrugName,
		DrugSpec:  boi.DrugSpec,
		Quantity:  boi.Quantity,
		Price:     boi.Price,
		Subtotal:  boi.Subtotal,
		CreatedAt: boi.CreatedAt,
	}
}

// 创建订单
func (r *orderRepo) CreateOrder(ctx context.Context, order *biz.MtOrder) error {
	do := r.toDataOrder(order)

	db := r.getDB(ctx)
	result := db.Create(do)
	if result.Error != nil {
		r.log.Errorf("创建订单失败: %v", result.Error)
		return result.Error
	}

	// 更新ID
	order.ID = do.ID
	return nil
}

// 根据ID查询订单
func (r *orderRepo) GetOrderByID(ctx context.Context, id int64) (*biz.MtOrder, error) {
	var do MtOrder
	db := r.getDB(ctx)
	result := db.Where("id = ?", id).First(&do)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.log.Errorf("查询订单失败: %v", result.Error)
		return nil, result.Error
	}

	return r.toBizOrder(&do), nil
}

// 根据订单号查询订单
func (r *orderRepo) GetOrderByOrderNo(ctx context.Context, orderNo string) (*biz.MtOrder, error) {
	var do MtOrder
	db := r.getDB(ctx)
	result := db.Where("order_no = ?", orderNo).First(&do)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.log.Errorf("查询订单失败: %v", result.Error)
		return nil, result.Error
	}

	return r.toBizOrder(&do), nil
}

// 更新订单状态
func (r *orderRepo) UpdateOrderStatus(ctx context.Context, orderNo string, status string, timestamp time.Time) error {
	updates := map[string]interface{}{
		"status": status,
	}

	// 根据状态设置对应的时间字段
	switch status {
	case "2": // 已支付
		if updates["pay_time"] == nil {
			updates["pay_time"] = timestamp
		}
	case "3": // 配药中
		updates["drug_time"] = timestamp
	case "4": // 已发货
		updates["send_time"] = timestamp
	case "5": // 已完成
		updates["finish_time"] = timestamp
	case "6": // 已取消
		updates["cancel_time"] = timestamp
	}

	db := r.getDB(ctx)
	result := db.Model(&MtOrder{}).
		Where("order_no = ?", orderNo).
		Updates(updates)

	if result.Error != nil {
		r.log.Errorf("更新订单状态失败: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.log.Warnf("更新订单状态未影响任何行: orderNo=%s", orderNo)
	}

	return nil
}

// 更新订单支付信息
func (r *orderRepo) UpdateOrderPayment(ctx context.Context, orderNo string, payType string, payTime time.Time) error {
	updates := map[string]interface{}{
		"pay_type": payType,
		"pay_time": payTime,
	}

	db := r.getDB(ctx)
	result := db.Model(&MtOrder{}).
		Where("order_no = ?", orderNo).
		Updates(updates)

	if result.Error != nil {
		r.log.Errorf("更新订单支付信息失败: %v", result.Error)
		return result.Error
	}

	return nil
}

// 查询用户订单列表
func (r *orderRepo) ListUserOrders(ctx context.Context, userID int64, page, pageSize int32) ([]*biz.MtOrder, int64, error) {
	var orders []MtOrder
	var total int64

	db := r.getDB(ctx)
	
	// 查询总数
	countResult := db.Model(&MtOrder{}).Where("user_id = ?", userID).Count(&total)
	if countResult.Error != nil {
		r.log.Errorf("查询订单总数失败: %v", countResult.Error)
		return nil, 0, countResult.Error
	}

	// 查询列表
	offset := (page - 1) * pageSize
	result := db.Where("user_id = ?", userID).
		Order("id DESC").
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&orders)

	if result.Error != nil {
		r.log.Errorf("查询订单列表失败: %v", result.Error)
		return nil, 0, result.Error
	}

	// 转换为业务模型
	bizOrders := make([]*biz.MtOrder, len(orders))
	for i, order := range orders {
		bizOrders[i] = r.toBizOrder(&order)
	}

	return bizOrders, total, nil
}

// 创建订单项
func (r *orderRepo) CreateOrderItems(ctx context.Context, items []*biz.MtOrderItem) error {
	if len(items) == 0 {
		return nil
	}

	// 转换为数据模型
	dataItems := make([]*MtOrderItem, len(items))
	for i, item := range items {
		dataItems[i] = r.toDataOrderItem(item)
	}

	db := r.getDB(ctx)
	result := db.Create(&dataItems)
	if result.Error != nil {
		r.log.Errorf("创建订单项失败: %v", result.Error)
		return result.Error
	}

	// 更新ID
	for i, dataItem := range dataItems {
		items[i].ID = dataItem.ID
	}

	return nil
}

// 获取订单项
func (r *orderRepo) GetOrderItems(ctx context.Context, orderID int64) ([]*biz.MtOrderItem, error) {
	var items []MtOrderItem
	db := r.getDB(ctx)
	result := db.Where("order_id = ?", orderID).
		Order("created_at ASC").
		Find(&items)

	if result.Error != nil {
		r.log.Errorf("查询订单项失败: %v", result.Error)
		return nil, result.Error
	}

	// 转换为业务模型
	bizItems := make([]*biz.MtOrderItem, len(items))
	for i, item := range items {
		bizItems[i] = r.toBizOrderItem(&item)
	}

	return bizItems, nil
}

// 事务支持
func (r *orderRepo) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.data.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建新的上下文，包含事务信息
		txCtx := context.WithValue(ctx, "tx", tx)
		return fn(txCtx)
	})
}

// 按状态查询订单列表
func (r *orderRepo) ListOrdersByStatus(ctx context.Context, status string, page, pageSize int32) ([]*biz.MtOrder, int64, error) {
	var orders []MtOrder
	var total int64

	db := r.getDB(ctx)
	
	// 查询总数
	countResult := db.Model(&MtOrder{}).Where("status = ?", status).Count(&total)
	if countResult.Error != nil {
		r.log.Errorf("查询订单总数失败: %v", countResult.Error)
		return nil, 0, countResult.Error
	}

	// 查询列表
	offset := (page - 1) * pageSize
	result := db.Where("status = ?", status).
		Order("id DESC").
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&orders)

	if result.Error != nil {
		r.log.Errorf("查询订单列表失败: %v", result.Error)
		return nil, 0, result.Error
	}

	// 转换为业务模型
	bizOrders := make([]*biz.MtOrder, len(orders))
	for i, order := range orders {
		bizOrders[i] = r.toBizOrder(&order)
	}

	return bizOrders, total, nil
}

// 统计指定状态的订单数量
func (r *orderRepo) CountOrdersByStatus(ctx context.Context, status string) (int64, error) {
	var count int64
	db := r.getDB(ctx)
	
	result := db.Model(&MtOrder{}).Where("status = ?", status).Count(&count)
	if result.Error != nil {
		r.log.Errorf("统计订单数量失败: %v", result.Error)
		return 0, result.Error
	}

	return count, nil
}

// 获取数据库连接（支持事务）
func (r *orderRepo) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return r.data.Db
}