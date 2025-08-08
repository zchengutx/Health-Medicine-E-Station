package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"

	pb "kratos_client/api/coupon/v1"
	"kratos_client/internal/biz"
)

// CouponService 优惠券服务
type CouponService struct {
	pb.UnimplementedCouponServiceServer

	couponUc *biz.CouponUsecase
	log      *log.Helper
}

// NewCouponService 创建优惠券服务
func NewCouponService(couponUc *biz.CouponUsecase, logger log.Logger) *CouponService {
	return &CouponService{
		couponUc: couponUc,
		log:      log.NewHelper(logger),
	}
}

// ListCoupons 获取优惠券列表
func (s *CouponService) ListCoupons(ctx context.Context, req *pb.ListCouponsRequest) (*pb.ListCouponsReply, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	coupons, total, err := s.couponUc.ListCoupons(ctx, req.StoreId, page, pageSize)
	if err != nil {
		s.log.Errorf("获取优惠券列表失败: %v", err)
		return nil, err
	}

	// 转换为响应格式
	pbCoupons := make([]*pb.Coupon, len(coupons))
	for i, coupon := range coupons {
		pbCoupons[i] = &pb.Coupon{
			Id:             coupon.ID,
			DiscountName:   coupon.DiscountName,
			Classify:       coupon.Classify,
			StoreId:        coupon.StoreID,
			DiscountAmount: coupon.DiscountAmount.String(),
			MinOrderAmount: coupon.MinOrderAmount.String(),
			StartTime:      coupon.StartTime.Format(time.RFC3339),
			EndTime:        coupon.EndTime.Format(time.RFC3339),
			MaxIssue:       coupon.MaxIssue,
			MaxPerUser:     coupon.MaxPerUser,
			IssuedCount:    coupon.IssuedCount,
			UsedCount:      coupon.UsedCount,
		}
	}

	return &pb.ListCouponsReply{
		Coupons: pbCoupons,
		Total:   total,
	}, nil
}

// ListUserCoupons 获取用户优惠券列表
func (s *CouponService) ListUserCoupons(ctx context.Context, req *pb.ListUserCouponsRequest) (*pb.ListUserCouponsReply, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	userCoupons, total, err := s.couponUc.ListUserCoupons(ctx, int32(req.UserId), page, pageSize)
	if err != nil {
		s.log.Errorf("获取用户优惠券列表失败: %v", err)
		return nil, err
	}

	// 转换为响应格式
	pbUserCoupons := make([]*pb.UserCoupon, len(userCoupons))
	for i, userCouponDetail := range userCoupons {
		userCoupon := userCouponDetail.DiscountUser
		coupon := userCouponDetail.Coupon

		// 根据实际的 MtDiscountUser 结构体字段进行映射
		pbUserCoupons[i] = &pb.UserCoupon{
			Id:         int64(userCoupon.ID),
			CouponId:   userCoupon.DiscountID,
			UserId:     int64(userCoupon.UserID),
			Status:     "available", // 默认状态，可以根据 userCouponDetail.IsUsed 等字段判断
			ClaimTime:  "", // MtDiscountUser 没有这个字段，可以设为空或从其他地方获取
			UseTime:    "", // MtDiscountUser 没有这个字段
			ExpireTime: coupon.EndTime.Format(time.RFC3339), // 使用优惠券的结束时间
			Coupon: &pb.Coupon{
				Id:             coupon.ID,
				DiscountName:   coupon.DiscountName,
				Classify:       coupon.Classify,
				StoreId:        coupon.StoreID,
				DiscountAmount: coupon.DiscountAmount.String(),
				MinOrderAmount: coupon.MinOrderAmount.String(),
				StartTime:      coupon.StartTime.Format(time.RFC3339),
				EndTime:        coupon.EndTime.Format(time.RFC3339),
				MaxIssue:       coupon.MaxIssue,
				MaxPerUser:     coupon.MaxPerUser,
			},
		}
	}

	return &pb.ListUserCouponsReply{
		Coupons: pbUserCoupons,
		Total:   total,
	}, nil
}

// GetCouponDetail 获取优惠券详情
func (s *CouponService) GetCouponDetail(ctx context.Context, req *pb.GetCouponDetailRequest) (*pb.GetCouponDetailReply, error) {
	detail, err := s.couponUc.GetCouponDetail(ctx, req.CouponId)
	if err != nil {
		s.log.Errorf("获取优惠券详情失败: %v", err)
		return nil, err
	}

	// 转换优惠券信息
	coupon := &pb.Coupon{
		Id:             detail.Coupon.ID,
		DiscountName:   detail.Coupon.DiscountName,
		Classify:       detail.Coupon.Classify,
		StoreId:        detail.Coupon.StoreID,
		DiscountAmount: detail.Coupon.DiscountAmount.String(),
		MinOrderAmount: detail.Coupon.MinOrderAmount.String(),
		StartTime:      detail.Coupon.StartTime.Format(time.RFC3339),
		EndTime:        detail.Coupon.EndTime.Format(time.RFC3339),
		MaxIssue:       detail.Coupon.MaxIssue,
		MaxPerUser:     detail.Coupon.MaxPerUser,
	}

	// 转换规则信息
	rules := make([]*pb.CouponRule, len(detail.Rules))
	for i, rule := range detail.Rules {
		rules[i] = &pb.CouponRule{
			Id:         rule.ID,
			DiscountId: rule.DiscountID,
			RuleKey:    rule.RuleKey,
			RuleValue:  rule.RuleValue,
			Platform:   rule.Platform,
			DrugId:     rule.DrugID,
			Astrict:    rule.Astrict,
		}
	}

	return &pb.GetCouponDetailReply{
		Coupon: coupon,
		Rules:  rules,
	}, nil
}

// ClaimCoupon 领取优惠券
func (s *CouponService) ClaimCoupon(ctx context.Context, req *pb.ClaimCouponRequest) (*pb.ClaimCouponReply, error) {
	userCoupon, err := s.couponUc.ClaimCoupon(ctx, req.CouponId, int32(req.UserId))
	if err != nil {
		s.log.Errorf("领取优惠券失败: %v", err)
		return &pb.ClaimCouponReply{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.ClaimCouponReply{
		Success:      true,
		Message:      "领取成功",
		UserCouponId: int64(userCoupon.ID),
	}, nil
}

// CalculateAvailableCoupons 计算可用优惠券
func (s *CouponService) CalculateAvailableCoupons(ctx context.Context, req *pb.CalculateAvailableCouponsRequest) (*pb.CalculateAvailableCouponsReply, error) {
	// 解析总金额
	totalAmount, err := decimal.NewFromString(req.TotalAmount)
	if err != nil {
		s.log.Errorf("解析总金额失败: %v", err)
		return nil, err
	}

	// 转换订单项
	items := make([]*biz.CouponCalculateItem, len(req.Items))
	for i, item := range req.Items {
		price, err := decimal.NewFromString(item.Price)
		if err != nil {
			s.log.Errorf("解析商品价格失败: %v", err)
			return nil, err
		}

		items[i] = &biz.CouponCalculateItem{
			DrugID:   item.DrugId,
			Quantity: item.Quantity,
			Price:    price,
		}
	}

	// 构造计算请求
	calculateReq := &biz.CouponCalculateRequest{
		UserID:      int32(req.UserId),
		Items:       items,
		TotalAmount: totalAmount,
		StoreID:     req.StoreId,
	}

	// 计算可用优惠券
	availableCoupons, err := s.couponUc.CalculateAvailableCoupons(ctx, calculateReq)
	if err != nil {
		s.log.Errorf("计算可用优惠券失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	pbAvailableCoupons := make([]*pb.AvailableCoupon, len(availableCoupons))
	var bestDiscountAmount decimal.Decimal
	var bestCouponID int64

	for i, availableCoupon := range availableCoupons {
		pbAvailableCoupons[i] = &pb.AvailableCoupon{
			UserCouponId:   availableCoupon.UserCouponID,
			DiscountAmount: availableCoupon.DiscountAmount.String(),
			FinalAmount:    availableCoupon.FinalAmount.String(),
			CanUse:         availableCoupon.CanUse,
			Reason:         availableCoupon.Reason,
			Coupon: &pb.Coupon{
				Id:             availableCoupon.Coupon.ID,
				DiscountName:   availableCoupon.Coupon.DiscountName,
				Classify:       availableCoupon.Coupon.Classify,
				StoreId:        availableCoupon.Coupon.StoreID,
				DiscountAmount: availableCoupon.Coupon.DiscountAmount.String(),
				MinOrderAmount: availableCoupon.Coupon.MinOrderAmount.String(),
				StartTime:      availableCoupon.Coupon.StartTime.Format(time.RFC3339),
				EndTime:        availableCoupon.Coupon.EndTime.Format(time.RFC3339),
				MaxIssue:       availableCoupon.Coupon.MaxIssue,
				MaxPerUser:     availableCoupon.Coupon.MaxPerUser,
			},
		}

		// 找出最优惠的券
		if availableCoupon.CanUse && availableCoupon.DiscountAmount.GreaterThan(bestDiscountAmount) {
			bestDiscountAmount = availableCoupon.DiscountAmount
			bestCouponID = availableCoupon.UserCouponID
		}
	}

	return &pb.CalculateAvailableCouponsReply{
		AvailableCoupons:   pbAvailableCoupons,
		BestDiscountAmount: bestDiscountAmount.String(),
		BestCouponId:       bestCouponID,
	}, nil
}