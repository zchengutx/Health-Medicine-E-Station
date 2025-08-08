# 修复审核API接口对接问题

## 问题分析

审核详情接口没有对接上，可能的原因：

1. **路由未正确注册** - 已修复
2. **数据库表未创建** - 需要检查
3. **认证问题** - 需要检查
4. **API路径问题** - 需要检查

## 修复步骤

### 1. 检查数据库表

运行以下SQL检查表是否存在：

```sql
-- 检查表是否存在
SHOW TABLES LIKE 'mt_doctor_approval';

-- 如果不存在，创建表
CREATE TABLE IF NOT EXISTS `mt_doctor_approval` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `doctor_id` bigint unsigned NOT NULL COMMENT '医生ID',
  `doctor_name` varchar(50) DEFAULT NULL COMMENT '医生姓名',
  `approval_status` varchar(20) NOT NULL COMMENT '审核状态：0-未通过，1-通过，2-待审核',
  `approval_time` datetime(3) DEFAULT NULL COMMENT '审核时间',
  `approver_id` bigint unsigned DEFAULT NULL COMMENT '审核人ID',
  `approver_name` varchar(50) DEFAULT NULL COMMENT '审核人姓名',
  `approval_reason` varchar(500) DEFAULT NULL COMMENT '审核理由',
  `reject_reason` varchar(500) DEFAULT NULL COMMENT '拒绝理由',
  `submit_time` datetime(3) DEFAULT NULL COMMENT '提交审核时间',
  `created_by` bigint unsigned DEFAULT NULL,
  `updated_by` bigint unsigned DEFAULT NULL,
  `deleted_by` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_mt_doctor_approval_deleted_at` (`deleted_at`),
  KEY `idx_mt_doctor_approval_doctor_id` (`doctor_id`),
  KEY `idx_mt_doctor_approval_status` (`approval_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='医生审核记录表';

-- 插入测试数据
INSERT INTO mt_doctor_approval (doctor_id, doctor_name, approval_status, approval_time, approver_id, approver_name, approval_reason, reject_reason, submit_time, created_at, updated_at) VALUES
(1, '张医生', '1', NOW() - INTERVAL 24 HOUR, 1, '管理员', '资质齐全，审核通过', NULL, NOW() - INTERVAL 25 HOUR, NOW(), NOW()),
(2, '李医生', '0', NOW() - INTERVAL 12 HOUR, 1, '管理员', NULL, '执业证书过期，需要重新提交', NOW() - INTERVAL 13 HOUR, NOW(), NOW()),
(3, '王医生', '2', NULL, NULL, NULL, NULL, NULL, NOW() - INTERVAL 6 HOUR, NOW(), NOW()),
(4, '赵医生', '1', NOW() - INTERVAL 2 HOUR, 1, '管理员', '所有材料符合要求，审核通过', NULL, NOW() - INTERVAL 3 HOUR, NOW(), NOW()),
(5, '刘医生', '0', NOW() - INTERVAL 1 HOUR, 1, '管理员', NULL, '缺少必要的执业证书复印件', NOW() - INTERVAL 2 HOUR, NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();
```

### 2. 检查路由注册

确保在 `router_biz.go` 中已添加：

```go
medicineRouter.InitMtDoctorApprovalRouter(privateGroup, publicGroup)
```

### 3. 测试API接口

使用以下方法测试：

1. **使用浏览器开发者工具**：
   - 打开医生列表页面
   - 点击"审核详情"
   - 查看Network标签页中的请求

2. **使用测试页面**：
   - 打开 `test_approval.html`
   - 点击测试按钮

3. **使用curl命令**：
```bash
# 测试获取审核记录列表
curl -X GET "http://localhost:8888/api/mtDoctorApproval/getMtDoctorApprovalList?page=1&pageSize=10" \
  -H "x-token: YOUR_TOKEN" \
  -H "x-user-id: YOUR_USER_ID"

# 测试根据医生ID获取审核记录
curl -X GET "http://localhost:8888/api/mtDoctorApproval/getMtDoctorApprovalByDoctorId?doctorId=1" \
  -H "x-token: YOUR_TOKEN" \
  -H "x-user-id: YOUR_USER_ID"
```

### 4. 检查认证

确保请求头包含正确的认证信息：
- `x-token`: 用户token
- `x-user-id`: 用户ID

### 5. 调试步骤

1. **检查后端日志**：
   - 启动后端服务
   - 查看控制台输出

2. **检查前端网络请求**：
   - 打开浏览器开发者工具
   - 查看Network标签页
   - 检查请求URL、参数和响应

3. **检查数据库连接**：
   - 确保数据库服务正常运行
   - 确保表已创建并有数据

### 6. 常见问题解决

1. **404错误**：路由未正确注册
2. **401错误**：认证失败，检查token
3. **500错误**：服务器内部错误，检查后端日志
4. **数据库连接错误**：检查数据库配置

### 7. 验证步骤

1. 重启后端服务
2. 重启前端服务
3. 登录系统
4. 进入医生列表页面
5. 点击"审核详情"
6. 检查弹框是否正常显示

## 预期结果

修复后，点击"审核详情"应该：
1. 正常发送API请求
2. 获取到审核数据
3. 显示审核详情弹框
4. 显示审核历史时间线 