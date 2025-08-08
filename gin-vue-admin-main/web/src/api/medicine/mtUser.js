import service from '@/utils/request'
// @Tags MtUser
// @Summary 创建mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtUser true "创建mtUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtUser/createMtUser [post]
export const createMtUser = (data) => {
  return service({
    url: '/mtUser/createMtUser',
    method: 'post',
    data
  })
}

// @Tags MtUser
// @Summary 删除mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtUser true "删除mtUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtUser/deleteMtUser [delete]
export const deleteMtUser = (params) => {
  return service({
    url: '/mtUser/deleteMtUser',
    method: 'delete',
    params
  })
}

// @Tags MtUser
// @Summary 批量删除mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mtUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtUser/deleteMtUser [delete]
export const deleteMtUserByIds = (params) => {
  return service({
    url: '/mtUser/deleteMtUserByIds',
    method: 'delete',
    params
  })
}

// @Tags MtUser
// @Summary 更新mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtUser true "更新mtUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtUser/updateMtUser [put]
export const updateMtUser = (data) => {
  return service({
    url: '/mtUser/updateMtUser',
    method: 'put',
    data
  })
}

// @Tags MtUser
// @Summary 用id查询mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MtUser true "用id查询mtUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtUser/findMtUser [get]
export const findMtUser = (params) => {
  return service({
    url: '/mtUser/findMtUser',
    method: 'get',
    params
  })
}

// @Tags MtUser
// @Summary 分页获取mtUser表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtUser/getMtUserList [get]
export const getMtUserList = (params) => {
  return service({
    url: '/mtUser/getMtUserList',
    method: 'get',
    params
  })
}

// @Tags MtUser
// @Summary 不需要鉴权的mtUser表接口
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtUserSearch true "分页获取mtUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtUser/getMtUserPublic [get]
export const getMtUserPublic = () => {
  return service({
    url: '/mtUser/getMtUserPublic',
    method: 'get',
  })
}
