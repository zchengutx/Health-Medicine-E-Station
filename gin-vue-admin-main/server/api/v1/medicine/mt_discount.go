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

type MtDiscountApi struct {}



// CreateMtDiscount 创建mtDiscount表
// @Tags MtDiscount
// @Summary 创建mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDiscount true "创建mtDiscount表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtDiscount/createMtDiscount [post]
func (mtDiscountApi *MtDiscountApi) CreateMtDiscount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var mtDiscount medicine.MtDiscount
	err := c.ShouldBindJSON(&mtDiscount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtDiscount.CreatedBy = utils.GetUserID(c)
	err = mtDiscountService.CreateMtDiscount(ctx,&mtDiscount)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMtDiscount 删除mtDiscount表
// @Tags MtDiscount
// @Summary 删除mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDiscount true "删除mtDiscount表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtDiscount/deleteMtDiscount [delete]
func (mtDiscountApi *MtDiscountApi) DeleteMtDiscount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := mtDiscountService.DeleteMtDiscount(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtDiscountByIds 批量删除mtDiscount表
// @Tags MtDiscount
// @Summary 批量删除mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtDiscount/deleteMtDiscountByIds [delete]
func (mtDiscountApi *MtDiscountApi) DeleteMtDiscountByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := mtDiscountService.DeleteMtDiscountByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtDiscount 更新mtDiscount表
// @Tags MtDiscount
// @Summary 更新mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDiscount true "更新mtDiscount表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtDiscount/updateMtDiscount [put]
func (mtDiscountApi *MtDiscountApi) UpdateMtDiscount(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var mtDiscount medicine.MtDiscount
	err := c.ShouldBindJSON(&mtDiscount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtDiscount.UpdatedBy = utils.GetUserID(c)
	err = mtDiscountService.UpdateMtDiscount(ctx,mtDiscount)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtDiscount 用id查询mtDiscount表
// @Tags MtDiscount
// @Summary 用id查询mtDiscount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mtDiscount表"
// @Success 200 {object} response.Response{data=medicine.MtDiscount,msg=string} "查询成功"
// @Router /mtDiscount/findMtDiscount [get]
func (mtDiscountApi *MtDiscountApi) FindMtDiscount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remtDiscount, err := mtDiscountService.GetMtDiscount(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remtDiscount, c)
}
// GetMtDiscountList 分页获取mtDiscount表列表
// @Tags MtDiscount
// @Summary 分页获取mtDiscount表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDiscountSearch true "分页获取mtDiscount表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtDiscount/getMtDiscountList [get]
func (mtDiscountApi *MtDiscountApi) GetMtDiscountList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo medicineReq.MtDiscountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtDiscountService.GetMtDiscountInfoList(ctx,pageInfo)
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

// GetMtDiscountPublic 不需要鉴权的mtDiscount表接口
// @Tags MtDiscount
// @Summary 不需要鉴权的mtDiscount表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtDiscount/getMtDiscountPublic [get]
func (mtDiscountApi *MtDiscountApi) GetMtDiscountPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    mtDiscountService.GetMtDiscountPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的mtDiscount表接口信息",
    }, "获取成功", c)
}
