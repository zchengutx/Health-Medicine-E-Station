package medicine

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/medicine"
	medicineReq "github.com/flipped-aurora/gin-vue-admin/server/model/medicine/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MtDoctorApprovalApi struct{}

var mtDoctorApprovalService = service.ServiceGroupApp.MedicineServiceGroup.MtDoctorApprovalService

// CreateMtDoctorApproval 创建mtDoctorApproval表
// @Tags MtDoctorApproval
// @Summary 创建mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctorApproval true "创建mtDoctorApproval表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mtDoctorApproval/createMtDoctorApproval [post]
func (mtDoctorApprovalApi *MtDoctorApprovalApi) CreateMtDoctorApproval(c *gin.Context) {
	var mtDoctorApproval medicine.MtDoctorApproval
	err := c.ShouldBindJSON(&mtDoctorApproval)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mtDoctorApprovalService.CreateMtDoctorApproval(c.Request.Context(), mtDoctorApproval); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMtDoctorApproval 删除mtDoctorApproval表
// @Tags MtDoctorApproval
// @Summary 删除mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctorApproval true "删除mtDoctorApproval表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDoctorApproval/deleteMtDoctorApproval [delete]
func (mtDoctorApprovalApi *MtDoctorApprovalApi) DeleteMtDoctorApproval(c *gin.Context) {
	ID := c.Query("ID")
	if err := mtDoctorApprovalService.DeleteMtDoctorApproval(c.Request.Context(), ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMtDoctorApprovalByIds 批量删除mtDoctorApproval表
// @Tags MtDoctorApproval
// @Summary 批量删除mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mtDoctorApproval/deleteMtDoctorApproval [delete]
func (mtDoctorApprovalApi *MtDoctorApprovalApi) DeleteMtDoctorApprovalByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := mtDoctorApprovalService.DeleteMtDoctorApprovalByIds(c.Request.Context(), IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMtDoctorApproval 更新mtDoctorApproval表
// @Tags MtDoctorApproval
// @Summary 更新mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body medicine.MtDoctorApproval true "更新mtDoctorApproval表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mtDoctorApproval/updateMtDoctorApproval [put]
func (mtDoctorApprovalApi *MtDoctorApprovalApi) UpdateMtDoctorApproval(c *gin.Context) {
	var mtDoctorApproval medicine.MtDoctorApproval
	err := c.ShouldBindJSON(&mtDoctorApproval)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mtDoctorApprovalService.UpdateMtDoctorApproval(c.Request.Context(), mtDoctorApproval); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMtDoctorApproval 用id查询mtDoctorApproval表
// @Tags MtDoctorApproval
// @Summary 用id查询mtDoctorApproval表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query medicine.MtDoctorApproval true "用id查询mtDoctorApproval表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mtDoctorApproval/findMtDoctorApproval [get]
func (mtDoctorApprovalApi *MtDoctorApprovalApi) FindMtDoctorApproval(c *gin.Context) {
	ID := c.Query("ID")
	remtDoctorApproval, err := mtDoctorApprovalService.GetMtDoctorApproval(c.Request.Context(), ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remtDoctorApproval": remtDoctorApproval}, c)
	}
}

// GetMtDoctorApprovalList 分页获取mtDoctorApproval表列表
// @Tags MtDoctorApproval
// @Summary 分页获取mtDoctorApproval表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query medicineReq.MtDoctorApprovalSearch true "分页获取mtDoctorApproval表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDoctorApproval/getMtDoctorApprovalList [get]
func (mtDoctorApprovalApi *MtDoctorApprovalApi) GetMtDoctorApprovalList(c *gin.Context) {
	var pageInfo medicineReq.MtDoctorApprovalSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mtDoctorApprovalService.GetMtDoctorApprovalInfoList(c.Request.Context(), pageInfo)
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

// GetMtDoctorApprovalByDoctorId 根据医生ID获取审核记录
// @Tags MtDoctorApproval
// @Summary 根据医生ID获取审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param doctorId query uint true "医生ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDoctorApproval/getMtDoctorApprovalByDoctorId [get]
func (mtDoctorApprovalApi *MtDoctorApprovalApi) GetMtDoctorApprovalByDoctorId(c *gin.Context) {
	doctorId := c.Query("doctorId")
	if doctorId == "" {
		response.FailWithMessage("医生ID不能为空", c)
		return
	}

	// 将字符串转换为uint
	var doctorIdUint uint
	_, err := fmt.Sscanf(doctorId, "%d", &doctorIdUint)
	if err != nil {
		response.FailWithMessage("医生ID格式错误", c)
		return
	}

	approval, err := mtDoctorApprovalService.GetMtDoctorApprovalByDoctorId(c.Request.Context(), doctorIdUint)
	if err != nil {
		global.GVA_LOG.Error("获取审核记录失败!", zap.Error(err))
		response.FailWithMessage("获取审核记录失败", c)
	} else {
		response.OkWithData(gin.H{"approval": approval}, c)
	}
}

// GetMtDoctorApprovalPublic 不需要鉴权的mtDoctorApproval表接口
// @Tags MtDoctorApproval
// @Summary 不需要鉴权的mtDoctorApproval表接口
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mtDoctorApproval/getMtDoctorApprovalPublic [get]
func (mtDoctorApprovalApi *MtDoctorApprovalApi) GetMtDoctorApprovalPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	mtDoctorApprovalService.GetMtDoctorApprovalPublic(c.Request.Context())
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的mtDoctorApproval表接口信息",
	}, "获取成功", c)
}
