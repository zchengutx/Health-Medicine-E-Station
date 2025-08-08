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

type MtDrugApi struct {}



// CreateMtDrug 创建mtDrug表
// @Tags MtDrug
// @Summary 创建mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrug true "创建mtDrug表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtDrug/createMtDrug [post]
func (mtDrugApi *MtDrugApi) CreateMtDrug(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var mtDrug medicine.MtDrug
	err := c.ShouldBindJSON(&mtDrug)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtDrug.CreatedBy = utils.GetUserID(c)
	err = mtDrugService.CreateMtDrug(ctx,&mtDrug)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMtDrug 删除mtDrug表
// @Tags MtDrug
// @Summary 删除mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrug true "删除mtDrug表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtDrug/deleteMtDrug [delete]
func (mtDrugApi *MtDrugApi) DeleteMtDrug(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := mtDrugService.DeleteMtDrug(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtDrugByIds 批量删除mtDrug表
// @Tags MtDrug
// @Summary 批量删除mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtDrug/deleteMtDrugByIds [delete]
func (mtDrugApi *MtDrugApi) DeleteMtDrugByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := mtDrugService.DeleteMtDrugByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtDrug 更新mtDrug表
// @Tags MtDrug
// @Summary 更新mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrug true "更新mtDrug表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtDrug/updateMtDrug [put]
func (mtDrugApi *MtDrugApi) UpdateMtDrug(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var mtDrug medicine.MtDrug
	err := c.ShouldBindJSON(&mtDrug)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtDrug.UpdatedBy = utils.GetUserID(c)
	err = mtDrugService.UpdateMtDrug(ctx,mtDrug)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtDrug 用id查询mtDrug表
// @Tags MtDrug
// @Summary 用id查询mtDrug表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mtDrug表"
// @Success 200 {object} response.Response{data=medicine.MtDrug,msg=string} "查询成功"
// @Router /mtDrug/findMtDrug [get]
func (mtDrugApi *MtDrugApi) FindMtDrug(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remtDrug, err := mtDrugService.GetMtDrug(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remtDrug, c)
}
// GetMtDrugList 分页获取mtDrug表列表
// @Tags MtDrug
// @Summary 分页获取mtDrug表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDrugSearch true "分页获取mtDrug表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtDrug/getMtDrugList [get]
func (mtDrugApi *MtDrugApi) GetMtDrugList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo medicineReq.MtDrugSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtDrugService.GetMtDrugInfoList(ctx,pageInfo)
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

// GetMtDrugPublic 不需要鉴权的mtDrug表接口
// @Tags MtDrug
// @Summary 不需要鉴权的mtDrug表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtDrug/getMtDrugPublic [get]
func (mtDrugApi *MtDrugApi) GetMtDrugPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    mtDrugService.GetMtDrugPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的mtDrug表接口信息",
    }, "获取成功", c)
}
