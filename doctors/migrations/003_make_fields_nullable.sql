-- 修改name和license_number字段为可空
ALTER TABLE doctors MODIFY COLUMN name varchar(50) DEFAULT NULL COMMENT '医生姓名';
ALTER TABLE doctors MODIFY COLUMN license_number varchar(50) DEFAULT NULL COMMENT '执业医师资格证号';

-- 删除可能存在的唯一约束（如果有的话）
ALTER TABLE doctors DROP INDEX IF EXISTS uk_license_number;

-- 重新创建唯一约束，但允许NULL值
ALTER TABLE doctors ADD CONSTRAINT uk_license_number UNIQUE (license_number);