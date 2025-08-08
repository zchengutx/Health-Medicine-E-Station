import service from '@/utils/request'

// 获取所有一级分类
export const getAllMtDrugTypeStair = () => {
  return service({
    url: '/mtDrugTypeStair/getMtDrugTypeStairPublic',
    method: 'get'
  })
}

// 分页获取一级分类列表
export const getMtDrugTypeStairList = (params) => {
  return service({
    url: '/mtDrugTypeStair/getMtDrugTypeStairList',
    method: 'get',
    params
  })
}

// 创建一级分类
export const createMtDrugTypeStair = (data) => {
  return service({
    url: '/mtDrugTypeStair/createMtDrugTypeStair',
    method: 'post',
    data
  })
}

// 更新一级分类
export const updateMtDrugTypeStair = (data) => {
  return service({
    url: '/mtDrugTypeStair/updateMtDrugTypeStair',
    method: 'put',
    data
  })
}

// 删除一级分类
export const deleteMtDrugTypeStair = (params) => {
  return service({
    url: '/mtDrugTypeStair/deleteMtDrugTypeStair',
    method: 'delete',
    params
  })
}

// 批量删除一级分类
export const deleteMtDrugTypeStairByIds = (params) => {
  return service({
    url: '/mtDrugTypeStair/deleteMtDrugTypeStairByIds',
    method: 'delete',
    params
  })
}

// 根据ID查询一级分类
export const findMtDrugTypeStair = (params) => {
  return service({
    url: '/mtDrugTypeStair/findMtDrugTypeStair',
    method: 'get',
    params
  })
} 