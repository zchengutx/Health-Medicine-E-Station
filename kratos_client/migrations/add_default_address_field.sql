-- 添加is_default字段到mt_address表
-- 如果字段不存在，则添加该字段

ALTER TABLE mt_address 
ADD COLUMN IF NOT EXISTS is_default TINYINT(1) DEFAULT 0 COMMENT '是否默认地址';

-- 为现有数据创建索引（可选）
CREATE INDEX IF NOT EXISTS idx_user_default ON mt_address(user_id, is_default);
