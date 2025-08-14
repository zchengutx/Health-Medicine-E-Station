const http = require('http');
const url = require('url');

// 模拟医生数据
const mockDoctors = [
    {
        name: "张振北",
        title: "执业医师",
        hospital: "鼎康慈桦互联网医院",
        consultationCount: 8455,
        rating: "100%",
        responseTime: "15s",
        specialties: "妇科（月经不调、妇科炎症的治疗）、男科（阳痿、早泄）、成人消化系统疾病（胃炎、胃溃疡等）",
        tags: ["上呼吸道感染", "胃肠炎", "阳痿", "早泄", "关节炎"],
        consultationMethod: "在线",
        avatar: "👨‍⚕️"
    },
    {
        name: "宋丽娜",
        title: "主治医师",
        hospital: "鼎康慈桦互联网医院",
        consultationCount: 237874,
        rating: "100%",
        responseTime: "12s",
        specialties: "擅长运用中西医结合方法治疗心脑血管病（高血压、糖尿病、失眠）、妇科疾病等",
        tags: ["高血压", "糖尿病", "失眠", "妇科疾病", "心脑血管"],
        consultationMethod: "在线",
        avatar: "👩‍⚕️"
    },
    {
        name: "李明华",
        title: "副主任医师",
        hospital: "鼎康慈桦互联网医院",
        consultationCount: 15623,
        rating: "99%",
        responseTime: "18s",
        specialties: "擅长：儿科常见病、多发病的诊治，儿童生长发育评估，儿童营养指导",
        tags: ["儿童感冒", "发热", "咳嗽", "生长发育", "营养指导"],
        consultationMethod: "在线",
        avatar: "👨‍⚕️"
    },
    {
        name: "王美玲",
        title: "主任医师",
        hospital: "鼎康慈桦互联网医院",
        consultationCount: 45678,
        rating: "100%",
        responseTime: "10s",
        specialties: "擅长：皮肤科常见病、过敏性皮肤病、痤疮、湿疹、银屑病等疾病的诊治",
        tags: ["痤疮", "湿疹", "银屑病", "过敏", "皮炎"],
        consultationMethod: "在线",
        avatar: "👩‍⚕️"
    }
];

// 创建HTTP服务器
const server = http.createServer((req, res) => {
    // 设置CORS头
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS');
    res.setHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization');
    
    // 处理预检请求
    if (req.method === 'OPTIONS') {
        res.writeHead(200);
        res.end();
        return;
    }
    
    const parsedUrl = url.parse(req.url, true);
    const path = parsedUrl.pathname;
    
    console.log(`${req.method} ${path}`);
    
    // 处理医生列表API
    if (path === '/v1/DoctorsList' && req.method === 'POST') {
        let body = '';
        
        req.on('data', chunk => {
            body += chunk.toString();
        });
        
        req.on('end', () => {
            try {
                // 解析请求体（如果有的话）
                const requestData = body ? JSON.parse(body) : {};
                console.log('收到请求数据:', requestData);
                
                // 返回医生列表数据
                const response = {
                    success: true,
                    message: "DoctorsList success",
                    data: mockDoctors
                };
                
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify(response));
                
            } catch (error) {
                console.error('处理请求时出错:', error);
                res.writeHead(400, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ error: 'Invalid JSON' }));
            }
        });
        
        return;
    }
    
    // 处理根路径
    if (path === '/' && req.method === 'GET') {
        res.writeHead(200, { 'Content-Type': 'text/plain' });
        res.end('Mock Doctor Server is running!\n\nAvailable endpoints:\n- POST /v1/DoctorsList - Get doctors list');
        return;
    }
    
    // 404处理
    res.writeHead(404, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({ error: 'Not Found' }));
});

const PORT = 8000;

server.listen(PORT, () => {
    console.log(`🚀 Mock Doctor Server 启动成功!`);
    console.log(`📍 服务地址: http://localhost:${PORT}`);
    console.log(`📋 可用接口:`);
    console.log(`   - POST /v1/DoctorsList - 获取医生列表`);
    console.log(`   - GET / - 服务器状态`);
    console.log(`\n💡 现在可以访问 http://localhost:8080/doctor-list.html 查看医生列表页面`);
});

// 优雅关闭
process.on('SIGINT', () => {
    console.log('\n🛑 正在关闭服务器...');
    server.close(() => {
        console.log('✅ 服务器已关闭');
        process.exit(0);
    });
});

