package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MtCity struct {
	Code    int32  `gorm:"column:code;type:int;comment:行政区划代码;primaryKey;not null;" json:"code"`  // 行政区划代码
	Name    string `gorm:"column:name;type:varchar(40);comment:行政区划名称;" json:"name"`              // 行政区划名称
	Pcode   int32  `gorm:"column:pcode;type:int;comment:上级区划代码;" json:"pcode"`                    // 上级区划代码
	Sname   string `gorm:"column:sname;type:varchar(40);comment:地名简称;" json:"sname"`              // 地名简称
	Level   int32  `gorm:"column:level;type:int;comment:行政区划等级（1：省、直辖市；2：市州；3：区县）;" json:"level"` // 行政区划等级（1：省、直辖市；2：市州；3：区县）
	Mername string `gorm:"column:mername;type:varchar(100);comment:组合名称;" json:"mername"`         // 组合名称
	Pinyin  string `gorm:"column:pinyin;type:varchar(100);comment:拼音;" json:"pinyin"`             // 拼音
}

type MtAddress struct {
	Id              int32  `gorm:"column:id;type:int;comment:主键id;primaryKey;" json:"id"`                           // 主键id
	UserId          int32  `gorm:"column:user_id;type:int;comment:用户id;" json:"user_id"`                            // 用户id
	Consignee       string `gorm:"column:consignee;type:varchar(20);comment:收货人;" json:"consignee"`                 // 收货人
	Mobile          string `gorm:"column:mobile;type:char(11);comment:联系电话;" json:"mobile"`                         // 联系电话
	CityId          int32  `gorm:"column:city_id;type:int;comment:城市id;" json:"city_id"`                            // 城市id
	ShippingAddress string `gorm:"column:shipping_address;type:varchar(100);comment:收货地址;" json:"shipping_address"` // 收货地址
	DoorplateFloor  string `gorm:"column:doorplate_floor;type:varchar(100);comment:门牌楼层;" json:"doorplate_floor"`   // 门牌楼层
	Label           string `gorm:"column:label;type:varchar(20);comment:标签;" json:"label"`                          // 标签
}

func (m *MtCity) TableName() string {
	return "mt_city"
}

func (m *MtAddress) TableName() string {
	return "mt_address"
}

type CityRepo interface {
	Find(ctx context.Context, m *MtCity) (*[]MtCity, error)
	LikeFind(ctx context.Context, m *MtCity) (*[]MtCity, error)
	Create(ctx context.Context, req *MtAddress) (*MtAddress, error)
}

type CityService struct {
	repo CityRepo
	log  *log.Helper
}

func NewCityUsecase(repo CityRepo, logger log.Logger) *CityService {
	return &CityService{repo: repo, log: log.NewHelper(logger)}
}

func (m *CityService) Find(ctx context.Context, req *MtCity) (*[]MtCity, error) {
	m.log.WithContext(ctx).Infof("MtCity %+v", req)
	return m.repo.Find(ctx, req)
}

func (m *CityService) LikeFind(ctx context.Context, req *MtCity) (*[]MtCity, error) {
	m.log.WithContext(ctx).Infof("MtCity %+v", req)
	return m.repo.LikeFind(ctx, req)
}

func (m *CityService) Create(ctx context.Context, req *MtAddress) (*MtAddress, error) {
	m.log.WithContext(ctx).Infof("MtCity %+v", req)
	return m.repo.Create(ctx, req)
}
