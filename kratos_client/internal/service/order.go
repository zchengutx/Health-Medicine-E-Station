package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"

	pb "kratos_client/api/order/v1"
	"kratos_client/internal/biz"
)

// OrderService 订单服务
type OrderService struct {
	pb.UnimplementedOrderServiceServer

	orderUc *biz.OrderUsecase
	log     *log.Helper
}

// NewOrderService 创建订单服务
func NewOrderService(orderUc *biz.OrderUsecase, logger log.Logger) *OrderService {
	return &OrderService{
		orderUc: orderUc,
		log:     log.NewHelper(logger),
	}
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	// 转换请求参数
	items := make([]*biz.CreateOrderItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = &biz.CreateOrderItem{
			DrugID:   item.DrugId,
			Quantity: item.Quantity,
		}
	}

	createReq := &biz.CreateOrderRequest{
		UserID:        req.UserId,
		UserName:      req.UserName,
		UserPhone:     req.UserPhone,
		DoctorID:      req.DoctorId,
		DoctorName:    req.DoctorName,
		AddressID:     req.AddressId,
		AddressDetail: req.AddressDetail,
		Items:         items,
		Remark:        req.Remark,
	}

	// 创建订单
	order, err := s.orderUc.CreateOrder(ctx, createReq)
	if err != nil {
		s.log.Errorf("创建订单失败: %v", err)
		return &pb.CreateOrderReply{
			Message: err.Error(),
		}, err
	}

	return &pb.CreateOrderReply{
		OrderNo:     order.OrderNo,
		TotalAmount: order.TotalAmount.String(),
		Status:      order.Status,
		Message:     "订单创建成功",
	}, nil
}

// GetOrder 获取订单详情
func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderReply, error) {
	orderDetail, err := s.orderUc.GetOrder(ctx, req.OrderNo)
	if err != nil {
		s.log.Errorf("获取订单详情失败: %v", err)
		return nil, err
	}

	// 转换订单信息
	order := &pb.Order{
		Id:            orderDetail.Order.ID,
		OrderNo:       orderDetail.Order.OrderNo,
		UserId:        orderDetail.Order.UserID,
		UserName:      orderDetail.Order.UserName,
		UserPhone:     orderDetail.Order.UserPhone,
		DoctorId:      orderDetail.Order.DoctorID,
		DoctorName:    orderDetail.Order.DoctorName,
		AddressId:     orderDetail.Order.AddressID,
		AddressDetail: orderDetail.Order.AddressDetail,
		TotalAmount:   orderDetail.Order.TotalAmount.String(),
		PayType:       orderDetail.Order.PayType,
		Status:        orderDetail.Order.Status,
		Remark:        orderDetail.Order.Remark,
	}

	// 转换时间字段
	if orderDetail.Order.PayTime != nil {
		order.PayTime = orderDetail.Order.PayTime.Format(time.RFC3339)
	}
	if orderDetail.Order.DrugTime != nil {
		order.DrugTime = orderDetail.Order.DrugTime.Format(time.RFC3339)
	}
	if orderDetail.Order.SendTime != nil {
		order.SendTime = orderDetail.Order.SendTime.Format(time.RFC3339)
	}
	if orderDetail.Order.FinishTime != nil {
		order.FinishTime = orderDetail.Order.FinishTime.Format(time.RFC3339)
	}
	if orderDetail.Order.CancelTime != nil {
		order.CancelTime = orderDetail.Order.CancelTime.Format(time.RFC3339)
	}

	// 转换订单项
	items := make([]*pb.OrderItemDetail, len(orderDetail.Items))
	for i, item := range orderDetail.Items {
		items[i] = &pb.OrderItemDetail{
			Id:       item.ID,
			OrderId:  item.OrderID,
			DrugId:   item.DrugID,
			DrugName: item.DrugName,
			DrugSpec: item.DrugSpec,
			Quantity: item.Quantity,
			Price:    item.Price.String(),
			Subtotal: item.Subtotal.String(),
		}
	}

	return &pb.GetOrderReply{
		Order: order,
		Items: items,
	}, nil
}

// ListUserOrders 获取用户订单列表
func (s *OrderService) ListUserOrders(ctx context.Context, req *pb.ListUserOrdersRequest) (*pb.ListUserOrdersReply, error) {
	// 设置默认分页参数
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	summaries, total, err := s.orderUc.ListUserOrders(ctx, req.UserId, page, pageSize)
	if err != nil {
		s.log.Errorf("获取用户订单列表失败: %v", err)
		return nil, err
	}

	// 转换订单摘要
	orders := make([]*pb.OrderSummary, len(summaries))
	for i, summary := range summaries {
		orders[i] = &pb.OrderSummary{
			OrderNo:     summary.OrderNo,
			TotalAmount: summary.TotalAmount.String(),
			Status:      summary.Status,
			ItemCount:   summary.ItemCount,
		}
	}

	return &pb.ListUserOrdersReply{
		Orders: orders,
		Total:  total,
	}, nil
}

// ProcessPayment 处理支付
func (s *OrderService) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentReply, error) {
	// 解析支付金额
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		s.log.Errorf("解析支付金额失败: %v", err)
		return &pb.ProcessPaymentReply{
			Success: false,
			Message: "支付金额格式错误",
		}, err
	}

	// 构造支付信息
	paymentInfo := &biz.PaymentInfo{
		PayType:     req.PayType,
		Amount:      amount,
		TradeNo:     req.TradeNo,
		PaymentTime: time.Now(),
	}

	// 处理支付
	err = s.orderUc.ProcessPayment(ctx, req.OrderNo, paymentInfo)
	if err != nil {
		s.log.Errorf("处理支付失败: %v", err)
		return &pb.ProcessPaymentReply{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.ProcessPaymentReply{
		Success: true,
		Message: "支付处理成功",
	}, nil
}