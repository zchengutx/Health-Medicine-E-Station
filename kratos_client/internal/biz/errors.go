package biz

import "errors"

var (
	// ErrInsufficientInventory 库存不足错误
	ErrInsufficientInventory = errors.New("insufficient inventory")
	
	// ErrCartItemNotFound 购物车项目不存在错误
	ErrCartItemNotFound = errors.New("cart item not found")
	
	// ErrDrugNotFound 药品不存在错误
	ErrDrugNotFound = errors.New("drug not found")
)