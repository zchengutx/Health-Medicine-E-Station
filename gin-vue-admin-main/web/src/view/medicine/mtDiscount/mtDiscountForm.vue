
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="优惠券名称:" prop="discountName">
    <el-input v-model="formData.discountName" :clearable="true" placeholder="请输入优惠券名称" />
</el-form-item>
        <el-form-item label="券的来源分类:" prop="classify">
    <el-select v-model="formData.classify" placeholder="请选择券的来源分类" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in sourceOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="优惠金额:" prop="discountAmout">
    <el-input-number v-model="formData.discountAmout" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="最低消费门槛:" prop="minOrderAmount">
    <el-input-number v-model="formData.minOrderAmount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="有效期开始:" prop="startTime">
    <el-date-picker v-model="formData.startTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="有效期结束:" prop="endTime">
    <el-date-picker v-model="formData.endTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="总发行量，0=不限:" prop="maxIssue">
    <el-input v-model.number="formData.maxIssue" :clearable="true" placeholder="请输入总发行量，0=不限" />
</el-form-item>
        <el-form-item label="每人限领:" prop="maxPerUser">
    <el-switch v-model="formData.maxPerUser" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createMtDiscount,
  updateMtDiscount,
  findMtDiscount
} from '@/api/medicine/mtDiscount'

defineOptions({
    name: 'MtDiscountForm'
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
const sourceOptions = ref([])
const formData = ref({
            discountName: '',
            classify: '',
            discountAmout: 0,
            minOrderAmount: 0,
            startTime: new Date(),
            endTime: new Date(),
            maxIssue: undefined,
            maxPerUser: false,
        })
// 验证规则
const rule = reactive({
               discountName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               classify : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               discountAmout : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               minOrderAmount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               endTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               maxIssue : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               maxPerUser : [{
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
      const res = await findMtDiscount({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    sourceOptions.value = await getDictFunc('source')
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
               res = await createMtDiscount(formData.value)
               break
             case 'update':
               res = await updateMtDiscount(formData.value)
               break
             default:
               res = await createMtDiscount(formData.value)
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
