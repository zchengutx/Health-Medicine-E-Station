// API配置文件
const API_CONFIG = {
  // 后端服务基础地址，请根据实际情况修改
  BASE_URL: 'http://localhost:8888',
  
  // API版本
  API_VERSION: 'v1',
  
  // 完整的API端点配置
  ENDPOINTS: {
    // 用户相关接口
    USER: {
    SEND_SMS: '/v1/sendSms',
    LOGIN: '/v1/login',
      UPLOAD_AVATAR: '/upload',
      GET_USER_INFO: '/v1/GetTargeted',
      SEARCH_CITIES: '/v1/SearchForCities',
      SELECT_CITY: '/v1/SelectTheCity'
    },
    
    // 医生相关接口
    DOCTORS: {
      LIST: '/v1/DoctorsList'
    },
    
    // 药品相关接口
    DRUGS: {
      LIST: '/v1/drug/list',
      DETAIL: '/v1/drug',
      SEARCH: '/v1/drug/search',
      HOT_SEARCH: '/v1/drug/hot-search'
    },
    
    // 购物车相关接口
    CART: {
      CREATE: '/v1/cart/create',
      UPDATE: '/v1/cart/update',
      DELETE: '/v1/cart/delete',
      LIST: '/v1/cart/list'
    },
    
    // 订单相关接口
    ORDER: {
      CREATE: '/v1/order/create',
      LIST: '/v1/order/list',
      DETAIL: '/v1/order/detail',
      UPDATE_STATUS: '/v1/order/update-status',
      CANCEL: '/v1/order/cancel'
    },
    
    // 支付相关接口
    PAYMENT: {
      CREATE: '/v1/payment/create',
      QUERY: '/v1/payment/query',
      REFUND: '/v1/payment/refund'
    },
    
    // 优惠券相关接口
    COUPON: {
      LIST: '/v1/coupon/list',
      APPLY: '/v1/coupon/apply',
      MY_COUPONS: '/v1/coupon/my-coupons'
    },
    
    // 处方相关接口
    PRESCRIPTION: {
      CREATE: '/v1/prescription/create',
      LIST: '/v1/prescription/list',
      DETAIL: '/v1/prescription/detail',
      APPROVE: '/v1/prescription/approve'
    },
    
    // 聊天相关接口
    CHAT: {
      HISTORY: '/v1/chat/history',
      SAVE_MESSAGE: '/v1/chat/save-message',
      ROOMS: '/v1/chat/rooms',
      MARK_READ: '/v1/chat/mark-read',
      WEBSOCKET: '/ws/chat'
    },
    
    // 评估相关接口
    ESTIMATE: {
      CREATE: '/v1/estimate/create',
      LIST: '/v1/estimate/list',
      DETAIL: '/v1/estimate/detail'
    },
    
    // 心跳检测接口
    HEARTBEAT: {
      USER_STATUS: '/api/heartbeat/user/{userId}/status',
      ONLINE_USERS: '/api/heartbeat/online-users',
      UPDATE: '/api/heartbeat/update'
    }
  },
  
  // 请求超时时间（毫秒）
  TIMEOUT: 10000,
  
  // 请求头配置
  HEADERS: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  },
  
  // 获取完整的API URL
  getFullUrl: function(endpoint) {
    return this.BASE_URL + endpoint;
  },
  
  // 获取带版本号的API URL
  getVersionedUrl: function(endpoint) {
    return this.BASE_URL + '/' + this.API_VERSION + endpoint;
  },
  
  // 替换URL中的参数
  replaceParams: function(url, params) {
    let result = url;
    for (const [key, value] of Object.entries(params)) {
      result = result.replace(`{${key}}`, value);
    }
    return result;
  }
};

// 导出配置到全局变量
window.API_CONFIG = API_CONFIG;
