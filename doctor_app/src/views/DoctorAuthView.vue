<template>
  <div class="doctor-auth-view">
    <h2 class="title">请如实填写您的个人信息</h2>
    <form class="auth-form" @submit.prevent="handleSubmit">
      <div class="form-group" v-for="item in formItems" :key="item.key">
        <label :for="item.key" class="form-label">
          <span v-if="item.required" class="required">*</span>{{ item.label }}
        </label>
        <input
          v-model="form[item.key]"
          :id="item.key"
          :placeholder="item.placeholder"
          :type="item.type || 'text'"
          class="form-input"
          :required="item.required"
        />
      </div>
      <FeedbackButton
        class="submit-btn"
        type="primary"
        block
        :loading="loading"
        text="下一步"
        @click="handleSubmit"
      />
    </form>
    <ToastMessage v-if="toast.visible" :message="toast.message" :type="toast.type" @close="toast.visible=false" />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import FeedbackButton from '@/components/FeedbackButton.vue'
import ToastMessage from '@/components/ToastMessage.vue'
import DoctorApiService from '@/api/doctor'
import { useRouter } from 'vue-router'

const router = useRouter()
const loading = ref(false)
const toast = reactive({ visible: false, message: '', type: 'info' })

const form = reactive({
  realName: '',
  hospital: '',
  department: '',
  title: '',
  specialty: '',
  profile: '',
  experience: ''
})

const formItems = [
  { key: 'realName', label: '真实姓名', placeholder: '请填写您的真实姓名', required: true },
  { key: 'hospital', label: '就职医院', placeholder: '请填写您目前所执业的医院', required: true },
  { key: 'department', label: '所属科室', placeholder: '请填写您所属的科室', required: true },
  { key: 'title', label: '职称', placeholder: '请填写您的职称', required: true },
  { key: 'specialty', label: '擅长领域', placeholder: '请填写您的擅长领域', required: false },
  { key: 'profile', label: '个人简介', placeholder: '请填写个人简介', required: false },
  { key: 'experience', label: '职业经历', placeholder: '请填写您的职业经历', required: false }
]

const doctorApi = new DoctorApiService()
const handleSubmit = async () => {
  if (loading.value) return
  for (const item of formItems) {
    if (item.required && !form[item.key]) {
      toast.message = `${item.label}为必填项`
      toast.type = 'error'
      toast.visible = true
      return
    }
  }
  loading.value = true
  try {
    await doctorApi.authentication({ ...form })
    toast.message = '认证信息提交成功！'
    toast.type = 'success'
    toast.visible = true
    setTimeout(() => {
      router.replace('/mine')
    }, 1200)
  } catch (e) {
    toast.message = '提交失败，请重试'
    toast.type = 'error'
    toast.visible = true
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.doctor-auth-view {
  padding: 20px 16px;
}
.title {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 18px;
  text-align: left;
}
.auth-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.form-label {
  font-size: 15px;
  color: #333;
}
.required {
  color: #ff4d4f;
  margin-right: 2px;
}
.form-input {
  padding: 8px 12px;
  border: 1px solid #e5e5e5;
  border-radius: 6px;
  font-size: 15px;
  outline: none;
}
.submit-btn {
  margin-top: 24px;
}
</style>