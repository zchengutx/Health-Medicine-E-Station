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

type MtUserApi struct {}



// CreateMtUser 创建mtUser表
// @Tags MtUser
// @Summary 创建mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtUser true "创建mtUser表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtUser/createMtUser [post]
func (mtUserApi *MtUserApi) CreateMtUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var mtUser medicine.MtUser
	err := c.ShouldBindJSON(&mtUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtUser.CreatedBy = utils.GetUserID(c)
	err = mtUserService.CreateMtUser(ctx,&mtUser)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMtUser 删除mtUser表
// @Tags MtUser
// @Summary 删除mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtUser true "删除mtUser表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtUser/deleteMtUser [delete]
func (mtUserApi *MtUserApi) DeleteMtUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := mtUserService.DeleteMtUser(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtUserByIds 批量删除mtUser表
// @Tags MtUser
// @Summary 批量删除mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtUser/deleteMtUserByIds [delete]
func (mtUserApi *MtUserApi) DeleteMtUserByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := mtUserService.DeleteMtUserByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtUser 更新mtUser表
// @Tags MtUser
// @Summary 更新mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtUser true "更新mtUser表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtUser/updateMtUser [put]
func (mtUserApi *MtUserApi) UpdateMtUser(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var mtUser medicine.MtUser
	err := c.ShouldBindJSON(&mtUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtUser.UpdatedBy = utils.GetUserID(c)
	err = mtUserService.UpdateMtUser(ctx,mtUser)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtUser 用id查询mtUser表
// @Tags MtUser
// @Summary 用id查询mtUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mtUser表"
// @Success 200 {object} response.Response{data=medicine.MtUser,msg=string} "查询成功"
// @Router /mtUser/findMtUser [get]
func (mtUserApi *MtUserApi) FindMtUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remtUser, err := mtUserService.GetMtUser(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remtUser, c)
}
// GetMtUserList 分页获取mtUser表列表
// @Tags MtUser
// @Summary 分页获取mtUser表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtUserSearch true "分页获取mtUser表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtUser/getMtUserList [get]
func (mtUserApi *MtUserApi) GetMtUserList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo medicineReq.MtUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtUserService.GetMtUserInfoList(ctx,pageInfo)
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

// GetMtUserPublic 不需要鉴权的mtUser表接口
// @Tags MtUser
// @Summary 不需要鉴权的mtUser表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtUser/getMtUserPublic [get]
func (mtUserApi *MtUserApi) GetMtUserPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    mtUserService.GetMtUserPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的mtUser表接口信息",
    }, "获取成功", c)
}
