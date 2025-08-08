package medicine

import (
	"gorm.io/gorm"
	"time"
)

// MtDepartments 对应数据库中的 mt_departments 表结构
type MtDepartments struct {
	ID             uint           `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:科室ID"`
	DepartmentCode string         `gorm:"column:department_code;type:varchar(32);not null;comment:科室编码"`
	HospitalID     uint           `gorm:"column:hospital_id;type:bigint unsigned;not null;comment:医院ID;index:idx_hospital_id"`
	Name           string         `gorm:"column:name;type:varchar(100);not null;comment:科室名称"`
	ShortName      string         `gorm:"column:short_name;type:varchar(50);default:null;comment:科室简称"`
	Type           string         `gorm:"column:type;type:varchar(20);default:null;comment:科室类型：内科，外科，妇科等;index:idx_type"`
	ParentID       uint           `gorm:"column:parent_id;type:bigint unsigned;default:0;comment:父科室ID;index:idx_parent_id"`
	Level          int            `gorm:"column:level;type:tinyint;default:1;comment:科室层级"`
	SortOrder      int            `gorm:"column:sort_order;type:int;default:0;comment:排序"`
	Description    string         `gorm:"column:description;type:text;default:null;comment:科室描述"`
	Location       string         `gorm:"column:location;type:varchar(100);default:null;comment:科室位置"`
	Phone          string         `gorm:"column:phone;type:varchar(20);default:null;comment:科室电话"`
	Status         string         `gorm:"column:status;type:varchar(10);not null;default:'启用';comment:状态：禁用/启用"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:datetime(6);not null;default:current_timestamp(6) on update current_timestamp(6);comment:创建时间"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime(6);not null;default:current_timestamp(6) on update current_timestamp(6);comment:更新时间"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);default:null;index;comment:删除时间（软删除标识）"`
}

// TableName 设置表名
func (MtDepartments) TableName() string {
	return "mt_departments"
}
