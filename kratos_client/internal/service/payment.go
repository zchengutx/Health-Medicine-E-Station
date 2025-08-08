package service

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	pb "kratos_client/api/payment/v1"
	"kratos_client/comment"
	"kratos_client/internal/biz"
	"kratos_client/internal/data"
)

// 支付服务
type PaymentService struct {
	pb.UnimplementedPaymentServer
	uc   *biz.PaymentUsecase
	data *data.Data
}

// 创建支付服务
func NewPaymentService(uc *biz.PaymentUsecase, data *data.Data) *PaymentService {
	return &PaymentService{
		uc:   uc,
		data: data,
	}
}

// 创建支付订单
func (s *PaymentService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentReply, error) {
	// 解析JWT token获取用户ID
	claims, errMsg := comment.GetToken(req.Token)
	if claims == nil || errMsg != "" {
		return &pb.CreatePaymentReply{
			Code:    401,
			Message: "token无效: " + errMsg,
		}, nil
	}

	userIDFloat, ok := claims["user"].(float64)
	if !ok {
		return &pb.CreatePaymentReply{
			Code:    401,
			Message: "token中用户ID格式错误",
		}, nil
	}
	userID := int32(userIDFloat)

	// 验证参数
	if req.Subject == "" || req.TotalAmount == "" || req.OrderType == "" || req.BusinessId == "" {
		return &pb.CreatePaymentReply{
			Code:    400,
			Message: "参数不完整",
		}, nil
	}

	// 验证金额格式
	if _, err := strconv.ParseFloat(req.TotalAmount, 64); err != nil {
		return &pb.CreatePaymentReply{
			Code:    400,
			Message: "金额格式错误",
		}, nil
	}

	// 创建支付订单
	order, err := s.uc.CreatePaymentOrder(ctx, userID, req.Subject, req.TotalAmount, req.OrderType, req.BusinessId, req.Description)
	if err != nil {
		return &pb.CreatePaymentReply{
			Code:    500,
			Message: "创建支付订单失败: " + err.Error(),
		}, nil
	}

	// 创建支付宝支付链接
	paymentURL, err := comment.CreateAlipayOrder(order.OrderID, req.Subject, req.TotalAmount, strconv.Itoa(int(userID)))
	if err != nil {
		return &pb.CreatePaymentReply{
			Code:    500,
			Message: "创建支付链接失败: " + err.Error(),
		}, nil
	}

	return &pb.CreatePaymentReply{
		Code:       0,
		Message:    "success",
		PaymentUrl: paymentURL,
		OrderId:    order.OrderID,
	}, nil
}

// 支付回调通知
func (s *PaymentService) PaymentNotify(ctx context.Context, req *pb.PaymentNotifyRequest) (*pb.PaymentNotifyReply, error) {
	// 验证签名
	ok, err := comment.VerifyAlipayCallback(req.Params)
	if err != nil {
		return &pb.PaymentNotifyReply{Result: "fail"}, nil
	}

	if !ok {
		return &pb.PaymentNotifyReply{Result: "fail"}, nil
	}

	// 解析回调结果
	result := comment.ParseAlipayCallback(req.Params)
	
	// 如果支付成功，更新订单状态
	if result.Success {
		err = s.uc.HandlePaymentSuccess(ctx, result.OrderID, result.TradeNo, result.TotalAmount)
		if err != nil {
			// 记录错误但仍返回success，避免支付宝重复通知
			fmt.Printf("处理支付成功回调失败: %v\n", err)
		}
	}

	return &pb.PaymentNotifyReply{Result: "success"}, nil
}

// 查询支付状态
func (s *PaymentService) QueryPayment(ctx context.Context, req *pb.QueryPaymentRequest) (*pb.QueryPaymentReply, error) {
	// 解析JWT token获取用户ID
	claims, errMsg := comment.GetToken(req.Token)
	if claims == nil || errMsg != "" {
		return &pb.QueryPaymentReply{
			Code:    401,
			Message: "token无效: " + errMsg,
		}, nil
	}

	userIDFloat, ok := claims["user"].(float64)
	if !ok {
		return &pb.QueryPaymentReply{
			Code:    401,
			Message: "token中用户ID格式错误",
		}, nil
	}
	userID := int32(userIDFloat)

	// 查询订单
	order, err := s.uc.GetPaymentOrder(ctx, req.OrderId)
	if err != nil {
		return &pb.QueryPaymentReply{
			Code:    500,
			Message: "查询订单失败: " + err.Error(),
		}, nil
	}

	if order == nil {
		return &pb.QueryPaymentReply{
			Code:    404,
			Message: "订单不存在",
		}, nil
	}

	// 验证订单所有者
	if order.UserID != userID {
		return &pb.QueryPaymentReply{
			Code:    403,
			Message: "无权限查询此订单",
		}, nil
	}

	// 如果订单状态为待支付，查询支付宝订单状态
	if order.Status == biz.PaymentStatusPending {
		alipayResp, err := comment.QueryAlipayOrder(order.OrderID)
		if err == nil && alipayResp != nil {
			// 根据支付宝返回的状态更新本地订单状态
			if alipayResp.TradeStatus == "TRADE_SUCCESS" || alipayResp.TradeStatus == "TRADE_FINISHED" {
				s.uc.UpdatePaymentOrderStatus(ctx, order.OrderID, biz.PaymentStatusPaid, alipayResp.TradeNo)
				order.Status = biz.PaymentStatusPaid
				order.TradeNo = alipayResp.TradeNo
			}
		}
	}

	// 构造返回结果
	paymentInfo := &pb.PaymentInfo{
		OrderId:     order.OrderID,
		TradeNo:     order.TradeNo,
		TotalAmount: order.TotalAmount,
		TradeStatus: order.Status,
		Subject:     order.Subject,
	}

	if !order.PayTime.IsZero() {
		paymentInfo.PayTime = order.PayTime.Format("2006-01-02 15:04:05")
	}

	return &pb.QueryPaymentReply{
		Code:        0,
		Message:     "success",
		PaymentInfo: paymentInfo,
	}, nil
}

// 支付返回页面
func (s *PaymentService) PaymentReturn(ctx context.Context, req *pb.PaymentReturnRequest) (*pb.PaymentReturnReply, error) {
	// 验证签名
	ok, err := comment.VerifyAlipayCallback(req.Params)
	if err != nil {
		return &pb.PaymentReturnReply{
			Code:        500,
			Message:     "验证签名失败",
			RedirectUrl: "/payment/error",
		}, nil
	}

	if !ok {
		return &pb.PaymentReturnReply{
			Code:        400,
			Message:     "签名验证失败",
			RedirectUrl: "/payment/error",
		}, nil
	}

	// 解析回调结果
	result := comment.ParseAlipayCallback(req.Params)
	
	// 构造重定向URL
	redirectURL := "/payment/result"
	params := url.Values{}
	params.Set("order_id", result.OrderID)
	params.Set("success", strconv.FormatBool(result.Success))
	params.Set("message", result.Message)
	
	if result.Success {
		params.Set("trade_no", result.TradeNo)
		params.Set("amount", result.TotalAmount)
	}

	redirectURL += "?" + params.Encode()

	return &pb.PaymentReturnReply{
		Code:        0,
		Message:     "success",
		RedirectUrl: redirectURL,
	}, nil
}

// 申请退款 - 暂时注释掉，因为proto定义缺失
/*
func (s *PaymentService) RefundPayment(ctx context.Context, req *pb.RefundPaymentRequest) (*pb.RefundPaymentReply, error) {
	// 解析JWT token获取用户ID
	claims, errMsg := comment.GetToken(req.Token)
	if claims == nil || errMsg != "" {
		return &pb.RefundPaymentReply{
			Code:    401,
			Message: "token无效: " + errMsg,
		}, nil
	}

	userIDFloat, ok := claims["user"].(float64)
	if !ok {
		return &pb.RefundPaymentReply{
			Code:    401,
			Message: "token中用户ID格式错误",
		}, nil
	}
	userID := int32(userIDFloat)

	// 查询订单
	order, err := s.uc.GetPaymentOrder(ctx, req.OrderId)
	if err != nil {
		return &pb.RefundPaymentReply{
			Code:    500,
			Message: "查询订单失败: " + err.Error(),
		}, nil
	}

	if order == nil {
		return &pb.RefundPaymentReply{
			Code:    404,
			Message: "订单不存在",
		}, nil
	}

	// 验证订单所有者
	if order.UserID != userID {
		return &pb.RefundPaymentReply{
			Code:    403,
			Message: "无权限操作此订单",
		}, nil
	}

	// 验证订单状态
	if order.Status != biz.PaymentStatusPaid {
		return &pb.RefundPaymentReply{
			Code:    400,
			Message: "订单状态不允许退款",
		}, nil
	}

	// 验证退款金额
	if _, err := strconv.ParseFloat(req.RefundAmount, 64); err != nil {
		return &pb.RefundPaymentReply{
			Code:    400,
			Message: "退款金额格式错误",
		}, nil
	}

	// 创建退款记录
	refundRecord, err := s.uc.CreateRefund(ctx, req.OrderId, req.RefundAmount, req.RefundReason)
	if err != nil {
		return &pb.RefundPaymentReply{
			Code:    500,
			Message: "创建退款记录失败: " + err.Error(),
		}, nil
	}

	// 调用支付宝退款接口
	alipayResp, err := comment.RefundAlipayOrder(req.OrderId, req.RefundAmount, req.RefundReason)
	if err != nil {
		return &pb.RefundPaymentReply{
			Code:    500,
			Message: "申请退款失败: " + err.Error(),
		}, nil
	}

	// 构造返回结果
	refundInfo := &pb.RefundInfo{
		RefundId:     refundRecord.RefundID,
		RefundAmount: req.RefundAmount,
		RefundStatus: "success",
	}

	if alipayResp != nil && alipayResp.Content.GmtRefundPay != "" {
		refundInfo.RefundTime = alipayResp.Content.GmtRefundPay
	}

	return &pb.RefundPaymentReply{
		Code:       0,
		Message:    "退款申请成功",
		RefundInfo: refundInfo,
	}, nil
}
*/

// 解析表单参数为map
func parseFormParams(formData string) map[string]string {
	params := make(map[string]string)
	pairs := strings.Split(formData, "&")
	
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			key, _ := url.QueryUnescape(kv[0])
			value, _ := url.QueryUnescape(kv[1])
			params[key] = value
		}
	}
	
	return params
}