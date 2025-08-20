// 药品服务 - 处理所有药品相关的API调用
import { API_CONFIG } from '../config/api.js';

class DrugService {
  constructor() {
    this.baseURL = API_CONFIG.BASE_URL;
  }

  // 获取药品列表
  async getDrugList(params = {}) {
    try {
      const {
        frist_category_id = 0,
        second_category_id = 0,
        keyword = ''
      } = params;

      console.log('🔍 调用药品列表API:', {
        url: `${this.baseURL}${API_CONFIG.ENDPOINTS.DRUGS.LIST}`,
        params: { frist_category_id, second_category_id, keyword }
      });

      const response = await fetch(`${this.baseURL}${API_CONFIG.ENDPOINTS.DRUGS.LIST}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
        body: JSON.stringify({
          frist_category_id,
          second_category_id,
          keyword
        })
      });

      console.log('📡 API响应状态:', response.status);

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }

      const data = await response.json();
      console.log('✅ 药品列表API响应:', data);

      if (data.code != 0) {
        throw new Error(data.msg || 'API调用失败');
      }

      return {
        success: true,
        data: data.drug || [],
        message: data.msg
      };

    } catch (error) {
      console.error('❌ 获取药品列表失败:', error);
      return {
        success: false,
        data: [],
        message: error.message
      };
    }
  }

  // 获取药品详情
  async getDrugDetail(drugId) {
    try {
      console.log('🔍 获取药品详情:', drugId);

      const response = await fetch(`${this.baseURL}${API_CONFIG.ENDPOINTS.DRUGS.DETAIL}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
        body: JSON.stringify({
          id: drugId
        })
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }

      const data = await response.json();
      console.log('✅ 药品详情API响应:', data);

      if (data.code != 0) {
        throw new Error(data.msg || 'API调用失败');
      }

      return {
        success: true,
        data: data.drug,
        message: data.msg
      };

    } catch (error) {
      console.error('❌ 获取药品详情失败:', error);
      return {
        success: false,
        data: null,
        message: error.message
      };
    }
  }

  // 搜索药品
  async searchDrugs(params = {}) {
    try {
      const {
        keyword = '',
        category_id = 0,
        page = 1,
        size = 20,
        only_in_stock = false,
        include_prescription = false
      } = params;

      console.log('🔍 搜索药品:', params);

      const response = await fetch(`${this.baseURL}${API_CONFIG.ENDPOINTS.DRUGS.SEARCH}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
        body: JSON.stringify({
          keyword,
          category_id,
          page,
          size,
          only_in_stock,
          include_prescription
        })
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }

      const data = await response.json();
      console.log('✅ 药品搜索API响应:', data);

      if (data.code != 0) {
        throw new Error(data.msg || 'API调用失败');
      }

      return {
        success: true,
        data: data.drugs || [],
        total: data.total || 0,
        message: data.msg
      };

    } catch (error) {
      console.error('❌ 搜索药品失败:', error);
      return {
        success: false,
        data: [],
        total: 0,
        message: error.message
      };
    }
  }

  // 获取热门搜索
  async getHotSearch(limit = 10) {
    try {
      console.log('🔍 获取热门搜索');

      const response = await fetch(`${this.baseURL}${API_CONFIG.ENDPOINTS.DRUGS.HOT_SEARCH}?limit=${limit}`, {
        method: 'GET',
        headers: {
          'Accept': 'application/json'
        }
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }

      const data = await response.json();
      console.log('✅ 热门搜索API响应:', data);

      if (data.code != 0) {
        throw new Error(data.msg || 'API调用失败');
      }

      return {
        success: true,
        keywords: data.hot_keywords || [],
        symptoms: data.hot_symptoms || [],
        questions: data.hot_questions || [],
        message: data.msg
      };

    } catch (error) {
      console.error('❌ 获取热门搜索失败:', error);
      return {
        success: false,
        keywords: [],
        symptoms: [],
        questions: [],
        message: error.message
      };
    }
  }

  // 格式化药品数据用于显示
  formatDrugForDisplay(drug) {
    return {
      id: drug.id || Math.random(),
      name: drug.drugName || drug.drug_name || '未知药品',
      description: drug.specification || '规格待确认',
      price: drug.price ? drug.price.toFixed(2) : '0.00',
      originalPrice: drug.price ? (drug.price * 1.2).toFixed(2) : '0.00',
      image: drug.exhibitionUrl || drug.exhibition_url || '/static/default-drug.jpg',
      inventory: drug.inventory || 0,
      salesVolume: drug.salesVolume || drug.sales_volume || 0,
      manufacturer: drug.manufacturer || '生产厂家'
    };
  }

  // 批量格式化药品列表
  formatDrugList(drugs) {
    if (!Array.isArray(drugs)) {
      console.warn('⚠️ 药品数据不是数组格式:', drugs);
      return [];
    }
    return drugs.map(drug => this.formatDrugForDisplay(drug));
  }
}

// 创建并导出单例实例
const drugService = new DrugService();
export default drugService;
