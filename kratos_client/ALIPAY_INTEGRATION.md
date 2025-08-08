# 支付宝沙箱支付集成指南

## 概述

本项目已集成支付宝沙箱支付功能，支持创建支付订单、处理支付回调、查询支付状态、申请退款等功能。

## 配置说明

### 需要的配置信息

你需要从支付宝开放平台沙箱获取以下3个配置：

1. **AppID** - 沙箱应用ID
2. **应用私钥** - 你生成的RSA私钥  
3. **支付宝公钥** - 支付宝提供的公钥

### 获取配置步骤

#### 1. 获取AppID
- 访问 [支付宝开放平台](https://open.alipay.com/)
- 登录后进入"开发者中心" -> "沙箱应用"
- 复制显示的AppID（如：2021000122671234）

#### 2. 生成应用密钥对
- 下载支付宝提供的[密钥生成工具](https://opendocs.alipay.com/common/02kipl)
- 生成RSA2(SHA256)密钥对
- 将**应用公钥**上传到支付宝开放平台
- 保存**应用私钥**用于配置

#### 3. 获取支付宝公钥
- 在支付宝开放平台上传应用公钥后
- 支付宝会生成对应的**支付宝公钥**
- 复制支付宝公钥用于配置

### 配置文件

编辑 `comment/alipay_config.go` 文件中的 `GetAlipayConfig()` 函数：

```go
func GetAlipayConfig() AlipayConfig {
    return AlipayConfig{
        AppID: "你的沙箱AppID",
        PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
你的应用私钥内容
-----END RSA PRIVATE KEY-----`,
        AlipayPublicKey: `-----BEGIN PUBLIC KEY-----
支付宝公钥内容
-----END PUBLIC KEY-----`,
        IsProduction: false,
        NotifyURL:    "http://localhost:8000/v1/payment/notify",
        ReturnURL:    "http://localhost:8000/v1/payment/return",
    }
}
```

## 快速测试

如果你想快速测试功能，可以使用内置的测试配置：

在 `comment/alipay_config.go` 中的 `getDefaultConfig()` 函数中：

```go
func getDefaultConfig() AlipayConfig {
    return GetTestAlipayConfig() // 使用测试配置
    // return GetAlipayConfig()  // 使用你自己的配置
}
```

## API接口

### 1. 创建支付订单

**接口**: `POST /v1/payment/create`

**请求参数**:
```json
{
    "token": "用户JWT token",
    "subject": "商品标题",
    "total_amount": "支付金额",
    "order_type": "订单类型",
    "business_id": "业务ID",
    "description": "订单描述"
}
```

**响应**:
```json
{
    "code": 0,
    "message": "success",
    "payment_url": "支付链接",
    "order_id": "订单号"
}
```

### 2. 支付回调通知

**接口**: `POST /v1/payment/notify`

支付宝会自动调用此接口通知支付结果。

### 3. 查询支付状态

**接口**: `GET /v1/payment/query?order_id=订单号&token=用户token`

**响应**:
```json
{
    "code": 0,
    "message": "success",
    "payment_info": {
        "order_id": "订单号",
        "trade_no": "支付宝交易号",
        "total_amount": "支付金额",
        "trade_status": "交易状态",
        "pay_time": "支付时间",
        "subject": "商品标题"
    }
}
```

### 4. 申请退款

**接口**: `POST /v1/payment/refund`

**请求参数**:
```json
{
    "token": "用户JWT token",
    "order_id": "订单号",
    "refund_amount": "退款金额",
    "refund_reason": "退款原因"
}
```

## 使用示例

### 创建支付订单

```bash
curl -X POST "http://localhost:8000/v1/payment/create" \
  -H "Content-Type: application/json" \
  -d '{
    "token": "your_jwt_token",
    "subject": "药品购买",
    "total_amount": "99.99",
    "order_type": "drug_order",
    "business_id": "drug_123",
    "description": "购买感冒药"
  }'
```

### 查询支付状态

```bash
curl "http://localhost:8000/v1/payment/query?order_id=ORDER_123&token=your_jwt_token"
```

## 测试流程

1. **启动服务器**
   ```bash
   cd kratos_client
   make build
   ./bin/kratos_client.exe -conf ./configs/config.yaml
   ```

2. **创建支付订单**
   - 调用创建支付接口
   - 获取支付链接

3. **完成支付**
   - 在浏览器中打开支付链接
   - 使用沙箱买家账号完成支付

4. **验证结果**
   - 查看支付回调日志
   - 调用查询接口确认支付状态

## 沙箱测试账号

支付宝沙箱提供测试买家账号：
- 买家账号：jjdltb4986@sandbox.com
- 登录密码：111111
- 支付密码：111111

## 注意事项

1. **环境配置**: 确保 `IsProduction` 设为 `false`
2. **回调地址**: NotifyURL 必须是外网可访问的地址
3. **签名验证**: 所有回调都会进行签名验证
4. **订单号**: 确保订单号唯一性
5. **金额格式**: 金额必须是字符串格式，保留两位小数

## 常见问题

### 1. 签名验证失败
- 检查应用私钥是否正确
- 确认支付宝公钥是否匹配
- 验证参数编码格式

### 2. 回调接收不到
- 确认NotifyURL是否外网可访问
- 检查防火墙设置
- 查看服务器日志

### 3. 支付页面打不开
- 检查AppID是否正确
- 确认沙箱环境配置
- 验证参数格式

## 生产环境部署

切换到生产环境时需要：

1. 将 `IsProduction` 设为 `true`
2. 使用正式环境的AppID和密钥
3. 配置正式的回调地址
4. 申请支付宝正式应用

## 技术支持

- [支付宝开放平台文档](https://opendocs.alipay.com/)
- [沙箱环境说明](https://opendocs.alipay.com/common/02kkv7)
- [API参考文档](https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay)