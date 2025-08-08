import service from '@/utils/request'

// @Tags MtDoctorApproval
// @Summary 创建mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDoctorApproval true "创建mtDoctorApproval表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtDoctorApproval/createMtDoctorApproval [post]
export const createMtDoctorApproval = (data) => {
  return service({
    url: '/mtDoctorApproval/createMtDoctorApproval',
    method: 'post',
    data
  })
}

// @Tags MtDoctorApproval
// @Summary 删除mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDoctorApproval true "删除mtDoctorApproval表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDoctorApproval/deleteMtDoctorApproval [delete]
export const deleteMtDoctorApproval = (params) => {
  return service({
    url: '/mtDoctorApproval/deleteMtDoctorApproval',
    method: 'delete',
    params
  })
}

// @Tags MtDoctorApproval
// @Summary 批量删除mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDoctorApproval/deleteMtDoctorApproval [delete]
export const deleteMtDoctorApprovalByIds = (params) => {
  return service({
    url: '/mtDoctorApproval/deleteMtDoctorApprovalByIds',
    method: 'delete',
    params
  })
}

// @Tags MtDoctorApproval
// @Summary 更新mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDoctorApproval true "更新mtDoctorApproval表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtDoctorApproval/updateMtDoctorApproval [put]
export const updateMtDoctorApproval = (data) => {
  return service({
    url: '/mtDoctorApproval/updateMtDoctorApproval',
    method: 'put',
    data
  })
}

// @Tags MtDoctorApproval
// @Summary 用id查询mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MtDoctorApproval true "用id查询mtDoctorApproval表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtDoctorApproval/findMtDoctorApproval [get]
export const findMtDoctorApproval = (params) => {
  return service({
    url: '/mtDoctorApproval/findMtDoctorApproval',
    method: 'get',
    params
  })
}

// @Tags MtDoctorApproval
// @Summary 分页获取mtDoctorApproval表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtDoctorApproval表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDoctorApproval/getMtDoctorApprovalList [get]
export const getMtDoctorApprovalList = (params) => {
  return service({
    url: '/mtDoctorApproval/getMtDoctorApprovalList',
    method: 'get',
    params
  })
}

// @Tags MtDoctorApproval
// @Summary 根据医生ID获取审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param doctorId query uint true "医生ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDoctorApproval/getMtDoctorApprovalByDoctorId [get]
export const getMtDoctorApprovalByDoctorId = (params) => {
  return service({
    url: '/mtDoctorApproval/getMtDoctorApprovalByDoctorId',
    method: 'get',
    params
  })
} 