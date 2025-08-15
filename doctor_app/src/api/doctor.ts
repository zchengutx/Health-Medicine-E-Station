import request from './request'
import type { ApiResponse } from './request'

// 发送短信请求参数
export interface SendSmsParams {
  Phone: string
  SendSmsCode: string // 短信类型：login, register, reset_password等
}

// 登录请求参数
export interface LoginParams {
  Phone: string
  Password: string
  SendSmsCode: string
}

// 注册请求参数
export interface RegisterParams {
  Phone: string
  Password: string
  SendSmsCode: string
}

// 修改密码请求参数
export interface ChangePasswordParams {
  DId: number
  OldPassword: string
  NewPassword: string
  ConfirmPassword: string
}

// 认证请求参数
export interface AuthenticationParams {
  DId: number
  Name: string
  Gender: string
  BirthDate: string
  Email: string
  Avatar: string
  LicenseNumber: string
  DepartmentId: number
  HospitalId: number
  Title: string
  Speciality: string
  PracticeScope: string
}

// 更新个人信息请求参数
export interface UpdateProfileParams {
  DId: number
  Name: string
  Gender: string
  BirthDate: string
  Email: string
  Avatar: string
  Title: string
  Speciality: string
  PracticeScope: string
}

// 获取个人信息请求参数
export interface GetProfileParams {
  doctor_id: number
}

// 登录响应数据
export interface LoginResponse extends ApiResponse {
  DId: number
  Token: string
}

// 医生个人信息
export interface DoctorProfile {
  DId: number
  DoctorCode: string
  Name: string
  Gender: string
  BirthDate: string
  Phone: string
  Email: string
  Avatar: string
  LicenseNumber: string
  DepartmentId: number
  HospitalId: number
  Title: string
  Speciality: string
  PracticeScope: string
  Status: string
  CreatedAt: string
  UpdatedAt: string
}

// 获取个人信息响应数据
export interface GetProfileResponse extends ApiResponse {
  Profile: DoctorProfile
}

// 短信验证码类型
export enum SmsCodeType {
  LOGIN = 'login',
  REGISTER = 'register',
  RESET_PASSWORD = 'reset_password',
  CHANGE_PHONE = 'change_phone'
}

// 错误信息映射
const ERROR_MESSAGES: Record<string, string> = {
  'INVALID_PHONE': '手机号格式不正确',
  'PHONE_NOT_FOUND': '手机号不存在，请先注册',
  'INVALID_SMS_CODE': '验证码错误，请重新输入',
  'SMS_CODE_EXPIRED': '验证码已过期，请重新获取',
  'SMS_SEND_FAILED': '验证码发送失败，请稍后重试',
  'SMS_SEND_TOO_FREQUENT': '验证码发送过于频繁，请稍后再试',
  'LOGIN_FAILED': '登录失败，请检查手机号和验证码',
  'ACCOUNT_LOCKED': '账户已被锁定，请联系客服',
  'NETWORK_ERROR': '网络连接失败，请检查网络设置',
  'SERVER_ERROR': '服务器繁忙，请稍后重试'
}

// API错误处理装饰器
function handleApiError<T extends any[], R>(
  target: any,
  propertyKey: string,
  descriptor: TypedPropertyDescriptor<(...args: T) => Promise<R>>
) {
  const originalMethod = descriptor.value!
  
  descriptor.value = async function (...args: T): Promise<R> {
    try {
      const result = await originalMethod.apply(this, args)
      return result
    } catch (error: any) {
      // 记录API调用错误
      console.error(`API Error [${propertyKey}]:`, error.message)
      
      // 处理特定的业务错误
      let enhancedError = error
      
      if (error.message) {
        // 尝试匹配错误信息并提供更友好的提示
        const errorKey = Object.keys(ERROR_MESSAGES).find(key => 
          error.message.toLowerCase().includes(key.toLowerCase().replace('_', ' '))
        )
        
        if (errorKey) {
          enhancedError = new Error(ERROR_MESSAGES[errorKey])
          enhancedError.originalError = error
        } else if (error.message.includes('验证码')) {
          enhancedError = new Error(ERROR_MESSAGES.INVALID_SMS_CODE)
        } else if (error.message.includes('手机号')) {
          enhancedError = new Error(ERROR_MESSAGES.PHONE_NOT_FOUND)
        } else if (error.message.includes('网络') || error.code === 'NETWORK_ERROR') {
          enhancedError = new Error(ERROR_MESSAGES.NETWORK_ERROR)
        } else if (error.response?.status >= 500) {
          enhancedError = new Error(ERROR_MESSAGES.SERVER_ERROR)
        }
      }
      
      throw enhancedError
    }
  }
  
  return descriptor
}

// 医生API接口类
class DoctorApiService {
  // 发送短信验证码
  @handleApiError
  async sendSms(params: SendSmsParams): Promise<ApiResponse> {
    return request.post('/api/v1/doctor/SendSms', params)
  }

  // 医生登录
  @handleApiError
  async login(params: LoginParams): Promise<LoginResponse> {
    return request.post('/api/v1/doctor/LoginDoctor', params)
  }

  // 医生注册
  @handleApiError
  async register(params: RegisterParams): Promise<ApiResponse> {
    return request.post('/api/v1/doctor/RegisterDoctor', params)
  }

  // 个人资料认证
  @handleApiError
  async authentication(params: AuthenticationParams): Promise<ApiResponse> {
    return request.post('/api/v1/doctor/Authentication', params)
  }

  // 获取医生个人信息
  @handleApiError
  async getProfile(params: GetProfileParams): Promise<GetProfileResponse> {
    return request.post('/api/v1/doctor/GetDoctorProfile', params)
  }

  // 更新医生个人信息
  @handleApiError
  async updateProfile(params: UpdateProfileParams): Promise<ApiResponse> {
    return request.post('/api/v1/doctor/UpdateDoctorProfile', params)
  }

  // 修改密码
  @handleApiError
  async changePassword(params: ChangePasswordParams): Promise<ApiResponse> {
    return request.post('/api/v1/doctor/ChangePassword', params)
  }

  // 便捷方法：发送登录验证码
  @handleApiError
  async sendLoginSms(phone: string): Promise<ApiResponse> {
    return this.sendSms({
      Phone: phone,
      SendSmsCode: SmsCodeType.LOGIN
    })
  }

  // 便捷方法：发送注册验证码
  @handleApiError
  async sendRegisterSms(phone: string): Promise<ApiResponse> {
    return this.sendSms({
      Phone: phone,
      SendSmsCode: SmsCodeType.REGISTER
    })
  }

  // 便捷方法：发送重置密码验证码
  @handleApiError
  async sendResetPasswordSms(phone: string): Promise<ApiResponse> {
    return this.sendSms({
      Phone: phone,
      SendSmsCode: SmsCodeType.RESET_PASSWORD
    })
  }

  // 注销账号
  @handleApiError
  async deleteAccount(DId: number): Promise<ApiResponse> {
    return request.post('/api/v1/doctor/DeleteAccount', { DId })
  }
}

// 导出API实例
export const doctorApi = new DoctorApiService()

// 导出类型
export type {
  ApiResponse,
  SendSmsParams,
  LoginParams,
  RegisterParams,
  ChangePasswordParams,
  AuthenticationParams,
  UpdateProfileParams,
  GetProfileParams,
  LoginResponse,
  DoctorProfile,
  GetProfileResponse
}

export default DoctorApiService;