-- 测试医生数据
-- 查看当前数据
SELECT id, name, phone, hospital_name, department_name, status, fans_count, created_at 
FROM mt_doctors 
ORDER BY created_at DESC 
LIMIT 10;

-- 插入测试数据
INSERT INTO mt_doctors (
    doctor_code, name, phone, department_id, department_name, title, status, 
    license_number, gender, hospital_name, fans_count, created_at, updated_at
) VALUES 
('TEST001', '测试医生1', '13800138001', 1, '内科', '主治医师', '1', 'TEST001', '1', '北大医院', 15, NOW(), NOW()),
('TEST002', '测试医生2', '13800138002', 2, '外科', '副主任医师', '0', 'TEST002', '2', '浦东医院', 8, NOW(), NOW()),
('TEST003', '测试医生3', '13800138003', 3, '骨科', '主任医师', '2', 'TEST003', '1', '上海医院', 25, NOW(), NOW());

-- 查看插入后的数据
SELECT id, name, phone, hospital_name, department_name, status, fans_count, created_at 
FROM mt_doctors 
ORDER BY created_at DESC 
LIMIT 10; 