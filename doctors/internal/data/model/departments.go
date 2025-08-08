package model

import (
	"time"

	"gorm.io/gorm"
)

type Departments struct {
	Id             uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:科室ID;primaryKey;not null;" json:"id"`                               // 科室ID
	DepartmentCode string         `gorm:"column:department_code;type:varchar(32);comment:科室编码;not null;" json:"department_code"`                    // 科室编码
	HospitalId     uint64         `gorm:"column:hospital_id;type:bigint UNSIGNED;comment:医院ID;not null;" json:"hospital_id"`                        // 医院ID
	Name           string         `gorm:"column:name;type:varchar(100);comment:科室名称;not null;" json:"name"`                                         // 科室名称
	ShortName      string         `gorm:"column:short_name;type:varchar(50);comment:科室简称;default:NULL;" json:"short_name"`                          // 科室简称
	Type           string         `gorm:"column:type;type:varchar(20);comment:科室类型：内科，外科，儿科等;default:NULL;" json:"type"`                            // 科室类型：内科，外科，儿科等
	ParentId       uint64         `gorm:"column:parent_id;type:bigint UNSIGNED;comment:父科室ID;default:0;" json:"parent_id"`                          // 父科室ID
	Level          int8           `gorm:"column:level;type:tinyint;comment:科室层级;default:1;" json:"level"`                                           // 科室层级
	SortOrder      int32          `gorm:"column:sort_order;type:int;comment:排序;default:0;" json:"sort_order"`                                       // 排序
	Description    string         `gorm:"column:description;type:text;comment:科室描述;default:NULL;" json:"description"`                               // 科室描述
	Location       string         `gorm:"column:location;type:varchar(100);comment:科室位置;default:NULL;" json:"location"`                             // 科室位置
	Phone          string         `gorm:"column:phone;type:varchar(20);comment:科室电话;default:NULL;" json:"phone"`                                    // 科室电话
	Status         string         `gorm:"column:status;type:varchar(10);comment:状态：禁用/启用;not null;default:启用;" json:"status"`                       // 状态：禁用/启用
	CreatedAt      time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"` // 创建时间
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (d *Departments) TableName() string {
	return "departments"
}
