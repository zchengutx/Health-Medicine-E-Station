package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

func main() {
	// 初始化路由
	router := initialize.Routers()

	// 测试路由是否注册
	testRoutes := []string{
		"/api/mtDoctorApproval/getMtDoctorApprovalList",
		"/api/mtDoctorApproval/getMtDoctorApprovalByDoctorId",
		"/api/mtDoctorApproval/findMtDoctorApproval",
		"/api/mtDoctorApproval/createMtDoctorApproval",
		"/api/mtDoctorApproval/updateMtDoctorApproval",
		"/api/mtDoctorApproval/deleteMtDoctorApproval",
	}

	fmt.Println("=== 测试审核API路由注册 ===")

	for _, route := range testRoutes {
		req, _ := http.NewRequest("GET", route, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		status := w.Code
		if status == 404 {
			fmt.Printf("❌ 路由未注册: %s (状态码: %d)\n", route, status)
		} else {
			fmt.Printf("✅ 路由已注册: %s (状态码: %d)\n", route, status)
		}
	}

	fmt.Println("\n=== 测试完成 ===")

	// 打印所有注册的路由
	fmt.Println("\n=== 所有注册的路由 ===")
	for _, route := range router.Routes() {
		if strings.Contains(route.Path, "mtDoctorApproval") {
			fmt.Printf("路由: %s %s\n", route.Method, route.Path)
		}
	}
}
