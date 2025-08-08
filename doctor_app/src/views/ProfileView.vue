<template>
  <div class="profile-view-scroll">
    <div class="profile-view">
      <h2 class="title">个人信息</h2>
      <div class="profile-form">
        <div class="form-group" v-for="item in formItems" :key="item.key">
          <label :for="item.key" class="form-label">{{ item.label }}</label>
          <template v-if="item.key === 'DepartmentId'">
            <select v-model="form.DepartmentId" class="form-input">
              <option v-for="dept in departments" :key="dept.id" :value="dept.id">{{ dept.name }}</option>
            </select>
          </template>
          <template v-else-if="item.key === 'HospitalId'">
            <select v-model="form.HospitalId" class="form-input">
              <option v-for="hos in hospitals" :key="hos.id" :value="hos.id">{{ hos.name }}</option>
            </select>
          </template>
          <template v-else-if="item.key === 'BirthDate'">
            <input
              v-model="form.BirthDate"
              :id="item.key"
              :placeholder="item.placeholder"
              type="date"
              class="form-input"
              :readonly="item.readonly"
            />
          </template>
          <template v-else>
            <input
              v-model="form[item.key]"
              :id="item.key"
              :placeholder="item.placeholder"
              :type="item.type || 'text'"
              class="form-input"
              :readonly="item.readonly"
            />
          </template>
        </div>
        <FeedbackButton
          class="save-btn"
          type="primary"
          block
          :loading="loading"
          text="保存"
          @click="handleSave"
        />
      </div>
      <ToastMessage v-if="toast.visible" :message="toast.message" :type="toast.type" @close="toast.visible=false" />
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import FeedbackButton from '@/components/FeedbackButton.vue'
import ToastMessage from '@/components/ToastMessage.vue'
import DoctorApiService from '@/api/doctor'
import { useRouter } from 'vue-router'
const router = useRouter()
const loading = ref(false)
const toast = reactive({ visible: false, message: '', type: 'info' })
const doctorApi = new DoctorApiService()
const departments = ref([
  { id: 1, name: '内科' },
  { id: 2, name: '外科' },
  { id: 3, name: '儿科' }
])
const hospitals = ref([
  { id: 1, name: '协和医院' },
  { id: 2, name: '同济医院' },
  { id: 3, name: '人民医院' }
])
const form = reactive({
  DId: 1,
  Name: '',
  Gender: '',
  BirthDate: '',
  Phone: '',
  Email: '',
  Avatar: '',
  LicenseNumber: '',
  DepartmentId: 1,
  HospitalId: 1,
  Title: '',
  Speciality: '',
  PracticeScope: ''
})
const formItems = [
  { key: 'Name', label: '姓名', placeholder: '请输入姓名', readonly: false },
  { key: 'Gender', label: '性别', placeholder: '请输入性别', readonly: false },
  { key: 'BirthDate', label: '出生日期', placeholder: '请输入出生日期', readonly: false },
  { key: 'Phone', label: '手机号', placeholder: '请输入手机号', readonly: false },
  { key: 'Email', label: '邮箱', placeholder: '请输入邮箱', readonly: false },
  { key: 'LicenseNumber', label: '执业证号', placeholder: '请输入执业证号', readonly: false },
  { key: 'DepartmentId', label: '科室', placeholder: '', readonly: false },
  { key: 'HospitalId', label: '医院', placeholder: '', readonly: false },
  { key: 'Title', label: '职称', placeholder: '请输入职称', readonly: false },
  { key: 'Speciality', label: '专长', placeholder: '请输入专长', readonly: false },
  { key: 'PracticeScope', label: '执业范围', placeholder: '请输入执业范围', readonly: false }
]
onMounted(fetchProfile)
function fetchProfile() {
  loading.value = true
  doctorApi.getProfile({ doctor_id: 1 }).then(res => {
    if (res && res.Profile) {
      Object.assign(form, res.Profile)
    }
  }).catch(() => {
    toast.message = '获取个人信息失败'
    toast.type = 'error'
    toast.visible = true
  }).finally(() => {
    loading.value = false
  })
}
async function handleSave() {
  loading.value = true
  try {
    await doctorApi.updateProfile({ ...form })
    toast.message = '保存成功'
    toast.type = 'success'
    toast.visible = true
    setTimeout(() => {
      router.replace('/mine')
    }, 1200)
  } catch (e) {
    toast.message = '保存失败，请重试'
    toast.type = 'error'
    toast.visible = true
  } finally {
    loading.value = false
  }
}
</script>
<style scoped>
.profile-view-scroll {
  height: 100vh;
  overflow-y: auto;
  background: #fff;
}
.profile-view {
  padding: 20px 16px;
}
.title {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 18px;
  text-align: left;
}
.profile-form {
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
.form-input {
  padding: 8px 12px;
  border: 1px solid #e5e5e5;
  border-radius: 6px;
  font-size: 15px;
  outline: none;
}
.save-btn {
  margin-top: 24px;
}
.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin-bottom: 16px;
}
.avatar-img {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  object-fit: cover;
  border: 1px solid #e5e5e5;
}
</style>
