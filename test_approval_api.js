// 测试审核API接口
const axios = require('axios');

const baseURL = 'http://localhost:8888/api';

// 测试获取审核记录列表
async function testGetApprovalList() {
  try {
    const response = await axios.get(`${baseURL}/mtDoctorApproval/getMtDoctorApprovalList`, {
      params: {
        page: 1,
        pageSize: 10
      }
    });
    console.log('获取审核记录列表成功:', response.data);
  } catch (error) {
    console.error('获取审核记录列表失败:', error.response?.data || error.message);
  }
}

// 测试根据医生ID获取审核记录
async function testGetApprovalByDoctorId() {
  try {
    const response = await axios.get(`${baseURL}/mtDoctorApproval/getMtDoctorApprovalByDoctorId`, {
      params: {
        doctorId: 1
      }
    });
    console.log('根据医生ID获取审核记录成功:', response.data);
  } catch (error) {
    console.error('根据医生ID获取审核记录失败:', error.response?.data || error.message);
  }
}

// 运行测试
async function runTests() {
  console.log('开始测试审核API接口...');
  
  await testGetApprovalList();
  console.log('---');
  await testGetApprovalByDoctorId();
  
  console.log('测试完成');
}

runTests(); 