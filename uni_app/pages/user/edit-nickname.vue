<template>
  <view class="edit-nickname-page">
    <!-- 顶部导航 -->
    <view class="nav-header">
      <view class="nav-left" @click="goBack">
        <text class="back-icon">‹</text>
      </view>
      <view class="nav-title">更改昵称</view>
      <view class="nav-right">
        <text class="more-icon">⋯</text>
      </view>
    </view>

    <!-- 昵称输入区域 -->
    <view class="input-section">
      <view class="input-container">
        <input 
          class="nickname-input" 
          v-model="nickname" 
          placeholder="请输入昵称"
          maxlength="12"
          @input="onInput"
        />
        <view class="clear-btn" @click="clearInput" v-if="nickname">
          <text class="clear-icon">×</text>
        </view>
      </view>
      <view class="input-hint">
        <text class="hint-text">2-12个字符,可输入汉字、数字、字母</text>
      </view>
    </view>

    <!-- 保存按钮 -->
    <view class="save-section">
      <button class="save-btn" @click="saveNickname" :disabled="!canSave">
        <text class="save-text">保存</text>
      </button>
    </view>
  </view>
</template>

<script>
import API_CONFIG from '@/src/config/api.js'
export default {
  name: 'EditNickname',
  data() {
    return {
      nickname: '',
      originalNickname: ''
    }
  },
  computed: {
    canSave() {
      return this.nickname.trim().length >= 2 && this.nickname.trim() !== this.originalNickname
    }
  },
  onLoad(options) {
    // 获取当前昵称
    const userInfo = uni.getStorageSync('userInfo')
    if (userInfo && userInfo.username) {
      this.nickname = userInfo.username
      this.originalNickname = userInfo.username
    }
  },
  methods: {
    // 返回上一页
    goBack() {
      uni.navigateBack()
    },

    // 输入处理
    onInput(e) {
      this.nickname = e.detail.value
    },

    // 清空输入
    clearInput() {
      this.nickname = ''
    },

    // 保存昵称
    async saveNickname() {
      const newNickname = this.nickname.trim()
      
      if (newNickname.length < 2) {
        uni.showToast({
          title: '昵称至少需要2个字符',
          icon: 'none'
        })
        return
      }

      if (newNickname.length > 12) {
        uni.showToast({
          title: '昵称不能超过12个字符',
          icon: 'none'
        })
        return
      }

      // 显示加载提示
      uni.showLoading({
        title: '保存中...'
      })

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

        // 调用后端API修改昵称
        const res = await uni.request({
          url: `${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.UPDATE_NICKNAME}`,
          method: 'POST',
          header: {
            'Content-Type': 'application/json',
            'Authorization': token
          },
          data: {
            nickName: newNickname
          }
        })

        uni.hideLoading()

        // 检查响应结果
        if (res.statusCode === 200) {
          const data = res.data
          if (data.success || data.code === 0 || (data.message && data.message.includes('success'))) {
            // 更新本地存储的用户信息
            const userInfo = uni.getStorageSync('userInfo')
            if (userInfo) {
              userInfo.username = newNickname
              uni.setStorageSync('userInfo', userInfo)
            }

            // 显示保存成功提示
            uni.showToast({
              title: '昵称修改成功',
              icon: 'success'
            })

            // 返回上一页
            setTimeout(() => {
              uni.navigateBack()
            }, 1500)
          } else {
            uni.showToast({
              title: data.message || '修改失败',
              icon: 'none'
            })
          }
        } else {
          uni.showToast({
            title: '网络错误',
            icon: 'none'
          })
        }
      } catch (error) {
        uni.hideLoading()
        console.error('修改昵称失败:', error)
        uni.showToast({
          title: '修改失败，请重试',
          icon: 'none'
        })
      }
    }
  }
}
</script>

<style scoped>
.edit-nickname-page {
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

.input-section {
  background: #fff;
  padding: 40rpx;
  margin-top: 20rpx;
}

.input-container {
  position: relative;
  display: flex;
  align-items: center;
}

.nickname-input {
  flex: 1;
  height: 80rpx;
  font-size: 32rpx;
  color: #333;
  border: none;
  outline: none;
  background: transparent;
}

.clear-btn {
  width: 60rpx;
  height: 60rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: 20rpx;
}

.clear-icon {
  font-size: 40rpx;
  color: #999;
}

.input-hint {
  margin-top: 20rpx;
}

.hint-text {
  font-size: 24rpx;
  color: #999;
}

.save-section {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 40rpx;
  background: #fff;
  border-top: 1rpx solid #eee;
}

.save-btn {
  width: 100%;
  height: 90rpx;
  background: #ff4d4f;
  border: none;
  border-radius: 45rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.save-btn:disabled {
  background: #ccc;
}

.save-text {
  font-size: 32rpx;
  color: #fff;
  font-weight: bold;
}
</style>
