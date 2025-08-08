# 医生API使用示例

## 接口说明

根据你提供的数据库表结构，我已经实现了完整的医生管理API业务逻辑。

## API接口列表

### 1. 发送短信验证码 - SendSms

**请求参数：**
```json
{
  "Phone": "13800138000",
  "SendSmsCode": "" // 此字段在发送短信时不需要
}
```

**响应：**
```json
{
  "Message": "短信发送成功",
  "Code": 200
}
```

### 2. 医生注册 - RegisterDoctor

**请求参数：**
```json
{
  "Phone": "13800138000",
  "Password": "123456",
  "SendSmsCode": "123456"
}
```

**响应：**
```json
{
  "Message": "注册成功",
  "Code": 200
}
```

### 3. 医生登录 - LoginDoctor

**请求参数：**
```json
{
  "Phone": "13800138000",
  "Password": "123456",
  "SendSmsCode": "123456"
}
```

**响应：**
```json
{
  "Message": "登录成功",
  "Code": 200,
  "DId": 1
}
```

### 4. 医生认证 - Authentication

**请求参数：**
```json
{
  "DId": 1,
  "Name": "张医生",
  "Gender": "男",
  "BirthData": "1980-01-01",
  "Email": "doctor@example.com",
  "Avatar": "http://example.com/avatar.jpg",
  "LicenseNumber": "110123456789",
  "DepartmentId": 1,
  "HospitalId": 1,
  "Title": "主治医师",
  "Speciality": "心血管内科",
  "PracticeScope": "内科诊疗"
}
```

**响应：**
```json
{
  "Message": "认证信息提交成功，请等待审核",
  "Code": 200
}
```

## 业务流程

1. **注册流程：**
   - 调用 `SendSms` 发送验证码
   - 调用 `RegisterDoctor` 完成注册

2. **登录流程：**
   - 调用 `SendSms` 发送验证码
   - 调用 `LoginDoctor` 完成登录

3. **认证流程：**
   - 登录后调用 `Authentication` 提交认证信息
   - 系统将状态设置为"待审核"

## 实现特性

### 安全特性
- 密码使用 bcrypt 加密存储
- 短信验证码5分钟过期
- 参数验证和错误处理

### 数据库特性
- 自动生成医生编码
- 支持软删除
- 记录最后登录时间和IP
- 完整的索引优化

### 缓存特性
- Redis存储验证码
- 支持医生信息缓存
- 自动过期机制

### 短信特性
- 集成阿里云短信服务
- 支持验证码发送
- 手机号格式验证

## 错误码说明

- `200`: 成功
- `400`: 参数错误
- `500`: 服务器内部错误

## 数据库状态说明

- `Status` 字段：
  - `"0"`: 禁用
  - `"1"`: 启用
  - `"2"`: 待审核

## 配置要求

确保 `configs/config.yaml` 中包含：

```yaml
data:
  database:
    driver: mysql
    source: "your_mysql_connection_string"
  redis:
    addr: "your_redis_address"
    password: "your_redis_password"
    db: 0
  aliYun:
    accessKeyID: "your_aliyun_access_key_id"
    accessKeySecret: "your_aliyun_access_key_secret"
```

## 注意事项

1. 短信模板需要在阿里云控制台配置
2. 数据库表已存在，无需自动迁移
3. 验证码在测试环境会在日志中显示
4. 生产环境建议使用HTTPS
5. 建议添加接口限流和防刷机制