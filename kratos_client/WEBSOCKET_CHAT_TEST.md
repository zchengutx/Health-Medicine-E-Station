# WebSocket 聊天功能测试文档

## 测试环境准备

### 1. 数据库准备

首先确保数据库中存在必要的表结构：

```sql
-- 创建聊天房间表
CREATE TABLE mt_chat_room (
    id                BIGINT AUTO_INCREMENT PRIMARY KEY,
    room_id           VARCHAR(50) NOT NULL UNIQUE COMMENT '房间ID',
    user1_id          INT DEFAULT 0 NOT NULL COMMENT '用户1ID',
    user2_id          INT DEFAULT 0 NOT NULL COMMENT '用户2ID',
    last_message      TEXT NULL COMMENT '最后一条消息',
    last_message_time DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) NULL COMMENT '最后消息时间',
    created_at        DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at        DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) NOT NULL
) COMMENT '聊天房间';

-- 创建聊天消息表
CREATE TABLE mt_chat_message (
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    from_id      INT NOT NULL COMMENT '发送者ID',
    to_id        INT NOT NULL COMMENT '接收者ID',
    content      TEXT COMMENT '消息内容',
    message_type VARCHAR(20) NOT NULL DEFAULT 'text' COMMENT '消息类型',
    room_id      VARCHAR(50) NOT NULL COMMENT '房间ID',
    is_read      TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否已读',
    created_at   DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at   DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    
    INDEX idx_room_id (room_id),
    INDEX idx_from_to (from_id, to_id),
    INDEX idx_created_at (created_at)
) COMMENT '聊天消息';
```

### 2. 启动服务

```bash
cd kratos_client
make build
./bin/kratos_client
```

服务默认启动在 `http://localhost:8000`

## 测试方法

### 方法一：使用内置测试页面（推荐）

#### 1. 访问测试页面
打开浏览器访问：`http://localhost:8000/chat`

#### 2. 模拟两个用户聊天
- **步骤1**: 打开两个浏览器标签页或窗口
- **步骤2**: 在第一个页面设置用户信息：
  - 用户ID: `1`
  - 用户名: `患者张三`
  - 用户角色: `患者`
  - 目标用户ID: `2`
  - 点击"连接"按钮

- **步骤3**: 在第二个页面设置用户信息：
  - 用户ID: `2`
  - 用户名: `药师李四`
  - 用户角色: `药师`
  - 目标用户ID: `1`
  - 点击"连接"按钮

#### 3. 测试消息发送
- 在任一页面输入消息并发送
- 观察另一页面是否实时收到消息
- 测试双向通信

### 方法二：使用WebSocket客户端工具

#### 1. 推荐工具
- **Postman** (支持WebSocket)
- **WebSocket King** (Chrome扩展)
- **wscat** (命令行工具)

#### 2. 连接参数
```
WebSocket URL: ws://localhost:8000/ws/chat
查询参数:
- user_id: 用户ID (如: 1)
- user_name: 用户名 (如: 张三)
- user_role: 用户角色 (patient/doctor)
- target_id: 目标用户ID (如: 2)
```

#### 3. 完整连接示例
```
ws://localhost:8000/ws/chat?user_id=1&user_name=张三&user_role=patient&target_id=2
```

#### 4. 发送消息格式
```json
{
  "type": "text",
  "content": "你好，我想咨询一下药品信息",
  "to_id": 2
}
```

### 方法三：使用JavaScript代码测试

#### 1. 创建测试HTML文件
```html
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket测试</title>
</head>
<body>
    <div id="messages"></div>
    <input type="text" id="messageInput" placeholder="输入消息">
    <button onclick="sendMessage()">发送</button>

    <script>
        const ws = new WebSocket('ws://localhost:8000/ws/chat?user_id=1&user_name=测试用户&user_role=patient&target_id=2');
        
        ws.onopen = function(event) {
            console.log('连接已建立');
            document.getElementById('messages').innerHTML += '<p>连接成功</p>';
        };
        
        ws.onmessage = function(event) {
            const message = JSON.parse(event.data);
            console.log('收到消息:', message);
            document.getElementById('messages').innerHTML += 
                `<p><strong>${message.from_name}:</strong> ${message.content}</p>`;
        };
        
        ws.onclose = function(event) {
            console.log('连接已关闭');
            document.getElementById('messages').innerHTML += '<p>连接已断开</p>';
        };
        
        function sendMessage() {
            const input = document.getElementById('messageInput');
            const message = {
                type: 'text',
                content: input.value,
                to_id: 2
            };
            ws.send(JSON.stringify(message));
            document.getElementById('messages').innerHTML += 
                `<p><strong>我:</strong> ${input.value}</p>`;
            input.value = '';
        }
    </script>
</body>
</html>
```

## 测试用例

### 基础功能测试

#### 测试用例1: WebSocket连接建立
- **测试目标**: 验证WebSocket连接能够正常建立
- **测试步骤**:
  1. 使用正确的参数连接WebSocket
  2. 观察连接状态
- **预期结果**: 连接成功，收到连接确认

#### 测试用例2: 消息发送与接收
- **测试目标**: 验证一对一消息传递功能
- **测试步骤**:
  1. 建立两个WebSocket连接（用户1和用户2）
  2. 用户1发送消息给用户2
  3. 观察用户2是否收到消息
- **预期结果**: 用户2实时收到用户1的消息

#### 测试用例3: 双向通信
- **测试目标**: 验证双向消息传递
- **测试步骤**:
  1. 用户1发送消息给用户2
  2. 用户2回复消息给用户1
- **预期结果**: 双方都能收到对方的消息

#### 测试用例4: 消息持久化
- **测试目标**: 验证消息保存到数据库
- **测试步骤**:
  1. 发送几条消息
  2. 查询数据库中的消息记录
- **预期结果**: 数据库中存在相应的消息记录

### 异常情况测试

#### 测试用例5: 参数缺失
- **测试目标**: 验证参数验证功能
- **测试步骤**:
  1. 使用缺少必要参数的URL连接
- **预期结果**: 连接被拒绝，返回400错误

#### 测试用例6: 无效用户ID
- **测试目标**: 验证用户ID验证
- **测试步骤**:
  1. 使用非数字的用户ID连接
- **预期结果**: 连接被拒绝，返回400错误

#### 测试用例7: 连接断开重连
- **测试目标**: 验证断线重连功能
- **测试步骤**:
  1. 建立连接后主动断开
  2. 重新连接
- **预期结果**: 能够重新建立连接

#### 测试用例8: 目标用户离线
- **测试目标**: 验证离线消息处理
- **测试步骤**:
  1. 只有用户1在线
  2. 用户1发送消息给离线的用户2
- **预期结果**: 消息保存到数据库，用户2上线后可以查看历史消息

### 性能测试

#### 测试用例9: 并发连接
- **测试目标**: 验证服务器并发处理能力
- **测试步骤**:
  1. 同时建立多个WebSocket连接
  2. 观察服务器性能
- **预期结果**: 服务器能够稳定处理多个连接

#### 测试用例10: 消息频率
- **测试目标**: 验证高频消息处理
- **测试步骤**:
  1. 快速连续发送多条消息
  2. 观察消息处理情况
- **预期结果**: 所有消息都能正确处理和传递

## HTTP API 测试

### 1. 获取聊天历史记录
```bash
curl "http://localhost:8000/v1/chat/history?user_id=1&target_id=2&page=1&page_size=10"
```

### 2. 保存聊天消息
```bash
curl -X POST "http://localhost:8000/v1/chat/message" \
  -H "Content-Type: application/json" \
  -d '{
    "from_id": 1,
    "to_id": 2,
    "content": "测试消息",
    "message_type": "text",
    "room_id": "room_1_2"
  }'
```

### 3. 获取用户聊天房间列表
```bash
curl "http://localhost:8000/v1/chat/rooms?user_id=1"
```

### 4. 获取房间在线用户
```bash
curl "http://localhost:8000/api/chat/room/room_1_2/users"
```

## 常见问题排查

### 1. 连接失败
- 检查服务是否正常启动
- 确认端口号是否正确
- 检查防火墙设置

### 2. 消息发送失败
- 检查消息格式是否正确
- 确认目标用户ID是否有效
- 查看服务器日志

### 3. 数据库相关问题
- 确认数据库连接正常
- 检查表结构是否正确创建
- 查看数据库日志

### 4. 前端页面问题
- 检查浏览器控制台错误
- 确认WebSocket支持
- 检查网络连接

## 测试报告模板

### 测试环境
- 操作系统: 
- 浏览器版本: 
- 服务器版本: 
- 数据库版本: 

### 测试结果
| 测试用例 | 测试结果 | 备注 |
|---------|---------|------|
| WebSocket连接建立 | ✅/❌ | |
| 消息发送与接收 | ✅/❌ | |
| 双向通信 | ✅/❌ | |
| 消息持久化 | ✅/❌ | |
| 参数验证 | ✅/❌ | |
| 异常处理 | ✅/❌ | |

### 发现的问题
1. 问题描述
2. 重现步骤
3. 预期结果
4. 实际结果
5. 严重程度

### 建议
- 功能改进建议
- 性能优化建议
- 用户体验改进建议

---

## 注意事项

1. **测试数据**: 使用测试环境，避免影响生产数据
2. **并发测试**: 注意服务器资源限制
3. **网络环境**: 确保网络连接稳定
4. **浏览器兼容性**: 测试不同浏览器的兼容性
5. **移动端测试**: 如需要，测试移动端浏览器

通过以上测试方法和用例，可以全面验证WebSocket聊天功能的正确性和稳定性。