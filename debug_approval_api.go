package main

import (
	"context"
	"fmt"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
)

func main() {
	// 初始化数据库连接
	initialize.Gorm()

	// 测试审核记录服务
	service := &medicine.MtDoctorApprovalService{}

	// 测试获取审核记录列表
	fmt.Println("=== 测试获取审核记录列表 ===")
	list, total, err := service.GetMtDoctorApprovalInfoList(context.Background(), medicineReq.MtDoctorApprovalSearch{
		PageInfo: struct {
			Page     int `json:"page" form:"page"`
			PageSize int `json:"pageSize" form:"pageSize"`
		}{
			Page:     1,
			PageSize: 10,
		},
	})

	if err != nil {
		log.Printf("获取审核记录列表失败: %v", err)
	} else {
		fmt.Printf("获取成功，总数: %d, 记录数: %d\n", total, len(list))
		for i, record := range list {
			fmt.Printf("记录 %d: 医生ID=%d, 医生姓名=%s, 审核状态=%s\n",
				i+1, *record.DoctorId, *record.DoctorName, *record.ApprovalStatus)
		}
	}

	// 测试根据医生ID获取审核记录
	fmt.Println("\n=== 测试根据医生ID获取审核记录 ===")
	approval, err := service.GetMtDoctorApprovalByDoctorId(context.Background(), 1)
	if err != nil {
		log.Printf("根据医生ID获取审核记录失败: %v", err)
	} else {
		fmt.Printf("获取成功: 医生ID=%d, 医生姓名=%s, 审核状态=%s\n",
			*approval.DoctorId, *approval.DoctorName, *approval.ApprovalStatus)
	}

	fmt.Println("\n=== 测试完成 ===")
}
