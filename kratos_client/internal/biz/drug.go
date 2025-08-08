package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type MtDrug struct {
	Id               int32     `gorm:"column:id;type:int;primaryKey;" json:"id"`
	DrugName         string    `gorm:"column:drug_name;type:varchar(20);comment:药品名称;" json:"drug_name"`                   // 药品名称
	FristCategoryId  int32     `gorm:"column:frist_category_id;type:mediumint;comment:一级分类Id;" json:"frist_category_id"`   // 一级分类Id
	SecondCategoryId int32     `gorm:"column:second_category_id;type:mediumint;comment:二级分类Id;" json:"second_category_id"` // 二级分类Id
	Guide            int16     `gorm:"column:guide;type:smallint;comment:用药指导id;" json:"guide"`                            // 用药指导id
	Explain          int16     `gorm:"column:explain;type:smallint;comment:说明书id;" json:"explain"`                          // 说明书id
	Specification    string    `gorm:"column:specification;type:varchar(20);comment:规格;" json:"specification"`               // 规格
	DrugStore        int16     `gorm:"column:drug_store;type:smallint;comment:店铺id;" json:"drug_store"`                      // 店铺id
	Price            float32   `gorm:"column:price;type:float;comment:价格;" json:"price"`                                     // 价格
	SalesVolume      float32   `gorm:"column:sales_volume;type:float;comment:销量;" json:"sales_volume"`                       // 销量
	Inventory        int16     `gorm:"column:inventory;type:smallint;comment:库存;" json:"inventory"`                          // 库存
	Manufacturer     string    `gorm:"column:manufacturer;type:varchar(100);comment:生产厂家;" json:"manufacturer"`            // 生产厂家
	ExhibitionId     int16     `gorm:"column:exhibition_id;type:smallint;comment:图片id;" json:"exhibition_id"`                // 图片id
	Status           int32     `gorm:"column:status;type:mediumint;comment:状态;" json:"status"`                               // 状态
	CreatedAt        time.Time `gorm:"column:created_at;type:datetime(3);" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:datetime(3);" json:"updated_at"`
	DeletedAt        time.Time `gorm:"column:deleted_at;type:datetime(3);" json:"deleted_at"`
	CreatedBy        uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;" json:"created_by"`    // 创建者
	UpdatedBy        uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;" json:"updated_by"`    // 更新者
	DeletedBy        uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;" json:"deleted_by"`    // 删除者
	DrugClassify     int16     `gorm:"column:drug_classify;type:smallint;comment:药品分类id;" json:"drug_classify"` // 药品分类id
}

type MtGuide struct {
	Id             int32  `gorm:"column:id;type:int;primaryKey;" json:"id"`
	MajorFunction  string `gorm:"column:major_function;type:varchar(100);comment:功能主治;" json:"major_function"`     // 功能主治
	UsageAndDosage string `gorm:"column:usage_and_dosage;type:varchar(100);comment:用法用量;" json:"usage_and_dosage"` // 用法用量
	Taboos         string `gorm:"column:taboos;type:varchar(100);comment:用药禁忌;" json:"taboos"`                     // 用药禁忌
	SpecialCrowd   string `gorm:"column:special_crowd;type:varchar(100);comment:特殊人群;" json:"special_crowd"`       // 特殊人群
	Condition      string `gorm:"column:condition;type:varchar(100);comment:贮藏条件;" json:"condition"`               // 贮藏条件
}

type MtExplain struct {
	Id                     int32  `gorm:"column:id;type:int;primaryKey;" json:"id"`
	CommonName             string `gorm:"column:common_name;type:varchar(50);comment:通用名称;" json:"common_name"`                          // 通用名称
	GoodsName              string `gorm:"column:goods_name;type:varchar(20);comment:商品名称;" json:"goods_name"`                            // 商品名称
	Component              string `gorm:"column:component;type:varchar(50);comment:药品成份;" json:"component"`                              // 药品成份
	Taboos                 string `gorm:"column:taboos;type:varchar(100);comment:用药禁忌;" json:"taboos"`                                   // 用药禁忌
	Function               string `gorm:"column:function;type:varchar(100);comment:功能主治;" json:"function"`                               // 功能主治
	UsageAndDosage         string `gorm:"column:usage_and_dosage;type:varchar(100);comment:用法用量;" json:"usage_and_dosage"`               // 用法用量
	Character              string `gorm:"column:character;type:varchar(100);comment:药品性状;" json:"character"`                             // 药品性状
	PackagingSpecification string `gorm:"column:packaging_specification;type:varchar(100);comment:包装规格;" json:"packaging_specification"` // 包装规格
	BadnessReaction        string `gorm:"column:badness_reaction;type:varchar(100);comment:不良反应;" json:"badness_reaction"`               // 不良反应
	Condition              string `gorm:"column:condition;type:varchar(100);comment:贮藏条件;" json:"condition"`                             // 贮藏条件
	ValidTime              string `gorm:"column:valid_time;type:varchar(20);comment:有效期;" json:"valid_time"`                              // 有效期
	Notice                 string `gorm:"column:notice;type:varchar(100);comment:注意事项;" json:"notice"`                                   // 注意事项
	Interaction            string `gorm:"column:interaction;type:varchar(100);comment:相互作用;" json:"interaction"`                         // 相互作用
	RatifyNumber           string `gorm:"column:ratify_number;type:varchar(100);comment:批准文号;" json:"ratify_number"`                     // 批准文号
	Manufacturer           string `gorm:"column:manufacturer;type:varchar(30);comment:生产厂商;" json:"manufacturer"`                        // 生产厂商
	StandardNumber         string `gorm:"column:standard_number;type:varchar(40);comment:执行标准号;" json:"standard_number"`                // 执行标准号
	Possessor              string `gorm:"column:possessor;type:varchar(40);comment:上市许可持有人;" json:"possessor"`                        // 上市许可持有人
	Address                string `gorm:"column:address;type:varchar(100);comment:上市许可持有人地址;" json:"address"`                       // 上市许可持有人地址
	Specification          string `gorm:"column:specification;type:varchar(40);comment:规格;" json:"specification"`                          // 规格
	DosageForm             string `gorm:"column:dosage_form;type:varchar(20);comment:剂型;" json:"dosage_form"`                              // 剂型
}

func (c *MtGuide) TableName() string {
	return "mt_guide"
}

func (c *MtDrug) TableName() string {
	return "mt_drug"
}

func (c *MtExplain) TableName() string {
	return "mt_explain"
}

// Elasticsearch文档模型
type DrugDocument struct {
	ID               int64     `json:"id"`
	DrugName         string    `json:"drug_name"`
	Specification    string    `json:"specification"`
	FirstCategoryID  int32     `json:"first_category_id"`
	SecondCategoryID int32     `json:"second_category_id"`
	Price            float64   `json:"price"`
	Inventory        int64     `json:"inventory"`
	Manufacturer     string    `json:"manufacturer"`
	Keywords         []string  `json:"keywords"`        // 搜索关键词
	Symptoms         []string  `json:"symptoms"`        // 适应症状
	DrugStoreID      int32     `json:"drug_store_id"`   // 药店ID
	IsPrescription   bool      `json:"is_prescription"` // 是否处方药
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// 处方药模型

func (m *MtPrescription) TableName() string {
	return "mt_prescription"
}

type PrescriptionStatus int

// 药店库存模型
type MtDrugInventory struct {
	ID             int64           `gorm:"column:id;type:bigint;primaryKey;autoIncrement;" json:"id"`
	DrugID         int64           `gorm:"column:drug_id;type:bigint;comment:药品ID;not null;" json:"drug_id"`
	DrugStoreID    int32           `gorm:"column:drug_store_id;type:int;comment:药店ID;not null;" json:"drug_store_id"`
	Quantity       int64           `gorm:"column:quantity;type:bigint;comment:库存数量;not null;default:0;" json:"quantity"`
	ReservedQty    int64           `gorm:"column:reserved_qty;type:bigint;comment:预留库存;not null;default:0;" json:"reserved_qty"`
	AlertThreshold int64           `gorm:"column:alert_threshold;type:bigint;comment:预警阈值;not null;default:10;" json:"alert_threshold"`
	Price          float64         `gorm:"column:price;type:decimal(10,2);comment:药店特定价格;not null;" json:"price"`
	Status         InventoryStatus `gorm:"column:status;type:tinyint;comment:库存状态;not null;default:0;" json:"status"`
	UpdatedAt      time.Time       `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"updated_at"`
}

func (m *MtDrugInventory) TableName() string {
	return "mt_drug_inventory"
}

type InventoryStatus int

type DrugRepo interface {
	ListDrug(ctx context.Context, fristCategoryId int32, secondCategoryId int32, keyword string) ([]*MtDrug, error)
	GetDrug(ctx context.Context, id int32) (*MtDrug, error)
	GetExplain(ctx context.Context, id int32) (*MtExplain, error)
	GetGuide(ctx context.Context, id int32) (*MtGuide, error)
}

// 搜索请求结构
type SearchRequest struct {
	Keyword             string      `json:"keyword"`
	CategoryID          int32       `json:"category_id"`
	PriceRange          *PriceRange `json:"price_range"`
	DrugStoreID         int32       `json:"drug_store_id"`
	OnlyInStock         bool        `json:"only_in_stock"`
	IncludePrescription bool        `json:"include_prescription"`
	Page                int         `json:"page"`
	Size                int         `json:"size"`
	SortBy              string      `json:"sort_by"` // price_asc, price_desc, sales_desc
}

type PriceRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// 搜索响应结构
type SearchResponse struct {
	Total  int64         `json:"total"`
	Drugs  []*DrugInfo   `json:"drugs"`
	Facets *SearchFacets `json:"facets"`
}

type DrugInfo struct {
	ID             int64   `json:"id"`
	DrugName       string  `json:"drug_name"`
	Specification  string  `json:"specification"`
	Price          float64 `json:"price"`
	Inventory      int64   `json:"inventory"`
	Manufacturer   string  `json:"manufacturer"`
	IsPrescription bool    `json:"is_prescription"`
}

type SearchFacets struct {
	Categories    []CategoryFacet     `json:"categories"`
	PriceRanges   []PriceFacet        `json:"price_ranges"`
	Manufacturers []ManufacturerFacet `json:"manufacturers"`
}

type CategoryFacet struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type PriceFacet struct {
	Range string `json:"range"`
	Count int64  `json:"count"`
}

type ManufacturerFacet struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// 热门搜索相关结构
type HotItem struct {
	Content string  `json:"content"`
	Count   int64   `json:"count"`
	Score   float64 `json:"score"`
}

type SearchType int

type DrugService struct {
	repo DrugRepo
	log  *log.Helper
}

// NewContentUsecase new a Content usecase.
func NewDrugService(repo DrugRepo, logger log.Logger) *DrugService {
	return &DrugService{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DrugService) ListDrug(ctx context.Context, fristCategoryId int32, secondCategoryId int32, keyword string) ([]*MtDrug, error) {
	uc.log.WithContext(ctx).Infof("List: %v+v", fristCategoryId, secondCategoryId, keyword)
	return uc.repo.ListDrug(ctx, fristCategoryId, secondCategoryId, keyword)
}
func (uc *DrugService) GetDrug(ctx context.Context, id int32) (*MtDrug, error) {
	uc.log.WithContext(ctx).Infof("GetDrug: %v+v", id)
	return uc.repo.GetDrug(ctx, id)
}
func (uc *DrugService) GetExplain(ctx context.Context, id int32) (*MtExplain, error) {
	uc.log.WithContext(ctx).Infof("GetExplain: %v+v", id)
	return uc.repo.GetExplain(ctx, id)
}

func (uc *DrugService) GetGuide(ctx context.Context, id int32) (*MtGuide, error) {
	uc.log.WithContext(ctx).Infof("GetGuide: %v+v", id)
	return uc.repo.GetGuide(ctx, id)
}

// 简化的搜索方法
func (uc *DrugService) SearchDrugs(ctx context.Context, keyword string, categoryId int32, page, size int) ([]*MtDrug, int64, error) {
	uc.log.WithContext(ctx).Infof("SearchDrugs: keyword=%s, categoryId=%d", keyword, categoryId)

	// 如果没有关键词，使用原有的ListDrug方法
	if keyword == "" {
		drugs, err := uc.repo.ListDrug(ctx, categoryId, 0, "")
		if err != nil {
			return nil, 0, err
		}
		return drugs, int64(len(drugs)), nil
	}

	// 使用药品名称进行搜索
	drugs, err := uc.repo.ListDrug(ctx, categoryId, 0, keyword)
	if err != nil {
		return nil, 0, err
	}

	// 简单分页
	total := int64(len(drugs))
	start := (page - 1) * size
	end := start + size

	if start >= len(drugs) {
		return []*MtDrug{}, total, nil
	}
	if end > len(drugs) {
		end = len(drugs)
	}

	return drugs[start:end], total, nil
}

// 获取热门关键词（简化版）
func (uc *DrugService) GetHotKeywords(ctx context.Context, limit int) ([]string, error) {
	uc.log.WithContext(ctx).Infof("GetHotKeywords: limit=%d", limit)
	// 返回一些默认的热门关键词
	keywords := []string{"感冒药", "止痛药", "消炎药", "维生素", "降压药", "胃药", "咳嗽药", "退烧药"}
	if limit > 0 && limit < len(keywords) {
		keywords = keywords[:limit]
	}
	return keywords, nil
}
