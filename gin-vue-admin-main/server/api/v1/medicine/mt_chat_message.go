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

type MtChatMessageApi struct {}



// CreateMtChatMessage 创建mtChatMessage表
// @Tags MtChatMessage
// @Summary 创建mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtChatMessage true "创建mtChatMessage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtChatMessage/createMtChatMessage [post]
func (mtChatMessageApi *MtChatMessageApi) CreateMtChatMessage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var mtChatMessage medicine.MtChatMessage
	err := c.ShouldBindJSON(&mtChatMessage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtChatMessage.CreatedBy = utils.GetUserID(c)
	err = mtChatMessageService.CreateMtChatMessage(ctx,&mtChatMessage)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMtChatMessage 删除mtChatMessage表
// @Tags MtChatMessage
// @Summary 删除mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtChatMessage true "删除mtChatMessage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtChatMessage/deleteMtChatMessage [delete]
func (mtChatMessageApi *MtChatMessageApi) DeleteMtChatMessage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := mtChatMessageService.DeleteMtChatMessage(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtChatMessageByIds 批量删除mtChatMessage表
// @Tags MtChatMessage
// @Summary 批量删除mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtChatMessage/deleteMtChatMessageByIds [delete]
func (mtChatMessageApi *MtChatMessageApi) DeleteMtChatMessageByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := mtChatMessageService.DeleteMtChatMessageByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtChatMessage 更新mtChatMessage表
// @Tags MtChatMessage
// @Summary 更新mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtChatMessage true "更新mtChatMessage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtChatMessage/updateMtChatMessage [put]
func (mtChatMessageApi *MtChatMessageApi) UpdateMtChatMessage(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var mtChatMessage medicine.MtChatMessage
	err := c.ShouldBindJSON(&mtChatMessage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtChatMessage.UpdatedBy = utils.GetUserID(c)
	err = mtChatMessageService.UpdateMtChatMessage(ctx,mtChatMessage)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtChatMessage 用id查询mtChatMessage表
// @Tags MtChatMessage
// @Summary 用id查询mtChatMessage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mtChatMessage表"
// @Success 200 {object} response.Response{data=medicine.MtChatMessage,msg=string} "查询成功"
// @Router /mtChatMessage/findMtChatMessage [get]
func (mtChatMessageApi *MtChatMessageApi) FindMtChatMessage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remtChatMessage, err := mtChatMessageService.GetMtChatMessage(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remtChatMessage, c)
}
// GetMtChatMessageList 分页获取mtChatMessage表列表
// @Tags MtChatMessage
// @Summary 分页获取mtChatMessage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtChatMessageSearch true "分页获取mtChatMessage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtChatMessage/getMtChatMessageList [get]
func (mtChatMessageApi *MtChatMessageApi) GetMtChatMessageList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo medicineReq.MtChatMessageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtChatMessageService.GetMtChatMessageInfoList(ctx,pageInfo)
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

// GetMtChatMessagePublic 不需要鉴权的mtChatMessage表接口
// @Tags MtChatMessage
// @Summary 不需要鉴权的mtChatMessage表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtChatMessage/getMtChatMessagePublic [get]
func (mtChatMessageApi *MtChatMessageApi) GetMtChatMessagePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    mtChatMessageService.GetMtChatMessagePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的mtChatMessage表接口信息",
    }, "获取成功", c)
}
