@echo off
echo 正在启动健康医疗驿站项目...
echo.
echo 项目将在浏览器中打开: http://localhost:8080
echo 按 Ctrl+C 停止服务器
echo.
python -m http.server 8080
pause
