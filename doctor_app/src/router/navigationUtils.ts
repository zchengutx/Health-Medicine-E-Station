import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { log } from '@/utils/logger'

// 导航选项接口
interface NavigationOptions {
  replace?: boolean
  fallback?: boolean
  delay?: number
  showToast?: boolean
}

/**
 * 导航工具类
 * 提供安全的页面导航功能
 */
export const useNavigationUtils = () => {
  const router = useRouter()
  const authStore = useAuthStore()

  /**
   * 安全跳转到指定路径
   * @param path 目标路径
   * @param options 跳转选项
   */
  const safePush = async (path: string, options?: NavigationOptions) => {
    try {
      // 如果设置了延迟，先等待
      if (options?.delay && options.delay > 0) {
        await new Promise(resolve => setTimeout(resolve, options.delay))
      }
      
      if (options?.replace) {
        await router.replace(path)
      } else {
        await router.push(path)
      }
      
      return true
    } catch (error) {
      if (options?.fallback !== false) {
        // 如果路由跳转失败，尝试使用备用方案
        handleNavigationFallback(path)
      }
      
      return false
    }
  }

  /**
   * 处理导航错误的备用方案
   * @param path 目标路径
   */
  const handleNavigationFallback = (path: string) => {
    try {
      // 尝试使用 replace 方式
      router.replace(path)
    } catch (replaceError) {
      // 最后使用 window.location
      window.location.href = path
    }
  }

  /**
   * 跳转到首页
   */
  const toHome = () => {
    safePush('/home')
  }

  /**
   * 强制跳转到首页（用于登录后）
   * 使用多重备用方案确保跳转成功
   */
  const forceToHome = () => {
    // 方案1: 尝试路由跳转
    router.push('/home').catch((error) => {
      // 方案2: 尝试replace跳转
      router.replace('/home').catch((replaceError) => {
        // 方案3: 使用window.location强制跳转
        window.location.href = '/home'
      })
    })
  }

  /**
   * 跳转到登录页
   */
  const toLogin = () => {
    safePush('/login')
  }

  /**
   * 跳转到注册页
   */
  const toRegister = () => {
    safePush('/register')
  }

  /**
   * 返回上一页
   */
  const goBack = () => {
    if (router.options.history.state.back) {
      router.back()
    } else {
      toHome()
    }
  }

  /**
   * 跳转到启动页
   */
  const toSplash = () => {
    safePush('/')
  }

  return {
    safePush,
    toHome,
    forceToHome,
    toLogin,
    toRegister,
    goBack,
    toSplash,
    handleNavigationFallback
  }
}

// 为了向后兼容，导出一个默认的导航工具实例
// 注意：这个实例只能在组件内部使用
export const navigationUtils = {
  safePush: (path: string) => {
    window.location.href = path
  },
  toHome: () => {
    window.location.href = '/home'
  },
  toLogin: () => {
    window.location.href = '/login'
  },
  toRegister: () => {
    window.location.href = '/register'
  },
  goBack: () => {
    window.history.back()
  },
  toSplash: () => {
    window.location.href = '/'
  },
  forceToHome: () => {
    window.location.href = '/home'
  }
} 