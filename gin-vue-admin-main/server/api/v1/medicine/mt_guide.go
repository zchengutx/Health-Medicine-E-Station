package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MtGuideApi struct{}

// FindMtGuide 用id查询用药指导表
// @Tags MtGuide
// @Summary 用id查询用药指导表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询用药指导表"
// @Success 200 {object} response.Response{data=medicine.MtGuide,msg=string} "查询成功"
// @Router /mtGuide/findMtGuide [get]
func (mtGuideApi *MtGuideApi) FindMtGuide(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")
	reMtGuide, err := mtGuideService.GetMtGuide(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reMtGuide, c)
}

// GetMtGuideList 分页获取用药指导表列表
// @Tags MtGuide
// @Summary 分页获取用药指导表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtGuideSearch true "分页获取用药指导表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtGuide/getMtGuideList [get]
func (mtGuideApi *MtGuideApi) GetMtGuideList(c *gin.Context) {
	ctx := c.Request.Context()

	var pageInfo medicineReq.MtGuideSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtGuideService.GetMtGuideInfoList(ctx, pageInfo)
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
