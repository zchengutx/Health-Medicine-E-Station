package biz

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

// 购物车项目结构
type CartItem struct {
	ID           int64   `json:"id"`
	UserID       int64   `json:"user_id"`
	DrugID       int64   `json:"drug_id"`
	Number       int64   `json:"number"`
	DrugName     string  `json:"drug_name"`
	Specification string `json:"specification"`
	Price        float64 `json:"price"`
	Inventory    int64   `json:"inventory"`
	ExhibitionURL string `json:"exhibition_url"`
}

// 购物车仓库接口
type CartRepo interface {
	// 创建购物车项目
	CreateCart(ctx context.Context, userID, drugID, number int64) error
	
	// 更新购物车项目数量
	UpdateCart(ctx context.Context, userID, drugID, number int64) error
	
	// 删除购物车项目（支持批量删除）
	DeleteCart(ctx context.Context, userID int64, drugIDs []int64) error
	
	// 获取用户购物车列表
	ListCart(ctx context.Context, userID int64) ([]*CartItem, error)
	
	// 检查药品库存
	CheckDrugInventory(ctx context.Context, drugID int64) (int64, error)
	
	// 获取药品信息
	GetDrugInfo(ctx context.Context, drugID int64) (*CartItem, error)
}

// 购物车服务
type CartService struct {
	repo CartRepo
	log  *log.Helper
}

// 创建购物车服务
func NewCartService(repo CartRepo, logger log.Logger) *CartService {
	return &CartService{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// 创建购物车项目
func (cs *CartService) CreateCart(ctx context.Context, userID, drugID, number int64) error {
	cs.log.WithContext(ctx).Infof("CreateCart: userID=%d, drugID=%d, number=%d", userID, drugID, number)
	
	// 检查库存
	inventory, err := cs.repo.CheckDrugInventory(ctx, drugID)
	if err != nil {
		return err
	}
	
	if number > inventory {
		return ErrInsufficientInventory
	}
	
	return cs.repo.CreateCart(ctx, userID, drugID, number)
}

// 更新购物车项目
func (cs *CartService) UpdateCart(ctx context.Context, userID, drugID, number int64) error {
	cs.log.WithContext(ctx).Infof("UpdateCart: userID=%d, drugID=%d, number=%d", userID, drugID, number)
	
	// 如果数量为0，则删除该项目
	if number == 0 {
		return cs.repo.DeleteCart(ctx, userID, []int64{drugID})
	}
	
	// 检查库存
	inventory, err := cs.repo.CheckDrugInventory(ctx, drugID)
	if err != nil {
		return err
	}
	
	if number > inventory {
		return ErrInsufficientInventory
	}
	
	return cs.repo.UpdateCart(ctx, userID, drugID, number)
}

// 删除购物车项目
func (cs *CartService) DeleteCart(ctx context.Context, userID int64, drugIDs []int64) error {
	cs.log.WithContext(ctx).Infof("DeleteCart: userID=%d, drugIDs=%v", userID, drugIDs)
	return cs.repo.DeleteCart(ctx, userID, drugIDs)
}

// 获取购物车列表
func (cs *CartService) ListCart(ctx context.Context, userID int64) ([]*CartItem, error) {
	cs.log.WithContext(ctx).Infof("ListCart: userID=%d", userID)
	return cs.repo.ListCart(ctx, userID)
}

// 生成购物车Redis键
func GenerateCartKey(userID int64) string {
	return "cart:user:" + strconv.FormatInt(userID, 10)
}

// 生成购物车项目Redis字段
func GenerateCartField(drugID int64) string {
	return "drug:" + strconv.FormatInt(drugID, 10)
}