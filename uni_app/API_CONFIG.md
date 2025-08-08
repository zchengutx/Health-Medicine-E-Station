# API 接口配置说明

## 后端接口地址配置

### 1. HTML 版本配置
在 `login.html` 文件中，修改以下配置：

```javascript
// 后端 API 基础 URL，请根据实际情况修改
const API_BASE_URL = 'http://localhost:8000';
```

### 2. uni-app 版本配置
在 `pages/login/login.vue` 文件中，修改以下配置：

```javascript
// 发送验证码接口
url: 'http://localhost:8000/v1/sendSms'

// 登录接口
url: 'http://localhost:8000/v1/login'
```

### 3. API 配置文件
在 `src/config/api.js` 文件中，修改以下配置：

```javascript
const API_CONFIG = {
  // 后端服务基础地址，请根据实际情况修改
  BASE_URL: 'http://localhost:8000',
  // ...
}
```

## 接口说明

### 发送短信验证码
- **接口地址**: `POST /v1/sendSms`
- **请求参数**:
  ```json
  {
    "mobile": "手机号码",
    "source": "login"
  }
  ```
- **响应格式**:
  ```json
  {
    "success": true,
    "message": "发送成功"
  }
  ```

### 用户登录
- **接口地址**: `POST /v1/login`
- **请求参数**:
  ```json
  {
    "mobile": "手机号码",
    "sendSmsCode": "验证码"
  }
  ```
- **响应格式**:
  ```json
  {
    "success": true,
    "data": {
      "username": "用户名",
      "phone": "手机号",
      "avatar": "头像地址",
      "token": "登录令牌"
    }
  }
  ```

## 注意事项

1. **跨域问题**: 如果前端和后端不在同一域名下，需要后端配置 CORS 支持
2. **HTTPS**: 生产环境建议使用 HTTPS 协议
3. **错误处理**: 前端已配置错误处理，会显示后端返回的错误信息
4. **网络超时**: 建议配置适当的网络超时时间

## 测试步骤

1. 确保后端服务已启动并监听在配置的端口上
2. 修改前端代码中的 API 地址为实际的后端地址
3. 打开浏览器开发者工具，查看网络请求和控制台日志
4. 测试发送验证码和登录功能

## 常见问题

### 1. 网络请求失败
- 检查后端服务是否正常运行
- 检查 API 地址是否正确
- 检查网络连接是否正常

### 2. 跨域错误
- 后端需要配置 CORS 头信息
- 或者使用代理服务器

### 3. 验证码发送失败
- 检查手机号格式是否正确
- 检查后端短信服务是否正常
- 查看后端日志获取详细错误信息
