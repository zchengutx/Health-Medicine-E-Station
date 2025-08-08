
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="药品名称:" prop="drugName">
    <el-input v-model="formData.drugName" :clearable="true" placeholder="请输入药品名称" />
</el-form-item>
        <el-form-item label="用药指导id:" prop="guide">
    <el-input v-model.number="formData.guide" :clearable="true" placeholder="请输入用药指导id" />
</el-form-item>
        <el-form-item label="说明书id:" prop="explain">
    <el-input v-model.number="formData.explain" :clearable="true" placeholder="请输入说明书id" />
</el-form-item>
        <el-form-item label="规格:" prop="specification">
    <el-input v-model="formData.specification" :clearable="true" placeholder="请输入规格" />
</el-form-item>
        <el-form-item label="价格:" prop="price">
    <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="销量:" prop="salesVolume">
    <el-input-number v-model="formData.salesVolume" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="库存:" prop="inventory">
    <el-input v-model.number="formData.inventory" :clearable="true" placeholder="请输入库存" />
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in drug_statusOptions" :key="key" :label="item.label" :value="item.value" />
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
  createMtDrug,
  updateMtDrug,
  findMtDrug
} from '@/api/medicine/mtDrug'

defineOptions({
    name: 'MtDrugForm'
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
const drug_statusOptions = ref([])
const formData = ref({
            drugName: '',
            guide: undefined,
            explain: undefined,
            specification: '',
            price: 0,
            salesVolume: 0,
            inventory: undefined,
            status: '',
        })
// 验证规则
const rule = reactive({
               drugName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               guide : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               explain : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               specification : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               price : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               salesVolume : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               inventory : [{
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
      const res = await findMtDrug({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    drug_statusOptions.value = await getDictFunc('drug_status')
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
               res = await createMtDrug(formData.value)
               break
             case 'update':
               res = await updateMtDrug(formData.value)
               break
             default:
               res = await createMtDrug(formData.value)
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
