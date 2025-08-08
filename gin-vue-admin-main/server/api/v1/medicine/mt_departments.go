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

type MtDepartmentsApi struct{}

var mtDepartmentsService = service.ServiceGroupApp.MedicineServiceGroup.MtDepartmentsService

// CreateMtDepartments 创建科室
// @Tags MtDepartments
// @Summary 创建科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtDepartments true "创建科室"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtDepartments/createMtDepartments [post]
func (mtDepartmentsApi *MtDepartmentsApi) CreateMtDepartments(c *gin.Context) {
	var mtDepartments medicine.MtDepartments
	err := c.ShouldBindJSON(&mtDepartments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mtDepartmentsService.CreateMtDepartments(mtDepartments)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMtDepartments 删除科室
// @Tags MtDepartments
// @Summary 删除科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtDepartments true "删除科室"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDepartments/deleteMtDepartments [delete]
func (mtDepartmentsApi *MtDepartmentsApi) DeleteMtDepartments(c *gin.Context) {
	ID := c.Query("ID")
	err := mtDepartmentsService.DeleteMtDepartments(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMtDepartmentsByIds 批量删除科室
// @Tags MtDepartments
// @Summary 批量删除科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDepartments/deleteMtDepartmentsByIds [delete]
func (mtDepartmentsApi *MtDepartmentsApi) DeleteMtDepartmentsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := mtDepartmentsService.DeleteMtDepartmentsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMtDepartments 更新科室
// @Tags MtDepartments
// @Summary 更新科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtDepartments true "更新科室"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtDepartments/updateMtDepartments [put]
func (mtDepartmentsApi *MtDepartmentsApi) UpdateMtDepartments(c *gin.Context) {
	var mtDepartments medicine.MtDepartments
	err := c.ShouldBindJSON(&mtDepartments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mtDepartmentsService.UpdateMtDepartments(mtDepartments)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMtDepartments 用id查询科室
// @Tags MtDepartments
// @Summary 用id查询科室
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query medicine.MtDepartments true "用id查询科室"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtDepartments/findMtDepartments [get]
func (mtDepartmentsApi *MtDepartmentsApi) FindMtDepartments(c *gin.Context) {
	ID := c.Query("ID")
	remtDepartments, err := mtDepartmentsService.GetMtDepartments(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remtDepartments": remtDepartments}, c)
	}
}

// GetMtDepartmentsList 分页获取科室列表
// @Tags MtDepartments
// @Summary 分页获取科室列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取科室列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDepartments/getMtDepartmentsList [get]
func (mtDepartmentsApi *MtDepartmentsApi) GetMtDepartmentsList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtDepartmentsService.GetMtDepartmentsInfoList(pageInfo)
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

// GetMtDepartmentsListPublic 获取科室列表（公开接口）
// @Tags MtDepartments
// @Summary 获取科室列表（公开接口）
// @accept application/json
// @Produce application/json
// @Param hospitalId query uint false "医院ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDepartments/getMtDepartmentsListPublic [get]
func (mtDepartmentsApi *MtDepartmentsApi) GetMtDepartmentsListPublic(c *gin.Context) {
	hospitalId := c.Query("hospitalId")
	list, err := mtDepartmentsService.GetMtDepartmentsListPublic(hospitalId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
