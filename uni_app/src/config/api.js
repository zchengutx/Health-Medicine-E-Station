// API 配置文件
const API_CONFIG = {
  // 后端服务基础地址，请根据实际情况修改
  BASE_URL: 'http://localhost:8000',
  
  // API 接口定义
  ENDPOINTS: {
    // 发送短信验证码
    SEND_SMS: '/v1/sendSms',
    // 用户登录
    LOGIN: '/v1/login',
    // 修改昵称
    UPDATE_NICKNAME: '/v1/updateNickName',
    // 获取用户信息
    GET_USER_INFO: '/v1/userInfo',
    // 修改头像
    UPDATE_AVATAR: '/upload',
    // 获取医生列表
    DOCTORS_LIST: '/v1/DoctorsList'
  }
}

// 通用请求方法
const request = async (url, options = {}) => {
  const fullUrl = `${API_CONFIG.BASE_URL}${url}`
  
  // 获取token
  const token = localStorage.getItem('token')
  
  const defaultOptions = {
    headers: {
      'Content-Type': 'application/json',
      ...(token && { 'Authorization': token }),
      ...options.headers
    }
  }
  
  const finalOptions = { ...defaultOptions, ...options }
  
  try {
    const response = await fetch(fullUrl, finalOptions)
    
    if (!response.ok) {
      const errorText = await response.text()
      throw new Error(`HTTP ${response.status}: ${errorText}`)
    }
    
    return await response.json()
  } catch (error) {
    console.error('API请求失败:', error)
    throw error
  }
}

// API 方法
export const api = {
  // 发送短信验证码
  sendSms: async (mobile, source = 'login') => {
    return request(API_CONFIG.ENDPOINTS.SEND_SMS, {
      method: 'POST',
      body: JSON.stringify({ mobile, source })
    })
  },
  
  // 用户登录
  login: async (mobile, sendSmsCode) => {
    return request(API_CONFIG.ENDPOINTS.LOGIN, {
      method: 'POST',
      body: JSON.stringify({ mobile, sendSmsCode })
    })
  },
  
  // 修改昵称
  updateNickname: async (nickname) => {
    return request(API_CONFIG.ENDPOINTS.UPDATE_NICKNAME, {
      method: 'POST',
      body: JSON.stringify({ nickName: nickname })
    })
  },
  
  // 获取用户信息
  getUserInfo: async () => {
    const response = await request(API_CONFIG.ENDPOINTS.GET_USER_INFO, {
      method: 'POST',
      body: JSON.stringify({}) // 空的请求体，根据你的proto定义
    })
    
    // 处理字段名映射 - 后端返回userName，前端期望username
    if (response && response.data) {
      const mappedData = {
        ...response.data,
        username: response.data.userName || response.data.username,
        phone: response.data.mobile || response.data.phone
      }
      return { ...response, data: mappedData }
    }
    
    return response
  },

  // 修改头像
  updateAvatar: async (avatarFile) => {
    // 创建FormData对象用于文件上传
    const formData = new FormData()
    formData.append('file', avatarFile)
    
    // 获取token
    const token = localStorage.getItem('token')
    
    const response = await fetch(`${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.UPDATE_AVATAR}`, {
      method: 'POST',
      headers: {
        'Authorization': token
        // 注意：不要设置Content-Type，让浏览器自动设置multipart/form-data
      },
      body: formData
    })
    
    if (!response.ok) {
      const errorText = await response.text()
      throw new Error(`HTTP ${response.status}: ${errorText}`)
    }
    
    return await response.json()
  },

  // 获取医生列表
  getDoctorsList: async () => {
    return request(API_CONFIG.ENDPOINTS.DOCTORS_LIST, {
      method: 'POST',
      body: JSON.stringify({}) // 空的请求体
    })
  }
}

export default API_CONFIG
