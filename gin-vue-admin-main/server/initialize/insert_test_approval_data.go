package initialize

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	"go.uber.org/zap"
)

// InsertTestApprovalData 插入测试审核数据
func InsertTestApprovalData() {
	// 检查是否已有数据
	var count int64
	global.GVA_DB.Model(&medicine.MtDoctorApproval{}).Count(&count)
	if count > 0 {
		global.GVA_LOG.Info("审核记录表已有数据，跳过测试数据插入")
		return
	}

	// 插入测试数据
	testData := []medicine.MtDoctorApproval{
		{
			DoctorId:       uintPtr(1),
			DoctorName:     stringPtr("张医生"),
			ApprovalStatus: stringPtr("1"),
			ApprovalTime:   timePtr(time.Now().Add(-24 * time.Hour)),
			ApproverId:     uintPtr(1),
			ApproverName:   stringPtr("管理员"),
			ApprovalReason: stringPtr("资质齐全，审核通过"),
			SubmitTime:     timePtr(time.Now().Add(-25 * time.Hour)),
		},
		{
			DoctorId:       uintPtr(2),
			DoctorName:     stringPtr("李医生"),
			ApprovalStatus: stringPtr("0"),
			ApprovalTime:   timePtr(time.Now().Add(-12 * time.Hour)),
			ApproverId:     uintPtr(1),
			ApproverName:   stringPtr("管理员"),
			RejectReason:   stringPtr("执业证书过期，需要重新提交"),
			SubmitTime:     timePtr(time.Now().Add(-13 * time.Hour)),
		},
		{
			DoctorId:       uintPtr(3),
			DoctorName:     stringPtr("王医生"),
			ApprovalStatus: stringPtr("2"),
			SubmitTime:     timePtr(time.Now().Add(-6 * time.Hour)),
		},
		{
			DoctorId:       uintPtr(4),
			DoctorName:     stringPtr("赵医生"),
			ApprovalStatus: stringPtr("1"),
			ApprovalTime:   timePtr(time.Now().Add(-2 * time.Hour)),
			ApproverId:     uintPtr(1),
			ApproverName:   stringPtr("管理员"),
			ApprovalReason: stringPtr("所有材料符合要求，审核通过"),
			SubmitTime:     timePtr(time.Now().Add(-3 * time.Hour)),
		},
		{
			DoctorId:       uintPtr(5),
			DoctorName:     stringPtr("刘医生"),
			ApprovalStatus: stringPtr("0"),
			ApprovalTime:   timePtr(time.Now().Add(-1 * time.Hour)),
			ApproverId:     uintPtr(1),
			ApproverName:   stringPtr("管理员"),
			RejectReason:   stringPtr("缺少必要的执业证书复印件"),
			SubmitTime:     timePtr(time.Now().Add(-2 * time.Hour)),
		},
	}

	for _, data := range testData {
		if err := global.GVA_DB.Create(&data).Error; err != nil {
			global.GVA_LOG.Error("插入测试审核数据失败", zap.Error(err))
		}
	}

	global.GVA_LOG.Info("测试审核数据插入完成")
}

// 辅助函数
func uintPtr(v uint) *uint {
	return &v
}

func stringPtr(v string) *string {
	return &v
}

func timePtr(v time.Time) *time.Time {
	return &v
}
