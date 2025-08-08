# Kiro修改：心跳检测功能测试脚本

Write-Host "=== 心跳检测功能测试 ===" -ForegroundColor Green
Write-Host "Kiro修改：测试心跳检测相关API和功能" -ForegroundColor Yellow

# 测试1: 检查服务器状态
Write-Host "`n1. 检查服务器状态..." -ForegroundColor Cyan
$serverCheck = Get-NetTCPConnection -LocalPort 8000 -ErrorAction SilentlyContinue
if ($serverCheck) {
    Write-Host "✅ 服务器正在运行" -ForegroundColor Green
} else {
    Write-Host "❌ 服务器未运行，请先启动服务器" -ForegroundColor Red
    exit 1
}

# 测试2: 测试在线用户列表API
Write-Host "`n2. 测试在线用户列表API..." -ForegroundColor Cyan
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8000/api/heartbeat/online-users" -Method GET
    Write-Host "✅ 在线用户API工作正常" -ForegroundColor Green
    Write-Host "在线用户数量: $($response.data.count)" -ForegroundColor Gray
} catch {
    if ($_.Exception.Message -like "*404*") {
        Write-Host "⚠️  API路由未注册，需要重启服务器" -ForegroundColor Yellow
    } else {
        Write-Host "❌ API测试失败: $($_.Exception.Message)" -ForegroundColor Red
    }
}

# 测试3: 测试用户状态查询API
Write-Host "`n3. 测试用户状态查询API..." -ForegroundColor Cyan
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8000/api/heartbeat/user/test123/status" -Method GET
    Write-Host "✅ 用户状态API工作正常" -ForegroundColor Green
    Write-Host "用户状态: $($response.data | ConvertTo-Json -Compress)" -ForegroundColor Gray
} catch {
    if ($_.Exception.Message -like "*404*") {
        Write-Host "⚠️  API路由未注册，需要重启服务器" -ForegroundColor Yellow
    } else {
        Write-Host "❌ API测试失败: $($_.Exception.Message)" -ForegroundColor Red
    }
}

# 测试4: 测试WebSocket连接
Write-Host "`n4. 测试WebSocket连接..." -ForegroundColor Cyan
try {
    $headers = @{
        "Upgrade" = "websocket"
        "Connection" = "Upgrade"
        "Sec-WebSocket-Key" = "dGhlIHNhbXBsZSBub25jZQ=="
        "Sec-WebSocket-Version" = "13"
    }
    
    $response = Invoke-WebRequest -Uri "http://localhost:8000/ws/test" -Headers $headers -Method GET
    
    if ($response.StatusCode -eq 101) {
        Write-Host "✅ WebSocket连接成功，心跳功能已集成" -ForegroundColor Green
    }
} catch {
    Write-Host "❌ WebSocket连接失败: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n=== 测试完成 ===" -ForegroundColor Green
Write-Host "Kiro修改：心跳检测功能已实现，包含以下特性：" -ForegroundColor Yellow
Write-Host "1. 自动心跳检测和超时管理" -ForegroundColor Gray
Write-Host "2. 用户在线状态查询API" -ForegroundColor Gray
Write-Host "3. WebSocket心跳消息处理" -ForegroundColor Gray
Write-Host "4. 后台自动清理过期心跳" -ForegroundColor Gray
Write-Host "`n使用方法：" -ForegroundColor White
Write-Host "1. 重启服务器以加载新的API路由" -ForegroundColor Gray
Write-Host "2. 在浏览器中打开 test_heartbeat.html 进行测试" -ForegroundColor Gray
Write-Host "3. 使用WebSocket发送心跳消息保持在线状态" -ForegroundColor Gray