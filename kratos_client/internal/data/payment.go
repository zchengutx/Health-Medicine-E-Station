package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos_client/internal/biz"
)

// 支付订单数据模型
type PaymentOrder struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID      string    `gorm:"uniqueIndex;size:64;not null" json:"order_id"`
	UserID       int32     `gorm:"index;not null" json:"user_id"`
	BusinessID   string    `gorm:"size:64;not null" json:"business_id"`
	OrderType    string    `gorm:"size:32;not null" json:"order_type"`
	Subject      string    `gorm:"size:256;not null" json:"subject"`
	Description  string    `gorm:"type:text" json:"description"`
	TotalAmount  string    `gorm:"size:16;not null" json:"total_amount"`
	Status       string    `gorm:"size:32;not null;default:'pending'" json:"status"`
	PaymentURL   string    `gorm:"type:text" json:"payment_url"`
	TradeNo      string    `gorm:"size:64" json:"trade_no"`
	PayTime      time.Time `json:"pay_time"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// 表名
func (PaymentOrder) TableName() string {
	return "payment_orders"
}

// 退款记录数据模型
type RefundRecord struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID      string    `gorm:"index;size:64;not null" json:"order_id"`
	RefundID     string    `gorm:"uniqueIndex;size:64;not null" json:"refund_id"`
	RefundAmount string    `gorm:"size:16;not null" json:"refund_amount"`
	RefundReason string    `gorm:"type:text" json:"refund_reason"`
	RefundStatus string    `gorm:"size:32;not null;default:'processing'" json:"refund_status"`
	RefundTime   time.Time `json:"refund_time"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// 表名
func (RefundRecord) TableName() string {
	return "refund_records"
}

// 支付仓储实现
type paymentRepo struct {
	data *Data
	log  *log.Helper
}

// 创建支付仓储
func NewPaymentRepo(data *Data, logger log.Logger) biz.PaymentRepo {
	return &paymentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 转换业务模型到数据模型
func (r *paymentRepo) toBizPaymentOrder(po *PaymentOrder) *biz.PaymentOrder {
	return &biz.PaymentOrder{
		ID:          po.ID,
		OrderID:     po.OrderID,
		UserID:      po.UserID,
		BusinessID:  po.BusinessID,
		OrderType:   po.OrderType,
		Subject:     po.Subject,
		Description: po.Description,
		TotalAmount: po.TotalAmount,
		Status:      po.Status,
		PaymentURL:  po.PaymentURL,
		TradeNo:     po.TradeNo,
		PayTime:     po.PayTime,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
	}
}

// 转换业务模型到数据模型
func (r *paymentRepo) toDataPaymentOrder(bo *biz.PaymentOrder) *PaymentOrder {
	return &PaymentOrder{
		ID:          bo.ID,
		OrderID:     bo.OrderID,
		UserID:      bo.UserID,
		BusinessID:  bo.BusinessID,
		OrderType:   bo.OrderType,
		Subject:     bo.Subject,
		Description: bo.Description,
		TotalAmount: bo.TotalAmount,
		Status:      bo.Status,
		PaymentURL:  bo.PaymentURL,
		TradeNo:     bo.TradeNo,
		PayTime:     bo.PayTime,
		CreatedAt:   bo.CreatedAt,
		UpdatedAt:   bo.UpdatedAt,
	}
}

// 转换退款记录
func (r *paymentRepo) toBizRefundRecord(rr *RefundRecord) *biz.RefundRecord {
	return &biz.RefundRecord{
		ID:           rr.ID,
		OrderID:      rr.OrderID,
		RefundID:     rr.RefundID,
		RefundAmount: rr.RefundAmount,
		RefundReason: rr.RefundReason,
		RefundStatus: rr.RefundStatus,
		RefundTime:   rr.RefundTime,
		CreatedAt:    rr.CreatedAt,
	}
}

// 转换退款记录
func (r *paymentRepo) toDataRefundRecord(br *biz.RefundRecord) *RefundRecord {
	return &RefundRecord{
		ID:           br.ID,
		OrderID:      br.OrderID,
		RefundID:     br.RefundID,
		RefundAmount: br.RefundAmount,
		RefundReason: br.RefundReason,
		RefundStatus: br.RefundStatus,
		RefundTime:   br.RefundTime,
		CreatedAt:    br.CreatedAt,
	}
}

// 创建支付订单
func (r *paymentRepo) CreatePaymentOrder(ctx context.Context, order *biz.PaymentOrder) error {
	po := r.toDataPaymentOrder(order)
	
	result := r.data.Db.WithContext(ctx).Create(po)
	if result.Error != nil {
		r.log.Errorf("创建支付订单失败: %v", result.Error)
		return result.Error
	}

	// 更新ID
	order.ID = po.ID
	return nil
}

// 根据订单号查询支付订单
func (r *paymentRepo) GetPaymentOrderByOrderID(ctx context.Context, orderID string) (*biz.PaymentOrder, error) {
	var po PaymentOrder
	result := r.data.Db.WithContext(ctx).Where("order_id = ?", orderID).First(&po)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.log.Errorf("查询支付订单失败: %v", result.Error)
		return nil, result.Error
	}

	return r.toBizPaymentOrder(&po), nil
}

// 根据用户ID查询支付订单列表
func (r *paymentRepo) GetPaymentOrdersByUserID(ctx context.Context, userID int32, page, pageSize int32) ([]*biz.PaymentOrder, int64, error) {
	var orders []PaymentOrder
	var total int64

	// 查询总数
	countResult := r.data.Db.WithContext(ctx).Model(&PaymentOrder{}).Where("user_id = ?", userID).Count(&total)
	if countResult.Error != nil {
		r.log.Errorf("查询支付订单总数失败: %v", countResult.Error)
		return nil, 0, countResult.Error
	}

	// 查询列表
	offset := (page - 1) * pageSize
	result := r.data.Db.WithContext(ctx).Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&orders)
	
	if result.Error != nil {
		r.log.Errorf("查询支付订单列表失败: %v", result.Error)
		return nil, 0, result.Error
	}

	// 转换为业务模型
	bizOrders := make([]*biz.PaymentOrder, len(orders))
	for i, order := range orders {
		bizOrders[i] = r.toBizPaymentOrder(&order)
	}

	return bizOrders, total, nil
}

// 更新支付订单状态
func (r *paymentRepo) UpdatePaymentOrderStatus(ctx context.Context, orderID, status, tradeNo string, payTime time.Time) error {
	updates := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}

	if tradeNo != "" {
		updates["trade_no"] = tradeNo
	}

	if !payTime.IsZero() {
		updates["pay_time"] = payTime
	}

	result := r.data.Db.WithContext(ctx).Model(&PaymentOrder{}).
		Where("order_id = ?", orderID).
		Updates(updates)

	if result.Error != nil {
		r.log.Errorf("更新支付订单状态失败: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.log.Warnf("更新支付订单状态未影响任何行: orderID=%s", orderID)
	}

	return nil
}

// 创建退款记录
func (r *paymentRepo) CreateRefundRecord(ctx context.Context, record *biz.RefundRecord) error {
	rr := r.toDataRefundRecord(record)
	
	result := r.data.Db.WithContext(ctx).Create(rr)
	if result.Error != nil {
		r.log.Errorf("创建退款记录失败: %v", result.Error)
		return result.Error
	}

	// 更新ID
	record.ID = rr.ID
	return nil
}

// 根据订单号查询退款记录
func (r *paymentRepo) GetRefundRecordsByOrderID(ctx context.Context, orderID string) ([]*biz.RefundRecord, error) {
	var records []RefundRecord
	result := r.data.Db.WithContext(ctx).Where("order_id = ?", orderID).
		Order("created_at DESC").
		Find(&records)
	
	if result.Error != nil {
		r.log.Errorf("查询退款记录失败: %v", result.Error)
		return nil, result.Error
	}

	// 转换为业务模型
	bizRecords := make([]*biz.RefundRecord, len(records))
	for i, record := range records {
		bizRecords[i] = r.toBizRefundRecord(&record)
	}

	return bizRecords, nil
}