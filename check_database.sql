-- 检查审核表是否存在
SELECT 
    CASE 
        WHEN COUNT(*) > 0 THEN '✅ 表存在'
        ELSE '❌ 表不存在'
    END as table_status
FROM information_schema.tables 
WHERE table_schema = DATABASE() 
AND table_name = 'mt_doctor_approval';

-- 如果表存在，检查表结构
DESCRIBE mt_doctor_approval;

-- 检查表中的数据
SELECT COUNT(*) as record_count FROM mt_doctor_approval;

-- 如果表不存在，创建表
CREATE TABLE IF NOT EXISTS `mt_doctor_approval` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `doctor_id` bigint unsigned NOT NULL COMMENT '医生ID',
  `doctor_name` varchar(50) DEFAULT NULL COMMENT '医生姓名',
  `approval_status` varchar(20) NOT NULL COMMENT '审核状态：0-未通过，1-通过，2-待审核',
  `approval_time` datetime(3) DEFAULT NULL COMMENT '审核时间',
  `approver_id` bigint unsigned DEFAULT NULL COMMENT '审核人ID',
  `approver_name` varchar(50) DEFAULT NULL COMMENT '审核人姓名',
  `approval_reason` varchar(500) DEFAULT NULL COMMENT '审核理由',
  `reject_reason` varchar(500) DEFAULT NULL COMMENT '拒绝理由',
  `submit_time` datetime(3) DEFAULT NULL COMMENT '提交审核时间',
  `created_by` bigint unsigned DEFAULT NULL,
  `updated_by` bigint unsigned DEFAULT NULL,
  `deleted_by` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_mt_doctor_approval_deleted_at` (`deleted_at`),
  KEY `idx_mt_doctor_approval_doctor_id` (`doctor_id`),
  KEY `idx_mt_doctor_approval_status` (`approval_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='医生审核记录表';

-- 插入测试数据（如果表为空）
INSERT INTO mt_doctor_approval (doctor_id, doctor_name, approval_status, approval_time, approver_id, approver_name, approval_reason, reject_reason, submit_time, created_at, updated_at) 
SELECT 1, '张医生', '1', NOW() - INTERVAL 24 HOUR, 1, '管理员', '资质齐全，审核通过', NULL, NOW() - INTERVAL 25 HOUR, NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM mt_doctor_approval WHERE doctor_id = 1);

INSERT INTO mt_doctor_approval (doctor_id, doctor_name, approval_status, approval_time, approver_id, approver_name, approval_reason, reject_reason, submit_time, created_at, updated_at) 
SELECT 2, '李医生', '0', NOW() - INTERVAL 12 HOUR, 1, '管理员', NULL, '执业证书过期，需要重新提交', NOW() - INTERVAL 13 HOUR, NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM mt_doctor_approval WHERE doctor_id = 2);

INSERT INTO mt_doctor_approval (doctor_id, doctor_name, approval_status, approval_time, approver_id, approver_name, approval_reason, reject_reason, submit_time, created_at, updated_at) 
SELECT 3, '王医生', '2', NULL, NULL, NULL, NULL, NULL, NOW() - INTERVAL 6 HOUR, NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM mt_doctor_approval WHERE doctor_id = 3);

-- 显示插入的数据
SELECT * FROM mt_doctor_approval ORDER BY created_at DESC; 