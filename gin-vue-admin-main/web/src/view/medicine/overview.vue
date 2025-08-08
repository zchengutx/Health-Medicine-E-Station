<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="goToPage('mtDrug')">药品管理</el-button>
        <el-button type="primary" icon="plus" @click="goToPage('mtDrugCategory')">药品分类</el-button>
        <el-button type="primary" icon="plus" @click="goToPage('mtDrugTypeStair')">一级分类</el-button>
        <el-button type="primary" icon="plus" @click="goToPage('mtDrugTypeLevel')">二级分类</el-button>
      </div>
      <div class="gva-card-list">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card class="box-card" @click="goToPage('mtDrug')">
              <template #header>
                <div class="card-header">
                  <span>药品管理</span>
                </div>
              </template>
              <div class="card-content">
                <el-icon size="40" color="#409EFF"><MedicineBox /></el-icon>
                <p>管理药品信息</p>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="box-card" @click="goToPage('mtDrugCategory')">
              <template #header>
                <div class="card-header">
                  <span>药品分类</span>
                </div>
              </template>
              <div class="card-content">
                <el-icon size="40" color="#67C23A"><Folder /></el-icon>
                <p>管理药品分类</p>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="box-card" @click="goToPage('mtDrugTypeStair')">
              <template #header>
                <div class="card-header">
                  <span>一级分类</span>
                </div>
              </template>
              <div class="card-content">
                <el-icon size="40" color="#E6A23C"><FolderOpened /></el-icon>
                <p>管理一级分类</p>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="box-card" @click="goToPage('mtDrugTypeLevel')">
              <template #header>
                <div class="card-header">
                  <span>二级分类</span>
                </div>
              </template>
              <div class="card-content">
                <el-icon size="40" color="#F56C6C"><FolderOpened /></el-icon>
                <p>管理二级分类</p>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

defineOptions({
  name: 'MedicineOverview'
})

const router = useRouter()
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  // 可以添加搜索逻辑
}

// 跳转到指定页面
const goToPage = (pageName) => {
  router.push({ name: pageName })
}
</script>

<style scoped>
.gva-card-list {
  margin-top: 20px;
}

.box-card {
  cursor: pointer;
  transition: all 0.3s;
}

.box-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-content {
  text-align: center;
  padding: 20px 0;
}

.card-content p {
  margin-top: 10px;
  color: #666;
}
</style> 