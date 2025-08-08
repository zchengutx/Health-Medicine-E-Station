<template>
  <div class="setting-container">
    <div class="setting-header-with-back">
      <span class="back-btn" @click="goBack">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M15 18L9 12L15 6" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </span>
      <span class="setting-title">设置</span>
    </div>
    <div class="setting-list">
      <div class="setting-list-item">更换手机号</div>
      <div class="setting-list-item">修改密码</div>
      <div class="setting-list-item">常用语</div>
      <div class="setting-list-item" @click="logout" style="color: #e53e3e; font-weight: bold;">退出登录</div>
      <div class="setting-list-item" @click="deleteAccount" style="color: #e53e3e;">注销账号</div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { showToast } from 'vant';
const router = useRouter();
const authStore = useAuthStore();

const goBack = () => {
  router.back();
};

const logout = () => {
  authStore.logout();
  showToast('已退出登录');
  router.replace('/login');
};

const deleteAccount = async () => {
  try {
    await authStore.deleteAccount();
    showToast('账号已注销');
    router.replace('/login');
  } catch (e) {
    showToast('注销失败，请稍后重试');
  }
};
</script>

<style scoped>
.setting-container {
  background: #f8f8f8;
  min-height: 100vh;
}
.setting-header-with-back {
  display: flex;
  align-items: center;
  background: #4a90e2;
  color: #fff;
  font-size: 22px;
  font-weight: bold;
  padding: 24px 16px 16px 16px;
  position: relative;
}
.back-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 8px;
  cursor: pointer;
  width: 32px;
  height: 32px;
}
.setting-title {
  flex: 1;
  text-align: left;
}
.setting-list {
  background: #fff;
  margin: 16px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.02);
}
.setting-list-item {
  padding: 18px 16px;
  border-bottom: 1px solid #f0f0f0;
  font-size: 16px;
  color: #333;
}
.setting-list-item:last-child {
  border-bottom: none;
}
</style>