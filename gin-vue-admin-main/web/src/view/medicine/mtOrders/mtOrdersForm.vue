
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单编号:" prop="orderNo">
    <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入订单编号" />
</el-form-item>
        <el-form-item label="患者姓名:" prop="userName">
    <el-input v-model="formData.userName" :clearable="true" placeholder="请输入患者姓名" />
</el-form-item>
        <el-form-item label="患者电话:" prop="userPhone">
    <el-input v-model="formData.userPhone" :clearable="true" placeholder="请输入患者电话" />
</el-form-item>
        <el-form-item label="医生姓名:" prop="doctorName">
    <el-input v-model="formData.doctorName" :clearable="true" placeholder="请输入医生姓名" />
</el-form-item>
        <el-form-item label="地址详情:" prop="addressDetail">
    <el-input v-model="formData.addressDetail" :clearable="true" placeholder="请输入地址详情" />
</el-form-item>
        <el-form-item label="总金额:" prop="totalAmount">
    <el-input-number v-model="formData.totalAmount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="支付方式：1-微信，2-支付宝，3-银行卡:" prop="payType">
    <el-select v-model="formData.payType" placeholder="请选择支付方式：1-微信，2-支付宝，3-银行卡" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in pay_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="订单状态：1-待支付，2-已支付，3-配药中，4-已发货，5-已完成，6-已取消:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择订单状态：1-待支付，2-已支付，3-配药中，4-已发货，5-已完成，6-已取消" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in order_statusOptions" :key="key" :label="item.label" :value="item.value" />
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
  createMtOrders,
  updateMtOrders,
  findMtOrders
} from '@/api/medicine/mtOrders'

defineOptions({
    name: 'MtOrdersForm'
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
const pay_typeOptions = ref([])
const order_statusOptions = ref([])
const formData = ref({
            orderNo: '',
            userName: '',
            userPhone: '',
            doctorName: '',
            addressDetail: '',
            totalAmount: 0,
            payType: '',
            status: '',
        })
// 验证规则
const rule = reactive({
               orderNo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userPhone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               doctorName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               addressDetail : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalAmount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               payType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
      const res = await findMtOrders({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    pay_typeOptions.value = await getDictFunc('pay_type')
    order_statusOptions.value = await getDictFunc('order_status')
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
               res = await createMtOrders(formData.value)
               break
             case 'update':
               res = await updateMtOrders(formData.value)
               break
             default:
               res = await createMtOrders(formData.value)
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
