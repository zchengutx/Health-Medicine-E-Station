package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MtDoctorPatientsApi struct{}

// CreateMtDoctorPatients 创建mtDoctorPatients表
// @Tags MtDoctorPatients
// @Summary 创建mtDoctorPatients表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctorPatients true "创建mtDoctorPatients表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtDoctorPatients/createMtDoctorPatients [post]
func (mtDoctorPatientsApi *MtDoctorPatientsApi) CreateMtDoctorPatients(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var mtDoctorPatients medicine.MtDoctorPatients
	err := c.ShouldBindJSON(&mtDoctorPatients)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mtDoctorPatients.CreatedBy = utils.GetUserID(c)
	err = mtDoctorPatientsService.CreateMtDoctorPatients(ctx, &mtDoctorPatients)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMtDoctorPatients 删除mtDoctorPatients表
// @Tags MtDoctorPatients
// @Summary 删除mtDoctorPatients表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctorPatients true "删除mtDoctorPatients表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtDoctorPatients/deleteMtDoctorPatients [delete]
func (mtDoctorPatientsApi *MtDoctorPatientsApi) DeleteMtDoctorPatients(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := mtDoctorPatientsService.DeleteMtDoctorPatients(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtDoctorPatientsByIds 批量删除mtDoctorPatients表
// @Tags MtDoctorPatients
// @Summary 批量删除mtDoctorPatients表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtDoctorPatients/deleteMtDoctorPatientsByIds [delete]
func (mtDoctorPatientsApi *MtDoctorPatientsApi) DeleteMtDoctorPatientsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := mtDoctorPatientsService.DeleteMtDoctorPatientsByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtDoctorPatients 更新mtDoctorPatients表
// @Tags MtDoctorPatients
// @Summary 更新mtDoctorPatients表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctorPatients true "更新mtDoctorPatients表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtDoctorPatients/updateMtDoctorPatients [put]
func (mtDoctorPatientsApi *MtDoctorPatientsApi) UpdateMtDoctorPatients(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var mtDoctorPatients medicine.MtDoctorPatients
	err := c.ShouldBindJSON(&mtDoctorPatients)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mtDoctorPatients.UpdatedBy = utils.GetUserID(c)
	err = mtDoctorPatientsService.UpdateMtDoctorPatients(ctx, mtDoctorPatients)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtDoctorPatients 用id查询mtDoctorPatients表
// @Tags MtDoctorPatients
// @Summary 用id查询mtDoctorPatients表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mtDoctorPatients表"
// @Success 200 {object} response.Response{data=medicine.MtDoctorPatients,msg=string} "查询成功"
// @Router /mtDoctorPatients/findMtDoctorPatients [get]
func (mtDoctorPatientsApi *MtDoctorPatientsApi) FindMtDoctorPatients(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	remtDoctorPatients, err := mtDoctorPatientsService.GetMtDoctorPatients(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remtDoctorPatients, c)
}

// GetMtDoctorPatientsList 分页获取mtDoctorPatients表列表
// @Tags MtDoctorPatients
// @Summary 分页获取mtDoctorPatients表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDoctorPatientsSearch true "分页获取mtDoctorPatients表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtDoctorPatients/getMtDoctorPatientsList [get]
func (mtDoctorPatientsApi *MtDoctorPatientsApi) GetMtDoctorPatientsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo medicineReq.MtDoctorPatientsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtDoctorPatientsService.GetMtDoctorPatientsInfoList(ctx, pageInfo)
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

// GetMtDoctorPatientsPublic 不需要鉴权的mtDoctorPatients表接口
// @Tags MtDoctorPatients
// @Summary 不需要鉴权的mtDoctorPatients表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtDoctorPatients/getMtDoctorPatientsPublic [get]
func (mtDoctorPatientsApi *MtDoctorPatientsApi) GetMtDoctorPatientsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	mtDoctorPatientsService.GetMtDoctorPatientsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的mtDoctorPatients表接口信息",
	}, "获取成功", c)
}
