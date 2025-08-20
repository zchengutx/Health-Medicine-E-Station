
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="发送者id:" prop="fromId">
    <el-input v-model.number="formData.fromId" :clearable="true" placeholder="请输入发送者id" />
</el-form-item>
        <el-form-item label="接收者id:" prop="toId">
    <el-input v-model.number="formData.toId" :clearable="true" placeholder="请输入接收者id" />
</el-form-item>
        <el-form-item label="消息内容:" prop="content">
    <el-input v-model="formData.content" :clearable="true" placeholder="请输入消息内容" />
</el-form-item>
        <el-form-item label="消息类型:" prop="messageType">
    <el-select v-model="formData.messageType" placeholder="请选择消息类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in message_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="房间id:" prop="roomId">
    <el-input v-model.number="formData.roomId" :clearable="true" placeholder="请输入房间id" />
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
  createMtChatMessage,
  updateMtChatMessage,
  findMtChatMessage
} from '@/api/medicine/mtChatMessage'

defineOptions({
    name: 'MtChatMessageForm'
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
const message_typeOptions = ref([])
const formData = ref({
            fromId: undefined,
            toId: undefined,
            content: '',
            messageType: '',
            roomId: undefined,
        })
// 验证规则
const rule = reactive({
               fromId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               toId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               messageType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roomId : [{
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
      const res = await findMtChatMessage({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    message_typeOptions.value = await getDictFunc('message_type')
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
               res = await createMtChatMessage(formData.value)
               break
             case 'update':
               res = await updateMtChatMessage(formData.value)
               break
             default:
               res = await createMtChatMessage(formData.value)
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
