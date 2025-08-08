import service from '@/utils/request'

// 获取说明书详情
export const getMtExplain = (params) => {
  return service({
    url: '/mtExplain/findMtExplain',
    method: 'get',
    params
  })
}

// 获取说明书列表
export const getMtExplainList = (params) => {
  return service({
    url: '/mtExplain/getMtExplainList',
    method: 'get',
    params
  })
}