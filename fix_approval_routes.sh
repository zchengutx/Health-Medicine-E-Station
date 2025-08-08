#!/bin/bash

echo "=== 修复审核API路由配置 ==="

# 1. 检查文件是否存在
echo "1. 检查必要文件..."
files=(
    "gin-vue-admin-main/server/model/medicine/mt_doctor_approval.go"
    "gin-vue-admin-main/server/model/medicine/request/mt_doctor_approval.go"
    "gin-vue-admin-main/server/service/medicine/mt_doctor_approval.go"
    "gin-vue-admin-main/server/api/v1/medicine/mt_doctor_approval.go"
    "gin-vue-admin-main/server/router/medicine/mt_doctor_approval.go"
)

for file in "${files[@]}"; do
    if [ -f "$file" ]; then
        echo "✅ $file 存在"
    else
        echo "❌ $file 不存在"
    fi
done

# 2. 检查路由注册
echo -e "\n2. 检查路由注册..."
if grep -q "InitMtDoctorApprovalRouter" gin-vue-admin-main/server/initialize/router_biz.go; then
    echo "✅ 路由已在 router_biz.go 中注册"
else
    echo "❌ 路由未在 router_biz.go 中注册"
fi

# 3. 检查API组配置
echo -e "\n3. 检查API组配置..."
if grep -q "MtDoctorApprovalApi" gin-vue-admin-main/server/api/v1/medicine/enter.go; then
    echo "✅ API已在 enter.go 中配置"
else
    echo "❌ API未在 enter.go 中配置"
fi

# 4. 检查服务组配置
echo -e "\n4. 检查服务组配置..."
if grep -q "MtDoctorApprovalService" gin-vue-admin-main/server/service/medicine/enter.go; then
    echo "✅ 服务已在 enter.go 中配置"
else
    echo "❌ 服务未在 enter.go 中配置"
fi

echo -e "\n=== 修复建议 ==="
echo "1. 重启后端服务: cd gin-vue-admin-main/server && go run main.go"
echo "2. 检查数据库表: 确保 mt_doctor_approval 表已创建"
echo "3. 测试API: 使用浏览器开发者工具检查网络请求"

echo -e "\n=== 修复完成 ===" 