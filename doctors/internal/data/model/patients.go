package model

import (
	"gorm.io/gorm"
	"time"
)

type Patients struct {
	Id               uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:患者ID;primaryKey;not null;" json:"id"`                                // 患者ID
	PatientCode      string         `gorm:"column:patient_code;type:varchar(32);comment:患者编码;not null;" json:"patient_code"`                          // 患者编码
	Name             string         `gorm:"column:name;type:varchar(50);comment:患者姓名;not null;" json:"name"`                                         // 患者姓名
	Gender           string         `gorm:"column:gender;type:varchar(10);comment:性别：未知/男/女;not null;default:未知;" json:"gender"`                      // 性别：未知/男/女
	BirthDate        time.Time      `gorm:"column:birth_date;type:date;comment:出生日期;default:NULL;" json:"birth_date"`                                // 出生日期
	Age              int32          `gorm:"column:age;type:int;comment:年龄;default:NULL;" json:"age"`                                                  // 年龄
	Phone            string         `gorm:"column:phone;type:varchar(20);comment:手机号码;default:NULL;" json:"phone"`                                    // 手机号码
	IdCard           string         `gorm:"column:id_card;type:varchar(18);comment:身份证号;default:NULL;" json:"id_card"`                                // 身份证号
	Address          string         `gorm:"column:address;type:varchar(200);comment:地址;default:NULL;" json:"address"`                                  // 地址
	EmergencyContact string         `gorm:"column:emergency_contact;type:varchar(50);comment:紧急联系人;default:NULL;" json:"emergency_contact"`            // 紧急联系人
	EmergencyPhone   string         `gorm:"column:emergency_phone;type:varchar(20);comment:紧急联系人电话;default:NULL;" json:"emergency_phone"`             // 紧急联系人电话
	BloodType        string         `gorm:"column:blood_type;type:varchar(10);comment:血型;default:NULL;" json:"blood_type"`                             // 血型
	AllergyHistory   string         `gorm:"column:allergy_history;type:text;comment:过敏史;default:NULL;" json:"allergy_history"`                        // 过敏史
	MedicalHistory   string         `gorm:"column:medical_history;type:text;comment:既往病史;default:NULL;" json:"medical_history"`                       // 既往病史
	FamilyHistory    string         `gorm:"column:family_history;type:text;comment:家族病史;default:NULL;" json:"family_history"`                         // 家族病史
	Occupation       string         `gorm:"column:occupation;type:varchar(100);comment:职业;default:NULL;" json:"occupation"`                           // 职业
	MaritalStatus    string         `gorm:"column:marital_status;type:varchar(20);comment:婚姻状况：未知/未婚/已婚/离异/丧偶;default:未知;" json:"marital_status"` // 婚姻状况：未知/未婚/已婚/离异/丧偶
	InsuranceType    string         `gorm:"column:insurance_type;type:varchar(50);comment:医保类型;default:NULL;" json:"insurance_type"`                   // 医保类型
	InsuranceNumber  string         `gorm:"column:insurance_number;type:varchar(50);comment:医保号;default:NULL;" json:"insurance_number"`                // 医保号
	Status           string         `gorm:"column:status;type:varchar(10);comment:状态：禁用/启用;not null;default:启用;" json:"status"`                       // 状态：禁用/启用
	CreatedAt        time.Time      `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`      // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`      // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;comment:删除时间;default:NULL;" json:"deleted_at"`                           // 删除时间
}

func (p *Patients)TableName() string {
	return "patients"
}
