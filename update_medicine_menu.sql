-- 更新药品管理菜单配置，使其显示为下拉菜单样式
-- 类似于驱动管理的下拉菜单样式

-- 1. 确保药品管理主菜单的组件路径正确（应该是routerHolder.vue来支持子菜单）
UPDATE sys_base_menus 
SET component = 'view/routerHolder.vue' 
WHERE name = 'medicine' AND path = 'medicine';

-- 2. 更新药品管理子菜单的配置，确保它们正确关联到父菜单
-- 药品概览
UPDATE sys_base_menus 
SET meta = JSON_SET(meta, '$.title', '药品概览'), 
    meta = JSON_SET(meta, '$.icon', 'odometer'),
    sort = 1,
    hidden = false
WHERE name = 'medicineOverview' AND parent_id = (SELECT id FROM sys_base_menus WHERE name = 'medicine');

-- 药品列表
UPDATE sys_base_menus 
SET meta = JSON_SET(meta, '$.title', '药品列表'), 
    meta = JSON_SET(meta, '$.icon', 'list'),
    sort = 2,
    hidden = false
WHERE name = 'mtDrug' AND parent_id = (SELECT id FROM sys_base_menus WHERE name = 'medicine');

-- 药品分类
UPDATE sys_base_menus 
SET meta = JSON_SET(meta, '$.title', '药品分类'), 
    meta = JSON_SET(meta, '$.icon', 'folder'),
    sort = 3,
    hidden = false
WHERE name = 'mtDrugCategory' AND parent_id = (SELECT id FROM sys_base_menus WHERE name = 'medicine');

-- 3. 确保所有子菜单都是可见的
UPDATE sys_base_menus 
SET hidden = false 
WHERE parent_id = (SELECT id FROM sys_base_menus WHERE name = 'medicine');

-- 4. 验证菜单结构
SELECT 
    m1.name as parent_name,
    m1.meta->>'$.title' as parent_title,
    m2.name as child_name,
    m2.meta->>'$.title' as child_title,
    m2.sort as child_sort,
    m2.hidden as child_hidden
FROM sys_base_menus m1
LEFT JOIN sys_base_menus m2 ON m1.id = m2.parent_id
WHERE m1.name = 'medicine'
ORDER BY m2.sort; 