-- 添加医生表新字段
ALTER TABLE mt_doctors 
ADD COLUMN department_name VARCHAR(100) COMMENT '科室名称',
ADD COLUMN hospital_name VARCHAR(100) COMMENT '医院名称',
ADD COLUMN fans_count INT DEFAULT 0 COMMENT '粉丝数量';

-- 更新现有数据的默认值
UPDATE mt_doctors SET 
department_name = '内科',
hospital_name = '北大医院',
fans_count = 12
WHERE department_name IS NULL OR hospital_name IS NULL;

-- 为现有数据随机分配医院和科室
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
END
WHERE hospital_name = '北大医院' OR department_name = '内科'; 