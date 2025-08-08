package medicine

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MtOrdersDrugApi struct{}

// CreateMtOrdersDrug 创建mtOrdersDrug表
// @Tags MtOrdersDrug
// @Summary 创建mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtOrdersDrug true "创建mtOrdersDrug表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtOrdersDrug/createMtOrdersDrug [post]
func (mtOrdersDrugApi *MtOrdersDrugApi) CreateMtOrdersDrug(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var mtOrdersDrug medicine.MtOrdersDrug
	err := c.ShouldBindJSON(&mtOrdersDrug)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mtOrdersDrug.CreatedBy = utils.GetUserID(c)
	err = mtOrdersDrugService.CreateMtOrdersDrug(ctx, &mtOrdersDrug)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMtOrdersDrug 删除mtOrdersDrug表
// @Tags MtOrdersDrug
// @Summary 删除mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtOrdersDrug true "删除mtOrdersDrug表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtOrdersDrug/deleteMtOrdersDrug [delete]
func (mtOrdersDrugApi *MtOrdersDrugApi) DeleteMtOrdersDrug(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := mtOrdersDrugService.DeleteMtOrdersDrug(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtOrdersDrugByIds 批量删除mtOrdersDrug表
// @Tags MtOrdersDrug
// @Summary 批量删除mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtOrdersDrug/deleteMtOrdersDrugByIds [delete]
func (mtOrdersDrugApi *MtOrdersDrugApi) DeleteMtOrdersDrugByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := mtOrdersDrugService.DeleteMtOrdersDrugByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtOrdersDrug 更新mtOrdersDrug表
// @Tags MtOrdersDrug
// @Summary 更新mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtOrdersDrug true "更新mtOrdersDrug表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtOrdersDrug/updateMtOrdersDrug [put]
func (mtOrdersDrugApi *MtOrdersDrugApi) UpdateMtOrdersDrug(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var mtOrdersDrug medicine.MtOrdersDrug
	err := c.ShouldBindJSON(&mtOrdersDrug)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mtOrdersDrug.UpdatedBy = utils.GetUserID(c)
	err = mtOrdersDrugService.UpdateMtOrdersDrug(ctx, mtOrdersDrug)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtOrdersDrug 用id查询mtOrdersDrug表
// @Tags MtOrdersDrug
// @Summary 用id查询mtOrdersDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mtOrdersDrug表"
// @Success 200 {object} response.Response{data=medicine.MtOrdersDrug,msg=string} "查询成功"
// @Router /mtOrdersDrug/findMtOrdersDrug [get]
func (mtOrdersDrugApi *MtOrdersDrugApi) FindMtOrdersDrug(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	remtOrdersDrug, err := mtOrdersDrugService.GetMtOrdersDrug(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remtOrdersDrug, c)
}

// GetMtOrdersDrugList 分页获取mtOrdersDrug表列表
// @Tags MtOrdersDrug
// @Summary 分页获取mtOrdersDrug表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtOrdersDrugSearch true "分页获取mtOrdersDrug表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtOrdersDrug/getMtOrdersDrugList [get]
func (mtOrdersDrugApi *MtOrdersDrugApi) GetMtOrdersDrugList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo medicineReq.MtOrdersDrugSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtOrdersDrugService.GetMtOrdersDrugInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetMtOrdersDrugPublic 不需要鉴权的mtOrdersDrug表接口
// @Tags MtOrdersDrug
// @Summary 不需要鉴权的mtOrdersDrug表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtOrdersDrug/getMtOrdersDrugPublic [get]
func (mtOrdersDrugApi *MtOrdersDrugApi) GetMtOrdersDrugPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	mtOrdersDrugService.GetMtOrdersDrugPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的mtOrdersDrug表接口信息",
	}, "获取成功", c)
}

// GetMtOrdersDrugDetail 获取订单详情信息
// @Tags MtOrdersDrug
// @Summary 获取订单详情信息
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "订单详情ID"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtOrdersDrug/getMtOrdersDrugDetail [get]
func (mtOrdersDrugApi *MtOrdersDrugApi) GetMtOrdersDrugDetail(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	detail, err := mtOrdersDrugService.GetMtOrdersDrugDetail(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("获取订单详情失败!", zap.Error(err))
		response.FailWithMessage("获取订单详情失败:"+err.Error(), c)
		return
	}
	response.OkWithData(detail, c)
}

// TestMtOrdersDrugDetail 测试接口
// @Tags MtOrdersDrug
// @Summary 测试接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtOrdersDrug/testMtOrdersDrugDetail [get]
func (mtOrdersDrugApi *MtOrdersDrugApi) TestMtOrdersDrugDetail(c *gin.Context) {
	response.OkWithData(gin.H{
		"message":   "接口正常工作",
		"timestamp": time.Now().Unix(),
	}, c)
}
