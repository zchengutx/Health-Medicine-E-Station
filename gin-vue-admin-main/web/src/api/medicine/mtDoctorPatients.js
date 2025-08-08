import service from '@/utils/request'

// @Tags MtDoctorPatients
// @Summary 创建mtDoctorPatients表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDoctorPatients true "创建mtDoctorPatients表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtDoctorPatients/createMtDoctorPatients [post]
export const createMtDoctorPatients = (data) => {
  return service({
    url: '/mtDoctorPatients/createMtDoctorPatients',
    method: 'post',
    data
  })
}

// @Tags MtDoctorPatients
// @Summary 删除mtDoctorPatients表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDoctorPatients true "删除mtDoctorPatients表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDoctorPatients/deleteMtDoctorPatients [delete]
export const deleteMtDoctorPatients = (params) => {
  return service({
    url: '/mtDoctorPatients/deleteMtDoctorPatients',
    method: 'delete',
    params
  })
}

// @Tags MtDoctorPatients
// @Summary 批量删除mtDoctorPatients表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mtDoctorPatients表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDoctorPatients/deleteMtDoctorPatientsByIds [delete]
export const deleteMtDoctorPatientsByIds = (params) => {
  return service({
    url: '/mtDoctorPatients/deleteMtDoctorPatientsByIds',
    method: 'delete',
    params
  })
}

// @Tags MtDoctorPatients
// @Summary 更新mtDoctorPatients表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDoctorPatients true "更新mtDoctorPatients表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtDoctorPatients/updateMtDoctorPatients [put]
export const updateMtDoctorPatients = (data) => {
  return service({
    url: '/mtDoctorPatients/updateMtDoctorPatients',
    method: 'put',
    data
  })
}

// @Tags MtDoctorPatients
// @Summary 用id查询mtDoctorPatients表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MtDoctorPatients true "用id查询mtDoctorPatients表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtDoctorPatients/findMtDoctorPatients [get]
export const findMtDoctorPatients = (params) => {
  return service({
    url: '/mtDoctorPatients/findMtDoctorPatients',
    method: 'get',
    params
  })
}

// @Tags MtDoctorPatients
// @Summary 分页获取mtDoctorPatients表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtDoctorPatients表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDoctorPatients/getMtDoctorPatientsList [get]
export const getMtDoctorPatientsList = (params) => {
  return service({
    url: '/mtDoctorPatients/getMtDoctorPatientsList',
    method: 'get',
    params
  })
}