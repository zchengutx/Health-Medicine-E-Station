-- 更新医生数据，添加医院和科室信息
-- 为现有医生数据随机分配医院和科室

UPDATE mt_doctors SET 
hospital_name = CASE 
  WHEN RAND() < 0.33 THEN '北大医院'
  WHEN RAND() < 0.66 THEN '浦东医院'
  ELSE '上海医院'
END,
department_name = CASE 
  WHEN RAND() < 0.2 THEN '内科'
  WHEN RAND() < 0.4 THEN '口腔科'
  WHEN RAND() < 0.6 THEN '外科'
  WHEN RAND() < 0.8 THEN '骨科'
  ELSE '急诊'
END,
fans_count = FLOOR(RAND() * 50) + 1
WHERE id > 0;

-- 查看更新后的数据
SELECT id, name, hospital_name, department_name, fans_count, status FROM mt_doctors LIMIT 10; 