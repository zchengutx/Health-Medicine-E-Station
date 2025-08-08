package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// 支付订单状态
const (
	PaymentStatusPending   = "pending"   // 待支付
	PaymentStatusPaid      = "paid"      // 已支付
	PaymentStatusCancelled = "cancelled" // 已取消
	PaymentStatusRefunded  = "refunded"  // 已退款
)

// 支付订单类型
const (
	OrderTypeDrug         = "drug_order"         // 药品订单
	OrderTypeConsultation = "consultation_order" // 咨询订单
)

// 支付订单模型
type PaymentOrder struct {
	ID           int64     `json:"id"`
	OrderID      string    `json:"order_id"`      // 订单号
	UserID       int32     `json:"user_id"`       // 用户ID
	BusinessID   string    `json:"business_id"`   // 业务ID
	OrderType    string    `json:"order_type"`    // 订单类型
	Subject      string    `json:"subject"`       // 商品标题
	Description  string    `json:"description"`   // 订单描述
	TotalAmount  string    `json:"total_amount"`  // 支付金额
	Status       string    `json:"status"`        // 支付状态
	PaymentURL   string    `json:"payment_url"`   // 支付链接
	TradeNo      string    `json:"trade_no"`      // 支付宝交易号
	PayTime      time.Time `json:"pay_time"`      // 支付时间
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
}

// 退款记录模型
type RefundRecord struct {
	ID           int64     `json:"id"`
	OrderID      string    `json:"order_id"`      // 订单号
	RefundID     string    `json:"refund_id"`     // 退款单号
	RefundAmount string    `json:"refund_amount"` // 退款金额
	RefundReason string    `json:"refund_reason"` // 退款原因
	RefundStatus string    `json:"refund_status"` // 退款状态
	RefundTime   time.Time `json:"refund_time"`   // 退款时间
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
}



// 支付仓储接口
type PaymentRepo interface {
	// 创建支付订单
	CreatePaymentOrder(ctx context.Context, order *PaymentOrder) error
	// 根据订单号查询支付订单
	GetPaymentOrderByOrderID(ctx context.Context, orderID string) (*PaymentOrder, error)
	// 根据用户ID查询支付订单列表
	GetPaymentOrdersByUserID(ctx context.Context, userID int32, page, pageSize int32) ([]*PaymentOrder, int64, error)
	// 更新支付订单状态
	UpdatePaymentOrderStatus(ctx context.Context, orderID, status, tradeNo string, payTime time.Time) error
	
	// 退款记录相关
	CreateRefundRecord(ctx context.Context, record *RefundRecord) error
	GetRefundRecordsByOrderID(ctx context.Context, orderID string) ([]*RefundRecord, error)
}

// 支付用例
type PaymentUsecase struct {
	repo PaymentRepo
	log  *log.Helper
}

// 创建支付用例
func NewPaymentUsecase(repo PaymentRepo, logger log.Logger) *PaymentUsecase {
	return &PaymentUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// 创建支付订单
func (uc *PaymentUsecase) CreatePaymentOrder(ctx context.Context, userID int32, subject, totalAmount, orderType, businessID, description string) (*PaymentOrder, error) {
	// 生成订单号
	orderID := fmt.Sprintf("PAY_%d_%d", userID, time.Now().UnixNano())

	// 创建支付订单
	order := &PaymentOrder{
		OrderID:     orderID,
		UserID:      userID,
		BusinessID:  businessID,
		OrderType:   orderType,
		Subject:     subject,
		Description: description,
		TotalAmount: totalAmount,
		Status:      PaymentStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 保存到数据库
	err := uc.repo.CreatePaymentOrder(ctx, order)
	if err != nil {
		uc.log.Errorf("创建支付订单失败: %v", err)
		return nil, err
	}

	uc.log.Infof("创建支付订单成功: orderID=%s, userID=%d, amount=%s", orderID, userID, totalAmount)
	return order, nil
}

// 查询支付订单
func (uc *PaymentUsecase) GetPaymentOrder(ctx context.Context, orderID string) (*PaymentOrder, error) {
	order, err := uc.repo.GetPaymentOrderByOrderID(ctx, orderID)
	if err != nil {
		uc.log.Errorf("查询支付订单失败: orderID=%s, error=%v", orderID, err)
		return nil, err
	}

	return order, nil
}

// 查询用户支付订单列表
func (uc *PaymentUsecase) GetUserPaymentOrders(ctx context.Context, userID int32, page, pageSize int32) ([]*PaymentOrder, int64, error) {
	orders, total, err := uc.repo.GetPaymentOrdersByUserID(ctx, userID, page, pageSize)
	if err != nil {
		uc.log.Errorf("查询用户支付订单列表失败: userID=%d, error=%v", userID, err)
		return nil, 0, err
	}

	return orders, total, nil
}

// 更新支付订单状态
func (uc *PaymentUsecase) UpdatePaymentOrderStatus(ctx context.Context, orderID, status, tradeNo string) error {
	var payTime time.Time
	if status == PaymentStatusPaid {
		payTime = time.Now()
	}

	err := uc.repo.UpdatePaymentOrderStatus(ctx, orderID, status, tradeNo, payTime)
	if err != nil {
		uc.log.Errorf("更新支付订单状态失败: orderID=%s, status=%s, error=%v", orderID, status, err)
		return err
	}

	uc.log.Infof("更新支付订单状态成功: orderID=%s, status=%s, tradeNo=%s", orderID, status, tradeNo)
	return nil
}

// 处理支付成功回调
func (uc *PaymentUsecase) HandlePaymentSuccess(ctx context.Context, orderID, tradeNo, totalAmount string) error {
	// 查询订单
	order, err := uc.repo.GetPaymentOrderByOrderID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("查询订单失败: %v", err)
	}

	// 验证金额
	if order.TotalAmount != totalAmount {
		return fmt.Errorf("订单金额不匹配: expected=%s, actual=%s", order.TotalAmount, totalAmount)
	}

	// 更新订单状态
	err = uc.UpdatePaymentOrderStatus(ctx, orderID, PaymentStatusPaid, tradeNo)
	if err != nil {
		return fmt.Errorf("更新订单状态失败: %v", err)
	}

	// 根据订单类型处理业务逻辑
	switch order.OrderType {
	case OrderTypeDrug:
		// 处理药品订单支付成功逻辑
		uc.log.Infof("药品订单支付成功: orderID=%s, businessID=%s", orderID, order.BusinessID)
		// TODO: 调用药品订单服务更新状态
	case OrderTypeConsultation:
		// 处理咨询订单支付成功逻辑
		uc.log.Infof("咨询订单支付成功: orderID=%s, businessID=%s", orderID, order.BusinessID)
		// TODO: 调用咨询订单服务更新状态
	}

	return nil
}

// 创建退款记录
func (uc *PaymentUsecase) CreateRefund(ctx context.Context, orderID, refundAmount, refundReason string) (*RefundRecord, error) {
	// 查询原订单
	order, err := uc.repo.GetPaymentOrderByOrderID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("查询原订单失败: %v", err)
	}

	if order.Status != PaymentStatusPaid {
		return nil, fmt.Errorf("订单状态不允许退款: status=%s", order.Status)
	}

	// 生成退款单号
	refundID := fmt.Sprintf("REFUND_%s_%d", orderID, time.Now().Unix())

	// 创建退款记录
	record := &RefundRecord{
		OrderID:      orderID,
		RefundID:     refundID,
		RefundAmount: refundAmount,
		RefundReason: refundReason,
		RefundStatus: "processing", // 处理中
		CreatedAt:    time.Now(),
	}

	// 保存退款记录
	err = uc.repo.CreateRefundRecord(ctx, record)
	if err != nil {
		uc.log.Errorf("创建退款记录失败: %v", err)
		return nil, err
	}

	uc.log.Infof("创建退款记录成功: refundID=%s, orderID=%s, amount=%s", refundID, orderID, refundAmount)
	return record, nil
}

// 查询退款记录
func (uc *PaymentUsecase) GetRefundRecords(ctx context.Context, orderID string) ([]*RefundRecord, error) {
	records, err := uc.repo.GetRefundRecordsByOrderID(ctx, orderID)
	if err != nil {
		uc.log.Errorf("查询退款记录失败: orderID=%s, error=%v", orderID, err)
		return nil, err
	}

	return records, nil
}