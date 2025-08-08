-- 更新药品管理主菜单的组件路径
UPDATE sys_base_menus 
SET component = 'view/routerHolder.vue' 
WHERE name = 'medicine' AND path = 'medicine';

-- 删除旧的药品管理子菜单
DELETE FROM sys_base_menus 
WHERE parent_id = (SELECT id FROM sys_base_menus WHERE name = 'medicine') 
AND name IN ('mtDrugTypeStair', 'mtDrugTypeLevel');

-- 更新药品管理子菜单的标题和图标
UPDATE sys_base_menus 
SET meta = JSON_SET(meta, '$.title', '药品列表'), 
    meta = JSON_SET(meta, '$.icon', 'list'),
    sort = 2
WHERE name = 'mtDrug' AND parent_id = (SELECT id FROM sys_base_menus WHERE name = 'medicine');

UPDATE sys_base_menus 
SET sort = 3
WHERE name = 'mtDrugCategory' AND parent_id = (SELECT id FROM sys_base_menus WHERE name = 'medicine');

-- 插入新的药品概览菜单
INSERT INTO sys_base_menus (
    created_at, updated_at, menu_level, parent_id, path, name, 
    component, sort, meta, hidden
) VALUES (
    NOW(), NOW(), 1, 
    (SELECT id FROM sys_base_menus WHERE name = 'medicine'), 
    'medicineOverview', 'medicineOverview', 
    'view/medicine/overview.vue', 1, 
    '{"title": "药品概览", "icon": "odometer"}', false
); 