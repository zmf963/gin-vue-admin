package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type PortInfoApi struct {
}

// CreatePortInfo 创建PortInfo
// @Tags PortInfo
// @Summary 创建PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.PortInfo true "创建PortInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /port_info/createPortInfo [post]
func (port_infoApi *PortInfoApi) CreatePortInfo(c *gin.Context) {
	var _port_info project.PortInfo
	_ = c.ShouldBindJSON(&_port_info)
	if err := port_infoService.CreatePortInfo(_port_info); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePortInfo 删除PortInfo
// @Tags PortInfo
// @Summary 删除PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除PortInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /port_info/deletePortInfo [delete]
func (port_infoApi *PortInfoApi) DeletePortInfo(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := port_infoService.DeletePortInfo(_id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePortInfoByIds 批量删除PortInfo
// @Tags PortInfo
// @Summary 批量删除PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除PortInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /port_info/deletePortInfoByIds [delete]
func (port_infoApi *PortInfoApi) DeletePortInfoByIds(c *gin.Context) {
	var IDS request.MIdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := port_infoService.DeletePortInfoByIds(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePortInfo 更新PortInfo
// @Tags PortInfo
// @Summary 更新PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.PortInfo true "更新PortInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /port_info/updatePortInfo [put]
func (port_infoApi *PortInfoApi) UpdatePortInfo(c *gin.Context) {
	var port_info project.PortInfo
	_ = c.ShouldBindJSON(&port_info)
	if err := port_infoService.UpdatePortInfo(port_info); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPortInfo 用id查询PortInfo
// @Tags PortInfo
// @Summary 用id查询PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询PortInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /port_info/findPortInfo [get]
func (port_infoApi *PortInfoApi) FindPortInfo(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if report_info, err := port_infoService.GetPortInfoById(_id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"report_info": report_info}, c)
	}
}

// GetPortInfoList 分页获取PortInfo列表
// @Tags PortInfo
// @Summary 分页获取PortInfo列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.PortInfoSearch true "分页获取PortInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /port_info/getPortInfoList [post]
func (port_infoApi *PortInfoApi) GetPortInfoList(c *gin.Context) {
	var pageInfo projectReq.PortInfoSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if list, total, err := port_infoService.GetPortInfoInfoList(pageInfo.PortInfo, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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
