-- 测试多表联查SQL
SELECT 
    mt_drug.id,
    mt_drug.created_at,
    mt_drug.updated_at,
    mt_drug.deleted_at,
    mt_drug.drug_name,
    mt_drug.guide,
    mt_drug.explain,
    mt_drug.specification,
    mt_drug.price,
    mt_drug.sales_volume,
    mt_drug.inventory,
    mt_drug.status,
    mt_drug.created_by,
    mt_drug.updated_by,
    mt_drug.deleted_by,
    COALESCE(mt_guide.major_function, '') as guides,
    COALESCE(mt_explain.common_name, '') as explains
FROM mt_drug
LEFT JOIN mt_guide ON mt_drug.guide = mt_guide.id
LEFT JOIN mt_explain ON mt_drug.explain = mt_explain.id
WHERE mt_drug.deleted_at IS NULL
ORDER BY mt_drug.created_at DESC
LIMIT 10;

-- 检查是否有数据
SELECT COUNT(*) as drug_count FROM mt_drug WHERE deleted_at IS NULL;
SELECT COUNT(*) as guide_count FROM mt_guide;
SELECT COUNT(*) as explain_count FROM mt_explain;

-- 检查关联情况
SELECT 
    COUNT(*) as total_drugs,
    COUNT(mt_drug.guide) as drugs_with_guide,
    COUNT(mt_drug.explain) as drugs_with_explain,
    COUNT(CASE WHEN mt_drug.guide IS NOT NULL AND mt_drug.explain IS NOT NULL THEN 1 END) as drugs_with_both
FROM mt_drug 
WHERE mt_drug.deleted_at IS NULL;