import service from '@/utils/request'
// @Tags MtDiscount
// @Summary 创建mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDiscount true "创建mtDiscount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtDiscount/createMtDiscount [post]
export const createMtDiscount = (data) => {
  return service({
    url: '/mtDiscount/createMtDiscount',
    method: 'post',
    data
  })
}

// @Tags MtDiscount
// @Summary 删除mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDiscount true "删除mtDiscount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDiscount/deleteMtDiscount [delete]
export const deleteMtDiscount = (params) => {
  return service({
    url: '/mtDiscount/deleteMtDiscount',
    method: 'delete',
    params
  })
}

// @Tags MtDiscount
// @Summary 批量删除mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mtDiscount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDiscount/deleteMtDiscount [delete]
export const deleteMtDiscountByIds = (params) => {
  return service({
    url: '/mtDiscount/deleteMtDiscountByIds',
    method: 'delete',
    params
  })
}

// @Tags MtDiscount
// @Summary 更新mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MtDiscount true "更新mtDiscount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtDiscount/updateMtDiscount [put]
export const updateMtDiscount = (data) => {
  return service({
    url: '/mtDiscount/updateMtDiscount',
    method: 'put',
    data
  })
}

// @Tags MtDiscount
// @Summary 用id查询mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MtDiscount true "用id查询mtDiscount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtDiscount/findMtDiscount [get]
export const findMtDiscount = (params) => {
  return service({
    url: '/mtDiscount/findMtDiscount',
    method: 'get',
    params
  })
}

// @Tags MtDiscount
// @Summary 分页获取mtDiscount表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mtDiscount表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDiscount/getMtDiscountList [get]
export const getMtDiscountList = (params) => {
  return service({
    url: '/mtDiscount/getMtDiscountList',
    method: 'get',
    params
  })
}

// @Tags MtDiscount
// @Summary 不需要鉴权的mtDiscount表接口
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDiscountSearch true "分页获取mtDiscount表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtDiscount/getMtDiscountPublic [get]
export const getMtDiscountPublic = () => {
  return service({
    url: '/mtDiscount/getMtDiscountPublic',
    method: 'get',
  })
}
