import service from '@/utils/request'
// @Tags MtDrug
// @Summary 创建mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDrug true "创建mtDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtDrug/createMtDrug [post]
export const createMtDrug = (data) => {
  return service({
    url: '/mtDrug/createMtDrug',
    method: 'post',
    data
  })
}

// @Tags MtDrug
// @Summary 删除mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDrug true "删除mtDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDrug/deleteMtDrug [delete]
export const deleteMtDrug = (params) => {
  return service({
    url: '/mtDrug/deleteMtDrug',
    method: 'delete',
    params
  })
}

// @Tags MtDrug
// @Summary 批量删除mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mtDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDrug/deleteMtDrug [delete]
export const deleteMtDrugByIds = (params) => {
  return service({
    url: '/mtDrug/deleteMtDrugByIds',
    method: 'delete',
    params
  })
}

// @Tags MtDrug
// @Summary 更新mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDrug true "更新mtDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtDrug/updateMtDrug [put]
export const updateMtDrug = (data) => {
  return service({
    url: '/mtDrug/updateMtDrug',
    method: 'put',
    data
  })
}

// @Tags MtDrug
// @Summary 用id查询mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MtDrug true "用id查询mtDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtDrug/findMtDrug [get]
export const findMtDrug = (params) => {
  return service({
    url: '/mtDrug/findMtDrug',
    method: 'get',
    params
  })
}

// @Tags MtDrug
// @Summary 分页获取mtDrug表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtDrug表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDrug/getMtDrugList [get]
export const getMtDrugList = (params) => {
  return service({
    url: '/mtDrug/getMtDrugList',
    method: 'get',
    params
  })
}

// @Tags MtDrug
// @Summary 不需要鉴权的mtDrug表接口
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDrugSearch true "分页获取mtDrug表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtDrug/getMtDrugPublic [get]
export const getMtDrugPublic = () => {
  return service({
    url: '/mtDrug/getMtDrugPublic',
    method: 'get',
  })
}
