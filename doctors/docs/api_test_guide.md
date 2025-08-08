# API 接口测试指南

## 服务信息
- **HTTP 服务**: http://localhost:8000
- **gRPC 服务**: localhost:9000

## 可用的 HTTP 接口

### 1. 发送短信验证码
**POST** `http://localhost:8000/api/v1/doctor/sms/send`

**请求体:**
```json
{
  "Phone": "13800138000"
}
```

**响应示例:**
```json
{
  "Message": "短信发送成功",
  "Code": 200
}
```

### 2. 医生注册
**POST** `http://localhost:8000/api/v1/doctor/register`

**请求体:**
```json
{
  "Phone": "13800138000",
  "Password": "123456",
  "SendSmsCode": "123456"
}
```

**响应示例:**
```json
{
  "Message": "注册成功",
  "Code": 200
}
```

### 3. 医生登录
**POST** `http://localhost:8000/api/v1/doctor/login`

**请求体:**
```json
{
  "Phone": "13800138000",
  "Password": "123456",
  "SendSmsCode": "123456"
}
```

**响应示例:**
```json
{
  "Message": "登录成功",
  "Code": 200,
  "DId": 1
}
```

### 4. 医生认证
**POST** `http://localhost:8000/api/v1/doctor/authentication`

**请求体:**
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

**响应示例:**
```json
{
  "Message": "认证信息提交成功，请等待审核",
  "Code": 200
}
```

## 使用工具测试

### 1. 使用 curl 测试

#### 发送短信验证码:
```bash
curl -X POST http://localhost:8000/api/v1/doctor/sms/send \
  -H "Content-Type: application/json" \
  -d '{"Phone": "13800138000"}'
```

#### 医生注册:
```bash
curl -X POST http://localhost:8000/api/v1/doctor/register \
  -H "Content-Type: application/json" \
  -d '{
    "Phone": "13800138000",
    "Password": "123456",
    "SendSmsCode": "123456"
  }'
```

#### 医生登录:
```bash
curl -X POST http://localhost:8000/api/v1/doctor/login \
  -H "Content-Type: application/json" \
  -d '{
    "Phone": "13800138000",
    "Password": "123456",
    "SendSmsCode": "123456"
  }'
```

### 2. 使用 Postman 测试

1. 创建新的 Collection
2. 添加请求，设置方法为 POST
3. 设置 URL 为上述接口地址
4. 在 Headers 中添加 `Content-Type: application/json`
5. 在 Body 中选择 raw -> JSON，粘贴请求体
6. 点击 Send 发送请求

### 3. 使用 Insomnia 测试

1. 创建新的请求
2. 设置方法为 POST
3. 输入 URL
4. 在 Body 选项卡中选择 JSON
5. 粘贴请求体内容
6. 点击 Send

### 4. 使用 VS Code REST Client 扩展

创建 `.http` 文件：

```http
### 发送短信验证码
POST http://localhost:8000/api/v1/doctor/sms/send
Content-Type: application/json

{
  "Phone": "13800138000"
}

### 医生注册
POST http://localhost:8000/api/v1/doctor/register
Content-Type: application/json

{
  "Phone": "13800138000",
  "Password": "123456",
  "SendSmsCode": "123456"
}

### 医生登录
POST http://localhost:8000/api/v1/doctor/login
Content-Type: application/json

{
  "Phone": "13800138000",
  "Password": "123456",
  "SendSmsCode": "123456"
}

### 医生认证
POST http://localhost:8000/api/v1/doctor/authentication
Content-Type: application/json

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

## 测试流程

1. **发送验证码**: 先调用发送短信接口，查看服务器日志获取验证码
2. **注册账号**: 使用手机号、密码和验证码注册
3. **登录获取ID**: 登录成功后会返回医生ID
4. **提交认证**: 使用医生ID提交认证信息

## 注意事项

1. **验证码获取**: 由于使用模拟发送，验证码会在服务器日志中显示
2. **数据库状态**: 注册后医生状态为"2"（待审核）
3. **密码加密**: 密码会自动使用 bcrypt 加密存储
4. **错误处理**: 接口会返回详细的错误信息
5. **Redis缓存**: 
   - 验证码存储5分钟自动过期
   - 登录成功后医生信息缓存1小时
   - 防止验证码频繁发送（1分钟内限制）
6. **验证码限制**: 同一手机号1分钟内只能发送一次验证码

## 日志查看

启动服务后，在终端中可以看到：
- 数据库连接状态
- Redis连接状态
- 短信验证码（模拟发送时显示）
- 各种操作的日志信息

## 常见错误码

- `200`: 成功
- `400`: 参数错误（如手机号格式不正确、密码太短等）
- `500`: 服务器内部错误（如数据库连接失败、验证码错误等）