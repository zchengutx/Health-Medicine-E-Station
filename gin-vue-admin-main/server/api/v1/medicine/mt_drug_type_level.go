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

type MtDrugTypeLevelApi struct{}

// CreateMtDrugTypeLevel 创建药品类型级别表
// @Tags MtDrugTypeLevel
// @Summary 创建药品类型级别表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrugTypeLevel true "创建药品类型级别表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtDrugTypeLevel/createMtDrugTypeLevel [post]
func (mtDrugTypeLevelApi *MtDrugTypeLevelApi) CreateMtDrugTypeLevel(c *gin.Context) {
	ctx := c.Request.Context()

	var mtDrugTypeLevel medicine.MtDrugTypeLevel
	err := c.ShouldBindJSON(&mtDrugTypeLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mtDrugTypeLevel.CreatedBy = utils.GetUserID(c)
	err = mtDrugTypeLevelService.CreateMtDrugTypeLevel(ctx, &mtDrugTypeLevel)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMtDrugTypeLevel 删除药品类型级别表
// @Tags MtDrugTypeLevel
// @Summary 删除药品类型级别表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrugTypeLevel true "删除药品类型级别表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtDrugTypeLevel/deleteMtDrugTypeLevel [delete]
func (mtDrugTypeLevelApi *MtDrugTypeLevelApi) DeleteMtDrugTypeLevel(c *gin.Context) {
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := mtDrugTypeLevelService.DeleteMtDrugTypeLevel(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtDrugTypeLevelByIds 批量删除药品类型级别表
// @Tags MtDrugTypeLevel
// @Summary 批量删除药品类型级别表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtDrugTypeLevel/deleteMtDrugTypeLevelByIds [delete]
func (mtDrugTypeLevelApi *MtDrugTypeLevelApi) DeleteMtDrugTypeLevelByIds(c *gin.Context) {
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := mtDrugTypeLevelService.DeleteMtDrugTypeLevelByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtDrugTypeLevel 更新药品类型级别表
// @Tags MtDrugTypeLevel
// @Summary 更新药品类型级别表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrugTypeLevel true "更新药品类型级别表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtDrugTypeLevel/updateMtDrugTypeLevel [put]
func (mtDrugTypeLevelApi *MtDrugTypeLevelApi) UpdateMtDrugTypeLevel(c *gin.Context) {
	ctx := c.Request.Context()

	var mtDrugTypeLevel medicine.MtDrugTypeLevel
	err := c.ShouldBindJSON(&mtDrugTypeLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mtDrugTypeLevel.UpdatedBy = utils.GetUserID(c)
	err = mtDrugTypeLevelService.UpdateMtDrugTypeLevel(ctx, mtDrugTypeLevel)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtDrugTypeLevel 用id查询药品类型级别表
// @Tags MtDrugTypeLevel
// @Summary 用id查询药品类型级别表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询药品类型级别表"
// @Success 200 {object} response.Response{data=medicine.MtDrugTypeLevel,msg=string} "查询成功"
// @Router /mtDrugTypeLevel/findMtDrugTypeLevel [get]
func (mtDrugTypeLevelApi *MtDrugTypeLevelApi) FindMtDrugTypeLevel(c *gin.Context) {
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reMtDrugTypeLevel, err := mtDrugTypeLevelService.GetMtDrugTypeLevel(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reMtDrugTypeLevel, c)
}

// GetMtDrugTypeLevelList 分页获取药品类型级别表列表
// @Tags MtDrugTypeLevel
// @Summary 分页获取药品类型级别表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDrugTypeLevelSearch true "分页获取药品类型级别表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtDrugTypeLevel/getMtDrugTypeLevelList [get]
func (mtDrugTypeLevelApi *MtDrugTypeLevelApi) GetMtDrugTypeLevelList(c *gin.Context) {
	ctx := c.Request.Context()

	var pageInfo medicineReq.MtDrugTypeLevelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtDrugTypeLevelService.GetMtDrugTypeLevelInfoList(ctx, pageInfo)
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

// GetAllMtDrugTypeLevel 获取所有药品类型级别
// @Tags MtDrugTypeLevel
// @Summary 获取所有药品类型级别
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]medicine.MtDrugTypeLevel,msg=string} "获取成功"
// @Router /mtDrugTypeLevel/getAllMtDrugTypeLevel [get]
func (mtDrugTypeLevelApi *MtDrugTypeLevelApi) GetAllMtDrugTypeLevel(c *gin.Context) {
	ctx := c.Request.Context()

	list, err := mtDrugTypeLevelService.GetAllMtDrugTypeLevel(ctx)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
