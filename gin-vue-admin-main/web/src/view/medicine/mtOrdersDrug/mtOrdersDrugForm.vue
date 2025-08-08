
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单id:" prop="orderId">
    <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入订单id" />
</el-form-item>
        <el-form-item label="药品id:" prop="drugId">
    <el-input v-model.number="formData.drugId" :clearable="true" placeholder="请输入药品id" />
</el-form-item>
        <el-form-item label="患者id:" prop="userId">
    <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入患者id" />
</el-form-item>
        <el-form-item label="数量:" prop="quantity">
    <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入数量" />
</el-form-item>
        <el-form-item label="订单状态:1-待发货，2-待收货，3-已收货:" prop="orderStatus">
    <el-select v-model="formData.orderStatus" placeholder="请选择订单状态:1-待发货，2-待收货，3-已收货" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in drug_orderOptions" :key="key" :label="item.label" :value="item.value" />
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
  createMtOrdersDrug,
  updateMtOrdersDrug,
  findMtOrdersDrug
} from '@/api/medicine/mtOrdersDrug'

defineOptions({
    name: 'MtOrdersDrugForm'
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
const drug_orderOptions = ref([])
const formData = ref({
            orderId: undefined,
            drugId: undefined,
            userId: undefined,
            quantity: undefined,
            orderStatus: '',
        })
// 验证规则
const rule = reactive({
               orderId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               drugId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               quantity : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               orderStatus : [{
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
      const res = await findMtOrdersDrug({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    drug_orderOptions.value = await getDictFunc('drug_order')
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
               res = await createMtOrdersDrug(formData.value)
               break
             case 'update':
               res = await updateMtOrdersDrug(formData.value)
               break
             default:
               res = await createMtOrdersDrug(formData.value)
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
