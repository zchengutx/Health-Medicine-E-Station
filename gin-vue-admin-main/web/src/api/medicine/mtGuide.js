import service from '@/utils/request'

// 获取用药指导详情
export const getMtGuide = (params) => {
  return service({
    url: '/mtGuide/findMtGuide',
    method: 'get',
    params
  })
}

// 获取用药指导列表
export const getMtGuideList = (params) => {
  return service({
    url: '/mtGuide/getMtGuideList',
    method: 'get',
    params
  })
}