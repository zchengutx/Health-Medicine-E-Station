#!/usr/bin/env node

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

console.log('🚀 开始设置医生端前端应用...\n');

// 检查Node.js版本
const nodeVersion = process.version;
const majorVersion = parseInt(nodeVersion.slice(1).split('.')[0]);

if (majorVersion < 16) {
  console.error('❌ 需要Node.js 16或更高版本');
  process.exit(1);
}

console.log(`✅ Node.js版本: ${nodeVersion}`);

// 检查包管理器
let packageManager = 'npm';
if (fs.existsSync('yarn.lock')) {
  packageManager = 'yarn';
} else if (fs.existsSync('pnpm-lock.yaml')) {
  packageManager = 'pnpm';
}

console.log(`📦 使用包管理器: ${packageManager}\n`);

try {
  // 安装依赖
  console.log('📥 安装依赖包...');
  execSync(`${packageManager} install`, { stdio: 'inherit' });
  
  console.log('\n✅ 项目设置完成！');
  console.log('\n🎉 可以使用以下命令启动开发服务器:');
  console.log(`   ${packageManager} run dev`);
  console.log('\n📖 更多信息请查看 README.md');
  
} catch (error) {
  console.error('\n❌ 安装失败:', error.message);
  process.exit(1);
}