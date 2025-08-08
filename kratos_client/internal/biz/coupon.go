package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
)

// 优惠券模型
type MtDiscount struct {
	ID             int32           `json:"id"`
	DiscountName   string          `json:"discount_name"`
	Classify       string          `json:"classify"`
	StoreID        int32           `json:"store_id"`
	DiscountAmount decimal.Decimal `json:"discount_amount"`
	MinOrderAmount decimal.Decimal `json:"min_order_amount"`
	StartTime      time.Time       `json:"start_time"`
	EndTime        time.Time       `json:"end_time"`
	MaxIssue       int32           `json:"max_issue"`
	MaxPerUser     int32           `json:"max_per_user"`
	IssuedCount    int32           `json:"issued_count"`
	UsedCount      int32           `json:"used_count"`
}

// 优惠券规则模型
type MtCouponRule struct {
	ID         int32  `json:"id"`
	DiscountID int32  `json:"discount_id"`
	RuleKey    int32  `json:"rule_key"`
	RuleValue  string `json:"rule_value"`
	Platform   string `json:"platform"`
	DrugID     int32  `json:"drug_id"`
	Astrict    string `json:"astrict"`
}

// 用户优惠券关联模型
type MtDiscountUser struct {
	ID         int32 `json:"id"`
	DiscountID int32 `json:"discount_id"`
	UserID     int32 `json:"user_id"`
}

// 订单优惠券使用记录模型
type MtOrderCoupon struct {
	ID             int64           `json:"id"`
	OrderNo        string          `json:"order_no"`
	DiscountUserID int32           `json:"discount_user_id"`
	DiscountID     int32           `json:"discount_id"`
	UserID         int32           `json:"user_id"`
	DiscountAmount decimal.Decimal `json:"discount_amount"`
	UseTime        time.Time       `json:"use_time"`
}

// 优惠券详情
type CouponDetail struct {
	Coupon *MtDiscount     `json:"coupon"`
	Rules  []*MtCouponRule `json:"rules"`
}

// 用户优惠券详情
type UserCouponDetail struct {
	DiscountUser *MtDiscountUser `json:"discount_user"`
	Coupon       *MtDiscount     `json:"coupon"`
	IsUsed       bool            `json:"is_used"`       // 是否已使用
	IsExpired    bool            `json:"is_expired"`    // 是否已过期
}

// 可用优惠券
type AvailableCoupon struct {
	UserCouponID   int64           `json:"user_coupon_id"`   // 对应 discount_user 表的 ID
	Coupon         *MtDiscount     `json:"coupon"`
	DiscountAmount decimal.Decimal `json:"discount_amount"`
	FinalAmount    decimal.Decimal `json:"final_amount"`
	CanUse         bool            `json:"can_use"`
	Reason         string          `json:"reason"`
}

// 优惠券计算请求
type CouponCalculateRequest struct {
	UserID      int32                   `json:"user_id"`
	Items       []*CouponCalculateItem  `json:"items"`
	TotalAmount decimal.Decimal         `json:"total_amount"`
	StoreID     int32                   `json:"store_id"`
}

// 优惠券计算项
type CouponCalculateItem struct {
	DrugID   int64           `json:"drug_id"`
	Quantity int32           `json:"quantity"`
	Price    decimal.Decimal `json:"price"`
}

// 优惠券仓储接口
type CouponRepo interface {
	// 优惠券CRUD
	ListCoupons(ctx context.Context, storeID int32, page, pageSize int32) ([]*MtDiscount, int64, error)
	GetCouponByID(ctx context.Context, id int32) (*MtDiscount, error)
	GetCouponRules(ctx context.Context, couponID int32) ([]*MtCouponRule, error)
	
	// 用户优惠券
	ListUserCoupons(ctx context.Context, userID int32, page, pageSize int32) ([]*UserCouponDetail, int64, error)
	ClaimCoupon(ctx context.Context, couponID int32, userID int32) (*MtDiscountUser, error)
	GetUserCoupon(ctx context.Context, discountUserID int32) (*MtDiscountUser, error)
	UseCoupon(ctx context.Context, orderNo string, discountUserID int32, discountAmount decimal.Decimal) error
	CheckCouponUsed(ctx context.Context, discountUserID int32) (bool, error)
	
	// 统计
	GetCouponStats(ctx context.Context, couponID int32) (issued int32, used int32, err error)
	CheckUserCouponLimit(ctx context.Context, couponID int32, userID int32) (bool, error)
}

// 优惠券用例
type CouponUsecase struct {
	couponRepo CouponRepo
	drugRepo   DrugRepo
	log        *log.Helper
}

// 创建优惠券用例
func NewCouponUsecase(couponRepo CouponRepo, drugRepo DrugRepo, logger log.Logger) *CouponUsecase {
	return &CouponUsecase{
		couponRepo: couponRepo,
		drugRepo:   drugRepo,
		log:        log.NewHelper(logger),
	}
}

// 获取优惠券列表
func (uc *CouponUsecase) ListCoupons(ctx context.Context, storeID int32, page, pageSize int32) ([]*MtDiscount, int64, error) {
	coupons, total, err := uc.couponRepo.ListCoupons(ctx, storeID, page, pageSize)
	if err != nil {
		uc.log.Errorf("获取优惠券列表失败: %v", err)
		return nil, 0, err
	}

	// 更新统计信息
	for _, coupon := range coupons {
		issued, used, err := uc.couponRepo.GetCouponStats(ctx, coupon.ID)
		if err != nil {
			uc.log.Warnf("获取优惠券统计失败: couponID=%d, error=%v", coupon.ID, err)
			continue
		}
		coupon.IssuedCount = issued
		coupon.UsedCount = used
	}

	return coupons, total, nil
}

// 获取用户优惠券列表
func (uc *CouponUsecase) ListUserCoupons(ctx context.Context, userID int32, page, pageSize int32) ([]*UserCouponDetail, int64, error) {
	return uc.couponRepo.ListUserCoupons(ctx, userID, page, pageSize)
}

// 获取优惠券详情
func (uc *CouponUsecase) GetCouponDetail(ctx context.Context, couponID int32) (*CouponDetail, error) {
	coupon, err := uc.couponRepo.GetCouponByID(ctx, couponID)
	if err != nil {
		uc.log.Errorf("获取优惠券详情失败: %v", err)
		return nil, err
	}

	rules, err := uc.couponRepo.GetCouponRules(ctx, couponID)
	if err != nil {
		uc.log.Errorf("获取优惠券规则失败: %v", err)
		return nil, err
	}

	return &CouponDetail{
		Coupon: coupon,
		Rules:  rules,
	}, nil
}

// 领取优惠券
func (uc *CouponUsecase) ClaimCoupon(ctx context.Context, couponID int32, userID int32) (*MtDiscountUser, error) {
	// 检查优惠券是否存在
	coupon, err := uc.couponRepo.GetCouponByID(ctx, couponID)
	if err != nil {
		return nil, err
	}

	// 检查优惠券是否在有效期内
	now := time.Now()
	if now.Before(coupon.StartTime) || now.After(coupon.EndTime) {
		return nil, fmt.Errorf("优惠券不在有效期内")
	}

	// 检查用户领取限制
	canClaim, err := uc.couponRepo.CheckUserCouponLimit(ctx, couponID, userID)
	if err != nil {
		return nil, err
	}
	if !canClaim {
		return nil, fmt.Errorf("已达到领取上限")
	}

	// 检查发行量限制
	if coupon.MaxIssue > 0 {
		issued, _, err := uc.couponRepo.GetCouponStats(ctx, couponID)
		if err != nil {
			return nil, err
		}
		if issued >= coupon.MaxIssue {
			return nil, fmt.Errorf("优惠券已领完")
		}
	}

	// 领取优惠券
	discountUser, err := uc.couponRepo.ClaimCoupon(ctx, couponID, userID)
	if err != nil {
		uc.log.Errorf("领取优惠券失败: %v", err)
		return nil, err
	}

	uc.log.Infof("用户领取优惠券成功: userID=%d, couponID=%d", userID, couponID)
	return discountUser, nil
}

// 计算可用优惠券
func (uc *CouponUsecase) CalculateAvailableCoupons(ctx context.Context, req *CouponCalculateRequest) ([]*AvailableCoupon, error) {
	// 获取用户可用优惠券
	userCoupons, _, err := uc.couponRepo.ListUserCoupons(ctx, req.UserID, 1, 100)
	if err != nil {
		return nil, err
	}

	var availableCoupons []*AvailableCoupon

	for _, userCouponDetail := range userCoupons {
		coupon := userCouponDetail.Coupon
		discountUser := userCouponDetail.DiscountUser

		// 检查优惠券是否过期
		if userCouponDetail.IsExpired {
			continue
		}

		// 检查是否已使用
		if userCouponDetail.IsUsed {
			continue
		}

		// 检查店铺限制
		if coupon.StoreID > 0 && coupon.StoreID != req.StoreID {
			availableCoupons = append(availableCoupons, &AvailableCoupon{
				UserCouponID:   int64(discountUser.ID),
				Coupon:         coupon,
				DiscountAmount: decimal.Zero,
				FinalAmount:    req.TotalAmount,
				CanUse:         false,
				Reason:         "不适用于当前店铺",
			})
			continue
		}

		// 检查最低消费门槛
		if req.TotalAmount.LessThan(coupon.MinOrderAmount) {
			availableCoupons = append(availableCoupons, &AvailableCoupon{
				UserCouponID:   int64(discountUser.ID),
				Coupon:         coupon,
				DiscountAmount: decimal.Zero,
				FinalAmount:    req.TotalAmount,
				CanUse:         false,
				Reason:         fmt.Sprintf("需满%s元才能使用", coupon.MinOrderAmount.String()),
			})
			continue
		}

		// 检查商品限制
		canUseForItems, err := uc.checkCouponItemRules(ctx, coupon.ID, req.Items)
		if err != nil {
			uc.log.Errorf("检查优惠券商品规则失败: %v", err)
			continue
		}

		if !canUseForItems {
			availableCoupons = append(availableCoupons, &AvailableCoupon{
				UserCouponID:   int64(discountUser.ID),
				Coupon:         coupon,
				DiscountAmount: decimal.Zero,
				FinalAmount:    req.TotalAmount,
				CanUse:         false,
				Reason:         "不适用于当前商品",
			})
			continue
		}

		// 计算优惠金额
		discountAmount := coupon.DiscountAmount
		finalAmount := req.TotalAmount.Sub(discountAmount)
		if finalAmount.LessThan(decimal.Zero) {
			finalAmount = decimal.Zero
		}

		availableCoupons = append(availableCoupons, &AvailableCoupon{
			UserCouponID:   int64(discountUser.ID),
			Coupon:         coupon,
			DiscountAmount: discountAmount,
			FinalAmount:    finalAmount,
			CanUse:         true,
			Reason:         "",
		})
	}

	return availableCoupons, nil
}

// 检查优惠券商品规则
func (uc *CouponUsecase) checkCouponItemRules(ctx context.Context, couponID int32, items []*CouponCalculateItem) (bool, error) {
	rules, err := uc.couponRepo.GetCouponRules(ctx, couponID)
	if err != nil {
		return false, err
	}

	// 如果没有商品规则，则适用于所有商品
	hasItemRules := false
	for _, rule := range rules {
		if rule.DrugID > 0 {
			hasItemRules = true
			break
		}
	}

	if !hasItemRules {
		return true, nil
	}

	// 检查是否有匹配的商品
	for _, item := range items {
		for _, rule := range rules {
			if rule.DrugID > 0 && int64(rule.DrugID) == item.DrugID {
				return true, nil
			}
		}
	}

	return false, nil
}

// 使用优惠券
func (uc *CouponUsecase) UseCoupon(ctx context.Context, orderNo string, discountUserID int32, discountAmount decimal.Decimal) error {
	return uc.couponRepo.UseCoupon(ctx, orderNo, discountUserID, discountAmount)
}

// 获取用户优惠券
func (uc *CouponUsecase) GetUserCoupon(ctx context.Context, discountUserID int32) (*MtDiscountUser, error) {
	return uc.couponRepo.GetUserCoupon(ctx, discountUserID)
}