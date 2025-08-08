import service from '@/utils/request'

// @Tags MtHospitals
// @Summary 创建医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtHospitals true "创建医院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtHospitals/createMtHospitals [post]
export const createMtHospitals = (data) => {
  return service({
    url: '/mtHospitals/createMtHospitals',
    method: 'post',
    data
  })
}

// @Tags MtHospitals
// @Summary 删除医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtHospitals true "删除医院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtHospitals/deleteMtHospitals [delete]
export const deleteMtHospitals = (params) => {
  return service({
    url: '/mtHospitals/deleteMtHospitals',
    method: 'delete',
    params
  })
}

// @Tags MtHospitals
// @Summary 批量删除医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtHospitals/deleteMtHospitals [delete]
export const deleteMtHospitalsByIds = (params) => {
  return service({
    url: '/mtHospitals/deleteMtHospitalsByIds',
    method: 'delete',
    params
  })
}

// @Tags MtHospitals
// @Summary 更新医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtHospitals true "更新医院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtHospitals/updateMtHospitals [put]
export const updateMtHospitals = (data) => {
  return service({
    url: '/mtHospitals/updateMtHospitals',
    method: 'put',
    data
  })
}

// @Tags MtHospitals
// @Summary 用id查询医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MtHospitals true "用id查询医院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtHospitals/findMtHospitals [get]
export const findMtHospitals = (params) => {
  return service({
    url: '/mtHospitals/findMtHospitals',
    method: 'get',
    params
  })
}

// @Tags MtHospitals
// @Summary 分页获取医院列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取医院列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtHospitals/getMtHospitalsList [get]
export const getMtHospitalsList = (params) => {
  return service({
    url: '/mtHospitals/getMtHospitalsList',
    method: 'get',
    params
  })
}

// @Tags MtHospitals
// @Summary 获取医院列表（公开接口）
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtHospitals/getMtHospitalsListPublic [get]
export const getMtHospitalsListPublic = (params) => {
  return service({
    url: '/mtHospitals/getMtHospitalsListPublic',
    method: 'get',
    params
  })
} 