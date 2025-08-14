<template>
  <view class="profile-page">
    <!-- 顶部导航 -->
    <view class="nav-header">
      <view class="nav-left" @click="goBack">
        <text class="back-icon">‹</text>
      </view>
      <view class="nav-title">用户信息管理</view>
      <view class="nav-right">
        <text class="more-icon">⋯</text>
      </view>
    </view>

    <!-- 用户信息列表 -->
    <view class="profile-list">
      <!-- 头像 -->
      <view class="profile-item" @click="changeAvatar">
        <view class="item-left">
          <text class="item-label">头像</text>
        </view>
        <view class="item-right">
          <image class="avatar" :src="getAvatarUrl(userInfo.avatar)" mode="aspectFill" @error="onAvatarError"></image>
          <text class="arrow">›</text>
        </view>
      </view>

      <!-- 昵称 -->
      <view class="profile-item" @click="changeNickname">
        <view class="item-left">
          <text class="item-label">昵称</text>
        </view>
        <view class="item-right">
          <text class="item-value">{{ userInfo.username || '未设置' }}</text>
          <text class="arrow">›</text>
        </view>
      </view>

      <!-- 手机号 -->
      <view class="profile-item" @click="changePhone">
        <view class="item-left">
          <text class="item-label">解绑手机号</text>
        </view>
        <view class="item-right">
          <text class="item-value">{{ maskPhone(userInfo.phone) }}</text>
          <text class="arrow">›</text>
        </view>
      </view>
    </view>

    <!-- 信息管理 -->
    <view class="section-title">信息管理</view>
    <view class="profile-list">
      <view class="profile-item" @click="handleInfo('inquiry')">
        <view class="item-left">
          <text class="item-label">个人信息查阅和管理</text>
        </view>
        <view class="item-right">
          <text class="arrow">›</text>
        </view>
      </view>

      <view class="profile-item" @click="handleInfo('download')">
        <view class="item-left">
          <text class="item-label">个人信息下载</text>
        </view>
        <view class="item-right">
          <text class="arrow">›</text>
        </view>
      </view>
    </view>

    <!-- 协议和设置 -->
    <view class="section-title">协议和设置</view>
    <view class="profile-list">
      <view class="profile-item" @click="handleInfo('agreement')">
        <view class="item-left">
          <text class="item-label">用户协议</text>
        </view>
        <view class="item-right">
          <text class="arrow">›</text>
        </view>
      </view>

      <view class="profile-item" @click="handleInfo('privacy')">
        <view class="item-left">
          <text class="item-label">隐私协议</text>
        </view>
        <view class="item-right">
          <text class="arrow">›</text>
        </view>
      </view>

      <view class="profile-item" @click="handleInfo('payment')">
        <view class="item-left">
          <text class="item-label">支付密码管理</text>
        </view>
        <view class="item-right">
          <text class="arrow">›</text>
        </view>
      </view>
    </view>

    <!-- 账号管理 -->
    <view class="section-title">账号管理</view>
    <view class="profile-list">
      <view class="profile-item danger" @click="showDeleteConfirm">
        <view class="item-left">
          <text class="item-label">注销账号</text>
        </view>
        <view class="item-right">
          <text class="arrow">›</text>
        </view>
      </view>
      <view class="delete-desc">
        <text class="desc-text">提交申请、删除所有数据,永久注销账号</text>
      </view>
    </view>
  </view>
</template>

<script>
import API_CONFIG from '@/config/api.js'
export default {
  name: 'UserProfile',
  data() {
    return {
      userInfo: {
        username: '',
        phone: '',
        avatar: ''
      }
    }
  },
  onLoad() {
    this.loadUserInfo()
  },
  
  // 页面显示时刷新用户信息（用于从编辑页面返回时更新显示）
  onShow() {
    this.loadUserInfo()
  },
  methods: {
    // 加载用户信息
    async loadUserInfo() {
      // 先从本地存储加载（快速显示）
      const localUserInfo = uni.getStorageSync('userInfo')
      if (localUserInfo) {
        this.userInfo = localUserInfo
      }
      
      // 然后从后端获取最新信息
      try {
        const token = uni.getStorageSync('token')
        if (!token) {
          return
        }
        
        console.log('正在从后端获取用户信息...')
        const res = await uni.request({
          url: `${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.GET_USER_INFO}`,
          method: 'POST',
          header: {
            'Content-Type': 'application/json',
            'Authorization': token
          },
          data: {} // 空的请求体
        })
        
        console.log('后端响应:', res)
        
        if (res.statusCode === 200) {
          const data = res.data
          console.log('后端数据:', data)
          
                      if (data.success || data.code === 0 || (data.message && data.message.includes('success'))) {
              // 更新用户信息 - 优先使用服务器数据
              const serverUserInfo = data.data || data
              if (serverUserInfo) {
                console.log('服务器用户信息:', serverUserInfo)
                
                // 处理字段名映射 - 后端返回userName，前端期望username
                const mappedUserInfo = {
                  username: serverUserInfo.userName || serverUserInfo.username,
                  phone: serverUserInfo.mobile || serverUserInfo.phone,
                  avatar: serverUserInfo.avatar
                }
                
                // 优先使用服务器数据，本地数据作为补充
                this.userInfo = {
                  ...localUserInfo, // 本地数据作为基础
                  ...mappedUserInfo  // 映射后的服务器数据覆盖
                }
                
                // 更新本地存储
                uni.setStorageSync('userInfo', this.userInfo)
                console.log('更新后的用户信息:', this.userInfo)
              }
            } else {
            console.log('后端返回错误:', data)
          }
        } else {
          console.log('HTTP错误:', res.statusCode)
        }
      } catch (error) {
        console.error('获取用户信息失败:', error)
        // 如果获取失败，继续使用本地数据
      }
    },

    // 返回上一页
    goBack() {
      uni.navigateBack()
    },

    // 手机号脱敏
    maskPhone(phone) {
      if (!phone) return '未绑定'
      return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
    },

    // 修改头像
    changeAvatar() {
      uni.chooseImage({
        count: 1,
        sizeType: ['compressed'],
        sourceType: ['album', 'camera'],
        success: (res) => {
          const tempFilePath = res.tempFilePaths[0]
          
          // 显示上传中提示
          uni.showLoading({
            title: '上传中...'
          })
          
          // 上传头像到服务器
          this.uploadAvatar(tempFilePath)
        },
        fail: (err) => {
          console.error('选择图片失败:', err)
          uni.showToast({
            title: '选择图片失败',
            icon: 'none'
          })
        }
      })
    },

    // 上传头像
    async uploadAvatar(filePath) {
      try {
        // 获取token
        const token = uni.getStorageSync('token')
        if (!token) {
          uni.hideLoading()
          uni.showToast({
            title: '请先登录',
            icon: 'none'
          })
          return
        }

        // 上传文件
        const uploadResult = await new Promise((resolve, reject) => {
          uni.uploadFile({
            url: `${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.UPDATE_AVATAR}`,
            filePath: filePath,
            name: 'file',
            header: {
              'Authorization': token
            },
            success: (res) => {
              console.log('上传成功:', res)
              try {
                const data = JSON.parse(res.data)
                resolve(data)
              } catch (e) {
                reject(new Error('解析响应失败'))
              }
            },
            fail: (err) => {
              console.error('上传失败:', err)
              reject(err)
            }
          })
        })

        uni.hideLoading()

        // 检查上传结果
        if (uploadResult.success || uploadResult.code === 0 || (uploadResult.message && uploadResult.message.includes('success'))) {
          // 更新本地用户信息
          let newAvatar = ''
          if (uploadResult.data && uploadResult.data.avatar) {
            newAvatar = uploadResult.data.avatar
          } else if (uploadResult.avatar) {
            newAvatar = uploadResult.avatar
          }
          
          // 强制更新头像 - 使用Vue的$set或直接赋值
          this.$set(this.userInfo, 'avatar', newAvatar)
          // 或者使用这种方式确保响应式更新
          this.userInfo = { ...this.userInfo, avatar: newAvatar }
          
          uni.setStorageSync('userInfo', this.userInfo)
          
          uni.showToast({
            title: '头像更新成功',
            icon: 'success'
          })
          
          console.log('头像更新成功:', newAvatar)
        } else {
          throw new Error(uploadResult.message || '上传失败')
        }
      } catch (error) {
        uni.hideLoading()
        console.error('上传头像失败:', error)
        uni.showToast({
          title: error.message || '上传失败',
          icon: 'none'
        })
      }
    },

              // 修改昵称
    changeNickname() {
      uni.navigateTo({
        url: '/pages/user/edit-nickname'
      })
    },

    // 获取头像URL
    getAvatarUrl(avatarUrl) {
      if (!avatarUrl || avatarUrl.trim() === '') {
        return '/static/default-avatar.png'
      }
      
      // 如果头像URL不是以http开头，则添加后端服务器地址
      if (!avatarUrl.startsWith('http')) {
        return `${API_CONFIG.BASE_URL}/uploads/${avatarUrl}`
      }
      
      return avatarUrl
    },
    
    // 头像加载失败处理
    onAvatarError() {
      console.log('头像加载失败，使用默认头像')
      // uni-app会自动使用默认图片
    },

    // 修改手机号
    changePhone() {
      uni.showModal({
        title: '解绑手机号',
        content: '确定要解绑当前手机号吗？',
        success: (res) => {
          if (res.confirm) {
            uni.showToast({
              title: '功能开发中',
              icon: 'none'
            })
          }
        }
      })
    },

    // 处理信息管理
    handleInfo(type) {
      const actions = {
        inquiry: '个人信息查阅和管理',
        download: '个人信息下载',
        agreement: '用户协议',
        privacy: '隐私协议',
        payment: '支付密码管理'
      }
      
      uni.showToast({
        title: `${actions[type]}功能开发中`,
        icon: 'none'
      })
    },

    // 显示注销确认
    showDeleteConfirm() {
      uni.showModal({
        title: '注销账号',
        content: '确定要注销账号吗？此操作不可恢复，所有数据将被永久删除。',
        confirmColor: '#ff4d4f',
        success: (res) => {
          if (res.confirm) {
            this.deleteAccount()
          }
        }
      })
    },

    // 注销账号
    deleteAccount() {
      uni.showLoading({
        title: '注销中...'
      })
      
      setTimeout(() => {
        uni.hideLoading()
        // 清除本地存储
        uni.removeStorageSync('token')
        uni.removeStorageSync('userInfo')
        
        uni.showToast({
          title: '账号已注销',
          icon: 'success'
        })
        
        // 跳转到登录页
        setTimeout(() => {
          uni.reLaunch({
            url: '/pages/login/login'
          })
        }, 1500)
      }, 2000)
    }
  }
}
</script>

<style scoped>
.profile-page {
  min-height: 100vh;
  background-color: #f5f5f5;
}

.nav-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20rpx 40rpx;
  background: #fff;
  border-bottom: 1rpx solid #eee;
}

.nav-left, .nav-right {
  width: 80rpx;
  display: flex;
  align-items: center;
}

.nav-left {
  justify-content: flex-start;
}

.nav-right {
  justify-content: flex-end;
}

.back-icon, .more-icon {
  font-size: 40rpx;
  color: #333;
}

.nav-title {
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
}

.section-title {
  padding: 30rpx 40rpx 20rpx;
  font-size: 28rpx;
  color: #999;
  background: #f5f5f5;
}

.profile-list {
  background: #fff;
  margin-bottom: 20rpx;
}

.profile-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 30rpx 40rpx;
  border-bottom: 1rpx solid #f0f0f0;
}

.profile-item:last-child {
  border-bottom: none;
}

.profile-item.danger .item-label {
  color: #ff4d4f;
}

.item-left {
  flex: 1;
}

.item-label {
  font-size: 30rpx;
  color: #333;
}

.item-right {
  display: flex;
  align-items: center;
}

.item-value {
  font-size: 28rpx;
  color: #666;
  margin-right: 20rpx;
}

.arrow {
  font-size: 32rpx;
  color: #ccc;
}

.avatar {
  width: 80rpx;
  height: 80rpx;
  border-radius: 40rpx;
  margin-right: 20rpx;
  background: #f0f0f0;
}

.delete-desc {
  padding: 0 40rpx 30rpx;
  background: #fff;
}

.desc-text {
  font-size: 24rpx;
  color: #999;
  line-height: 1.5;
}
</style>
