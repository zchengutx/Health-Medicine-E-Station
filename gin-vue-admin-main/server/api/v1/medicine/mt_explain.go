package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MtExplainApi struct{}

// FindMtExplain 用id查询说明书表
// @Tags MtExplain
// @Summary 用id查询说明书表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询说明书表"
// @Success 200 {object} response.Response{data=medicine.MtExplain,msg=string} "查询成功"
// @Router /mtExplain/findMtExplain [get]
func (mtExplainApi *MtExplainApi) FindMtExplain(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")
	reMtExplain, err := mtExplainService.GetMtExplain(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reMtExplain, c)
}

// GetMtExplainList 分页获取说明书表列表
// @Tags MtExplain
// @Summary 分页获取说明书表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtExplainSearch true "分页获取说明书表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtExplain/getMtExplainList [get]
func (mtExplainApi *MtExplainApi) GetMtExplainList(c *gin.Context) {
	ctx := c.Request.Context()

	var pageInfo medicineReq.MtExplainSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtExplainService.GetMtExplainInfoList(ctx, pageInfo)
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
