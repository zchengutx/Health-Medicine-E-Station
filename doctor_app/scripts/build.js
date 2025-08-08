#!/usr/bin/env node

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

// 颜色定义
const colors = {
  reset: '\x1b[0m',
  red: '\x1b[31m',
  green: '\x1b[32m',
  yellow: '\x1b[33m',
  blue: '\x1b[34m',
  magenta: '\x1b[35m',
  cyan: '\x1b[36m'
};

function log(message, color = 'reset') {
  console.log(`${colors[color]}${message}${colors.reset}`);
}

function formatBytes(bytes) {
  if (bytes === 0) return '0 Bytes';
  const k = 1024;
  const sizes = ['Bytes', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
}

function getDirectorySize(dirPath) {
  let totalSize = 0;
  
  function calculateSize(currentPath) {
    const stats = fs.statSync(currentPath);
    
    if (stats.isDirectory()) {
      const files = fs.readdirSync(currentPath);
      files.forEach(file => {
        calculateSize(path.join(currentPath, file));
      });
    } else {
      totalSize += stats.size;
    }
  }
  
  if (fs.existsSync(dirPath)) {
    calculateSize(dirPath);
  }
  
  return totalSize;
}

function main() {
  const startTime = Date.now();
  const environment = process.argv[2] || 'production';
  const analyze = process.argv.includes('--analyze');
  
  log('🚀 开始构建医生端前端应用', 'blue');
  log(`📦 环境: ${environment}`, 'cyan');
  log(`⏰ 开始时间: ${new Date().toLocaleString()}`, 'cyan');
  
  try {
    // 清理之前的构建
    log('🧹 清理之前的构建...', 'yellow');
    if (fs.existsSync('dist')) {
      execSync('rm -rf dist', { stdio: 'inherit' });
    }
    
    // 安装依赖
    log('📦 检查依赖...', 'yellow');
    if (!fs.existsSync('node_modules')) {
      log('📥 安装依赖...', 'yellow');
      execSync('npm ci', { stdio: 'inherit' });
    }
    
    // 类型检查
    log('🔍 类型检查...', 'yellow');
    execSync('npm run type-check', { stdio: 'inherit' });
    
    // 运行测试
    log('🧪 运行测试...', 'yellow');
    execSync('npm run test:run', { stdio: 'inherit' });
    
    // 构建应用
    log('🏗️ 构建应用...', 'yellow');
    const buildCommand = environment === 'production' 
      ? 'npm run build:prod' 
      : 'npm run build:dev';
    
    execSync(buildCommand, { stdio: 'inherit' });
    
    // 分析构建结果
    if (analyze) {
      log('📊 分析构建结果...', 'yellow');
      execSync('npm run build:analyze', { stdio: 'inherit' });
    }
    
    // 构建统计
    const endTime = Date.now();
    const buildTime = ((endTime - startTime) / 1000).toFixed(2);
    const distSize = getDirectorySize('dist');
    
    log('✅ 构建完成！', 'green');
    log(`⏱️ 构建时间: ${buildTime}s`, 'green');
    log(`📦 构建大小: ${formatBytes(distSize)}`, 'green');
    log(`🎯 环境: ${environment}`, 'green');
    
    // 生成构建报告
    const buildReport = {
      timestamp: new Date().toISOString(),
      environment,
      buildTime: `${buildTime}s`,
      size: formatBytes(distSize),
      sizeBytes: distSize,
      nodeVersion: process.version,
      success: true
    };
    
    fs.writeFileSync('dist/build-report.json', JSON.stringify(buildReport, null, 2));
    log('📋 构建报告已生成: dist/build-report.json', 'cyan');
    
    // 检查关键文件
    const criticalFiles = ['index.html', 'assets'];
    const missingFiles = criticalFiles.filter(file => 
      !fs.existsSync(path.join('dist', file))
    );
    
    if (missingFiles.length > 0) {
      log(`⚠️ 警告: 缺少关键文件: ${missingFiles.join(', ')}`, 'yellow');
    }
    
    // 输出文件列表
    log('📁 构建文件:', 'cyan');
    execSync('find dist -type f -name "*.js" -o -name "*.css" -o -name "*.html" | head -10', { 
      stdio: 'inherit' 
    });
    
  } catch (error) {
    log('❌ 构建失败', 'red');
    log(error.message, 'red');
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}