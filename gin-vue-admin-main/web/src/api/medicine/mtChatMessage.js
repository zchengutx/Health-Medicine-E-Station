import service from '@/utils/request'
// @Tags MtChatMessage
// @Summary 创建mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtChatMessage true "创建mtChatMessage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtChatMessage/createMtChatMessage [post]
export const createMtChatMessage = (data) => {
  return service({
    url: '/mtChatMessage/createMtChatMessage',
    method: 'post',
    data
  })
}

// @Tags MtChatMessage
// @Summary 删除mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtChatMessage true "删除mtChatMessage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtChatMessage/deleteMtChatMessage [delete]
export const deleteMtChatMessage = (params) => {
  return service({
    url: '/mtChatMessage/deleteMtChatMessage',
    method: 'delete',
    params
  })
}

// @Tags MtChatMessage
// @Summary 批量删除mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mtChatMessage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtChatMessage/deleteMtChatMessage [delete]
export const deleteMtChatMessageByIds = (params) => {
  return service({
    url: '/mtChatMessage/deleteMtChatMessageByIds',
    method: 'delete',
    params
  })
}

// @Tags MtChatMessage
// @Summary 更新mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtChatMessage true "更新mtChatMessage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtChatMessage/updateMtChatMessage [put]
export const updateMtChatMessage = (data) => {
  return service({
    url: '/mtChatMessage/updateMtChatMessage',
    method: 'put',
    data
  })
}

// @Tags MtChatMessage
// @Summary 用id查询mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MtChatMessage true "用id查询mtChatMessage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtChatMessage/findMtChatMessage [get]
export const findMtChatMessage = (params) => {
  return service({
    url: '/mtChatMessage/findMtChatMessage',
    method: 'get',
    params
  })
}

// @Tags MtChatMessage
// @Summary 分页获取mtChatMessage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtChatMessage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtChatMessage/getMtChatMessageList [get]
export const getMtChatMessageList = (params) => {
  return service({
    url: '/mtChatMessage/getMtChatMessageList',
    method: 'get',
    params
  })
}

// @Tags MtChatMessage
// @Summary 不需要鉴权的mtChatMessage表接口
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtChatMessageSearch true "分页获取mtChatMessage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtChatMessage/getMtChatMessagePublic [get]
export const getMtChatMessagePublic = () => {
  return service({
    url: '/mtChatMessage/getMtChatMessagePublic',
    method: 'get',
  })
}
