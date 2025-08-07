# WebSocket 一对一聊天功能

## 功能概述

实现了药师与患者之间的一对一实时聊天功能，基于WebSocket协议。

## 主要特性

- ✅ 一对一实时聊天
- ✅ 消息持久化存储
- ✅ 聊天历史记录查询
- ✅ 用户在线状态管理
- ✅ 消息已读状态跟踪
- ✅ 支持文本消息

## API 接口

### WebSocket 连接
```
GET /ws/chat?user_id={用户ID}&user_name={用户名}&user_role={角色}&target_id={目标用户ID}
```

参数说明：
- `user_id`: 当前用户ID
- `user_name`: 当前用户名
- `user_role`: 用户角色 (doctor/patient)
- `target_id`: 聊天对象的用户ID

### HTTP API

#### 1. 获取聊天历史记录
```
GET /v1/chat/history?user_id={用户ID}&target_id={目标用户ID}&page={页码}&page_size={每页数量}
```

#### 2. 保存聊天消息
```
POST /v1/chat/message
Content-Type: application/json

{
  "from_id": 1,
  "to_id": 2,
  "content": "消息内容",
  "message_type": "text",
  "room_id": "room_1_2"
}
```

#### 3. 获取用户聊天房间列表
```
GET /v1/chat/rooms?user_id={用户ID}
```

#### 4. 获取房间在线用户
```
GET /api/chat/room/{roomId}/users
```

## 数据库表结构

### 聊天消息表 (mt_chat_message)
```sql
CREATE TABLE mt_chat_message (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    from_id INT NOT NULL COMMENT '发送者ID',
    to_id INT NOT NULL COMMENT '接收者ID',
    content TEXT COMMENT '消息内容',
    message_type VARCHAR(20) NOT NULL DEFAULT 'text' COMMENT '消息类型',
    room_id VARCHAR(50) NOT NULL COMMENT '房间ID',
    is_read TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否已读',
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)
);
```

### 聊天房间表 (mt_chat_room)
```sql
CREATE TABLE mt_chat_room (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    room_id VARCHAR(50) NOT NULL UNIQUE COMMENT '房间ID',
    user1_id INT NOT NULL COMMENT '用户1ID',
    user2_id INT NOT NULL COMMENT '用户2ID',
    last_message TEXT COMMENT '最后一条消息',
    last_message_time DATETIME(6) COMMENT '最后消息时间',
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)
);
```

## 使用方法

### 1. 启动服务
```bash
cd kratos_client
make build
./bin/kratos_client
```

### 2. 测试聊天功能
访问: `http://localhost:8000/chat`

在测试页面中：
1. 填写用户信息（用户ID、用户名、角色、目标用户ID）
2. 点击"连接"按钮建立WebSocket连接
3. 在输入框中输入消息并发送
4. 可以开启多个浏览器标签页模拟不同用户

### 3. WebSocket 消息格式

#### 发送消息格式
```json
{
  "type": "text",
  "content": "消息内容",
  "to_id": 2
}
```

#### 接收消息格式
```json
{
  "type": "text",
  "content": "消息内容",
  "from_id": 1,
  "to_id": 2,
  "from_name": "发送者姓名",
  "timestamp": "2024-01-01T12:00:00Z",
  "room_id": "room_1_2"
}
```

## 房间ID规则

房间ID由两个用户ID生成，格式为 `room_{较小ID}_{较大ID}`，确保同一对用户始终使用相同的房间ID。

例如：用户1和用户2的房间ID为 `room_1_2`

## 注意事项

1. 确保数据库中存在相应的表结构
2. WebSocket连接需要提供完整的用户信息
3. 消息会自动保存到数据库中
4. 支持断线重连，历史消息可通过API获取
5. 生产环境中建议添加用户认证和权限验证