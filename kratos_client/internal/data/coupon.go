package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"kratos_client/internal/biz"
)

// 优惠券数据模型
type MtDiscount struct {
	ID             int32           `gorm:"primaryKey;autoIncrement" json:"id"`
	DiscountName   string          `gorm:"column:discount_name;size:50" json:"discount_name"`
	Classify       string          `gorm:"column:classify;size:20" json:"classify"`
	StoreID        int32           `gorm:"column:store_id;default:0;not null" json:"store_id"`
	DiscountAmount decimal.Decimal `gorm:"column:discount_amout;type:decimal(10,2)" json:"discount_amount"`
	MinOrderAmount decimal.Decimal `gorm:"column:min_order_amount;type:decimal(10,2)" json:"min_order_amount"`
	StartTime      time.Time       `gorm:"column:start_time" json:"start_time"`
	EndTime        time.Time       `gorm:"column:end_time" json:"end_time"`
	MaxIssue       int32           `gorm:"column:max_issue" json:"max_issue"`
	MaxPerUser     int32           `gorm:"column:max_per_user;default:1" json:"max_per_user"`
}

func (MtDiscount) TableName() string {
	return "mt_discount"
}

// 优惠券规则数据模型
type MtCouponRule struct {
	ID         int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	DiscountID int32  `gorm:"column:discount_id;default:0;not null" json:"discount_id"`
	RuleKey    int32  `gorm:"column:rule_key;default:0;not null" json:"rule_key"`
	RuleValue  string `gorm:"column:rule_value;size:30" json:"rule_value"`
	Platform   string `gorm:"column:platform;size:50" json:"platform"`
	DrugID     int32  `gorm:"column:drug_id;default:0;not null" json:"drug_id"`
	Astrict    string `gorm:"column:astrict;size:20" json:"astrict"`
}

func (MtCouponRule) TableName() string {
	return "mt_coupon_rule"
}

// 用户优惠券关联数据模型
type MtDiscountUser struct {
	ID         int32 `gorm:"primaryKey;autoIncrement" json:"id"`
	DiscountID int32 `gorm:"column:discount_id;default:0;not null" json:"discount_id"`
	UserID     int32 `gorm:"column:user_id;default:0;not null" json:"user_id"`
}

func (MtDiscountUser) TableName() string {
	return "mt_discount_user"
}

// 订单优惠券使用记录数据模型
type MtOrderCoupon struct {
	ID             int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderNo        string          `gorm:"column:order_no;size:50;not null" json:"order_no"`
	DiscountUserID int32           `gorm:"column:discount_user_id;not null" json:"discount_user_id"`
	DiscountID     int32           `gorm:"column:discount_id;not null" json:"discount_id"`
	UserID         int32           `gorm:"column:user_id;not null" json:"user_id"`
	DiscountAmount decimal.Decimal `gorm:"column:discount_amount;type:decimal(10,2);not null" json:"discount_amount"`
	UseTime        time.Time       `gorm:"column:use_time;autoCreateTime" json:"use_time"`
}

func (MtOrderCoupon) TableName() string {
	return "mt_order_coupon"
}

// 优惠券仓储实现
type couponRepo struct {
	data *Data
	log  *log.Helper
}

// 创建优惠券仓储
func NewCouponRepo(data *Data, logger log.Logger) biz.CouponRepo {
	return &couponRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 转换数据模型到业务模型
func (r *couponRepo) toBizDiscount(do *MtDiscount) *biz.MtDiscount {
	return &biz.MtDiscount{
		ID:             do.ID,
		DiscountName:   do.DiscountName,
		Classify:       do.Classify,
		StoreID:        do.StoreID,
		DiscountAmount: do.DiscountAmount,
		MinOrderAmount: do.MinOrderAmount,
		StartTime:      do.StartTime,
		EndTime:        do.EndTime,
		MaxIssue:       do.MaxIssue,
		MaxPerUser:     do.MaxPerUser,
	}
}

// 转换优惠券规则
func (r *couponRepo) toBizCouponRule(do *MtCouponRule) *biz.MtCouponRule {
	return &biz.MtCouponRule{
		ID:         do.ID,
		DiscountID: do.DiscountID,
		RuleKey:    do.RuleKey,
		RuleValue:  do.RuleValue,
		Platform:   do.Platform,
		DrugID:     do.DrugID,
		Astrict:    do.Astrict,
	}
}

// 转换用户优惠券关联
func (r *couponRepo) toBizDiscountUser(do *MtDiscountUser) *biz.MtDiscountUser {
	return &biz.MtDiscountUser{
		ID:         do.ID,
		DiscountID: do.DiscountID,
		UserID:     do.UserID,
	}
}

// 获取优惠券列表
func (r *couponRepo) ListCoupons(ctx context.Context, storeID int32, page, pageSize int32) ([]*biz.MtDiscount, int64, error) {
	var coupons []MtDiscount
	var total int64

	db := r.data.Db.WithContext(ctx).Model(&MtDiscount{})
	
	// 添加店铺过滤条件
	if storeID > 0 {
		db = db.Where("store_id = ? OR store_id = 0", storeID)
	}

	// 只显示有效期内的优惠券
	now := time.Now()
	db = db.Where("start_time <= ? AND end_time >= ?", now, now)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		r.log.Errorf("查询优惠券总数失败: %v", err)
		return nil, 0, err
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := db.Order("id DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&coupons).Error; err != nil {
		r.log.Errorf("查询优惠券列表失败: %v", err)
		return nil, 0, err
	}

	// 转换为业务模型
	bizCoupons := make([]*biz.MtDiscount, len(coupons))
	for i, coupon := range coupons {
		bizCoupons[i] = r.toBizDiscount(&coupon)
	}

	return bizCoupons, total, nil
}

// 根据ID获取优惠券
func (r *couponRepo) GetCouponByID(ctx context.Context, id int32) (*biz.MtDiscount, error) {
	var coupon MtDiscount
	if err := r.data.Db.WithContext(ctx).Where("id = ?", id).First(&coupon).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.log.Errorf("查询优惠券失败: %v", err)
		return nil, err
	}

	return r.toBizDiscount(&coupon), nil
}

// 获取优惠券规则
func (r *couponRepo) GetCouponRules(ctx context.Context, couponID int32) ([]*biz.MtCouponRule, error) {
	var rules []MtCouponRule
	if err := r.data.Db.WithContext(ctx).Where("discount_id = ?", couponID).Find(&rules).Error; err != nil {
		r.log.Errorf("查询优惠券规则失败: %v", err)
		return nil, err
	}

	bizRules := make([]*biz.MtCouponRule, len(rules))
	for i, rule := range rules {
		bizRules[i] = r.toBizCouponRule(&rule)
	}

	return bizRules, nil
}

// 获取用户优惠券列表
func (r *couponRepo) ListUserCoupons(ctx context.Context, userID int32, page, pageSize int32) ([]*biz.UserCouponDetail, int64, error) {
	var discountUsers []MtDiscountUser
	var total int64

	db := r.data.Db.WithContext(ctx).Model(&MtDiscountUser{}).Where("user_id = ?", userID)

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		r.log.Errorf("查询用户优惠券总数失败: %v", err)
		return nil, 0, err
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := db.Order("id DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&discountUsers).Error; err != nil {
		r.log.Errorf("查询用户优惠券列表失败: %v", err)
		return nil, 0, err
	}

	// 获取优惠券详情
	var details []*biz.UserCouponDetail
	for _, discountUser := range discountUsers {
		coupon, err := r.GetCouponByID(ctx, discountUser.DiscountID)
		if err != nil {
			r.log.Errorf("查询优惠券详情失败: %v", err)
			continue
		}

		// 检查是否已使用
		isUsed, err := r.CheckCouponUsed(ctx, discountUser.ID)
		if err != nil {
			r.log.Errorf("检查优惠券使用状态失败: %v", err)
			isUsed = false
		}

		// 检查是否已过期
		isExpired := time.Now().After(coupon.EndTime)

		details = append(details, &biz.UserCouponDetail{
			DiscountUser: r.toBizDiscountUser(&discountUser),
			Coupon:       coupon,
			IsUsed:       isUsed,
			IsExpired:    isExpired,
		})
	}

	return details, total, nil
}

// 领取优惠券
func (r *couponRepo) ClaimCoupon(ctx context.Context, couponID int32, userID int32) (*biz.MtDiscountUser, error) {
	discountUser := &MtDiscountUser{
		DiscountID: couponID,
		UserID:     userID,
	}

	if err := r.data.Db.WithContext(ctx).Create(discountUser).Error; err != nil {
		r.log.Errorf("创建用户优惠券关联失败: %v", err)
		return nil, err
	}

	return r.toBizDiscountUser(discountUser), nil
}

// 获取用户优惠券
func (r *couponRepo) GetUserCoupon(ctx context.Context, discountUserID int32) (*biz.MtDiscountUser, error) {
	var discountUser MtDiscountUser
	if err := r.data.Db.WithContext(ctx).Where("id = ?", discountUserID).First(&discountUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.log.Errorf("查询用户优惠券关联失败: %v", err)
		return nil, err
	}

	return r.toBizDiscountUser(&discountUser), nil
}

// 使用优惠券
func (r *couponRepo) UseCoupon(ctx context.Context, orderNo string, discountUserID int32, discountAmount decimal.Decimal) error {
	// 获取用户优惠券关联信息
	discountUser, err := r.GetUserCoupon(ctx, discountUserID)
	if err != nil {
		return err
	}

	// 创建使用记录
	orderCoupon := &MtOrderCoupon{
		OrderNo:        orderNo,
		DiscountUserID: discountUserID,
		DiscountID:     discountUser.DiscountID,
		UserID:         discountUser.UserID,
		DiscountAmount: discountAmount,
		UseTime:        time.Now(),
	}

	if err := r.data.Db.WithContext(ctx).Create(orderCoupon).Error; err != nil {
		r.log.Errorf("创建优惠券使用记录失败: %v", err)
		return err
	}

	return nil
}

// 检查优惠券是否已使用
func (r *couponRepo) CheckCouponUsed(ctx context.Context, discountUserID int32) (bool, error) {
	var count int64
	if err := r.data.Db.WithContext(ctx).Model(&MtOrderCoupon{}).Where("discount_user_id = ?", discountUserID).Count(&count).Error; err != nil {
		r.log.Errorf("检查优惠券使用状态失败: %v", err)
		return false, err
	}

	return count > 0, nil
}

// 获取优惠券统计
func (r *couponRepo) GetCouponStats(ctx context.Context, couponID int32) (issued int32, used int32, err error) {
	var issuedCount int64
	var usedCount int64

	// 统计已发放数量
	if err := r.data.Db.WithContext(ctx).Model(&MtDiscountUser{}).Where("discount_id = ?", couponID).Count(&issuedCount).Error; err != nil {
		r.log.Errorf("统计优惠券发放数量失败: %v", err)
		return 0, 0, err
	}

	// 统计已使用数量
	if err := r.data.Db.WithContext(ctx).Model(&MtOrderCoupon{}).Where("discount_id = ?", couponID).Count(&usedCount).Error; err != nil {
		r.log.Errorf("统计优惠券使用数量失败: %v", err)
		return 0, 0, err
	}

	return int32(issuedCount), int32(usedCount), nil
}

// 检查用户优惠券领取限制
func (r *couponRepo) CheckUserCouponLimit(ctx context.Context, couponID int32, userID int32) (bool, error) {
	// 获取优惠券信息
	coupon, err := r.GetCouponByID(ctx, couponID)
	if err != nil {
		return false, err
	}

	// 统计用户已领取数量
	var count int64
	if err := r.data.Db.WithContext(ctx).Model(&MtDiscountUser{}).Where("discount_id = ? AND user_id = ?", couponID, userID).Count(&count).Error; err != nil {
		r.log.Errorf("统计用户优惠券领取数量失败: %v", err)
		return false, err
	}

	return int32(count) < coupon.MaxPerUser, nil
}