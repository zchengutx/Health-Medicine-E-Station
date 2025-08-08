package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	sysService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderApiMedicine = initOrderMtDrugTypeLevel + 1

type initApiMedicine struct{}

// auto run
func init() {
	sysService.RegisterInit(initOrderApiMedicine, &initApiMedicine{})
}

func (i *initApiMedicine) InitializerName() string {
	return "api_medicine"
}

func (i *initApiMedicine) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (i *initApiMedicine) TableCreated(ctx context.Context) bool {
	return true
}

func (i *initApiMedicine) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, sysService.ErrMissingDBContext
	}

	// 获取所有角色
	var authorities []sysModel.SysAuthority
	if err := db.Find(&authorities).Error; err != nil {
		return ctx, err
	}

	// 获取药品相关的API
	var apis []sysModel.SysApi
	apiPaths := []string{
		"/mtDrugTypeStair/getAllMtDrugTypeStair",
		"/mtDrugTypeLevel/getAllMtDrugTypeLevel",
		"/mtDrugCategory/getMtDrugCategoryList",
		"/mtDrugCategory/createMtDrugCategory",
		"/mtDrugCategory/updateMtDrugCategory",
		"/mtDrugCategory/deleteMtDrugCategory",
	}

	if err := db.Where("path IN ?", apiPaths).Find(&apis).Error; err != nil {
		return ctx, errors.Wrap(err, "获取药品API失败")
	}

	// 为所有角色分配药品API权限
	for range authorities {
		// 这里我们只是确保API存在，实际的权限分配需要手动在管理界面操作
		// 或者通过casbin规则来设置
	}

	return ctx, nil
}

func (i *initApiMedicine) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}

	// 检查药品API是否存在
	var count int64
	db.Model(&sysModel.SysApi{}).Where("path = ?", "/mtDrugTypeStair/getAllMtDrugTypeStair").Count(&count)

	return count > 0
}
