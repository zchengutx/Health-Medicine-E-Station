package medicine

type MtExplain struct {
	Id                     int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	CommonName             string `gorm:"column:common_name;type:varchar(50);comment:通用名称;default:NULL;" json:"common_name"`                          // 通用名称
	GoodsName              string `gorm:"column:goods_name;type:varchar(20);comment:商品名称;default:NULL;" json:"goods_name"`                            // 商品名称
	Component              string `gorm:"column:component;type:varchar(50);comment:药品成份;default:NULL;" json:"component"`                              // 药品成份
	Taboos                 string `gorm:"column:taboos;type:varchar(100);comment:用药禁忌;default:NULL;" json:"taboos"`                                   // 用药禁忌
	Function               string `gorm:"column:function;type:varchar(100);comment:功能主治;default:NULL;" json:"function"`                               // 功能主治
	UsageAndDosage         string `gorm:"column:usage_and_dosage;type:varchar(100);comment:用法用量;default:NULL;" json:"usage_and_dosage"`               // 用法用量
	Character              string `gorm:"column:character;type:varchar(100);comment:药品性状;default:NULL;" json:"character"`                             // 药品性状
	PackagingSpecification string `gorm:"column:packaging_specification;type:varchar(100);comment:包装规格;default:NULL;" json:"packaging_specification"` // 包装规格
	BadnessReaction        string `gorm:"column:badness_reaction;type:varchar(100);comment:不良反应;default:NULL;" json:"badness_reaction"`               // 不良反应
	Condition              string `gorm:"column:condition;type:varchar(100);comment:贮藏条件;default:NULL;" json:"condition"`                             // 贮藏条件
	ValidTime              string `gorm:"column:valid_time;type:varchar(20);comment:有效期;default:NULL;" json:"valid_time"`                             // 有效期
	Notice                 string `gorm:"column:notice;type:varchar(100);comment:注意事项;default:NULL;" json:"notice"`                                   // 注意事项
	Interaction            string `gorm:"column:interaction;type:varchar(100);comment:相互作用;default:NULL;" json:"interaction"`                         // 相互作用
	RatifyNumber           string `gorm:"column:ratify_number;type:varchar(100);comment:批准文号;default:NULL;" json:"ratify_number"`                     // 批准文号
	Manufacturer           string `gorm:"column:manufacturer;type:varchar(30);comment:生产厂商;default:NULL;" json:"manufacturer"`                        // 生产厂商
	StandardNumber         string `gorm:"column:standard_number;type:varchar(40);comment:执行标准号;default:NULL;" json:"standard_number"`                 // 执行标准号
	Possessor              string `gorm:"column:possessor;type:varchar(40);comment:上市许可持有人;default:NULL;" json:"possessor"`                           // 上市许可持有人
	Address                string `gorm:"column:address;type:varchar(100);comment:上市许可持有人地址;default:NULL;" json:"address"`                            // 上市许可持有人地址
	Specification          string `gorm:"column:specification;type:varchar(40);comment:规格;default:NULL;" json:"specification"`                        // 规格
	DosageForm             string `gorm:"column:dosage_form;type:varchar(20);comment:剂型;default:NULL;" json:"dosage_form"`                            // 剂型
}

func (MtExplain) TableName() string {
	return "mt_explain"
}
