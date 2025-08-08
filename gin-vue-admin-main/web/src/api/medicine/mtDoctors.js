import service from '@/utils/request'
// @Tags MtDoctors
// @Summary 创建mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDoctors true "创建mtDoctors表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtDoctors/createMtDoctors [post]
export const createMtDoctors = (data) => {
  return service({
    url: '/mtDoctors/createMtDoctors',
    method: 'post',
    data
  })
}

// @Tags MtDoctors
// @Summary 删除mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDoctors true "删除mtDoctors表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDoctors/deleteMtDoctors [delete]
export const deleteMtDoctors = (params) => {
  return service({
    url: '/mtDoctors/deleteMtDoctors',
    method: 'delete',
    params
  })
}

// @Tags MtDoctors
// @Summary 批量删除mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mtDoctors表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDoctors/deleteMtDoctors [delete]
export const deleteMtDoctorsByIds = (params) => {
  return service({
    url: '/mtDoctors/deleteMtDoctorsByIds',
    method: 'delete',
    params
  })
}

// @Tags MtDoctors
// @Summary 更新mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDoctors true "更新mtDoctors表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtDoctors/updateMtDoctors [put]
export const updateMtDoctors = (data) => {
  return service({
    url: '/mtDoctors/updateMtDoctors',
    method: 'put',
    data
  })
}

// @Tags MtDoctors
// @Summary 用id查询mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MtDoctors true "用id查询mtDoctors表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtDoctors/findMtDoctors [get]
export const findMtDoctors = (params) => {
  return service({
    url: '/mtDoctors/findMtDoctors',
    method: 'get',
    params
  })
}

// @Tags MtDoctors
// @Summary 分页获取mtDoctors表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtDoctors表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDoctors/getMtDoctorsList [get]
export const getMtDoctorsList = (params) => {
  return service({
    url: '/mtDoctors/getMtDoctorsList',
    method: 'get',
    params
  })
}

// @Tags MtDoctors
// @Summary 不需要鉴权的mtDoctors表接口
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDoctorsSearch true "分页获取mtDoctors表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtDoctors/getMtDoctorsPublic [get]
export const getMtDoctorsPublic = () => {
  return service({
    url: '/mtDoctors/getMtDoctorsPublic',
    method: 'get',
  })
}
