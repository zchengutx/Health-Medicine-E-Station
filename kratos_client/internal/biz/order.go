package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
)



// 订单主表模型
type MtOrder struct {
	ID            uint64          `json:"id"`
	OrderNo       string          `json:"order_no"`       // 订单编号
	UserID        int64           `json:"user_id"`        // 患者id
	UserName      string          `json:"user_name"`      // 患者姓名
	UserPhone     string          `json:"user_phone"`     // 患者电话
	DoctorID      int64           `json:"doctor_id"`      // 医生id
	DoctorName    string          `json:"doctor_name"`    // 医生姓名
	AddressID     int64           `json:"address_id"`     // 地址id
	AddressDetail string          `json:"address_detail"` // 地址详情
	TotalAmount   decimal.Decimal `json:"total_amount"`   // 总金额
	PayType       string          `json:"pay_type"`       // 支付方式
	Status           string          `json:"status"`            // 订单状态
	PayTime          *time.Time      `json:"pay_time"`          // 支付时间
	DrugTime         *time.Time      `json:"drug_time"`         // 配药时间
	SendTime         *time.Time      `json:"send_time"`         // 发货时间
	FinishTime       *time.Time      `json:"finish_time"`       // 完成时间
	CancelTime       *time.Time      `json:"cancel_time"`       // 取消时间
	UserCouponID     *int64          `json:"user_coupon_id"`    // 使用的优惠券ID
	DiscountAmount   decimal.Decimal `json:"discount_amount"`   // 优惠金额
	OriginalAmount   decimal.Decimal `json:"original_amount"`   // 原始金额
	Remark           string          `json:"remark"`            // 备注
}

// 订单项模型
type MtOrderItem struct {
	ID        int64           `json:"id"`
	OrderID   int64           `json:"order_id"`   // 订单ID
	DrugID    int64           `json:"drug_id"`    // 药品ID
	DrugName  string          `json:"drug_name"`  // 药品名称
	DrugSpec  string          `json:"drug_spec"`  // 药品规格
	Quantity  int32           `json:"quantity"`   // 数量
	Price     decimal.Decimal `json:"price"`      // 单价
	Subtotal  decimal.Decimal `json:"subtotal"`   // 小计
	CreatedAt time.Time       `json:"created_at"` // 创建时间
}

// 创建订单请求
type CreateOrderRequest struct {
	UserID        int64                `json:"user_id" validate:"required"`
	UserName      string               `json:"user_name" validate:"required"`
	UserPhone     string               `json:"user_phone" validate:"required"`
	DoctorID      int64                `json:"doctor_id"`
	DoctorName    string               `json:"doctor_name"`
	AddressID     int64                `json:"address_id" validate:"required"`
	AddressDetail string               `json:"address_detail" validate:"required"`
	Items         []*CreateOrderItem   `json:"items" validate:"required,min=1"`
	UserCouponID  *int64               `json:"user_coupon_id"` // 使用的优惠券ID
	Remark        string               `json:"remark"`
}

// 创建订单项请求
type CreateOrderItem struct {
	DrugID   int64 `json:"drug_id" validate:"required"`
	Quantity int32 `json:"quantity" validate:"required,min=1"`
}

// 支付信息
type PaymentInfo struct {
	PayType     string          `json:"pay_type" validate:"required,oneof=1 2 3"`
	Amount      decimal.Decimal `json:"amount" validate:"required"`
	TradeNo     string          `json:"trade_no"`
	PaymentTime time.Time       `json:"payment_time"`
}

// 订单详情响应
type OrderDetail struct {
	Order *MtOrder       `json:"order"`
	Items []*MtOrderItem `json:"items"`
}

// 订单摘要响应
type OrderSummary struct {
	OrderNo     string          `json:"order_no"`
	TotalAmount decimal.Decimal `json:"total_amount"`
	Status      string          `json:"status"`
	ItemCount   int32           `json:"item_count"`
}

// 订单仓储接口
type OrderRepo interface {
	// 订单CRUD操作
	CreateOrder(ctx context.Context, order *MtOrder) error
	GetOrderByID(ctx context.Context, id int64) (*MtOrder, error)
	GetOrderByOrderNo(ctx context.Context, orderNo string) (*MtOrder, error)
	UpdateOrderStatus(ctx context.Context, orderNo string, status string, timestamp time.Time) error
	UpdateOrderPayment(ctx context.Context, orderNo string, payType string, payTime time.Time) error
	ListUserOrders(ctx context.Context, userID int64, page, pageSize int32) ([]*MtOrder, int64, error)
	ListOrdersByStatus(ctx context.Context, status string, page, pageSize int32) ([]*MtOrder, int64, error)
	CountOrdersByStatus(ctx context.Context, status string) (int64, error)

	// 订单项操作
	CreateOrderItems(ctx context.Context, items []*MtOrderItem) error
	GetOrderItems(ctx context.Context, orderID int64) ([]*MtOrderItem, error)

	// 事务支持
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

// 药品库存仓储接口
type DrugInventoryRepo interface {
	CheckInventory(ctx context.Context, drugID int64, quantity int32) (bool, error)
	ReserveInventory(ctx context.Context, drugID int64, quantity int32) error
	ReduceInventory(ctx context.Context, drugID int64, quantity int32) error
	ReleaseReservedInventory(ctx context.Context, drugID int64, quantity int32) error
}

// 订单用例
type OrderUsecase struct {
	orderRepo     OrderRepo
	drugRepo      DrugRepo
	inventoryRepo DrugInventoryRepo
	log           *log.Helper
}

// 创建订单用例
func NewOrderUsecase(
	orderRepo OrderRepo,
	drugRepo DrugRepo,
	inventoryRepo DrugInventoryRepo,
	logger log.Logger,
) *OrderUsecase {
	return &OrderUsecase{
		orderRepo:     orderRepo,
		drugRepo:      drugRepo,
		inventoryRepo: inventoryRepo,
		log:           log.NewHelper(logger),
	}
}

// 创建订单
func (uc *OrderUsecase) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*MtOrder, error) {
	// 生成订单号
	orderNo := fmt.Sprintf("ORD_%d_%d", req.UserID, time.Now().UnixNano())

	// 验证药品信息并计算总金额
	var totalAmount decimal.Decimal
	var orderItems []*MtOrderItem

	for _, item := range req.Items {
		// 获取药品信息
		drug, err := uc.drugRepo.GetDrug(ctx, int32(item.DrugID))
		if err != nil {
			uc.log.Errorf("获取药品信息失败: drugID=%d, error=%v", item.DrugID, err)
			return nil, fmt.Errorf("药品不存在: drugID=%d", item.DrugID)
		}

		// 检查库存
		hasStock, err := uc.inventoryRepo.CheckInventory(ctx, item.DrugID, item.Quantity)
		if err != nil {
			uc.log.Errorf("检查库存失败: drugID=%d, error=%v", item.DrugID, err)
			return nil, fmt.Errorf("检查库存失败")
		}
		if !hasStock {
			return nil, fmt.Errorf("药品库存不足: %s", drug.DrugName)
		}

		// 计算小计
		price := decimal.NewFromFloat32(drug.Price)
		quantity := decimal.NewFromInt32(item.Quantity)
		subtotal := price.Mul(quantity)
		totalAmount = totalAmount.Add(subtotal)

		// 创建订单项
		orderItem := &MtOrderItem{
			DrugID:   item.DrugID,
			DrugName: drug.DrugName,
			DrugSpec: drug.Specification,
			Quantity: item.Quantity,
			Price:    price,
			Subtotal: subtotal,
		}
		orderItems = append(orderItems, orderItem)
	}

	// 创建订单
	order := &MtOrder{
		OrderNo:       orderNo,
		UserID:        req.UserID,
		UserName:      req.UserName,
		UserPhone:     req.UserPhone,
		DoctorID:      req.DoctorID,
		DoctorName:    req.DoctorName,
		AddressID:     req.AddressID,
		AddressDetail: req.AddressDetail,
		TotalAmount:   totalAmount,
		Status:        "1", // 待支付
		Remark:        req.Remark,
	}

	// 使用事务创建订单和订单项
	err := uc.orderRepo.WithTx(ctx, func(ctx context.Context) error {
		// 创建订单
		if err := uc.orderRepo.CreateOrder(ctx, order); err != nil {
			return fmt.Errorf("创建订单失败: %v", err)
		}

		// 设置订单项的订单ID
		for _, item := range orderItems {
			item.OrderID = int64(order.ID)
		}

		// 创建订单项
		if err := uc.orderRepo.CreateOrderItems(ctx, orderItems); err != nil {
			return fmt.Errorf("创建订单项失败: %v", err)
		}

		// 预留库存
		for _, item := range req.Items {
			if err := uc.inventoryRepo.ReserveInventory(ctx, item.DrugID, item.Quantity); err != nil {
				return fmt.Errorf("预留库存失败: %v", err)
			}
		}

		return nil
	})

	if err != nil {
		uc.log.Errorf("创建订单事务失败: %v", err)
		return nil, err
	}

	uc.log.Infof("创建订单成功: orderNo=%s, userID=%d, amount=%s", orderNo, req.UserID, totalAmount.String())
	return order, nil
}

// 处理支付
func (uc *OrderUsecase) ProcessPayment(ctx context.Context, orderNo string, paymentInfo *PaymentInfo) error {
	// 查询订单
	order, err := uc.orderRepo.GetOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return fmt.Errorf("查询订单失败: %v", err)
	}

	if order == nil {
		return fmt.Errorf("订单不存在: %s", orderNo)
	}

	if order.Status != "1" { // 待支付
		return fmt.Errorf("订单状态不允许支付: %s", order.Status)
	}

	// 验证支付金额
	if !paymentInfo.Amount.Equal(order.TotalAmount) {
		return fmt.Errorf("支付金额不匹配: expected=%s, actual=%s", 
			order.TotalAmount.String(), paymentInfo.Amount.String())
	}

	// 使用事务处理支付
	err = uc.orderRepo.WithTx(ctx, func(ctx context.Context) error {
		// 更新订单支付信息
		if err := uc.orderRepo.UpdateOrderPayment(ctx, orderNo, paymentInfo.PayType, paymentInfo.PaymentTime); err != nil {
			return fmt.Errorf("更新订单支付信息失败: %v", err)
		}

		// 更新订单状态为已支付
		if err := uc.orderRepo.UpdateOrderStatus(ctx, orderNo, "2", paymentInfo.PaymentTime); err != nil { // 已支付
			return fmt.Errorf("更新订单状态失败: %v", err)
		}

		// 获取订单项并减少库存
		orderItems, err := uc.orderRepo.GetOrderItems(ctx, int64(order.ID))
		if err != nil {
			return fmt.Errorf("获取订单项失败: %v", err)
		}

		for _, item := range orderItems {
			if err := uc.inventoryRepo.ReduceInventory(ctx, item.DrugID, item.Quantity); err != nil {
				return fmt.Errorf("减少库存失败: drugID=%d, error=%v", item.DrugID, err)
			}
		}

		return nil
	})

	if err != nil {
		uc.log.Errorf("处理支付事务失败: orderNo=%s, error=%v", orderNo, err)
		return err
	}

	uc.log.Infof("处理支付成功: orderNo=%s, amount=%s", orderNo, paymentInfo.Amount.String())
	return nil
}

// 更新订单状态
func (uc *OrderUsecase) UpdateOrderStatus(ctx context.Context, orderNo string, status string) error {
	// 验证状态转换的合法性
	currentOrder, err := uc.orderRepo.GetOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return fmt.Errorf("查询订单失败: %v", err)
	}

	if currentOrder == nil {
		return fmt.Errorf("订单不存在: %s", orderNo)
	}

	// 验证状态转换
	if !uc.isValidStatusTransition(currentOrder.Status, status) {
		return fmt.Errorf("无效的状态转换: %s -> %s", currentOrder.Status, status)
	}

	// 更新状态
	timestamp := time.Now()
	err = uc.orderRepo.UpdateOrderStatus(ctx, orderNo, status, timestamp)
	if err != nil {
		uc.log.Errorf("更新订单状态失败: orderNo=%s, status=%s, error=%v", orderNo, status, err)
		return err
	}

	uc.log.Infof("更新订单状态成功: orderNo=%s, status=%s", orderNo, status)
	return nil
}

// 取消订单
func (uc *OrderUsecase) CancelOrder(ctx context.Context, orderNo string, reason string) error {
	order, err := uc.orderRepo.GetOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return fmt.Errorf("查询订单失败: %v", err)
	}

	if order == nil {
		return fmt.Errorf("订单不存在: %s", orderNo)
	}

	// 只有待支付状态的订单可以取消
	if order.Status != "1" { // 待支付
		return fmt.Errorf("订单状态不允许取消: %s", order.Status)
	}

	// 使用事务取消订单
	err = uc.orderRepo.WithTx(ctx, func(ctx context.Context) error {
		// 更新订单状态为已取消
		if err := uc.orderRepo.UpdateOrderStatus(ctx, orderNo, "6", time.Now()); err != nil { // 已取消
			return fmt.Errorf("更新订单状态失败: %v", err)
		}

		// 释放预留库存
		orderItems, err := uc.orderRepo.GetOrderItems(ctx, int64(order.ID))
		if err != nil {
			return fmt.Errorf("获取订单项失败: %v", err)
		}

		for _, item := range orderItems {
			if err := uc.inventoryRepo.ReleaseReservedInventory(ctx, item.DrugID, item.Quantity); err != nil {
				return fmt.Errorf("释放预留库存失败: drugID=%d, error=%v", item.DrugID, err)
			}
		}

		return nil
	})

	if err != nil {
		uc.log.Errorf("取消订单事务失败: orderNo=%s, error=%v", orderNo, err)
		return err
	}

	uc.log.Infof("取消订单成功: orderNo=%s, reason=%s", orderNo, reason)
	return nil
}

// 获取订单详情
func (uc *OrderUsecase) GetOrder(ctx context.Context, orderNo string) (*OrderDetail, error) {
	order, err := uc.orderRepo.GetOrderByOrderNo(ctx, orderNo)
	if err != nil {
		uc.log.Errorf("查询订单失败: orderNo=%s, error=%v", orderNo, err)
		return nil, err
	}

	if order == nil {
		return nil, fmt.Errorf("订单不存在: %s", orderNo)
	}

	items, err := uc.orderRepo.GetOrderItems(ctx, int64(order.ID))
	if err != nil {
		uc.log.Errorf("查询订单项失败: orderID=%d, error=%v", order.ID, err)
		return nil, err
	}

	return &OrderDetail{
		Order: order,
		Items: items,
	}, nil
}

// 获取用户订单列表
func (uc *OrderUsecase) ListUserOrders(ctx context.Context, userID int64, page, pageSize int32) ([]*OrderSummary, int64, error) {
	orders, total, err := uc.orderRepo.ListUserOrders(ctx, userID, page, pageSize)
	if err != nil {
		uc.log.Errorf("查询用户订单列表失败: userID=%d, error=%v", userID, err)
		return nil, 0, err
	}

	// 转换为摘要格式
	summaries := make([]*OrderSummary, len(orders))
	for i, order := range orders {
		// 获取订单项数量
		items, err := uc.orderRepo.GetOrderItems(ctx, int64(order.ID))
		if err != nil {
			uc.log.Warnf("获取订单项数量失败: orderID=%d, error=%v", order.ID, err)
		}

		summaries[i] = &OrderSummary{
			OrderNo:     order.OrderNo,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
			ItemCount:   int32(len(items)),
		}
	}

	return summaries, total, nil
}

// 获取指定状态的订单列表
func (uc *OrderUsecase) ListOrdersByStatus(ctx context.Context, status string, page, pageSize int32) ([]*OrderSummary, int64, error) {
	orders, total, err := uc.orderRepo.ListOrdersByStatus(ctx, status, page, pageSize)
	if err != nil {
		uc.log.Errorf("查询状态订单列表失败: status=%s, error=%v", status, err)
		return nil, 0, err
	}

	// 转换为摘要格式
	summaries := make([]*OrderSummary, len(orders))
	for i, order := range orders {
		// 获取订单项数量
		items, err := uc.orderRepo.GetOrderItems(ctx, int64(order.ID))
		if err != nil {
			uc.log.Warnf("获取订单项数量失败: orderID=%d, error=%v", order.ID, err)
		}

		summaries[i] = &OrderSummary{
			OrderNo:     order.OrderNo,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
			ItemCount:   int32(len(items)),
		}
	}

	return summaries, total, nil
}

// 获取订单统计信息
func (uc *OrderUsecase) GetOrderStats(ctx context.Context) (map[string]int64, error) {
	stats := make(map[string]int64)
	
	// 统计各状态订单数量
	statuses := []string{"1", "2", "3", "4", "5", "6"} // 待支付、已支付、配药中、已发货、已完成、已取消
	statusNames := []string{"pending", "paid", "preparing", "shipped", "completed", "cancelled"}
	
	for i, status := range statuses {
		count, err := uc.orderRepo.CountOrdersByStatus(ctx, status)
		if err != nil {
			uc.log.Errorf("统计订单状态失败: status=%s, error=%v", status, err)
			return nil, err
		}
		stats[statusNames[i]] = count
	}
	
	return stats, nil
}

// 验证状态转换是否合法 - 简化版本，具体业务规则由后台管理系统维护
func (uc *OrderUsecase) isValidStatusTransition(from, to string) bool {
	// 简单验证：不允许从完成或取消状态转换
	if from == "5" || from == "6" { // 已完成或已取消
		return false
	}
	return true
}