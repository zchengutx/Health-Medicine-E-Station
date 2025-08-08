package medicine

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
    medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type MtOrdersApi struct {}



// CreateMtOrders 创建mtOrders表
// @Tags MtOrders
// @Summary 创建mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtOrders true "创建mtOrders表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtOrders/createMtOrders [post]
func (mtOrdersApi *MtOrdersApi) CreateMtOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var mtOrders medicine.MtOrders
	err := c.ShouldBindJSON(&mtOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtOrders.CreatedBy = utils.GetUserID(c)
	err = mtOrdersService.CreateMtOrders(ctx,&mtOrders)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMtOrders 删除mtOrders表
// @Tags MtOrders
// @Summary 删除mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtOrders true "删除mtOrders表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtOrders/deleteMtOrders [delete]
func (mtOrdersApi *MtOrdersApi) DeleteMtOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := mtOrdersService.DeleteMtOrders(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtOrdersByIds 批量删除mtOrders表
// @Tags MtOrders
// @Summary 批量删除mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtOrders/deleteMtOrdersByIds [delete]
func (mtOrdersApi *MtOrdersApi) DeleteMtOrdersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := mtOrdersService.DeleteMtOrdersByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtOrders 更新mtOrders表
// @Tags MtOrders
// @Summary 更新mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtOrders true "更新mtOrders表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtOrders/updateMtOrders [put]
func (mtOrdersApi *MtOrdersApi) UpdateMtOrders(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var mtOrders medicine.MtOrders
	err := c.ShouldBindJSON(&mtOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtOrders.UpdatedBy = utils.GetUserID(c)
	err = mtOrdersService.UpdateMtOrders(ctx,mtOrders)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtOrders 用id查询mtOrders表
// @Tags MtOrders
// @Summary 用id查询mtOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mtOrders表"
// @Success 200 {object} response.Response{data=medicine.MtOrders,msg=string} "查询成功"
// @Router /mtOrders/findMtOrders [get]
func (mtOrdersApi *MtOrdersApi) FindMtOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remtOrders, err := mtOrdersService.GetMtOrders(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remtOrders, c)
}
// GetMtOrdersList 分页获取mtOrders表列表
// @Tags MtOrders
// @Summary 分页获取mtOrders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtOrdersSearch true "分页获取mtOrders表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtOrders/getMtOrdersList [get]
func (mtOrdersApi *MtOrdersApi) GetMtOrdersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo medicineReq.MtOrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtOrdersService.GetMtOrdersInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetMtOrdersPublic 不需要鉴权的mtOrders表接口
// @Tags MtOrders
// @Summary 不需要鉴权的mtOrders表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtOrders/getMtOrdersPublic [get]
func (mtOrdersApi *MtOrdersApi) GetMtOrdersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    mtOrdersService.GetMtOrdersPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的mtOrders表接口信息",
    }, "获取成功", c)
}
