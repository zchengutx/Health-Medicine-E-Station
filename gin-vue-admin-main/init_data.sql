-- 初始化药品状态字典
INSERT INTO sys_dictionaries (name, type, status, desc, created_at, updated_at) 
VALUES ('药品状态', 'drug_status', 1, '药品状态字典', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

-- 获取药品状态字典ID
SET @dict_id = (SELECT id FROM sys_dictionaries WHERE type = 'drug_status' LIMIT 1);

-- 插入药品状态字典详情
INSERT INTO sys_dictionary_details (label, value, status, sort, sys_dictionary_id, created_at, updated_at) 
VALUES 
('禁用', '0', 1, 1, @dict_id, NOW(), NOW()),
('启用', '1', 1, 2, @dict_id, NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

-- 插入一级分类数据
INSERT INTO mt_drug_type_stairs (stair_name, created_at, updated_at) 
VALUES 
('处方药', NOW(), NOW()),
('非处方药', NOW(), NOW()),
('保健品', NOW(), NOW()),
('医疗器械', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

-- 插入二级分类数据
INSERT INTO mt_drug_type_levels (level_name, created_at, updated_at) 
VALUES 
('感冒药', NOW(), NOW()),
('消炎药', NOW(), NOW()),
('维生素', NOW(), NOW()),
('钙片', NOW(), NOW()),
('血压计', NOW(), NOW()),
('体温计', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

-- 验证数据
SELECT '药品状态字典' as table_name, COUNT(*) as count FROM sys_dictionaries WHERE type = 'drug_status'
UNION ALL
SELECT '药品状态字典详情' as table_name, COUNT(*) as count FROM sys_dictionary_details WHERE sys_dictionary_id = @dict_id
UNION ALL
SELECT '一级分类' as table_name, COUNT(*) as count FROM mt_drug_type_stairs
UNION ALL
SELECT '二级分类' as table_name, COUNT(*) as count FROM mt_drug_type_levels; 