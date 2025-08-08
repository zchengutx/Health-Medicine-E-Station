import service from '@/utils/request'
// @Tags MtOrders
// @Summary 创建mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtOrders true "创建mtOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtOrders/createMtOrders [post]
export const createMtOrders = (data) => {
  return service({
    url: '/mtOrders/createMtOrders',
    method: 'post',
    data
  })
}

// @Tags MtOrders
// @Summary 删除mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtOrders true "删除mtOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtOrders/deleteMtOrders [delete]
export const deleteMtOrders = (params) => {
  return service({
    url: '/mtOrders/deleteMtOrders',
    method: 'delete',
    params
  })
}

// @Tags MtOrders
// @Summary 批量删除mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mtOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtOrders/deleteMtOrders [delete]
export const deleteMtOrdersByIds = (params) => {
  return service({
    url: '/mtOrders/deleteMtOrdersByIds',
    method: 'delete',
    params
  })
}

// @Tags MtOrders
// @Summary 更新mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtOrders true "更新mtOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtOrders/updateMtOrders [put]
export const updateMtOrders = (data) => {
  return service({
    url: '/mtOrders/updateMtOrders',
    method: 'put',
    data
  })
}

// @Tags MtOrders
// @Summary 用id查询mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MtOrders true "用id查询mtOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtOrders/findMtOrders [get]
export const findMtOrders = (params) => {
  return service({
    url: '/mtOrders/findMtOrders',
    method: 'get',
    params
  })
}

// @Tags MtOrders
// @Summary 分页获取mtOrders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtOrders表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtOrders/getMtOrdersList [get]
export const getMtOrdersList = (params) => {
  return service({
    url: '/mtOrders/getMtOrdersList',
    method: 'get',
    params
  })
}

// @Tags MtOrders
// @Summary 不需要鉴权的mtOrders表接口
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtOrdersSearch true "分页获取mtOrders表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtOrders/getMtOrdersPublic [get]
export const getMtOrdersPublic = () => {
  return service({
    url: '/mtOrders/getMtOrdersPublic',
    method: 'get',
  })
}
