-- 检查药品分类数据
SELECT 
    id,
    mt_drug_type_stair_id,
    mt_drug_type_level_id,
    status,
    created_at,
    updated_at
FROM mt_drug_category;

-- 检查字典数据
SELECT 
    id,
    name,
    type,
    status,
    desc
FROM sys_dictionaries 
WHERE type = 'drug_status';

-- 检查字典详情
SELECT 
    id,
    label,
    value,
    status,
    sort,
    sys_dictionary_id
FROM sys_dictionary_details 
WHERE sys_dictionary_id = (
    SELECT id FROM sys_dictionaries WHERE type = 'drug_status'
);

-- 检查一级分类数据
SELECT 
    id,
    stair_name,
    created_at,
    updated_at
FROM mt_drug_type_stairs;

-- 检查二级分类数据
SELECT 
    id,
    level_name,
    created_at,
    updated_at
FROM mt_drug_type_levels; 