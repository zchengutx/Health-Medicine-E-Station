import service from '@/utils/request'

// 获取所有二级分类
export const getAllMtDrugTypeLevel = () => {
  return service({
    url: '/mtDrugTypeLevel/getMtDrugTypeLevelPublic',
    method: 'get'
  })
}

// 分页获取二级分类列表
export const getMtDrugTypeLevelList = (params) => {
  return service({
    url: '/mtDrugTypeLevel/getMtDrugTypeLevelList',
    method: 'get',
    params
  })
}

// 创建二级分类
export const createMtDrugTypeLevel = (data) => {
  return service({
    url: '/mtDrugTypeLevel/createMtDrugTypeLevel',
    method: 'post',
    data
  })
}

// 更新二级分类
export const updateMtDrugTypeLevel = (data) => {
  return service({
    url: '/mtDrugTypeLevel/updateMtDrugTypeLevel',
    method: 'put',
    data
  })
}

// 删除二级分类
export const deleteMtDrugTypeLevel = (params) => {
  return service({
    url: '/mtDrugTypeLevel/deleteMtDrugTypeLevel',
    method: 'delete',
    params
  })
}

// 批量删除二级分类
export const deleteMtDrugTypeLevelByIds = (params) => {
  return service({
    url: '/mtDrugTypeLevel/deleteMtDrugTypeLevelByIds',
    method: 'delete',
    params
  })
}

// 根据ID查询二级分类
export const findMtDrugTypeLevel = (params) => {
  return service({
    url: '/mtDrugTypeLevel/findMtDrugTypeLevel',
    method: 'get',
    params
  })
} 