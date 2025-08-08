package comment

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"kratos_client/internal/biz"
)

// CouponManager 优惠券管理器
type CouponManager struct {
	couponUc *biz.CouponUsecase
}

// NewCouponManager 创建优惠券管理器
func NewCouponManager(couponUc *biz.CouponUsecase) *CouponManager {
	return &CouponManager{
		couponUc: couponUc,
	}
}

// OrderCouponInfo 订单优惠券信息
type OrderCouponInfo struct {
	UserCouponID   int64           `json:"user_coupon_id"`
	CouponID       int32           `json:"coupon_id"`
	DiscountAmount decimal.Decimal `json:"discount_amount"`
	OriginalAmount decimal.Decimal `json:"original_amount"`
	FinalAmount    decimal.Decimal `json:"final_amount"`
}

// ValidateCouponForOrder 验证优惠券是否可用于订单
func (cm *CouponManager) ValidateCouponForOrder(ctx context.Context, discountUserID int32, orderItems []*biz.CouponCalculateItem, totalAmount decimal.Decimal, storeID int32) (*OrderCouponInfo, error) {
	// 获取用户优惠券信息
	discountUser, err := cm.couponUc.GetUserCoupon(ctx, discountUserID)
	if err != nil {
		return nil, fmt.Errorf("获取用户优惠券失败: %v", err)
	}

	if discountUser == nil {
		return nil, fmt.Errorf("优惠券不存在")
	}

	// 获取优惠券详情
	couponDetail, err := cm.couponUc.GetCouponDetail(ctx, discountUser.DiscountID)
	if err != nil {
		return nil, fmt.Errorf("获取优惠券详情失败: %v", err)
	}

	coupon := couponDetail.Coupon

	// 检查是否过期
	if time.Now().After(coupon.EndTime) {
		return nil, fmt.Errorf("优惠券已过期")
	}

	// 检查店铺限制
	if coupon.StoreID > 0 && coupon.StoreID != storeID {
		return nil, fmt.Errorf("优惠券不适用于当前店铺")
	}

	// 检查最低消费门槛
	if totalAmount.LessThan(coupon.MinOrderAmount) {
		return nil, fmt.Errorf("订单金额不满足优惠券使用条件，需满%s元", coupon.MinOrderAmount.String())
	}

	// 检查商品限制
	canUseForItems, err := cm.checkCouponItemRules(ctx, coupon.ID, orderItems)
	if err != nil {
		return nil, fmt.Errorf("检查优惠券商品规则失败: %v", err)
	}

	if !canUseForItems {
		return nil, fmt.Errorf("优惠券不适用于当前商品")
	}

	// 计算优惠金额
	discountAmount := coupon.DiscountAmount
	finalAmount := totalAmount.Sub(discountAmount)
	if finalAmount.LessThan(decimal.Zero) {
		finalAmount = decimal.Zero
	}

	return &OrderCouponInfo{
		UserCouponID:   int64(discountUserID),
		CouponID:       coupon.ID,
		DiscountAmount: discountAmount,
		OriginalAmount: totalAmount,
		FinalAmount:    finalAmount,
	}, nil
}

// ApplyCouponToOrder 将优惠券应用到订单
func (cm *CouponManager) ApplyCouponToOrder(ctx context.Context, orderNo string, discountUserID int32, discountAmount decimal.Decimal) error {
	return cm.couponUc.UseCoupon(ctx, orderNo, discountUserID, discountAmount)
}

// GetBestCouponForOrder 获取订单的最优优惠券
func (cm *CouponManager) GetBestCouponForOrder(ctx context.Context, userID int32, orderItems []*biz.CouponCalculateItem, totalAmount decimal.Decimal, storeID int32) (*biz.AvailableCoupon, error) {
	req := &biz.CouponCalculateRequest{
		UserID:      userID,
		Items:       orderItems,
		TotalAmount: totalAmount,
		StoreID:     storeID,
	}

	availableCoupons, err := cm.couponUc.CalculateAvailableCoupons(ctx, req)
	if err != nil {
		return nil, err
	}

	var bestCoupon *biz.AvailableCoupon
	var maxDiscount decimal.Decimal

	for _, coupon := range availableCoupons {
		if coupon.CanUse && coupon.DiscountAmount.GreaterThan(maxDiscount) {
			maxDiscount = coupon.DiscountAmount
			bestCoupon = coupon
		}
	}

	return bestCoupon, nil
}

// CalculateOrderDiscount 计算订单优惠
func (cm *CouponManager) CalculateOrderDiscount(ctx context.Context, userID int64, orderItems []*biz.CouponCalculateItem, totalAmount decimal.Decimal, storeID int32, userCouponID *int64) (*OrderCouponInfo, error) {
	if userCouponID == nil {
		// 如果没有指定优惠券，尝试找最优的
		bestCoupon, err := cm.GetBestCouponForOrder(ctx, int32(userID), orderItems, totalAmount, storeID)
		if err != nil {
			return nil, err
		}

		if bestCoupon == nil || !bestCoupon.CanUse {
			// 没有可用优惠券
			return &OrderCouponInfo{
				OriginalAmount: totalAmount,
				FinalAmount:    totalAmount,
				DiscountAmount: decimal.Zero,
			}, nil
		}

		return &OrderCouponInfo{
			UserCouponID:   bestCoupon.UserCouponID,
			CouponID:       bestCoupon.Coupon.ID,
			DiscountAmount: bestCoupon.DiscountAmount,
			OriginalAmount: totalAmount,
			FinalAmount:    bestCoupon.FinalAmount,
		}, nil
	}

	// 验证指定的优惠券
	return cm.ValidateCouponForOrder(ctx, int32(*userCouponID), orderItems, totalAmount, storeID)
}

// checkCouponItemRules 检查优惠券商品规则
func (cm *CouponManager) checkCouponItemRules(ctx context.Context, couponID int32, items []*biz.CouponCalculateItem) (bool, error) {
	couponDetail, err := cm.couponUc.GetCouponDetail(ctx, couponID)
	if err != nil {
		return false, err
	}

	rules := couponDetail.Rules

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

// CouponRule 优惠券规则常量
const (
	// 规则维度
	RuleKeyStore = 1 // 店铺规则
	RuleKeyCity  = 2 // 城市规则
	RuleKeyDrug  = 3 // 商品规则

	// 优惠券状态
	CouponStatusAvailable = "available" // 可用
	CouponStatusUsed      = "used"      // 已使用
	CouponStatusExpired   = "expired"   // 已过期

	// 优惠券分类
	CouponClassifyPlatform = "platform" // 平台券
	CouponClassifyStore    = "store"    // 店铺券
	CouponClassifyNewUser  = "new_user" // 新用户券
)

// GetCouponStatusText 获取优惠券状态文本
func GetCouponStatusText(status string) string {
	switch status {
	case CouponStatusAvailable:
		return "可使用"
	case CouponStatusUsed:
		return "已使用"
	case CouponStatusExpired:
		return "已过期"
	default:
		return "未知状态"
	}
}

// GetCouponClassifyText 获取优惠券分类文本
func GetCouponClassifyText(classify string) string {
	switch classify {
	case CouponClassifyPlatform:
		return "平台券"
	case CouponClassifyStore:
		return "店铺券"
	case CouponClassifyNewUser:
		return "新用户券"
	default:
		return "其他"
	}
}