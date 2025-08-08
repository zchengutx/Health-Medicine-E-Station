import service from '@/utils/request'
// @Tags MtOrdersDrug
// @Summary 创建mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtOrdersDrug true "创建mtOrdersDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtOrdersDrug/createMtOrdersDrug [post]
export const createMtOrdersDrug = (data) => {
  return service({
    url: '/mtOrdersDrug/createMtOrdersDrug',
    method: 'post',
    data
  })
}

// @Tags MtOrdersDrug
// @Summary 删除mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtOrdersDrug true "删除mtOrdersDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtOrdersDrug/deleteMtOrdersDrug [delete]
export const deleteMtOrdersDrug = (params) => {
  return service({
    url: '/mtOrdersDrug/deleteMtOrdersDrug',
    method: 'delete',
    params
  })
}

// @Tags MtOrdersDrug
// @Summary 批量删除mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mtOrdersDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtOrdersDrug/deleteMtOrdersDrug [delete]
export const deleteMtOrdersDrugByIds = (params) => {
  return service({
    url: '/mtOrdersDrug/deleteMtOrdersDrugByIds',
    method: 'delete',
    params
  })
}

// @Tags MtOrdersDrug
// @Summary 更新mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtOrdersDrug true "更新mtOrdersDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtOrdersDrug/updateMtOrdersDrug [put]
export const updateMtOrdersDrug = (data) => {
  return service({
    url: '/mtOrdersDrug/updateMtOrdersDrug',
    method: 'put',
    data
  })
}

// @Tags MtOrdersDrug
// @Summary 用id查询mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MtOrdersDrug true "用id查询mtOrdersDrug表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtOrdersDrug/findMtOrdersDrug [get]
export const findMtOrdersDrug = (params) => {
  return service({
    url: '/mtOrdersDrug/findMtOrdersDrug',
    method: 'get',
    params
  })
}

// @Tags MtOrdersDrug
// @Summary 分页获取mtOrdersDrug表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtOrdersDrug表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtOrdersDrug/getMtOrdersDrugList [get]
export const getMtOrdersDrugList = (params) => {
  return service({
    url: '/mtOrdersDrug/getMtOrdersDrugList',
    method: 'get',
    params
  })
}

// @Tags MtOrdersDrug
// @Summary 不需要鉴权的mtOrdersDrug表接口
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtOrdersDrugSearch true "分页获取mtOrdersDrug表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtOrdersDrug/getMtOrdersDrugPublic [get]
export const getMtOrdersDrugPublic = () => {
  return service({
    url: '/mtOrdersDrug/getMtOrdersDrugPublic',
    method: 'get',
  })
}

// @Tags MtOrdersDrug
// @Summary 获取订单详情信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "订单详情ID"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtOrdersDrug/getMtOrdersDrugDetail [get]
export const getMtOrdersDrugDetail = (params) => {
  return service({
    url: '/mtOrdersDrug/getMtOrdersDrugDetail',
    method: 'get',
    params
  })
}
