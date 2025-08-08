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

type MtDoctorsApi struct {}



// CreateMtDoctors 创建mtDoctors表
// @Tags MtDoctors
// @Summary 创建mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctors true "创建mtDoctors表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mtDoctors/createMtDoctors [post]
func (mtDoctorsApi *MtDoctorsApi) CreateMtDoctors(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var mtDoctors medicine.MtDoctors
	err := c.ShouldBindJSON(&mtDoctors)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtDoctors.CreatedBy = utils.GetUserID(c)
	err = mtDoctorsService.CreateMtDoctors(ctx,&mtDoctors)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMtDoctors 删除mtDoctors表
// @Tags MtDoctors
// @Summary 删除mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctors true "删除mtDoctors表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mtDoctors/deleteMtDoctors [delete]
func (mtDoctorsApi *MtDoctorsApi) DeleteMtDoctors(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := mtDoctorsService.DeleteMtDoctors(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtDoctorsByIds 批量删除mtDoctors表
// @Tags MtDoctors
// @Summary 批量删除mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mtDoctors/deleteMtDoctorsByIds [delete]
func (mtDoctorsApi *MtDoctorsApi) DeleteMtDoctorsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := mtDoctorsService.DeleteMtDoctorsByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtDoctors 更新mtDoctors表
// @Tags MtDoctors
// @Summary 更新mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctors true "更新mtDoctors表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mtDoctors/updateMtDoctors [put]
func (mtDoctorsApi *MtDoctorsApi) UpdateMtDoctors(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var mtDoctors medicine.MtDoctors
	err := c.ShouldBindJSON(&mtDoctors)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    mtDoctors.UpdatedBy = utils.GetUserID(c)
	err = mtDoctorsService.UpdateMtDoctors(ctx,mtDoctors)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtDoctors 用id查询mtDoctors表
// @Tags MtDoctors
// @Summary 用id查询mtDoctors表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mtDoctors表"
// @Success 200 {object} response.Response{data=medicine.MtDoctors,msg=string} "查询成功"
// @Router /mtDoctors/findMtDoctors [get]
func (mtDoctorsApi *MtDoctorsApi) FindMtDoctors(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remtDoctors, err := mtDoctorsService.GetMtDoctors(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remtDoctors, c)
}
// GetMtDoctorsList 分页获取mtDoctors表列表
// @Tags MtDoctors
// @Summary 分页获取mtDoctors表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDoctorsSearch true "分页获取mtDoctors表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mtDoctors/getMtDoctorsList [get]
func (mtDoctorsApi *MtDoctorsApi) GetMtDoctorsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo medicineReq.MtDoctorsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtDoctorsService.GetMtDoctorsInfoList(ctx,pageInfo)
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

// GetMtDoctorsPublic 不需要鉴权的mtDoctors表接口
// @Tags MtDoctors
// @Summary 不需要鉴权的mtDoctors表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mtDoctors/getMtDoctorsPublic [get]
func (mtDoctorsApi *MtDoctorsApi) GetMtDoctorsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    mtDoctorsService.GetMtDoctorsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的mtDoctors表接口信息",
    }, "获取成功", c)
}
