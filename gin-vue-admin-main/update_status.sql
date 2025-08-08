-- 更新现有药品分类的状态字段
-- 如果status字段为NULL，设置为1（启用）
UPDATE mt_drug_category 
SET status = 1 
WHERE status IS NULL;

-- 如果status字段为0，保持不变（禁用）
-- 如果status字段为1，保持不变（启用）

-- 验证更新结果
SELECT 
    id,
    mt_drug_type_stair_id,
    mt_drug_type_level_id,
    status,
    CASE 
        WHEN status = 0 THEN '禁用'
        WHEN status = 1 THEN '启用'
        ELSE '未知'
    END as status_text,
    created_at,
    updated_at
FROM mt_drug_category; 