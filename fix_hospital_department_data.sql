-- 检查并创建医院表
CREATE TABLE IF NOT EXISTS `mt_hospitals` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `hospital_code` varchar(32) NOT NULL COMMENT '医院编码',
  `name` varchar(100) NOT NULL COMMENT '医院名称',
  `short_name` varchar(50) DEFAULT NULL COMMENT '医院简称',
  `level` varchar(20) DEFAULT NULL COMMENT '医院等级：三甲，三乙，二甲等',
  `type` varchar(20) DEFAULT NULL COMMENT '医院类型：综合，专科等',
  `address` varchar(200) DEFAULT NULL COMMENT '医院地址',
  `phone` varchar(20) DEFAULT NULL COMMENT '联系电话',
  `website` varchar(100) DEFAULT NULL COMMENT '官方网站',
  `introduction` text DEFAULT NULL COMMENT '医院介绍',
  `license_number` varchar(50) DEFAULT NULL COMMENT '医疗机构执业许可证号',
  `status` varchar(10) NOT NULL COMMENT '状态：禁用/启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_hospital_code` (`hospital_code`),
  KEY `idx_name` (`name`),
  KEY `idx_level` (`level`),
  KEY `idx_type` (`type`),
  KEY `idx_mt_hospitals_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='医院表';

-- 检查并创建科室表
CREATE TABLE IF NOT EXISTS `mt_departments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `department_code` varchar(32) NOT NULL COMMENT '科室编码',
  `hospital_id` bigint unsigned NOT NULL COMMENT '医院ID',
  `name` varchar(100) NOT NULL COMMENT '科室名称',
  `short_name` varchar(50) DEFAULT NULL COMMENT '科室简称',
  `type` varchar(20) DEFAULT NULL COMMENT '科室类型：内科，外科，妇科等',
  `parent_id` bigint unsigned DEFAULT 0 COMMENT '父科室ID',
  `level` tinyint DEFAULT 1 COMMENT '科室层级',
  `sort_order` int DEFAULT 0 COMMENT '排序',
  `description` text DEFAULT NULL COMMENT '科室描述',
  `location` varchar(100) DEFAULT NULL COMMENT '科室位置',
  `phone` varchar(20) DEFAULT NULL COMMENT '科室电话',
  `status` varchar(10) NOT NULL DEFAULT '启用' COMMENT '状态：禁用/启用',
  PRIMARY KEY (`id`),
  KEY `idx_hospital_id` (`hospital_id`),
  KEY `idx_type` (`type`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_mt_departments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='科室表';

-- 检查医生表是否有医院和科室字段
ALTER TABLE mt_doctors 
ADD COLUMN IF NOT EXISTS hospital_id INT COMMENT '医院ID',
ADD COLUMN IF NOT EXISTS department_id INT COMMENT '科室ID';

-- 插入医院测试数据
INSERT INTO mt_hospitals (hospital_code, name, short_name, level, type, address, phone, website, introduction, license_number, status, created_at, updated_at) VALUES
('H001', '北京大学第一医院', '北大一院', '三甲', '综合', '北京市西城区西什库大街8号', '010-83572211', 'http://www.bddyyy.com.cn', '北京大学第一医院是一所集医疗、教学、科研、预防、保健为一体的现代化综合性医院', 'PDY00001X', '启用', NOW(), NOW()),
('H002', '上海交通大学医学院附属瑞金医院', '瑞金医院', '三甲', '综合', '上海市黄浦区瑞金二路197号', '021-64370045', 'http://www.rjh.com.cn', '上海交通大学医学院附属瑞金医院是一所集医疗、教学、科研、预防、保健为一体的现代化综合性医院', 'PDY00002X', '启用', NOW(), NOW()),
('H003', '复旦大学附属华山医院', '华山医院', '三甲', '综合', '上海市静安区乌鲁木齐中路12号', '021-52889999', 'http://www.huashan.org.cn', '复旦大学附属华山医院是一所集医疗、教学、科研、预防、保健为一体的现代化综合性医院', 'PDY00003X', '启用', NOW(), NOW()),
('H004', '中山大学附属第一医院', '中山一院', '三甲', '综合', '广州市越秀区中山二路1号', '020-28823388', 'http://www.gzsums.net', '中山大学附属第一医院是一所集医疗、教学、科研、预防、保健为一体的现代化综合性医院', 'PDY00004X', '启用', NOW(), NOW()),
('H005', '四川大学华西医院', '华西医院', '三甲', '综合', '成都市武侯区国学巷37号', '028-85422114', 'http://www.cd120.com', '四川大学华西医院是一所集医疗、教学、科研、预防、保健为一体的现代化综合性医院', 'PDY00005X', '启用', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

-- 插入科室测试数据
INSERT INTO mt_departments (department_code, hospital_id, name, short_name, type, parent_id, level, sort_order, description, location, phone, status, created_at, updated_at) VALUES
-- 北大一院科室
('D001', 1, '内科', '内科', '内科', 0, 1, 1, '内科是医院的主要科室之一，负责内科疾病的诊断和治疗', '门诊楼1层', '010-83572211', '启用', NOW(), NOW()),
('D002', 1, '外科', '外科', '外科', 0, 1, 2, '外科是医院的主要科室之一，负责外科疾病的诊断和治疗', '门诊楼2层', '010-83572212', '启用', NOW(), NOW()),
('D003', 1, '妇产科', '妇产科', '妇产科', 0, 1, 3, '妇产科是医院的主要科室之一，负责妇科和产科疾病的诊断和治疗', '门诊楼3层', '010-83572213', '启用', NOW(), NOW()),
('D004', 1, '儿科', '儿科', '儿科', 0, 1, 4, '儿科是医院的主要科室之一，负责儿童疾病的诊断和治疗', '门诊楼4层', '010-83572214', '启用', NOW(), NOW()),
('D005', 1, '骨科', '骨科', '外科', 2, 2, 1, '骨科是外科的分支，专门治疗骨骼、关节、肌肉等疾病', '门诊楼2层', '010-83572215', '启用', NOW(), NOW()),

-- 瑞金医院科室
('D006', 2, '内科', '内科', '内科', 0, 1, 1, '内科是医院的主要科室之一，负责内科疾病的诊断和治疗', '门诊楼1层', '021-64370045', '启用', NOW(), NOW()),
('D007', 2, '外科', '外科', '外科', 0, 1, 2, '外科是医院的主要科室之一，负责外科疾病的诊断和治疗', '门诊楼2层', '021-64370046', '启用', NOW(), NOW()),
('D008', 2, '心内科', '心内科', '内科', 6, 2, 1, '心内科是内科的分支，专门治疗心血管疾病', '门诊楼1层', '021-64370047', '启用', NOW(), NOW()),
('D009', 2, '神经内科', '神经内科', '内科', 6, 2, 2, '神经内科是内科的分支，专门治疗神经系统疾病', '门诊楼1层', '021-64370048', '启用', NOW(), NOW()),

-- 华山医院科室
('D010', 3, '内科', '内科', '内科', 0, 1, 1, '内科是医院的主要科室之一，负责内科疾病的诊断和治疗', '门诊楼1层', '021-52889999', '启用', NOW(), NOW()),
('D011', 3, '外科', '外科', '外科', 0, 1, 2, '外科是医院的主要科室之一，负责外科疾病的诊断和治疗', '门诊楼2层', '021-52889998', '启用', NOW(), NOW()),
('D012', 3, '神经外科', '神经外科', '外科', 11, 2, 1, '神经外科是外科的分支，专门治疗神经系统疾病', '门诊楼2层', '021-52889997', '启用', NOW(), NOW()),
('D013', 3, '皮肤科', '皮肤科', '皮肤科', 0, 1, 3, '皮肤科是医院的主要科室之一，专门治疗皮肤疾病', '门诊楼3层', '021-52889996', '启用', NOW(), NOW()),

-- 中山一院科室
('D014', 4, '内科', '内科', '内科', 0, 1, 1, '内科是医院的主要科室之一，负责内科疾病的诊断和治疗', '门诊楼1层', '020-28823388', '启用', NOW(), NOW()),
('D015', 4, '外科', '外科', '外科', 0, 1, 2, '外科是医院的主要科室之一，负责外科疾病的诊断和治疗', '门诊楼2层', '020-28823387', '启用', NOW(), NOW()),
('D016', 4, '泌尿外科', '泌尿外科', '外科', 15, 2, 1, '泌尿外科是外科的分支，专门治疗泌尿系统疾病', '门诊楼2层', '020-28823386', '启用', NOW(), NOW()),

-- 华西医院科室
('D017', 5, '内科', '内科', '内科', 0, 1, 1, '内科是医院的主要科室之一，负责内科疾病的诊断和治疗', '门诊楼1层', '028-85422114', '启用', NOW(), NOW()),
('D018', 5, '外科', '外科', '外科', 0, 1, 2, '外科是医院的主要科室之一，负责外科疾病的诊断和治疗', '门诊楼2层', '028-85422113', '启用', NOW(), NOW()),
('D019', 5, '口腔科', '口腔科', '口腔科', 0, 1, 3, '口腔科是医院的主要科室之一，专门治疗口腔疾病', '门诊楼3层', '028-85422112', '启用', NOW(), NOW()),
('D020', 5, '眼科', '眼科', '眼科', 0, 1, 4, '眼科是医院的主要科室之一，专门治疗眼部疾病', '门诊楼4层', '028-85422111', '启用', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

-- 为现有医生数据随机分配医院和科室
UPDATE mt_doctors SET 
hospital_id = CASE 
  WHEN RAND() < 0.2 THEN 1
  WHEN RAND() < 0.4 THEN 2
  WHEN RAND() < 0.6 THEN 3
  WHEN RAND() < 0.8 THEN 4
  ELSE 5
END,
department_id = CASE 
  WHEN hospital_id = 1 THEN CASE WHEN RAND() < 0.5 THEN 1 ELSE 2 END
  WHEN hospital_id = 2 THEN CASE WHEN RAND() < 0.5 THEN 6 ELSE 7 END
  WHEN hospital_id = 3 THEN CASE WHEN RAND() < 0.5 THEN 10 ELSE 11 END
  WHEN hospital_id = 4 THEN CASE WHEN RAND() < 0.5 THEN 14 ELSE 15 END
  WHEN hospital_id = 5 THEN CASE WHEN RAND() < 0.5 THEN 17 ELSE 18 END
END
WHERE hospital_id IS NULL OR department_id IS NULL;

-- 查看修复后的数据
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