-- 地址管理功能数据库迁移脚本

-- 检查并添加is_default字段
ALTER TABLE mt_address 
ADD COLUMN IF NOT EXISTS is_default TINYINT(1) DEFAULT 0 COMMENT '是否默认地址';

-- 为现有数据创建索引
CREATE INDEX IF NOT EXISTS idx_user_default ON mt_address(user_id, is_default);

-- 显示表结构确认
DESCRIBE mt_address;
