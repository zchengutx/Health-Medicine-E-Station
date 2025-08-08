
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="医生编码:" prop="doctorCode">
    <el-input v-model="formData.doctorCode" :clearable="true" placeholder="请输入医生编码" />
</el-form-item>
        <el-form-item label="医生姓名:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入医生姓名" />
</el-form-item>
        <el-form-item label="性别：1-男，2-女:" prop="gender">
    <el-select v-model="formData.gender" placeholder="请选择性别：1-男，2-女" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="医院:" prop="hospitalId">
    <el-select v-model="formData.hospitalId" placeholder="请选择医院" style="width:100%" filterable :clearable="true" @change="onHospitalChange">
        <el-option v-for="item in hospitalOptions" :key="item.ID" :label="item.name" :value="item.ID" />
    </el-select>
</el-form-item>
        <el-form-item label="科室:" prop="departmentId">
    <el-select v-model="formData.departmentId" placeholder="请选择科室" style="width:100%" filterable :clearable="true" :disabled="!formData.hospitalId">
        <el-option v-for="item in departmentOptions" :key="item.ID" :label="item.name" :value="item.ID" />
    </el-select>
</el-form-item>
        <el-form-item label="职称:" prop="title">
    <el-select v-model="formData.title" placeholder="请选择职称" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in job_titleOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="审核状态：0-未通过，1-已通过，2-未审核:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择审核状态：0-未通过，1-已通过，2-未审核" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in doctor_statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="服务审核：0-未审核，1-已审核，2-待审核:" prop="serviceAudit">
    <el-select v-model="formData.serviceAudit" placeholder="请选择服务审核：0-未审核，1-已审核，2-待审核" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in service_auditOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createMtDoctors,
  updateMtDoctors,
  findMtDoctors
} from '@/api/medicine/mtDoctors'
import { getMtHospitalsListPublic } from '@/api/medicine/mtHospitals'
import { getMtDepartmentsListPublic } from '@/api/medicine/mtDepartments'

defineOptions({
    name: 'MtDoctorsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const doctor_statusOptions = ref([])
const service_auditOptions = ref([])
const genderOptions = ref([])
const job_titleOptions = ref([])
const hospitalOptions = ref([])
const departmentOptions = ref([])
const formData = ref({
            doctorCode: '',
            name: '',
            gender: '',
            departmentId: undefined,
            hospitalId: undefined,
            title: '',
            status: '',
            serviceAudit: '',
        })
// 验证规则
const rule = reactive({
               doctorCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gender : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               departmentId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               hospitalId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serviceAudit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMtDoctors({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    doctor_statusOptions.value = await getDictFunc('doctor_status')
    service_auditOptions.value = await getDictFunc('service_audit')
    genderOptions.value = await getDictFunc('gender')
    job_titleOptions.value = await getDictFunc('job_title')
    
    // 获取医院列表
    try {
      const hospitalRes = await getMtHospitalsListPublic()
      if (hospitalRes.code === 0) {
        hospitalOptions.value = hospitalRes.data.list || []
      }
    } catch (error) {
      console.error('获取医院列表失败:', error)
    }
}

// 医院选择变化时获取对应科室
const onHospitalChange = async (hospitalId) => {
  if (!hospitalId) {
    departmentOptions.value = []
    formData.value.departmentId = undefined
    return
  }
  
  try {
    const departmentRes = await getMtDepartmentsListPublic({ hospitalId: hospitalId })
    if (departmentRes.code === 0) {
      departmentOptions.value = departmentRes.data.list || []
    }
  } catch (error) {
    console.error('获取科室列表失败:', error)
    departmentOptions.value = []
  }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createMtDoctors(formData.value)
               break
             case 'update':
               res = await updateMtDoctors(formData.value)
               break
             default:
               res = await createMtDoctors(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
