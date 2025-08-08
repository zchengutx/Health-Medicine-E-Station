# 健康医疗驿站 (Health Medicine E-Station)

一个基于 uni-app 框架开发的健康医疗管理应用。

## 项目简介

健康医疗驿站是一个综合性的健康管理平台，为用户提供全方位的健康服务，包括健康咨询、药品管理、预约挂号、健康档案等功能。

## 功能特性

- 🏥 **健康咨询** - 专业医生在线咨询
- 💊 **药品管理** - 智能用药提醒和管理
- 📅 **预约挂号** - 在线医院预约服务
- 📋 **健康档案** - 完整的健康记录管理
- 🔐 **用户登录** - 手机号验证码登录

## 技术栈

- **框架**: uni-app
- **前端**: Vue 3 + TypeScript
- **构建工具**: Vite
- **支持平台**: H5、微信小程序、App

## 项目结构

```
uni_app/
├── pages/              # 页面文件
│   ├── index/         # 首页
│   └── login/         # 登录页
├── src/               # 源代码
│   └── utils/         # 工具函数
├── static/            # 静态资源
├── App.vue            # 根组件
├── main.js            # 入口文件
├── manifest.json      # 应用配置
├── pages.json         # 页面配置
├── vite.config.js     # Vite 配置
├── index.html         # 演示页面
├── login.html         # 登录演示页面
└── package.json       # 项目依赖
```

## 快速开始

### 方法一：使用演示页面（推荐）

1. 双击运行 `start.bat` 文件
2. 在浏览器中打开 `http://localhost:8080`
3. 查看项目演示页面
4. 点击右上角"登录"按钮查看登录页面

### 方法二：使用 HBuilderX（完整功能）

1. 下载并安装 [HBuilderX](https://www.dcloud.io/hbuilderx.html)
2. 用 HBuilderX 打开项目文件夹
3. 选择运行到浏览器或模拟器
4. 在首页点击"登录"按钮进入登录页面

### 方法三：命令行启动

```bash
# 安装依赖
npm install

# 启动演示服务器
python -m http.server 8080

# 或使用 Node.js
npx http-server -p 8080
```

## 页面说明

### 首页 (`pages/index/index.vue`)
- 展示应用主要功能模块
- 右上角有登录按钮
- 响应式设计，支持多种设备

### 登录页 (`pages/login/login.vue`)
- 手机号验证码登录
- 包含用户协议确认
- 表单验证和交互反馈

### 演示页面
- `index.html` - 项目功能展示页面
- `login.html` - 登录页面演示（包含 API 调用示例）

## 开发说明

### 当前状态

- ✅ 项目基础架构已完成
- ✅ 页面配置和路由设置完成
- ✅ 基础组件和样式已实现
- ✅ 演示页面可正常运行
- ✅ 登录页面功能完整
- ✅ 用户信息管理页面已完成
- ✅ 修改昵称功能已对接后端API
- 🔄 完整功能开发中

### API 接口说明

#### 1. 发送短信验证码
- **接口地址**: `POST /v1/sendSms`
- **请求参数**: 
  ```json
  {
    "mobile": "手机号",
    "source": "login"
  }
  ```
- **响应格式**: 
  ```json
  {
    "message": "sendSms success"
  }
  ```

#### 2. 用户登录
- **接口地址**: `POST /v1/login`
- **请求参数**: 
  ```json
  {
    "mobile": "手机号",
    "sendSmsCode": "验证码"
  }
  ```
- **响应格式**: 
  ```json
  {
    "message": "login success"
  }
  ```

#### 3. 修改昵称
- **接口地址**: `POST /v1/updateNickName`
- **请求头**: 
  ```
  Content-Type: application/json
  Authorization: <token>
  ```
- **请求参数**: 
  ```json
  {
    "nickName": "新昵称"
  }
  ```
- **响应格式**: 
  ```json
  {
    "message": "updateNickName success"
  }
  ```

#### 4. 获取用户信息
- **接口地址**: `POST /v1/userInfo`
- **请求头**: 
  ```
  Content-Type: application/json
  Authorization: <token>
  ```
- **请求参数**: 
  ```json
  {}
  ```
- **响应格式**: 
  ```json
  {
    "message": "userInfo success",
    "data": {
      "username": "用户昵称",
      "phone": "手机号",
      "avatar": "头像地址"
    }
  }
  ```

        #### 5. 修改头像
        - **接口地址**: `POST /upload`
- **请求头**: 
  ```
  Authorization: <token>
  ```
- **请求参数**: 
  ```
  FormData:
  file: 图片文件
  ```
- **响应格式**: 
  ```json
  {
    "message": "updateAvatar success",
    "data": {
      "avatar": "头像文件名"
    }
  }
  ```

### 注意事项

1. 这是一个 uni-app 项目，完整功能需要在 HBuilderX 中运行
2. 当前提供的 HTML 演示页面展示了项目的基本功能和界面
3. 登录页面中的 API 调用是示例代码，实际没有后端服务
4. 要开发完整功能，建议使用 HBuilderX 开发工具

## 浏览器兼容性

- Chrome 60+
- Firefox 55+
- Safari 12+
- Edge 79+

## 许可证

ISC License

## 联系方式

如有问题或建议，请联系开发团队。
