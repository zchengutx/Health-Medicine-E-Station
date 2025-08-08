package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMtDrugTypeStair = initOrderMenu + 1

type initMtDrugTypeStair struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMtDrugTypeStair, &initMtDrugTypeStair{})
}

func (i *initMtDrugTypeStair) InitializerName() string {
	return medicine.MtDrugTypeStair{}.TableName()
}

func (i *initMtDrugTypeStair) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&medicine.MtDrugTypeStair{})
}

func (i *initMtDrugTypeStair) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&medicine.MtDrugTypeStair{})
}

func (i *initMtDrugTypeStair) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 添加初始的一级分类数据
	stairName1 := "处方药"
	stairName2 := "非处方药"
	stairName3 := "保健品"
	stairName4 := "医疗器械"

	code1 := "RX"
	code2 := "OTC"
	code3 := "HEALTH"
	code4 := "DEVICE"

	sort1 := 1
	sort2 := 2
	sort3 := 3
	sort4 := 4

	status := "active"

	entities := []medicine.MtDrugTypeStair{
		{Name: &stairName1, Code: &code1, Sort: &sort1, Status: &status},
		{Name: &stairName2, Code: &code2, Sort: &sort2, Status: &status},
		{Name: &stairName3, Code: &code3, Sort: &sort3, Status: &status},
		{Name: &stairName4, Code: &code4, Sort: &sort4, Status: &status},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, medicine.MtDrugTypeStair{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMtDrugTypeStair) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "处方药").First(&medicine.MtDrugTypeStair{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
