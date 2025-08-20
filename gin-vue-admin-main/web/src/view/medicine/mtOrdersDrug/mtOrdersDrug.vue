
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
        
            <el-table-column align="left" label="订单编号" prop="order" width="120" />

            <el-table-column align="left" label="药品名称" prop="drug" width="120" />

            <el-table-column align="left" label="患者名称" prop="user" width="120" />

            <el-table-column align="left" label="数量" prop="quantity" width="120" />

            <el-table-column align="left" label="订单状态" prop="orderStatus" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.orderStatus,drug_orderOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>详情</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateMtOrdersDrugFunc(scope.row)">编辑</el-button>
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="订单详情">
      <div class="order-detail-container">
        <!-- 订单信息 -->
        <div class="order-section">
          <h3 class="section-title">订单信息</h3>
          <div class="order-info">
            <div class="info-item">
              <span class="label">订单编号：</span>
              <span class="value">{{ orderDetail.orderInfo?.orderNo || '-' }}</span>
            </div>
            <div class="info-item">
              <span class="label">订单金额：</span>
              <span class="value">¥{{ orderDetail.orderInfo?.totalAmount || '0' }}</span>
            </div>
            <div class="info-item">
              <span class="label">状态：</span>
              <span class="value">{{ getStatusText(orderDetail.orderInfo?.status) }}</span>
            </div>
            <div class="info-item">
              <span class="label">提交时间：</span>
              <span class="value">{{ formatDate(orderDetail.orderInfo?.payTime) }}</span>
            </div>
          </div>
        </div>

        <!-- 货物信息 -->
        <div class="order-section">
          <h3 class="section-title">货物信息</h3>
          <div class="drug-list">
            <div class="drug-item" v-if="orderDetail.drugInfo">
              <div class="drug-image">
                <img src="/src/assets/logo.png" alt="药品图片" />
                <div class="rx-label">Rx</div>
              </div>
              <div class="drug-info">
                <div class="drug-name">{{ orderDetail.drugInfo.drugName }}</div>
                <div class="drug-spec">规格：{{ orderDetail.drugInfo.specification }}</div>
                <div class="drug-dosage">口服：{{ getDosageText(orderDetail.drugInfo.drugName) }}</div>
                <div class="drug-quantity">数量：x{{ orderDetail.drugInfo.quantity }}</div>
              </div>
            </div>
          </div>
        </div>

        <!-- 收货人信息 -->
        <div class="order-section">
          <h3 class="section-title">收货人信息</h3>
          <div class="recipient-info">
            <div class="info-item">
              <span class="label">收货人：</span>
              <span class="value">{{ orderDetail.userInfo?.nickName || '-' }}</span>
            </div>
            <div class="info-item">
              <span class="label">电话：</span>
              <span class="value">{{ orderDetail.userInfo?.mobile || '-' }}</span>
            </div>
            <div class="info-item">
              <span class="label">收货地址：</span>
              <span class="value">{{ orderDetail.addressInfo?.addressDetail || '-' }}</span>
            </div>
          </div>
        </div>
      </div>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createMtOrdersDrug,
  deleteMtOrdersDrug,
  deleteMtOrdersDrugByIds,
  updateMtOrdersDrug,
  findMtOrdersDrug,
  getMtOrdersDrugList,
  getMtOrdersDrugDetail
} from '@/api/medicine/mtOrdersDrug'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'MtOrdersDrug'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
              ],
               drugId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               quantity : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               orderStatus : [{
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
  const table = await getMtOrdersDrugList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    drug_orderOptions.value = await getDictFunc('drug_order')
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
            deleteMtOrdersDrugFunc(row)
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
      const res = await deleteMtOrdersDrugByIds({ IDs })
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
const updateMtOrdersDrugFunc = async(row) => {
    const res = await findMtOrdersDrug({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMtOrdersDrugFunc = async (row) => {
    const res = await deleteMtOrdersDrug({ ID: row.ID })
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
        orderId: undefined,
        drugId: undefined,
        userId: undefined,
        quantity: undefined,
        orderStatus: '',
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
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})
const orderDetail = ref({})

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await getMtOrdersDrugDetail({ ID: row.ID })
  if (res.code === 0) {
    orderDetail.value = res.data
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  orderDetail.value = {}
}

// 获取状态文本
const getStatusText = (status) => {
  const statusMap = {
    '1': '待支付',
    '2': '已支付',
    '3': '配药中',
    '4': '已发货',
    '5': '已完成',
    '6': '已取消'
  }
  return statusMap[status] || '未知状态'
}

// 获取用药指导文本
const getDosageText = (drugName) => {
  if (!drugName) return '请遵医嘱'
  
  const dosageMap = {
    '苯磺酸氨氯地平片': '成人一次5mg(1片),每天一次',
    '琥珀酸美托洛尔缓释片': '成人一次47.5mg(1片),每天一次'
  }
  
  for (const [key, value] of Object.entries(dosageMap)) {
    if (drugName.includes(key)) {
      return value
    }
  }
  
  return '请遵医嘱'
}


</script>

<style scoped>
.order-detail-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.order-section {
  margin-bottom: 30px;
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.section-title {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin-bottom: 15px;
  padding-bottom: 10px;
  border-bottom: 2px solid #409eff;
}

.order-info, .recipient-info {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-item {
  display: flex;
  align-items: center;
  padding: 8px 0;
}

.label {
  font-weight: 600;
  color: #666;
  min-width: 100px;
  margin-right: 10px;
}

.value {
  color: #333;
  flex: 1;
}

.drug-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.drug-item {
  display: flex;
  align-items: flex-start;
  padding: 15px;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.drug-image {
  position: relative;
  margin-right: 15px;
  flex-shrink: 0;
}

.drug-image img {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: 6px;
  border: 1px solid #ddd;
}

.rx-label {
  position: absolute;
  top: -5px;
  right: -5px;
  background: #ff4757;
  color: white;
  font-size: 12px;
  font-weight: bold;
  padding: 2px 6px;
  border-radius: 10px;
  border: 2px solid white;
}

.drug-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.drug-name {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  line-height: 1.4;
}

.drug-spec, .drug-dosage, .drug-quantity {
  font-size: 14px;
  color: #666;
  line-height: 1.3;
}

.drug-quantity {
  color: #409eff;
  font-weight: 600;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .order-detail-container {
    padding: 15px;
  }
  
  .order-section {
    padding: 15px;
  }
  
  .drug-item {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
  
  .drug-image {
    margin-right: 0;
    margin-bottom: 10px;
  }
  
  .info-item {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .label {
    min-width: auto;
    margin-right: 0;
    margin-bottom: 5px;
  }
}
</style>
