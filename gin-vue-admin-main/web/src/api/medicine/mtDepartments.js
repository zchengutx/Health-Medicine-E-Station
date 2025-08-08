import service from '@/utils/request'

// @Tags MtDepartments
// @Summary 创建科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDepartments true "创建科室"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtDepartments/createMtDepartments [post]
export const createMtDepartments = (data) => {
  return service({
    url: '/mtDepartments/createMtDepartments',
    method: 'post',
    data
  })
}

// @Tags MtDepartments
// @Summary 删除科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDepartments true "删除科室"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDepartments/deleteMtDepartments [delete]
export const deleteMtDepartments = (params) => {
  return service({
    url: '/mtDepartments/deleteMtDepartments',
    method: 'delete',
    params
  })
}

// @Tags MtDepartments
// @Summary 批量删除科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDepartments/deleteMtDepartments [delete]
export const deleteMtDepartmentsByIds = (params) => {
  return service({
    url: '/mtDepartments/deleteMtDepartmentsByIds',
    method: 'delete',
    params
  })
}

// @Tags MtDepartments
// @Summary 更新科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MtDepartments true "更新科室"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtDepartments/updateMtDepartments [put]
export const updateMtDepartments = (data) => {
  return service({
    url: '/mtDepartments/updateMtDepartments',
    method: 'put',
    data
  })
}

// @Tags MtDepartments
// @Summary 用id查询科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MtDepartments true "用id查询科室"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtDepartments/findMtDepartments [get]
export const findMtDepartments = (params) => {
  return service({
    url: '/mtDepartments/findMtDepartments',
    method: 'get',
    params
  })
}

// @Tags MtDepartments
// @Summary 分页获取科室列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取科室列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDepartments/getMtDepartmentsList [get]
export const getMtDepartmentsList = (params) => {
  return service({
    url: '/mtDepartments/getMtDepartmentsList',
    method: 'get',
    params
  })
}

// @Tags MtDepartments
// @Summary 获取科室列表（公开接口）
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDepartments/getMtDepartmentsListPublic [get]
export const getMtDepartmentsListPublic = (params) => {
  return service({
    url: '/mtDepartments/getMtDepartmentsListPublic',
    method: 'get',
    params
  })
} 