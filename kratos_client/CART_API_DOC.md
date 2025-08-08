# 购物车API文档

## 功能概述

购物车模块提供了完整的购物车管理功能，包括添加商品、修改数量、删除商品和查看购物车列表。所有购物车数据存储在Redis中，使用哈希类型存储，确保高性能和数据一致性。

## 主要特性

- ✅ 添加商品到购物车（支持数量累加）
- ✅ 修改购物车商品数量
- ✅ 删除购物车商品（支持单个和批量删除）
- ✅ 查看购物车列表
- ✅ 库存检查和数据一致性保证
- ✅ Redis哈希存储，高性能访问

## API接口

### 1. 添加商品到购物车

**接口地址**: `POST /v1/cart/create`

**请求参数**:
```json
{
  "user_id": 1,
  "drug_id": 123,
  "number": 2
}
```

**响应示例**:
```json
{
  "code": 0,
  "msg": "添加购物车成功"
}
```

**错误响应**:
```json
{
  "code": 400,
  "msg": "库存不足"
}
```

### 2. 修改购物车商品数量

**接口地址**: `POST /v1/cart/update`

**请求参数**:
```json
{
  "user_id": 1,
  "drug_id": 123,
  "number": 5
}
```

**响应示例**:
```json
{
  "code": 0,
  "msg": "更新购物车成功"
}
```

**特殊说明**:
- 当 `number` 为 0 时，会自动删除该商品
- 会检查库存，确保不超过可用库存

### 3. 删除购物车商品

**接口地址**: `POST /v1/cart/delete`

**请求参数**:
```json
{
  "user_id": 1,
  "drug_ids": [123, 456, 789]
}
```

**响应示例**:
```json
{
  "code": 0,
  "msg": "删除购物车成功"
}
```

**特殊说明**:
- 支持批量删除，传入多个药品ID
- 支持单个删除，传入单个药品ID

### 4. 获取购物车列表

**接口地址**: `POST /v1/cart/list`

**请求参数**:
```json
{
  "user_id": 1
}
```

**响应示例**:
```json
{
  "code": 0,
  "msg": "获取购物车列表成功",
  "cart": [
    {
      "id": 123,
      "user_id": 1,
      "drug_id": 123,
      "number": 2,
      "drug_name": "阿莫西林胶囊",
      "specification": "0.25g*24粒",
      "price": 15.80,
      "inventory": 100,
      "exhibition_url": "https://example.com/drug123.jpg"
    },
    {
      "id": 456,
      "user_id": 1,
      "drug_id": 456,
      "number": 1,
      "drug_name": "感冒灵颗粒",
      "specification": "10g*12袋",
      "price": 28.50,
      "inventory": 50,
      "exhibition_url": "https://example.com/drug456.jpg"
    }
  ]
}
```

## Redis存储结构

### 存储方式
- **存储类型**: Redis Hash
- **Key格式**: `cart:user:{user_id}`
- **Field格式**: `drug:{drug_id}`
- **Value格式**: JSON字符串

### 存储示例
```
Key: cart:user:1
Fields:
  drug:123 -> {"id":123,"user_id":1,"drug_id":123,"number":2,"drug_name":"阿莫西林胶囊",...}
  drug:456 -> {"id":456,"user_id":1,"drug_id":456,"number":1,"drug_name":"感冒灵颗粒",...}
```

## 业务逻辑

### 1. 添加商品到购物车
1. 检查药品是否存在
2. 检查库存是否充足
3. 如果购物车中已存在该商品，则累加数量
4. 如果是新商品，则直接添加
5. 存储到Redis哈希中

### 2. 修改商品数量
1. 检查购物车中是否存在该商品
2. 检查新数量是否超过库存
3. 如果数量为0，则删除该商品
4. 更新Redis中的数据

### 3. 删除商品
1. 根据用户ID和药品ID列表
2. 从Redis哈希中删除对应的字段
3. 支持批量删除操作

### 4. 获取购物车列表
1. 从Redis哈希中获取用户的所有购物车数据
2. 实时获取最新的药品信息（价格、库存等）
3. 确保数据一致性
4. 返回完整的购物车列表

## 数据一致性保证

### 库存检查
- 添加商品时检查库存
- 修改数量时检查库存
- 确保不会超卖

### 实时数据更新
- 获取购物车列表时，实时从数据库获取最新的药品信息
- 确保价格、库存等信息的准确性

### 错误处理
- 库存不足时返回明确的错误信息
- 商品不存在时返回相应错误
- 购物车项目不存在时返回相应错误

## 测试示例

### 使用curl测试

#### 1. 添加商品到购物车
```bash
curl -X POST "http://localhost:8000/v1/cart/create" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "drug_id": 123,
    "number": 2
  }'
```

#### 2. 修改商品数量
```bash
curl -X POST "http://localhost:8000/v1/cart/update" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "drug_id": 123,
    "number": 5
  }'
```

#### 3. 删除商品
```bash
curl -X POST "http://localhost:8000/v1/cart/delete" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "drug_ids": [123, 456]
  }'
```

#### 4. 获取购物车列表
```bash
curl -X POST "http://localhost:8000/v1/cart/list" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1
  }'
```

## 注意事项

1. **Redis连接**: 确保Redis服务正常运行，配置信息在 `configs/config.yaml` 中
2. **数据库依赖**: 需要 `mt_drug` 表存在，用于获取药品信息和检查库存
3. **并发安全**: Redis操作是原子性的，支持并发访问
4. **数据过期**: 可以根据需要设置购物车数据的过期时间
5. **用户认证**: 建议在实际使用中添加用户认证中间件

## 扩展功能建议

1. **购物车数据过期**: 可以设置购物车数据的TTL
2. **购物车统计**: 添加购物车商品总数、总价等统计信息
3. **购物车同步**: 支持多端购物车数据同步
4. **购物车分享**: 支持购物车分享功能
5. **批量操作**: 支持批量修改商品数量