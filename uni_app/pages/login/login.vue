<template>
  <view class="container">
    <view class="header">
      <text class="title">登录</text>
    </view>
    <view class="form">
      <view class="input-item">
        <text class="label">手机号码</text>
        <input class="input" type="number" v-model="phone" placeholder="请输入手机号码" maxlength="11"/>
      </view>
      <view class="input-item">
        <text class="label">验证码</text>
        <input class="input" type="number" v-model="code" placeholder="输入验证码" maxlength="6"/>
        <button class="code-btn" :disabled="countDown > 0" @click="getCode">{{ countDown > 0 ? `${countDown}s` : '获取验证码' }}</button>
      </view>
      <button class="login-btn" @click="login">登录</button>
      <text class="quick-login">手机号快捷登录</text>
    </view>
    <view class="agreement">
      <checkbox v-model="agree" class="checkbox"/>
      <text class="text">我已认真阅读，理解并同意《叮当快药用户协议》及《叮当快药隐私协议》。未注册的手机号验证成功后会为您自动创建叮当账户</text>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      phone: '',
      code: '',
      countDown: 0,
      agree: false
    }
  },
  methods: {
    getCode() {
      if (!this.phone || this.phone.length !== 11) {
        uni.showToast({
          title: '请输入正确的手机号码',
          icon: 'none'
        })
        return
      }
      
      // 开始倒计时
      this.countDown = 60
      const timer = setInterval(() => {
        this.countDown--
        if (this.countDown <= 0) {
          clearInterval(timer)
        }
      }, 1000)
      
      // 调用真实的后端API发送验证码
      uni.request({
        url: 'http://localhost:8000/v1/sendSms',
        method: 'POST',
        header: {
          'Content-Type': 'application/json'
        },
        data: {
          mobile: this.phone,
          source: 'login'
        },
        success: (res) => {
          console.log('发送验证码响应:', res.data)
          // 检查多种可能的成功响应格式
          if (res.data.success || res.data.code === 0 || (res.data.message && res.data.message.includes('success'))) {
            uni.showToast({
              title: '验证码发送成功',
              icon: 'success'
            })
          } else {
            uni.showToast({
              title: res.data.message || res.data.msg || '发送失败',
              icon: 'none'
            })
            // 重置倒计时
            this.countDown = 0
            clearInterval(timer)
          }
        },
        fail: (err) => {
          console.error('发送验证码失败:', err)
          uni.showToast({
            title: '网络错误，请重试',
            icon: 'none'
          })
          // 重置倒计时
          this.countDown = 0
          clearInterval(timer)
        }
      })
    },
    login() {
      if (!this.phone || this.phone.length !== 11) {
        uni.showToast({
          title: '请输入正确的手机号码',
          icon: 'none'
        })
        return
      }
      if (!this.code || this.code.length !== 6) {
        uni.showToast({
          title: '请输入正确的验证码',
          icon: 'none'
        })
        return
      }
      if (!this.agree) {
        uni.showToast({
          title: '请同意用户协议和隐私协议',
          icon: 'none'
        })
        return
      }
      
      // 调用真实的后端API进行登录
      uni.request({
        url: 'http://localhost:8000/v1/login',
        method: 'POST',
        header: {
          'Content-Type': 'application/json'
        },
        data: {
          mobile: this.phone,
          sendSmsCode: this.code
        },
        success: (res) => {
          console.log('登录响应:', res.data)
          // 检查多种可能的成功响应格式
          if (res.data.success || res.data.code === 0 || (res.data.message && res.data.message.includes('success'))) {
            // 保存用户信息到本地存储
            const userInfo = {
              username: res.data.data?.username || '用户' + this.phone.substring(7),
              phone: this.phone,
              avatar: res.data.data?.avatar || '',
              token: res.data.data?.token || res.data.token || 'token_' + Date.now()
            }
            
            // 保存登录状态
            uni.setStorageSync('token', userInfo.token)
            uni.setStorageSync('userInfo', userInfo)
            
            // 直接跳转到用户页面，不显示弹框
            uni.switchTab({
              url: '/pages/user/user'
            })
          } else {
            uni.showToast({
              title: res.data.message || res.data.msg || '登录失败',
              icon: 'none'
            })
          }
        },
        fail: (err) => {
          console.error('登录失败:', err)
          uni.showToast({
            title: '网络错误，请重试',
            icon: 'none'
          })
        }
      })
    }
  }
}
</script>

<style scoped>
.container {
  padding: 40rpx;
  box-sizing: border-box;
  min-height: 100vh;
  background-color: #ffffff;
}
.header {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 60rpx 0;
}
.title {
  font-size: 36rpx;
  font-weight: bold;
  color: #333333;
}
.form {
  margin-top: 80rpx;
}
.input-item {
  display: flex;
  align-items: center;
  padding: 20rpx 0;
  border-bottom: 1rpx solid #eeeeee;
  margin-bottom: 30rpx;
}
.label {
  width: 160rpx;
  font-size: 28rpx;
  color: #333333;
}
.input {
  flex: 1;
  font-size: 28rpx;
  color: #333333;
}
.code-btn {
  width: 180rpx;
  height: 60rpx;
  line-height: 60rpx;
  text-align: center;
  font-size: 26rpx;
  color: #ff4d4f;
  background-color: transparent;
  border: none;
  padding: 0;
  margin: 0;
}
.login-btn {
  width: 100%;
  height: 90rpx;
  line-height: 90rpx;
  text-align: center;
  font-size: 32rpx;
  color: #ffffff;
  background-color: #ff4d4f;
  border-radius: 45rpx;
  margin-top: 60rpx;
  border: none;
}
.quick-login {
  display: block;
  text-align: center;
  font-size: 26rpx;
  color: #999999;
  margin-top: 30rpx;
}
.agreement {
  display: flex;
  align-items: flex-start;
  margin-top: 120rpx;
}
.checkbox {
  margin-top: 8rpx;
  margin-right: 10rpx;
  transform: scale(0.8);
}
.text {
  font-size: 24rpx;
  color: #999999;
  line-height: 1.5;
}
</style>