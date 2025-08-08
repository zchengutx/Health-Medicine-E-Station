package service

import (
	"context"

	cartv1 "kratos_client/api/cart/v1"
	"kratos_client/internal/biz"
	"kratos_client/internal/data"
)

type CartService struct {
	cartv1.UnimplementedCartServer
	data *data.Data
	uc   *biz.CartService
}

func NewCartService(uc *biz.CartService, d *data.Data) *CartService {
	return &CartService{
		UnimplementedCartServer: cartv1.UnimplementedCartServer{},
		data:                    d,
		uc:                      uc,
	}
}

// 创建购物车项目
func (s *CartService) CreateCart(ctx context.Context, req *cartv1.CreateCartRequest) (*cartv1.CreateCartReply, error) {
	err := s.uc.CreateCart(ctx, req.UserId, req.DrugId, req.Number)
	if err != nil {
		if err == biz.ErrInsufficientInventory {
			return &cartv1.CreateCartReply{
				Code: 400,
				Msg:  "库存不足",
			}, nil
		}
		return &cartv1.CreateCartReply{
			Code: 500,
			Msg:  "添加购物车失败",
		}, err
	}

	return &cartv1.CreateCartReply{
		Code: 0,
		Msg:  "添加购物车成功",
	}, nil
}

// 更新购物车项目
func (s *CartService) UpdateCart(ctx context.Context, req *cartv1.UpdateCartRequest) (*cartv1.UpdateCartReply, error) {
	err := s.uc.UpdateCart(ctx, req.UserId, req.DrugId, req.Number)
	if err != nil {
		if err == biz.ErrInsufficientInventory {
			return &cartv1.UpdateCartReply{
				Code: 400,
				Msg:  "库存不足",
			}, nil
		}
		if err == biz.ErrCartItemNotFound {
			return &cartv1.UpdateCartReply{
				Code: 404,
				Msg:  "购物车项目不存在",
			}, nil
		}
		return &cartv1.UpdateCartReply{
			Code: 500,
			Msg:  "更新购物车失败",
		}, err
	}

	return &cartv1.UpdateCartReply{
		Code: 0,
		Msg:  "更新购物车成功",
	}, nil
}

// 删除购物车项目
func (s *CartService) DeleteCart(ctx context.Context, req *cartv1.DeleteCartRequest) (*cartv1.DeleteCartReply, error) {
	err := s.uc.DeleteCart(ctx, req.UserId, req.DrugIds)
	if err != nil {
		return &cartv1.DeleteCartReply{
			Code: 500,
			Msg:  "删除购物车失败",
		}, err
	}

	return &cartv1.DeleteCartReply{
		Code: 0,
		Msg:  "删除购物车成功",
	}, nil
}

// 获取购物车列表
func (s *CartService) ListCart(ctx context.Context, req *cartv1.ListCartRequest) (*cartv1.ListCartReply, error) {
	cartItems, err := s.uc.ListCart(ctx, req.UserId)
	if err != nil {
		return &cartv1.ListCartReply{
			Code: 500,
			Msg:  "获取购物车列表失败",
		}, err
	}

	var cartList []*cartv1.InfoCart
	for _, item := range cartItems {
		cartList = append(cartList, &cartv1.InfoCart{
			Id:            item.ID,
			UserId:        item.UserID,
			DrugId:        item.DrugID,
			Number:        item.Number,
			DrugName:      item.DrugName,
			Specification: item.Specification,
			Price:         item.Price,
			Inventory:     item.Inventory,
			ExhibitionUrl: item.ExhibitionURL,
		})
	}

	return &cartv1.ListCartReply{
		Code: 0,
		Msg:  "获取购物车列表成功",
		Cart: cartList,
	}, nil
}