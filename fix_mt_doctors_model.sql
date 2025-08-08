-- 检查MtDoctors表结构
DESCRIBE mt_doctors;

-- 检查是否有Email和Avatar字段
SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT 
FROM INFORMATION_SCHEMA.COLUMNS 
WHERE TABLE_NAME = 'mt_doctors' 
AND TABLE_SCHEMA = DATABASE();

-- 如果存在Email和Avatar字段，检查它们的约束
SELECT 
    COLUMN_NAME,
    IS_NULLABLE,
    COLUMN_DEFAULT,
    COLUMN_TYPE
FROM INFORMATION_SCHEMA.COLUMNS 
WHERE TABLE_NAME = 'mt_doctors' 
AND COLUMN_NAME IN ('email', 'avatar', 'Email', 'Avatar')
AND TABLE_SCHEMA = DATABASE();

-- 如果Email和Avatar字段是必填的，但模型中没有定义，可以选择以下方案之一：

-- 方案1：将Email和Avatar字段设为可空（推荐）
-- ALTER TABLE mt_doctors MODIFY COLUMN email VARCHAR(255) NULL;
-- ALTER TABLE mt_doctors MODIFY COLUMN avatar VARCHAR(255) NULL;

-- 方案2：删除Email和Avatar字段（如果不需要）
-- ALTER TABLE mt_doctors DROP COLUMN email;
-- ALTER TABLE mt_doctors DROP COLUMN avatar;

-- 方案3：在模型中添加Email和Avatar字段（需要修改Go代码）
-- 这需要在mt_doctors.go模型中添加：
-- Email  *string `json:"email" form:"email" gorm:"comment:邮箱;column:email;size:255;"`
-- Avatar *string `json:"avatar" form:"avatar" gorm:"comment:头像;column:avatar;size:255;"`

-- 查看当前表的所有字段
SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT, COLUMN_COMMENT
FROM INFORMATION_SCHEMA.COLUMNS 
WHERE TABLE_NAME = 'mt_doctors' 
AND TABLE_SCHEMA = DATABASE()
ORDER BY ORDINAL_POSITION; 