# 订单 API 测试文档

## 1. 创建订单

### 请求
```bash
curl -X POST http://localhost:8000/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1001,
    "user_name": "张三",
    "user_phone": "13800138000",
    "doctor_id": 2001,
    "doctor_name": "李医生",
    "address_id": 3001,
    "address_detail": "北京市朝阳区某某街道123号",
    "items": [
      {
        "drug_id": 4001,
        "quantity": 2
      },
      {
        "drug_id": 4002,
        "quantity": 1
      }
    ],
    "remark": "请尽快配送"
  }'
```

### 响应
```json
{
  "order_no": "ORD_1001_1234567890",
  "total_amount": "99.99",
  "status": "1",
  "message": "订单创建成功"
}
```

## 2. 获取订单详情

### 请求
```bash
curl -X GET http://localhost:8000/api/v1/orders/ORD_1001_1234567890
```

### 响应
```json
{
  "order": {
    "id": 1,
    "order_no": "ORD_1001_1234567890",
    "user_id": 1001,
    "user_name": "张三",
    "user_phone": "13800138000",
    "doctor_id": 2001,
    "doctor_name": "李医生",
    "address_id": 3001,
    "address_detail": "北京市朝阳区某某街道123号",
    "total_amount": "99.99",
    "pay_type": "",
    "status": "1",
    "remark": "请尽快配送"
  },
  "items": [
    {
      "id": 1,
      "order_id": 1,
      "drug_id": 4001,
      "drug_name": "感冒灵颗粒",
      "drug_spec": "10g*12袋",
      "quantity": 2,
      "price": "29.99",
      "subtotal": "59.98"
    },
    {
      "id": 2,
      "order_id": 1,
      "drug_id": 4002,
      "drug_name": "维生素C片",
      "drug_spec": "100mg*30片",
      "quantity": 1,
      "price": "39.99",
      "subtotal": "39.99"
    }
  ]
}
```

## 3. 获取用户订单列表

### 请求
```bash
curl -X GET "http://localhost:8000/api/v1/users/1001/orders?page=1&page_size=10"
```

### 响应
```json
{
  "orders": [
    {
      "order_no": "ORD_1001_1234567890",
      "total_amount": "99.99",
      "status": "1",
      "item_count": 2
    }
  ],
  "total": 1
}
```

## 4. 处理支付

### 请求
```bash
curl -X POST http://localhost:8000/api/v1/orders/ORD_1001_1234567890/payment \
  -H "Content-Type: application/json" \
  -d '{
    "pay_type": "2",
    "amount": "99.99",
    "trade_no": "ALIPAY_TRADE_123456789"
  }'
```

### 响应
```json
{
  "success": true,
  "message": "支付处理成功"
}
```

## 订单状态说明

- `1`: 待支付
- `2`: 已支付
- `3`: 配药中
- `4`: 已发货
- `5`: 已完成
- `6`: 已取消

## 支付方式说明

- `1`: 微信支付
- `2`: 支付宝
- `3`: 银行卡