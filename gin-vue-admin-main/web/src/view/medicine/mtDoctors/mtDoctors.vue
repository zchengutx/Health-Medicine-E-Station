
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
        
            <el-table-column align="left" label="医生编码" prop="doctorCode" width="120" />

            <el-table-column align="left" label="医生姓名" prop="name" width="120" />

            <el-table-column align="left" label="性别" prop="gender" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.gender,genderOptions) }}
    </template>
</el-table-column>
          <el-table-column label="科室名称" prop="department" />
          
          <el-table-column label="医院名称" prop="hospital" />

            <el-table-column align="left" label="职称" prop="title" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.title,job_titleOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="审核状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,doctor_statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="服务审核" prop="serviceAudit" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.serviceAudit,service_auditOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateMtDoctorsFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="科室ID:" prop="departmentId">
    <el-input v-model.number="formData.departmentId" :clearable="true" placeholder="请输入科室ID" />
</el-form-item>
            <el-form-item label="医院ID:" prop="hospitalId">
    <el-input v-model.number="formData.hospitalId" :clearable="true" placeholder="请输入医院ID" />
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="医生编码">
    {{ detailFrom.doctorCode }}
</el-descriptions-item>
                    <el-descriptions-item label="医生姓名">
    {{ detailFrom.name }}
</el-descriptions-item>
                    <el-descriptions-item label="性别：1-男，2-女">
    {{ detailFrom.gender }}
</el-descriptions-item>
                    <el-descriptions-item label="科室ID">
    {{ detailFrom.departmentId }}
</el-descriptions-item>
                    <el-descriptions-item label="医院ID">
    {{ detailFrom.hospitalId }}
</el-descriptions-item>
                    <el-descriptions-item label="职称">
    {{ detailFrom.title }}
</el-descriptions-item>
                    <el-descriptions-item label="审核状态：0-未通过，1-已通过，2-未审核">
    {{ detailFrom.status }}
</el-descriptions-item>
                    <el-descriptions-item label="服务审核：0-未审核，1-已审核，2-待审核">
    {{ detailFrom.serviceAudit }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createMtDoctors,
  deleteMtDoctors,
  deleteMtDoctorsByIds,
  updateMtDoctors,
  findMtDoctors,
  getMtDoctorsList
} from '@/api/medicine/mtDoctors'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'MtDoctors'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const doctor_statusOptions = ref([])
const service_auditOptions = ref([])
const genderOptions = ref([])
const job_titleOptions = ref([])
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               name : [{
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
               gender : [{
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
               departmentId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               hospitalId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               title : [{
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
               serviceAudit : [{
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
  const table = await getMtDoctorsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    doctor_statusOptions.value = await getDictFunc('doctor_status')
    service_auditOptions.value = await getDictFunc('service_audit')
    genderOptions.value = await getDictFunc('gender')
    job_titleOptions.value = await getDictFunc('job_title')
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
            deleteMtDoctorsFunc(row)
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
      const res = await deleteMtDoctorsByIds({ IDs })
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
const updateMtDoctorsFunc = async(row) => {
    const res = await findMtDoctors({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMtDoctorsFunc = async (row) => {
    const res = await deleteMtDoctors({ ID: row.ID })
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
        doctorCode: '',
        name: '',
        gender: '',
        departmentId: undefined,
        hospitalId: undefined,
        title: '',
        status: '',
        serviceAudit: '',
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
  const res = await findMtDoctors({ ID: row.ID })
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
