package medicine

type MtGuide struct {
	Id             int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	MajorFunction  string `gorm:"column:major_function;type:varchar(100);comment:功能主治;default:NULL;" json:"major_function"`     // 功能主治
	UsageAndDosage string `gorm:"column:usage_and_dosage;type:varchar(100);comment:用法用量;default:NULL;" json:"usage_and_dosage"` // 用法用量
	Taboos         string `gorm:"column:taboos;type:varchar(100);comment:用药禁忌;default:NULL;" json:"taboos"`                     // 用药禁忌
	SpecialCrowd   string `gorm:"column:special_crowd;type:varchar(100);comment:特殊人群;default:NULL;" json:"special_crowd"`       // 特殊人群
	Condition      string `gorm:"column:condition;type:varchar(100);comment:贮藏条件;default:NULL;" json:"condition"`               // 贮藏条件
}

func (MtGuide) TableName() string {
	return "mt_guide"
}
