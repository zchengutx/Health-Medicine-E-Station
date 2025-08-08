-- 修复数据库表问题
-- 请先连接到数据库：mysql -h 14.103.134.228 -u root -p Health_Medicine_E_Station

-- 1. 删除可能存在的表（如果存在）
DROP TABLE IF EXISTS `mt_drug_type_stair`;
DROP TABLE IF EXISTS `mt_drug_type_level`;
DROP TABLE IF EXISTS `mt_drug_category`;
DROP TABLE IF EXISTS `mt_drug`;
DROP TABLE IF EXISTS `mt_orders_drug`;
DROP TABLE IF EXISTS `mt_user`;
DROP TABLE IF EXISTS `mt_orders`;

-- 2. 创建表（按依赖顺序）
-- 首先创建基础表
CREATE TABLE `mt_drug_type_stair` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `stair_name` varchar(40) DEFAULT NULL COMMENT '一级分类名称',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint unsigned DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`),
  KEY `idx_mt_drug_type_stair_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `mt_drug_type_level` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `level_name` varchar(40) DEFAULT NULL COMMENT '二级分类名称',
  `stair_id` bigint DEFAULT NULL COMMENT '一级分类id',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint unsigned DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`),
  KEY `idx_mt_drug_type_level_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `mt_drug_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `mt_drug_type_stair_id` bigint DEFAULT NULL COMMENT '一级分类id',
  `mt_drug_type_level_id` bigint DEFAULT NULL COMMENT '二级分类id',
  `status` int DEFAULT NULL COMMENT '状态（0-禁用，1-启用）',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint unsigned DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`),
  KEY `idx_mt_drug_category_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `mt_drug` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `drug_name` varchar(20) DEFAULT NULL COMMENT '药品名称',
  `guide` int DEFAULT NULL COMMENT '用药指导id',
  `explain` int DEFAULT NULL COMMENT '说明书id',
  `specification` varchar(20) DEFAULT NULL COMMENT '规格',
  `price` decimal(10,2) DEFAULT NULL COMMENT '价格',
  `sales_volume` decimal(5,2) DEFAULT NULL COMMENT '销量',
  `inventory` int DEFAULT NULL COMMENT '库存',
  `status` varchar(10) DEFAULT NULL COMMENT '状态',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint unsigned DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`),
  KEY `idx_mt_drug_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `mt_orders_drug` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `order_id` bigint DEFAULT NULL COMMENT '订单id',
  `drug_id` bigint DEFAULT NULL COMMENT '药品id',
  `user_id` bigint DEFAULT NULL COMMENT '患者id',
  `quantity` int DEFAULT NULL COMMENT '数量',
  `order_status` varchar(10) DEFAULT NULL COMMENT '订单状态:1-待发货，2-待收货，3-已收货',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint unsigned DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`),
  KEY `idx_mt_orders_drug_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `mt_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(20) DEFAULT NULL COMMENT '昵称',
  `mobile` char(11) DEFAULT NULL COMMENT '手机号',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `mt_orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_no` varchar(50) NOT NULL COMMENT '订单编号',
  `user_id` bigint DEFAULT NULL COMMENT '患者id',
  `user_name` varchar(50) DEFAULT NULL COMMENT '患者姓名',
  `user_phone` varchar(20) DEFAULT NULL COMMENT '患者电话',
  `doctor_id` bigint DEFAULT NULL COMMENT '医生id',
  `doctor_name` varchar(50) DEFAULT NULL COMMENT '医生姓名',
  `address_id` bigint DEFAULT NULL COMMENT '地址id',
  `address_detail` varchar(200) DEFAULT NULL COMMENT '地址详情',
  `total_amount` decimal(10,2) NOT NULL COMMENT '总金额',
  `pay_type` varchar(10) DEFAULT NULL COMMENT '支付方式：1-微信，2-支付宝，3-银行卡',
  `status` varchar(20) DEFAULT NULL COMMENT '订单状态：1-待支付，2-已支付，3-配药中，4-已发货，5-已完成，6-已取消',
  `pay_time` datetime(3) DEFAULT NULL COMMENT '支付时间',
  `drug_time` datetime(3) DEFAULT NULL COMMENT '配药时间',
  `send_time` datetime(3) DEFAULT NULL COMMENT '发货时间',
  `finish_time` datetime(3) DEFAULT NULL COMMENT '完成时间',
  `cancel_time` datetime(3) DEFAULT NULL COMMENT '取消时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 3. 插入测试数据
INSERT INTO `mt_drug_type_stair` (`stair_name`) VALUES 
('感冒药'),
('消炎药'),
('维生素');

INSERT INTO `mt_drug_type_level` (`level_name`, `stair_id`) VALUES 
('感冒药', 1),
('消炎药', 2),
('维生素C', 3);

INSERT INTO `mt_drug` (`drug_name`, `specification`, `price`, `sales_volume`, `inventory`, `status`) VALUES 
('999感冒灵', '10袋/盒', 15.00, 100.00, 50, '1'),
('阿莫西林', '24粒/盒', 25.00, 80.00, 30, '1'),
('维生素C片', '100片/瓶', 30.00, 120.00, 40, '1');

-- 4. 验证表创建
SELECT '表创建完成' as message;
SHOW TABLES LIKE 'mt_%'; 