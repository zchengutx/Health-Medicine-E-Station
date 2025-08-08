package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MtHospitalsApi struct{}

var mtHospitalsService = service.ServiceGroupApp.MedicineServiceGroup.MtHospitalsService

// CreateMtHospitals 创建医院
// @Tags MtHospitals
// @Summary 创建医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtHospitals true "创建医院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtHospitals/createMtHospitals [post]
func (mtHospitalsApi *MtHospitalsApi) CreateMtHospitals(c *gin.Context) {
	var mtHospitals medicine.MtHospitals
	err := c.ShouldBindJSON(&mtHospitals)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mtHospitalsService.CreateMtHospitals(mtHospitals)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMtHospitals 删除医院
// @Tags MtHospitals
// @Summary 删除医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtHospitals true "删除医院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtHospitals/deleteMtHospitals [delete]
func (mtHospitalsApi *MtHospitalsApi) DeleteMtHospitals(c *gin.Context) {
	ID := c.Query("ID")
	err := mtHospitalsService.DeleteMtHospitals(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtHospitalsByIds 批量删除医院
// @Tags MtHospitals
// @Summary 批量删除医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtHospitals/deleteMtHospitalsByIds [delete]
func (mtHospitalsApi *MtHospitalsApi) DeleteMtHospitalsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := mtHospitalsService.DeleteMtHospitalsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtHospitals 更新医院
// @Tags MtHospitals
// @Summary 更新医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtHospitals true "更新医院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtHospitals/updateMtHospitals [put]
func (mtHospitalsApi *MtHospitalsApi) UpdateMtHospitals(c *gin.Context) {
	var mtHospitals medicine.MtHospitals
	err := c.ShouldBindJSON(&mtHospitals)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mtHospitalsService.UpdateMtHospitals(mtHospitals)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtHospitals 用id查询医院
// @Tags MtHospitals
// @Summary 用id查询医院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query medicine.MtHospitals true "用id查询医院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtHospitals/findMtHospitals [get]
func (mtHospitalsApi *MtHospitalsApi) FindMtHospitals(c *gin.Context) {
	ID := c.Query("ID")
	remtHospitals, err := mtHospitalsService.GetMtHospitals(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remtHospitals": remtHospitals}, c)
	}
}

// GetMtHospitalsList 分页获取医院列表
// @Tags MtHospitals
// @Summary 分页获取医院列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取医院列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtHospitals/getMtHospitalsList [get]
func (mtHospitalsApi *MtHospitalsApi) GetMtHospitalsList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtHospitalsService.GetMtHospitalsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetMtHospitalsListPublic 获取医院列表（公开接口）
// @Tags MtHospitals
// @Summary 获取医院列表（公开接口）
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtHospitals/getMtHospitalsListPublic [get]
func (mtHospitalsApi *MtHospitalsApi) GetMtHospitalsListPublic(c *gin.Context) {
	list, err := mtHospitalsService.GetMtHospitalsListPublic()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
