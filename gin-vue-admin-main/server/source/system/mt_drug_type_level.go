package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMtDrugTypeLevel = initOrderMtDrugTypeStair + 1

type initMtDrugTypeLevel struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMtDrugTypeLevel, &initMtDrugTypeLevel{})
}

func (i *initMtDrugTypeLevel) InitializerName() string {
	return medicine.MtDrugTypeLevel{}.TableName()
}

func (i *initMtDrugTypeLevel) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&medicine.MtDrugTypeLevel{})
}

func (i *initMtDrugTypeLevel) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&medicine.MtDrugTypeLevel{})
}

func (i *initMtDrugTypeLevel) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 添加初始的二级分类数据
	levelName1 := "感冒药"
	levelName2 := "消炎药"
	levelName3 := "维生素"
	levelName4 := "钙片"
	levelName5 := "血压计"
	levelName6 := "体温计"

	code1 := "COLD"
	code2 := "ANTI"
	code3 := "VIT"
	code4 := "CAL"
	code5 := "BP"
	code6 := "TEMP"

	sort1 := 1
	sort2 := 2
	sort3 := 3
	sort4 := 4
	sort5 := 5
	sort6 := 6

	status := "active"

	entities := []medicine.MtDrugTypeLevel{
		{Name: &levelName1, Code: &code1, Sort: &sort1, Status: &status},
		{Name: &levelName2, Code: &code2, Sort: &sort2, Status: &status},
		{Name: &levelName3, Code: &code3, Sort: &sort3, Status: &status},
		{Name: &levelName4, Code: &code4, Sort: &sort4, Status: &status},
		{Name: &levelName5, Code: &code5, Sort: &sort5, Status: &status},
		{Name: &levelName6, Code: &code6, Sort: &sort6, Status: &status},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, medicine.MtDrugTypeLevel{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMtDrugTypeLevel) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "感冒药").First(&medicine.MtDrugTypeLevel{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
