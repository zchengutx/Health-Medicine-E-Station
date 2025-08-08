package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

const initOrderCasbinMedicine = initOrderApiMedicine + 1

type initCasbinMedicine struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCasbinMedicine, &initCasbinMedicine{})
}

func (i *initCasbinMedicine) InitializerName() string {
	return "casbin_medicine"
}

func (i *initCasbinMedicine) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (i *initCasbinMedicine) TableCreated(ctx context.Context) bool {
	return true
}

func (i *initCasbinMedicine) InitializeData(ctx context.Context) (next context.Context, err error) {
	_, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 获取casbin实例
	e := utils.GetCasbin()
	if e == nil {
		return ctx, nil
	}

	// 为所有角色添加药品相关API权限
	authorityIds := []string{"888", "8881", "9528"} // 超级管理员、普通用户、测试角色

	// 药品相关API路径
	medicineApis := []string{
		"/mtDrugTypeStair/getAllMtDrugTypeStair",
		"/mtDrugTypeLevel/getAllMtDrugTypeLevel",
		"/mtDrugCategory/getMtDrugCategoryList",
		"/mtDrugCategory/createMtDrugCategory",
		"/mtDrugCategory/updateMtDrugCategory",
		"/mtDrugCategory/deleteMtDrugCategory",
		"/mtDrugCategory/deleteMtDrugCategoryByIds",
		"/mtDrugCategory/findMtDrugCategory",
		"/mtDrugCategory/getAllMtDrugTypeStair",
		"/mtDrugCategory/getAllMtDrugTypeLevel",
		"/mtDrugTypeStair/getMtDrugTypeStairList",
		"/mtDrugTypeStair/createMtDrugTypeStair",
		"/mtDrugTypeStair/updateMtDrugTypeStair",
		"/mtDrugTypeStair/deleteMtDrugTypeStair",
		"/mtDrugTypeStair/deleteMtDrugTypeStairByIds",
		"/mtDrugTypeStair/findMtDrugTypeStair",
		"/mtDrugTypeLevel/getMtDrugTypeLevelList",
		"/mtDrugTypeLevel/createMtDrugTypeLevel",
		"/mtDrugTypeLevel/updateMtDrugTypeLevel",
		"/mtDrugTypeLevel/deleteMtDrugTypeLevel",
		"/mtDrugTypeLevel/deleteMtDrugTypeLevelByIds",
		"/mtDrugTypeLevel/findMtDrugTypeLevel",
	}

	// 为每个角色添加API权限
	for _, authorityId := range authorityIds {
		for _, api := range medicineApis {
			// 添加GET权限
			if api == "/mtDrugTypeStair/getAllMtDrugTypeStair" ||
				api == "/mtDrugTypeLevel/getAllMtDrugTypeLevel" ||
				api == "/mtDrugCategory/getMtDrugCategoryList" ||
				api == "/mtDrugCategory/findMtDrugCategory" ||
				api == "/mtDrugCategory/getAllMtDrugTypeStair" ||
				api == "/mtDrugCategory/getAllMtDrugTypeLevel" ||
				api == "/mtDrugTypeStair/getMtDrugTypeStairList" ||
				api == "/mtDrugTypeStair/findMtDrugTypeStair" ||
				api == "/mtDrugTypeLevel/getMtDrugTypeLevelList" ||
				api == "/mtDrugTypeLevel/findMtDrugTypeLevel" {
				_, err := e.AddPolicy(authorityId, api, "GET")
				if err != nil {
					// 忽略重复策略错误
				}
			}

			// 添加POST权限
			if api == "/mtDrugCategory/createMtDrugCategory" ||
				api == "/mtDrugTypeStair/createMtDrugTypeStair" ||
				api == "/mtDrugTypeLevel/createMtDrugTypeLevel" {
				_, err := e.AddPolicy(authorityId, api, "POST")
				if err != nil {
					// 忽略重复策略错误
				}
			}

			// 添加PUT权限
			if api == "/mtDrugCategory/updateMtDrugCategory" ||
				api == "/mtDrugTypeStair/updateMtDrugTypeStair" ||
				api == "/mtDrugTypeLevel/updateMtDrugTypeLevel" {
				_, err := e.AddPolicy(authorityId, api, "PUT")
				if err != nil {
					// 忽略重复策略错误
				}
			}

			// 添加DELETE权限
			if api == "/mtDrugCategory/deleteMtDrugCategory" ||
				api == "/mtDrugCategory/deleteMtDrugCategoryByIds" ||
				api == "/mtDrugTypeStair/deleteMtDrugTypeStair" ||
				api == "/mtDrugTypeStair/deleteMtDrugTypeStairByIds" ||
				api == "/mtDrugTypeLevel/deleteMtDrugTypeLevel" ||
				api == "/mtDrugTypeLevel/deleteMtDrugTypeLevelByIds" {
				_, err := e.AddPolicy(authorityId, api, "DELETE")
				if err != nil {
					// 忽略重复策略错误
				}
			}
		}
	}

	return ctx, nil
}

func (i *initCasbinMedicine) DataInserted(ctx context.Context) bool {
	e := utils.GetCasbin()
	if e == nil {
		return false
	}

	// 检查是否有药品API权限
	success, _ := e.Enforce("888", "/mtDrugTypeStair/getAllMtDrugTypeStair", "GET")
	return success
}
