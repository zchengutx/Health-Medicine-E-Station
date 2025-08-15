/**
 * 个人信息页面诊断工具
 * 用于判断问题是出现在前端还是后端
 */

import { useAuthStore } from '@/stores/auth'
import { doctorApi } from '@/api/doctor'
import { log } from '@/utils/logger'

export interface DiagnosticResult {
  step: string
  success: boolean
  error?: string
  data?: any
}

export class ProfileDiagnostic {
  private results: DiagnosticResult[] = []

  async runDiagnostic(): Promise<DiagnosticResult[]> {
    this.results = []
    
    // 步骤1: 检查认证状态
    await this.checkAuthState()
    
    // 步骤2: 检查本地存储
    await this.checkLocalStorage()
    
    // 步骤3: 检查API连接
    await this.checkApiConnection()
    
    // 步骤4: 测试获取个人信息API
    await this.testGetProfileApi()
    
    return this.results
  }

  private async checkAuthState(): Promise<void> {
    try {
      const authStore = useAuthStore()
      
      const result: DiagnosticResult = {
        step: '检查认证状态',
        success: false,
        data: {
          isLoggedIn: authStore.isLoggedIn,
          hasToken: !!authStore.token,
          hasUserInfo: !!authStore.doctorInfo,
          doctorId: authStore.loginState.doctorId,
          isInitialized: authStore.loginState.isInitialized
        }
      }
      
      if (authStore.isLoggedIn && authStore.loginState.doctorId) {
        result.success = true
      } else {
        result.error = '用户未登录或缺少医生ID'
      }
      
      this.results.push(result)
      log.debug('认证状态检查', result)
    } catch (error) {
      this.results.push({
        step: '检查认证状态',
        success: false,
        error: `认证状态检查失败: ${error}`
      })
    }
  }

  private async checkLocalStorage(): Promise<void> {
    try {
      const token = localStorage.getItem('doctor_token')
      const userInfo = localStorage.getItem('doctor_info')
      
      const result: DiagnosticResult = {
        step: '检查本地存储',
        success: false,
        data: {
          hasToken: !!token,
          hasUserInfo: !!userInfo,
          tokenLength: token?.length || 0,
          userInfoValid: false
        }
      }
      
      if (userInfo) {
        try {
          const parsed = JSON.parse(userInfo)
          result.data.userInfoValid = !!parsed.DId
          result.data.doctorId = parsed.DId
        } catch (e) {
          result.error = '用户信息JSON解析失败'
        }
      }
      
      if (token && result.data.userInfoValid) {
        result.success = true
      } else {
        result.error = '本地存储缺少必要信息'
      }
      
      this.results.push(result)
      log.debug('本地存储检查', result)
    } catch (error) {
      this.results.push({
        step: '检查本地存储',
        success: false,
        error: `本地存储检查失败: ${error}`
      })
    }
  }

  private async checkApiConnection(): Promise<void> {
    try {
      // 尝试发送一个简单的请求来测试API连接
      const response = await fetch('/api/v1/doctor/SendSms', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Phone: '13800000000',
          SendSmsCode: 'test'
        })
      })
      
      const result: DiagnosticResult = {
        step: '检查API连接',
        success: false,
        data: {
          status: response.status,
          statusText: response.statusText,
          url: response.url
        }
      }
      
      if (response.status < 500) {
        // 即使是400错误也说明API是可达的
        result.success = true
      } else {
        result.error = `API服务器错误: ${response.status} ${response.statusText}`
      }
      
      this.results.push(result)
      log.debug('API连接检查', result)
    } catch (error) {
      this.results.push({
        step: '检查API连接',
        success: false,
        error: `API连接失败: ${error}`
      })
    }
  }

  private async testGetProfileApi(): Promise<void> {
    try {
      const authStore = useAuthStore()
      const doctorId = authStore.loginState.doctorId
      
      if (!doctorId) {
        this.results.push({
          step: '测试获取个人信息API',
          success: false,
          error: '缺少医生ID，无法测试API'
        })
        return
      }
      
      const response = await doctorApi.getProfile({ doctor_id: doctorId })
      
      const result: DiagnosticResult = {
        step: '测试获取个人信息API',
        success: false,
        data: {
          hasResponse: !!response,
          hasProfile: !!response?.Profile,
          profileData: response?.Profile ? {
            DId: response.Profile.DId,
            Name: response.Profile.Name,
            Phone: response.Profile.Phone,
            hasRequiredFields: !!(response.Profile.DId && response.Profile.Name)
          } : null
        }
      }
      
      if (response && response.Profile && response.Profile.DId) {
        result.success = true
      } else {
        result.error = 'API返回数据格式异常'
      }
      
      this.results.push(result)
      log.debug('获取个人信息API测试', result)
    } catch (error: any) {
      this.results.push({
        step: '测试获取个人信息API',
        success: false,
        error: `API调用失败: ${error.message}`,
        data: {
          errorType: error.name,
          errorMessage: error.message,
          originalError: error.originalError?.message
        }
      })
    }
  }

  // 生成诊断报告
  generateReport(): string {
    let report = '=== 个人信息页面诊断报告 ===\n\n'
    
    this.results.forEach((result, index) => {
      report += `${index + 1}. ${result.step}: ${result.success ? '✅ 通过' : '❌ 失败'}\n`
      
      if (result.error) {
        report += `   错误: ${result.error}\n`
      }
      
      if (result.data) {
        report += `   数据: ${JSON.stringify(result.data, null, 2)}\n`
      }
      
      report += '\n'
    })
    
    // 分析问题类型
    const failedSteps = this.results.filter(r => !r.success)
    
    if (failedSteps.length === 0) {
      report += '🎉 所有检查都通过了，问题可能是临时的\n'
    } else {
      report += '🔍 问题分析:\n'
      
      const authFailed = failedSteps.some(r => r.step.includes('认证'))
      const storageFailed = failedSteps.some(r => r.step.includes('存储'))
      const apiFailed = failedSteps.some(r => r.step.includes('API'))
      
      if (authFailed || storageFailed) {
        report += '- 前端问题: 认证状态或本地存储异常\n'
        report += '- 建议: 重新登录或清理浏览器缓存\n'
      }
      
      if (apiFailed) {
        report += '- 后端问题: API服务异常\n'
        report += '- 建议: 检查后端服务状态和数据库连接\n'
      }
    }
    
    return report
  }
}

// 导出便捷函数
export const runProfileDiagnostic = async (): Promise<string> => {
  const diagnostic = new ProfileDiagnostic()
  await diagnostic.runDiagnostic()
  return diagnostic.generateReport()
}