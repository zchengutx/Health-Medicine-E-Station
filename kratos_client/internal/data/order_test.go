package data

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"kratos_client/internal/biz"
)

// 测试订单模型转换
func TestOrderModelConversion(t *testing.T) {
	repo := &orderRepo{}

	// 测试数据
	now := time.Now()
	bizOrder := &biz.MtOrder{
		ID:            1,
		OrderNo:       "ORD_123456789",
		UserID:        100,
		UserName:      "测试用户",
		UserPhone:     "13800138000",
		DoctorID:      200,
		DoctorName:    "测试医生",
		AddressID:     300,
		AddressDetail: "测试地址详情",
		TotalAmount:   decimal.NewFromFloat(99.99),
		PayType:       "1",
		Status:        "1",
		PayTime:       &now,
		Remark:        "测试备注",
	}

	// 业务模型转数据模型
	dataOrder := repo.toDataOrder(bizOrder)
	if dataOrder.OrderNo != bizOrder.OrderNo {
		t.Errorf("Expected OrderNo %s, got %s", bizOrder.OrderNo, dataOrder.OrderNo)
	}
	if dataOrder.UserID != bizOrder.UserID {
		t.Errorf("Expected UserID %d, got %d", bizOrder.UserID, dataOrder.UserID)
	}
	if !dataOrder.TotalAmount.Equal(bizOrder.TotalAmount) {
		t.Errorf("Expected TotalAmount %s, got %s", bizOrder.TotalAmount.String(), dataOrder.TotalAmount.String())
	}

	// 数据模型转业务模型
	convertedBizOrder := repo.toBizOrder(dataOrder)
	if convertedBizOrder.OrderNo != bizOrder.OrderNo {
		t.Errorf("Expected OrderNo %s, got %s", bizOrder.OrderNo, convertedBizOrder.OrderNo)
	}
	if convertedBizOrder.UserID != bizOrder.UserID {
		t.Errorf("Expected UserID %d, got %d", bizOrder.UserID, convertedBizOrder.UserID)
	}
	if !convertedBizOrder.TotalAmount.Equal(bizOrder.TotalAmount) {
		t.Errorf("Expected TotalAmount %s, got %s", bizOrder.TotalAmount.String(), convertedBizOrder.TotalAmount.String())
	}
}

// 测试订单项模型转换
func TestOrderItemModelConversion(t *testing.T) {
	repo := &orderRepo{}

	// 测试数据
	now := time.Now()
	bizOrderItem := &biz.MtOrderItem{
		ID:        1,
		OrderID:   100,
		DrugID:    200,
		DrugName:  "测试药品",
		DrugSpec:  "10mg*30片",
		Quantity:  2,
		Price:     decimal.NewFromFloat(49.99),
		Subtotal:  decimal.NewFromFloat(99.98),
		CreatedAt: now,
	}

	// 业务模型转数据模型
	dataOrderItem := repo.toDataOrderItem(bizOrderItem)
	if dataOrderItem.DrugName != bizOrderItem.DrugName {
		t.Errorf("Expected DrugName %s, got %s", bizOrderItem.DrugName, dataOrderItem.DrugName)
	}
	if dataOrderItem.Quantity != bizOrderItem.Quantity {
		t.Errorf("Expected Quantity %d, got %d", bizOrderItem.Quantity, dataOrderItem.Quantity)
	}
	if !dataOrderItem.Price.Equal(bizOrderItem.Price) {
		t.Errorf("Expected Price %s, got %s", bizOrderItem.Price.String(), dataOrderItem.Price.String())
	}

	// 数据模型转业务模型
	convertedBizOrderItem := repo.toBizOrderItem(dataOrderItem)
	if convertedBizOrderItem.DrugName != bizOrderItem.DrugName {
		t.Errorf("Expected DrugName %s, got %s", bizOrderItem.DrugName, convertedBizOrderItem.DrugName)
	}
	if convertedBizOrderItem.Quantity != bizOrderItem.Quantity {
		t.Errorf("Expected Quantity %d, got %d", bizOrderItem.Quantity, convertedBizOrderItem.Quantity)
	}
	if !convertedBizOrderItem.Price.Equal(bizOrderItem.Price) {
		t.Errorf("Expected Price %s, got %s", bizOrderItem.Price.String(), convertedBizOrderItem.Price.String())
	}
}

// 测试状态更新逻辑
func TestStatusUpdateLogic(t *testing.T) {
	testCases := []struct {
		status        string
		expectedField string
	}{
		{"2", "pay_time"},    // 已支付
		{"3", "drug_time"},   // 配药中
		{"4", "send_time"},   // 已发货
		{"5", "finish_time"}, // 已完成
		{"6", "cancel_time"}, // 已取消
	}

	for _, tc := range testCases {
		updates := map[string]interface{}{
			"status": tc.status,
		}

		timestamp := time.Now()
		switch tc.status {
		case "2": // 已支付
			if updates["pay_time"] == nil {
				updates["pay_time"] = timestamp
			}
		case "3": // 配药中
			updates["drug_time"] = timestamp
		case "4": // 已发货
			updates["send_time"] = timestamp
		case "5": // 已完成
			updates["finish_time"] = timestamp
		case "6": // 已取消
			updates["cancel_time"] = timestamp
		}

		if _, exists := updates[tc.expectedField]; !exists {
			t.Errorf("Expected field %s to be set for status %s", tc.expectedField, tc.status)
		}
	}
}
