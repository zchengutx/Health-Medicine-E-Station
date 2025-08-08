-- 检查医院表是否存在
SHOW TABLES LIKE 'mt_hospitals';

-- 检查科室表是否存在
SHOW TABLES LIKE 'mt_departments';

-- 检查医生表结构
DESCRIBE mt_doctors;

-- 检查医院数据
SELECT id, name, level, status FROM mt_hospitals ORDER BY id;

-- 检查科室数据
SELECT id, hospital_id, name, type, status FROM mt_departments ORDER BY hospital_id, id;

-- 检查医生数据
SELECT id, name, hospital_id, department_id, status FROM mt_doctors ORDER BY id;

-- 检查医生表是否有医院和科室的关联字段
SELECT 
    d.id,
    d.name as doctor_name,
    d.hospital_id,
    h.name as hospital_name,
    d.department_id,
    dept.name as department_name
FROM mt_doctors d
LEFT JOIN mt_hospitals h ON d.hospital_id = h.id
LEFT JOIN mt_departments dept ON d.department_id = dept.id
ORDER BY d.id; 