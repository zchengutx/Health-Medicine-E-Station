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
        
            <el-table-column align="left" label="医生ID" prop="doctorId" width="120" />

            <el-table-column align="left" label="患者ID" prop="patientId" width="120" />

            <el-table-column align="left" label="关系类型" prop="relationshipType" width="120" />

            <el-table-column align="left" label="患者标签" prop="tags" width="120" />

            <el-table-column align="left" label="备注" prop="notes" width="120" />

            <el-table-column align="left" label="首次就诊时间" prop="firstVisitTime" width="180">
                <template #default="scope">{{ formatDate(scope.row.firstVisitTime) }}</template>
            </el-table-column>

            <el-table-column align="left" label="最后就诊时间" prop="lastVisitTime" width="180">
                <template #default="scope">{{ formatDate(scope.row.lastVisitTime) }}</template>
            </el-table-column>

            <el-table-column align="left" label="就诊次数" prop="visitCount" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateMtDoctorPatientsFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="医生ID:" prop="doctorId">
    <el-input v-model.number="formData.doctorId" :clearable="true" placeholder="请输入医生ID" />
</el-form-item>
            <el-form-item label="患者ID:" prop="patientId">
    <el-input v-model.number="formData.patientId" :clearable="true" placeholder="请输入患者ID" />
</el-form-item>
            <el-form-item label="关系类型:" prop="relationshipType">
    <el-input v-model="formData.relationshipType" :clearable="true" placeholder="请输入关系类型: 普通/关注/VIP" />
</el-form-item>
            <el-form-item label="患者标签:" prop="tags">
    <el-input v-model="formData.tags" :clearable="true" placeholder="请输入患者标签，逗号分隔" />
</el-form-item>
            <el-form-item label="备注:" prop="notes">
    <el-input v-model="formData.notes" :clearable="true" placeholder="请输入备注" />
</el-form-item>
            <el-form-item label="首次就诊时间:" prop="firstVisitTime">
    <el-date-picker v-model="formData.firstVisitTime" type="datetime" style="width:100%" placeholder="选择日期时间" :clearable="true" />
</el-form-item>
            <el-form-item label="最后就诊时间:" prop="lastVisitTime">
    <el-date-picker v-model="formData.lastVisitTime" type="datetime" style="width:100%" placeholder="选择日期时间" :clearable="true" />
</el-form-item>
            <el-form-item label="就诊次数:" prop="visitCount">
    <el-input v-model.number="formData.visitCount" :clearable="true" placeholder="请输入就诊次数" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="医生ID">
    {{ detailFrom.doctorId }}
</el-descriptions-item>
                    <el-descriptions-item label="患者ID">
    {{ detailFrom.patientId }}
</el-descriptions-item>
                    <el-descriptions-item label="关系类型">
    {{ detailFrom.relationshipType }}
</el-descriptions-item>
                    <el-descriptions-item label="患者标签">
    {{ detailFrom.tags }}
</el-descriptions-item>
                    <el-descriptions-item label="备注">
    {{ detailFrom.notes }}
</el-descriptions-item>
                    <el-descriptions-item label="首次就诊时间">
    {{ formatDate(detailFrom.firstVisitTime) }}
</el-descriptions-item>
                    <el-descriptions-item label="最后就诊时间">
    {{ formatDate(detailFrom.lastVisitTime) }}
</el-descriptions-item>
                    <el-descriptions-item label="就诊次数">
    {{ detailFrom.visitCount }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createMtDoctorPatients,
  deleteMtDoctorPatients,
  deleteMtDoctorPatientsByIds,
  updateMtDoctorPatients,
  findMtDoctorPatients,
  getMtDoctorPatientsList
} from '@/api/medicine/mtDoctorPatients'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'MtDoctorPatients'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            doctorId: undefined,
            patientId: undefined,
            relationshipType: '',
            tags: '',
            notes: '',
            firstVisitTime: new Date(),
            lastVisitTime: new Date(),
            visitCount: undefined,
        })



// 验证规则
const rule = reactive({
               doctorId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               patientId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               firstVisitTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
  const table = await getMtDoctorPatientsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

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
            deleteMtDoctorPatientsFunc(row)
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
      const res = await deleteMtDoctorPatientsByIds({ IDs })
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
const updateMtDoctorPatientsFunc = async(row) => {
    const res = await findMtDoctorPatients({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMtDoctorPatientsFunc = async (row) => {
    const res = await deleteMtDoctorPatients({ ID: row.ID })
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
        doctorId: undefined,
        patientId: undefined,
        relationshipType: '',
        tags: '',
        notes: '',
        firstVisitTime: new Date(),
        lastVisitTime: new Date(),
        visitCount: undefined,
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
                  res = await createMtDoctorPatients(formData.value)
                  break
                case 'update':
                  res = await updateMtDoctorPatients(formData.value)
                  break
                default:
                  res = await createMtDoctorPatients(formData.value)
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
  const res = await findMtDoctorPatients({ ID: row.ID })
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