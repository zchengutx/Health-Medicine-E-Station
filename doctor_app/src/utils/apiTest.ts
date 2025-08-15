/**
 * API测试工具
 * 用于测试后端API是否正常工作
 */

import { log } from '@/utils/logger'

export interface ApiTestResult {
  endpoint: string
  method: string
  success: boolean
  status?: number
  data?: any
  error?: string
  responseTime?: number
}

export class ApiTester {
  private baseUrl = '/api/v1/doctor'

  async testGetDoctorProfile(doctorId: number): Promise<ApiTestResult> {
    const startTime = Date.now()
    
    try {
      const response = await fetch(`${this.baseUrl}/GetDoctorProfile`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('doctor_token') || ''}`
        },
        body: JSON.stringify({
          doctor_id: doctorId
        })
      })
      
      const responseTime = Date.now() - startTime
      const data = await response.json()
      
      const result: ApiTestResult = {
        endpoint: '/GetDoctorProfile',
        method: 'POST',
        success: response.ok,
        status: response.status,
        data,
        responseTime
      }
      
      if (!response.ok) {
        result.error = `HTTP ${response.status}: ${data.Message || response.statusText}`
      }
      
      log.debug('API测试结果', result)
      return result
      
    } catch (error: any) {
      return {
        endpoint: '/GetDoctorProfile',
        method: 'POST',
        success: false,
        error: error.message,
        responseTime: Date.now() - startTime
      }
    }
  }

  async testApiConnection(): Promise<ApiTestResult> {
    const startTime = Date.now()
    
    try {
      // 测试一个简单的API端点
      const response = await fetch(`${this.baseUrl}/SendSms`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Phone: '13800000000',
          SendSmsCode: 'test'
        })
      })
      
      const responseTime = Date.now() - startTime
      const data = await response.json()
      
      return {
        endpoint: '/SendSms',
        method: 'POST',
        success: response.status < 500, // 即使是400错误也说明API可达
        status: response.status,
        data,
        responseTime
      }
      
    } catch (error: any) {
      return {
        endpoint: '/SendSms',
        method: 'POST',
        success: false,
        error: error.message,
        responseTime: Date.now() - startTime
      }
    }
  }

  async runFullTest(doctorId?: number): Promise<ApiTestResult[]> {
    const results: ApiTestResult[] = []
    
    // 测试API连接
    results.push(await this.testApiConnection())
    
    // 如果有医生ID，测试获取个人信息
    if (doctorId) {
      results.push(await this.testGetDoctorProfile(doctorId))
    }
    
    return results
  }

  generateTestReport(results: ApiTestResult[]): string {
    let report = '=== API测试报告 ===\n\n'
    
    results.forEach((result, index) => {
      report += `${index + 1}. ${result.method} ${result.endpoint}\n`
      report += `   状态: ${result.success ? '✅ 成功' : '❌ 失败'}\n`
      report += `   HTTP状态码: ${result.status || 'N/A'}\n`
      report += `   响应时间: ${result.responseTime || 0}ms\n`
      
      if (result.error) {
        report += `   错误: ${result.error}\n`
      }
      
      if (result.data) {
        report += `   响应数据: ${JSON.stringify(result.data, null, 2)}\n`
      }
      
      report += '\n'
    })
    
    const successCount = results.filter(r => r.success).length
    const totalCount = results.length
    
    report += `总结: ${successCount}/${totalCount} 个测试通过\n`
    
    if (successCount === totalCount) {
      report += '🎉 所有API测试都通过了！\n'
    } else {
      report += '⚠️ 部分API测试失败，请检查后端服务\n'
    }
    
    return report
  }
}

// 导出便捷函数
export const testProfileApi = async (doctorId?: number): Promise<string> => {
  const tester = new ApiTester()
  const results = await tester.runFullTest(doctorId)
  return tester.generateTestReport(results)
}