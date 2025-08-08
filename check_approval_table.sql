-- 检查审核表是否存在
SHOW TABLES LIKE 'mt_doctor_approval';

-- 如果表存在，查看表结构
DESCRIBE mt_doctor_approval;

-- 查看表中的数据
SELECT * FROM mt_doctor_approval LIMIT 5;

-- 如果没有数据，插入一些测试数据
INSERT INTO mt_doctor_approval (doctor_id, doctor_name, approval_status, approval_time, approver_id, approver_name, approval_reason, reject_reason, submit_time, created_at, updated_at) VALUES
(1, '张医生', '1', NOW() - INTERVAL 24 HOUR, 1, '管理员', '资质齐全，审核通过', NULL, NOW() - INTERVAL 25 HOUR, NOW(), NOW()),
(2, '李医生', '0', NOW() - INTERVAL 12 HOUR, 1, '管理员', NULL, '执业证书过期，需要重新提交', NOW() - INTERVAL 13 HOUR, NOW(), NOW()),
(3, '王医生', '2', NULL, NULL, NULL, NULL, NULL, NOW() - INTERVAL 6 HOUR, NOW(), NOW()),
(4, '赵医生', '1', NOW() - INTERVAL 2 HOUR, 1, '管理员', '所有材料符合要求，审核通过', NULL, NOW() - INTERVAL 3 HOUR, NOW(), NOW()),
(5, '刘医生', '0', NOW() - INTERVAL 1 HOUR, 1, '管理员', NULL, '缺少必要的执业证书复印件', NOW() - INTERVAL 2 HOUR, NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW(); 