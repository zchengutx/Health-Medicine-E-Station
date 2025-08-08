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

type MtDrugTypeStairApi struct{}

// CreateMtDrugTypeStair 创建药品类型阶梯表
// @Tags MtDrugTypeStair
// @Summary 创建药品类型阶梯表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrugTypeStair true "创建药品类型阶梯表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtDrugTypeStair/createMtDrugTypeStair [post]
func (mtDrugTypeStairApi *MtDrugTypeStairApi) CreateMtDrugTypeStair(c *gin.Context) {
	ctx := c.Request.Context()

	var mtDrugTypeStair medicine.MtDrugTypeStair
	err := c.ShouldBindJSON(&mtDrugTypeStair)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mtDrugTypeStair.CreatedBy = utils.GetUserID(c)
	err = mtDrugTypeStairService.CreateMtDrugTypeStair(ctx, &mtDrugTypeStair)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMtDrugTypeStair 删除药品类型阶梯表
// @Tags MtDrugTypeStair
// @Summary 删除药品类型阶梯表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrugTypeStair true "删除药品类型阶梯表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtDrugTypeStair/deleteMtDrugTypeStair [delete]
func (mtDrugTypeStairApi *MtDrugTypeStairApi) DeleteMtDrugTypeStair(c *gin.Context) {
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := mtDrugTypeStairService.DeleteMtDrugTypeStair(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtDrugTypeStairByIds 批量删除药品类型阶梯表
// @Tags MtDrugTypeStair
// @Summary 批量删除药品类型阶梯表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtDrugTypeStair/deleteMtDrugTypeStairByIds [delete]
func (mtDrugTypeStairApi *MtDrugTypeStairApi) DeleteMtDrugTypeStairByIds(c *gin.Context) {
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := mtDrugTypeStairService.DeleteMtDrugTypeStairByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtDrugTypeStair 更新药品类型阶梯表
// @Tags MtDrugTypeStair
// @Summary 更新药品类型阶梯表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDrugTypeStair true "更新药品类型阶梯表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtDrugTypeStair/updateMtDrugTypeStair [put]
func (mtDrugTypeStairApi *MtDrugTypeStairApi) UpdateMtDrugTypeStair(c *gin.Context) {
	ctx := c.Request.Context()

	var mtDrugTypeStair medicine.MtDrugTypeStair
	err := c.ShouldBindJSON(&mtDrugTypeStair)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mtDrugTypeStair.UpdatedBy = utils.GetUserID(c)
	err = mtDrugTypeStairService.UpdateMtDrugTypeStair(ctx, mtDrugTypeStair)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtDrugTypeStair 用id查询药品类型阶梯表
// @Tags MtDrugTypeStair
// @Summary 用id查询药品类型阶梯表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询药品类型阶梯表"
// @Success 200 {object} response.Response{data=medicine.MtDrugTypeStair,msg=string} "查询成功"
// @Router /mtDrugTypeStair/findMtDrugTypeStair [get]
func (mtDrugTypeStairApi *MtDrugTypeStairApi) FindMtDrugTypeStair(c *gin.Context) {
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reMtDrugTypeStair, err := mtDrugTypeStairService.GetMtDrugTypeStair(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reMtDrugTypeStair, c)
}

// GetMtDrugTypeStairList 分页获取药品类型阶梯表列表
// @Tags MtDrugTypeStair
// @Summary 分页获取药品类型阶梯表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDrugTypeStairSearch true "分页获取药品类型阶梯表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtDrugTypeStair/getMtDrugTypeStairList [get]
func (mtDrugTypeStairApi *MtDrugTypeStairApi) GetMtDrugTypeStairList(c *gin.Context) {
	ctx := c.Request.Context()

	var pageInfo medicineReq.MtDrugTypeStairSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtDrugTypeStairService.GetMtDrugTypeStairInfoList(ctx, pageInfo)
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

// GetAllMtDrugTypeStair 获取所有药品类型阶梯
// @Tags MtDrugTypeStair
// @Summary 获取所有药品类型阶梯
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]medicine.MtDrugTypeStair,msg=string} "获取成功"
// @Router /mtDrugTypeStair/getAllMtDrugTypeStair [get]
func (mtDrugTypeStairApi *MtDrugTypeStairApi) GetAllMtDrugTypeStair(c *gin.Context) {
	ctx := c.Request.Context()

	list, err := mtDrugTypeStairService.GetAllMtDrugTypeStair(ctx)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
