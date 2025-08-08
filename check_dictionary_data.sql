-- 检查字典数据是否存在
SELECT * FROM sys_dictionaries WHERE type = 'drug_status';

-- 检查字典详情数据
SELECT * FROM sys_dictionary_details WHERE sys_dictionary_id IN (
    SELECT id FROM sys_dictionaries WHERE type = 'drug_status'
);

-- 如果没有数据，插入基础的药品状态字典
INSERT IGNORE INTO sys_dictionaries (created_at, updated_at, name, type, status, `desc`) 
VALUES (NOW(), NOW(), '药品状态', 'drug_status', 1, '药品状态字典');

-- 获取刚插入的字典ID
SET @dict_id = (SELECT id FROM sys_dictionaries WHERE type = 'drug_status' LIMIT 1);

-- 插入字典详情数据
INSERT IGNORE INTO sys_dictionary_details (created_at, updated_at, label, value, status, sort, sys_dictionary_id)
VALUES 
(NOW(), NOW(), '正常', 'normal', 1, 1, @dict_id),
(NOW(), NOW(), '停用', 'disabled', 1, 2, @dict_id),
(NOW(), NOW(), '缺货', 'out_of_stock', 1, 3, @dict_id);

-- 验证插入结果
SELECT 
    d.type,
    d.name,
    dd.label,
    dd.value,
    dd.sort
FROM sys_dictionaries d
LEFT JOIN sys_dictionary_details dd ON d.id = dd.sys_dictionary_id
WHERE d.type = 'drug_status'
ORDER BY dd.sort;