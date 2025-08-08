<template>
  <div class="mine-container">
    <div class="mine-header">
      <div class="mine-title">我的</div>
      <div class="mine-header-icons">
        <div class="setting-icon" @click="goSetting">
          <svg width="32" height="32" viewBox="0 0 32 32" fill="none">
            <circle cx="16" cy="16" r="16" fill="#fff"/>
            <path d="M21.5 16a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0Zm-8.5 0a3 3 0 1 0 6 0 3 3 0 0 0-6 0Z" fill="#4a90e2"/>
            <rect x="15" y="8" width="2" height="4" rx="1" fill="#4a90e2"/>
            <rect x="15" y="20" width="2" height="4" rx="1" fill="#4a90e2"/>
            <rect x="8" y="15" width="4" height="2" rx="1" fill="#4a90e2"/>
            <rect x="20" y="15" width="4" height="2" rx="1" fill="#4a90e2"/>
          </svg>
        </div>
      </div>
      <div class="mine-profile" @click="goDoctorAuth">
        <img class="avatar" :src="doctorAvatar" alt="头像" />
        <div class="profile-info">
          <div class="name-title">
            <span class="name">{{ doctorName }}</span>
            <span class="title">{{ doctorTitle }}</span>
          </div>
          <div class="hospital">{{ doctorHospital }}</div>
        </div>
        <i class="iconfont icon-arrow-right"></i>
      </div>
      <div class="edit-profile-btn-wrapper">
        <FeedbackButton text="编辑个人信息" type="primary" block @click="goEditProfile" />
      </div>
    </div>
    <div class="mine-list">
      <div class="mine-list-item" v-for="item in menuList" :key="item.text">
        <i :class="item.icon"></i>
        <span>{{ item.text }}</span>
        <i class="iconfont icon-arrow-right"></i>
      </div>
    </div>
    <div class="mine-footer"></div>
    <div class="mine-tabbar">
      <div class="tabbar-item" :class="{active: tab==='home'}" @click="goHome">
        <i class="iconfont icon-home"></i>
        <span>首页</span>
      </div>
      <div class="tabbar-item" :class="{active: tab==='patient'}" @click="goPatient">
        <i class="iconfont icon-patient"></i>
        <span>患者</span>
      </div>
      <div class="tabbar-item active">
        <i class="iconfont icon-mine"></i>
        <span>我的</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import FeedbackButton from '@/components/FeedbackButton.vue';
const tab = ref('mine');
const router = useRouter();
const menuList = [
  { icon: 'iconfont icon-auth', text: '我的认证' },
  { icon: 'iconfont icon-income', text: '我的收入' },
  { icon: 'iconfont icon-bankcard', text: '我的银行卡' },
  { icon: 'iconfont icon-evaluate', text: '患者评价' },
  { icon: 'iconfont icon-setting', text: '服务设置' },
  { icon: 'iconfont icon-help', text: '帮助中心' }
];
const authStore = useAuthStore();
const doctorName = authStore.doctorName || '未登录';
const doctorTitle = authStore.doctorInfo?.Title || '医师';
const doctorHospital = authStore.doctorInfo?.Speciality || '暂无医院信息';
const doctorAvatar = authStore.doctorAvatar && /^data:image/.test(authStore.doctorAvatar) ? authStore.doctorAvatar : '/default-avatar.svg';
const goSetting = () => {
  router.push('/setting');
};
const goHome = () => {
  tab.value = 'home';
  router.push('/');
};
const goPatient = () => {
  tab.value = 'patient';
  router.push('/patient');
};
const goDoctorAuth = () => {
  router.push('/doctor-auth');
};
const goEditProfile = () => {
  router.push('/profile');
};
</script>

<style scoped>
.mine-container {
  background: #f8f8f8;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
.mine-header {
  background: linear-gradient(180deg, #4a90e2 0%, #357ae8 100%);
  padding: 32px 16px 16px 16px;
  position: relative;
}
.mine-title {
  color: #fff;
  font-size: 24px;
  font-weight: bold;
}
.mine-header-icons {
  position: absolute;
  top: 32px;
  right: 16px;
  display: flex;
  gap: 16px;
  align-items: center;
}
.mine-profile {
  display: flex;
  align-items: center;
  margin-top: 24px;
  background: #fff;
  border-radius: 8px;
  padding: 12px 16px;
}
.avatar {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  margin-right: 16px;
}
.profile-info {
  flex: 1;
}
.name-title {
  display: flex;
  align-items: center;
  gap: 8px;
}
.name {
  font-size: 18px;
  font-weight: bold;
  color: #333;
}
.title {
  font-size: 14px;
  color: #888;
}
.hospital {
  font-size: 14px;
  color: #357ae8;
  margin-top: 4px;
}
.icon-arrow-right {
  color: #ccc;
  font-size: 20px;
}
.mine-list {
  background: #fff;
  margin: 16px 0;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.02);
}
.mine-list-item {
  display: flex;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
  font-size: 16px;
  color: #333;
}
.mine-list-item:last-child {
  border-bottom: none;
}
.mine-list-item i {
  font-size: 20px;
  color: #357ae8;
  margin-right: 12px;
}
.mine-footer {
  flex: 1;
  background: #f5f5f5;
}
.mine-tabbar {
  display: flex;
  justify-content: space-around;
  align-items: center;
  background: #fff;
  border-top: 1px solid #eee;
  height: 64px;
}
.tabbar-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  color: #bbb;
  font-size: 12px;
}
.tabbar-item.active {
  color: #357ae8;
}
.tabbar-item i {
  font-size: 24px;
  margin-bottom: 4px;
}
.setting-icon {
  width: 32px;
  height: 32px;
  background: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(74,144,226,0.08);
  cursor: pointer;
}
</style>