package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	"go.uber.org/zap"
)

func bizModel() error {
	db := global.GVA_DB
	global.GVA_LOG.Info("开始初始化业务表...")
	tables := []interface{}{medicine.MtDrugTypeStair{}, medicine.MtDrugTypeLevel{}, medicine.MtDrug{}, medicine.MtDoctors{}, medicine.MtOrdersDrug{}, medicine.MtUser{}, medicine.MtGuide{}, medicine.MtExplain{}, medicine.MtHospitals{}, medicine.MtDepartments{}, medicine.MtDoctorApproval{}}
	for _, table := range tables {
		global.GVA_LOG.Info("正在创建表: " + getTableName(table))
		err := db.AutoMigrate(table, medicine.MtDrug{}, medicine.MtOrders{}, medicine.MtChatMessage{}, medicine.MtDiscount{})
		if err != nil {
			global.GVA_LOG.Error("创建表失败: "+getTableName(table), zap.Error(err))
			return err
		}
		global.GVA_LOG.Info("表创建成功: " + getTableName(table))
	}
	global.GVA_LOG.Info("所有业务表初始化完成")
	return nil
}
func getTableName(table interface{}) string {
	if t, ok := table.(interface{ TableName() string }); ok {
		return t.TableName()
	}
	return "unknown"
}
