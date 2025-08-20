// API服务层
// 使用全局变量而不是ES6模块
const API_CONFIG = window.API_CONFIG || {
  BASE_URL: 'http://localhost:8888',
  TIMEOUT: 10000
};

// 通用请求方法
class ApiService {
  constructor() {
    this.baseURL = API_CONFIG.BASE_URL;
    this.timeout = API_CONFIG.TIMEOUT;
  }

  // 获取请求头
  getHeaders(includeAuth = true) {
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    };

    if (includeAuth) {
      const token = localStorage.getItem('token');
      if (token) {
        headers['Authorization'] = `Bearer ${token}`;
      }
    }

    return headers;
  }

  // 通用请求方法
  async request(endpoint, options = {}) {
    const url = this.baseURL + endpoint;
    const defaultOptions = {
      headers: this.getHeaders(),
      timeout: this.timeout
    };

    const finalOptions = { ...defaultOptions, ...options };

    try {
      const controller = new AbortController();
      const timeoutId = setTimeout(() => controller.abort(), this.timeout);

      const response = await fetch(url, {
        ...finalOptions,
        signal: controller.signal
      });

      clearTimeout(timeoutId);

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`HTTP ${response.status}: ${errorText}`);
      }

      return await response.json();
    } catch (error) {
      if (error.name === 'AbortError') {
        throw new Error('请求超时');
      }
      console.error('API请求失败:', error);
      throw error;
    }
  }
}

// 创建API服务实例
const apiService = new ApiService();

// 用户相关API
const userApi = {
  // 发送短信验证码
  sendSms: (mobile, source = 'login') => {
    return apiService.request(API_CONFIG.ENDPOINTS.USER.SEND_SMS, {
      method: 'POST',
      body: JSON.stringify({ mobile, source })
    });
  },

  // 用户登录
  login: (mobile, sendSmsCode) => {
    return apiService.request(API_CONFIG.ENDPOINTS.USER.LOGIN, {
      method: 'POST',
      body: JSON.stringify({ mobile, sendSmsCode })
    });
  },

  // 获取用户信息
  getUserInfo: () => {
    return apiService.request(API_CONFIG.ENDPOINTS.USER.GET_USER_INFO, {
      method: 'POST',
      body: JSON.stringify({})
    });
  }
};

// 导出所有API到全局变量
window.userApi = userApi;
window.apiService = apiService;
