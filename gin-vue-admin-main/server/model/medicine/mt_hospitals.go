package medicine

import (
	"gorm.io/gorm"
	"time"
)

// MtHospitals 对应数据库中的 mt_hospitals 表
type MtHospitals struct {
	ID            uint           `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:医院ID"`
	HospitalCode  string         `gorm:"column:hospital_code;type:varchar(32);not null;uniqueIndex:uk_hospital_code;comment:医院编码"`
	Name          string         `gorm:"column:name;type:varchar(100);not null;index:idx_name;comment:医院名称"`
	ShortName     string         `gorm:"column:short_name;type:varchar(50);default:null;comment:医院简称"`
	Level         string         `gorm:"column:level;type:varchar(20);default:null;index:idx_level;comment:医院等级：三甲，三乙，二甲等"`
	Type          string         `gorm:"column:type;type:varchar(20);default:null;index:idx_type;comment:医院类型：综合，专科等"`
	Address       string         `gorm:"column:address;type:varchar(200);default:null;comment:医院地址"`
	Phone         string         `gorm:"column:phone;type:varchar(20);default:null;comment:联系电话"`
	Website       string         `gorm:"column:website;type:varchar(100);default:null;comment:官方网站"`
	Introduction  string         `gorm:"column:introduction;type:text;default:null;comment:医院介绍"`
	LicenseNumber string         `gorm:"column:license_number;type:varchar(50);default:null;comment:医疗机构执业许可证号"`
	Status        string         `gorm:"column:status;type:varchar(10);not null;comment:状态：禁用/启用"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(6);not null;default:current_timestamp(6) on update current_timestamp(6);comment:创建时间"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(6);not null;default:current_timestamp(6) on update current_timestamp(6);comment:更新时间"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);default:null;index;comment:删除时间（软删除标识）"`
}

// TableName 显式指定表名，避免GORM自动复数化
func (MtHospitals) TableName() string {
	return "mt_hospitals"
}
