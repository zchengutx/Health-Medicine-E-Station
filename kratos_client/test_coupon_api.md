# 优惠券 API 测试文档

## 1. 获取优惠券列表

### 请求
```bash
curl -X GET "http://localhost:8000/api/v1/coupons?store_id=1&page=1&page_size=10"
```

### 响应
```json
{
  "coupons": [
    {
      "id": 1,
      "discount_name": "满30减5元",
      "classify": "platform",
      "store_id": 0,
      "discount_amount": "5.00",
      "min_order_amount": "30.00",
      "start_time": "2025-01-01T00:00:00Z",
      "end_time": "2025-12-31T23:59:59Z",
      "max_issue": 1000,
      "max_per_user": 1,
      "issued_count": 50,
      "used_count": 10
    }
  ],
  "total": 1
}
```

## 2. 获取用户优惠券列表

### 请求
```bash
curl -X GET "http://localhost:8000/api/v1/users/1001/coupons?page=1&page_size=10"
```

### 响应
```json
{
  "coupons": [
    {
      "id": 1,
      "coupon_id": 1,
      "user_id": 1001,
      "status": "available",
      "claim_time": "2025-01-15T10:00:00Z",
      "use_time": "",
      "expire_time": "2025-12-31T23:59:59Z",
      "coupon": {
        "id": 1,
        "discount_name": "满30减5元",
        "classify": "platform",
        "store_id": 0,
        "discount_amount": "5.00",
        "min_order_amount": "30.00",
        "start_time": "2025-01-01T00:00:00Z",
        "end_time": "2025-12-31T23:59:59Z",
        "max_issue": 1000,
        "max_per_user": 1
      }
    }
  ],
  "total": 1
}
```

## 3. 获取优惠券详情和使用规则

### 请求
```bash
curl -X GET "http://localhost:8000/api/v1/coupons/1"
```

### 响应
```json
{
  "coupon": {
    "id": 1,
    "discount_name": "满30减5元",
    "classify": "platform",
    "store_id": 0,
    "discount_amount": "5.00",
    "min_order_amount": "30.00",
    "start_time": "2025-01-01T00:00:00Z",
    "end_time": "2025-12-31T23:59:59Z",
    "max_issue": 1000,
    "max_per_user": 1
  },
  "rules": [
    {
      "id": 1,
      "discount_id": 1,
      "rule_key": 1,
      "rule_value": "1",
      "platform": "APP",
      "drug_id": 0,
      "astrict": "新用户专享"
    }
  ]
}
```

## 4. 领取优惠券

### 请求
```bash
curl -X POST http://localhost:8000/api/v1/coupons/1/claim \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1001
  }'
```

### 响应
```json
{
  "success": true,
  "message": "领取成功",
  "user_coupon_id": 123
}
```

## 5. 计算订单可用优惠券

### 请求
```bash
curl -X POST http://localhost:8000/api/v1/coupons/calculate \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1001,
    "items": [
      {
        "drug_id": 4001,
        "quantity": 2,
        "price": "15.00"
      },
      {
        "drug_id": 4002,
        "quantity": 1,
        "price": "20.00"
      }
    ],
    "total_amount": "50.00",
    "store_id": 1
  }'
```

### 响应
```json
{
  "available_coupons": [
    {
      "user_coupon_id": 123,
      "discount_amount": "5.00",
      "final_amount": "45.00",
      "can_use": true,
      "reason": "",
      "coupon": {
        "id": 1,
        "discount_name": "满30减5元",
        "classify": "platform",
        "store_id": 0,
        "discount_amount": "5.00",
        "min_order_amount": "30.00",
        "start_time": "2025-01-01T00:00:00Z",
        "end_time": "2025-12-31T23:59:59Z",
        "max_issue": 1000,
        "max_per_user": 1
      }
    }
  ],
  "best_discount_amount": "5.00",
  "best_coupon_id": 123
}
```

## 优惠券使用规则说明

### 规则维度 (rule_key)
- `1`: 店铺规则 - 限制在特定店铺使用
- `2`: 城市规则 - 限制在特定城市使用  
- `3`: 商品规则 - 限制特定商品使用

### 优惠券分类 (classify)
- `platform`: 平台券 - 全平台通用
- `store`: 店铺券 - 特定店铺发放
- `new_user`: 新用户券 - 新用户专享

### 使用限制
1. 每个订单只能使用一张优惠券
2. 必须满足最低消费门槛
3. 必须在有效期内使用
4. 必须符合商品和店铺限制规则
5. 每人限领数量限制

### 优惠券状态
- `available`: 可使用
- `used`: 已使用  
- `expired`: 已过期

## 支付时使用优惠券

在创建订单时，可以传入 `user_coupon_id` 参数来使用优惠券：

```bash
curl -X POST http://localhost:8000/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1001,
    "user_name": "张三",
    "user_phone": "13800138000",
    "address_id": 3001,
    "address_detail": "北京市朝阳区某某街道123号",
    "items": [
      {
        "drug_id": 4001,
        "quantity": 2
      }
    ],
    "user_coupon_id": 123,
    "remark": "使用优惠券下单"
  }'
```

系统会自动验证优惠券是否可用，并计算优惠后的金额。