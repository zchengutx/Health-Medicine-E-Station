
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="w-[380px]"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="订单编号" prop="orderNo" width="120" />

            <el-table-column align="left" label="患者姓名" prop="userName" width="120" />

            <el-table-column align="left" label="患者电话" prop="userPhone" width="120" />

            <el-table-column align="left" label="医生姓名" prop="doctorName" width="120" />

            <el-table-column align="left" label="地址详情" prop="addressDetail" width="120" />

            <el-table-column align="left" label="总金额" prop="totalAmount" width="120" />

            <el-table-column align="left" label="支付方式：1-微信，2-支付宝，3-银行卡" prop="payType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.payType,pay_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="订单状态：1-待支付，2-已支付，3-配药中，4-已发货，5-已完成，6-已取消" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,order_statusOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateMtOrdersFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="订单编号">
    {{ detailFrom.orderNo }}
</el-descriptions-item>
                    <el-descriptions-item label="患者姓名">
    {{ detailFrom.userName }}
</el-descriptions-item>
                    <el-descriptions-item label="患者电话">
    {{ detailFrom.userPhone }}
</el-descriptions-item>
                    <el-descriptions-item label="医生姓名">
    {{ detailFrom.doctorName }}
</el-descriptions-item>
                    <el-descriptions-item label="地址详情">
    {{ detailFrom.addressDetail }}
</el-descriptions-item>
                    <el-descriptions-item label="总金额">
    {{ detailFrom.totalAmount }}
</el-descriptions-item>
                    <el-descriptions-item label="支付方式：1-微信，2-支付宝，3-银行卡">
    {{ detailFrom.payType }}
</el-descriptions-item>
                    <el-descriptions-item label="订单状态：1-待支付，2-已支付，3-配药中，4-已发货，5-已完成，6-已取消">
    {{ detailFrom.status }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createMtOrders,
  deleteMtOrders,
  deleteMtOrdersByIds,
  updateMtOrders,
  findMtOrders,
  getMtOrdersList
} from '@/api/medicine/mtOrders'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'MtOrders'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               userName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               userPhone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               doctorName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               addressDetail : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               totalAmount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               payType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMtOrdersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    pay_typeOptions.value = await getDictFunc('pay_type')
    order_statusOptions.value = await getDictFunc('order_status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMtOrdersFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteMtOrdersByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMtOrdersFunc = async(row) => {
    const res = await findMtOrders({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMtOrdersFunc = async (row) => {
    const res = await deleteMtOrders({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        orderNo: '',
        userName: '',
        userPhone: '',
        doctorName: '',
        addressDetail: '',
        totalAmount: 0,
        payType: '',
        status: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findMtOrders({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
