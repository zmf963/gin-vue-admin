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

type TargetRelationApi struct {
}

// CreateTargetRelation 创建TargetRelation
// @Tags TargetRelation
// @Summary 创建TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.TargetRelation true "创建TargetRelation"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /target_relation/createTargetRelation [post]
func (target_relationApi *TargetRelationApi) CreateTargetRelation(c *gin.Context) {
	var _target_relation project.TargetRelation
	_ = c.ShouldBindJSON(&_target_relation)
	if err := target_relationService.CreateTargetRelation(_target_relation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTargetRelation 删除TargetRelation
// @Tags TargetRelation
// @Summary 删除TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除TargetRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target_relation/deleteTargetRelation [delete]
func (target_relationApi *TargetRelationApi) DeleteTargetRelation(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := target_relationService.DeleteTargetRelation(_id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTargetRelationByIds 批量删除TargetRelation
// @Tags TargetRelation
// @Summary 批量删除TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除TargetRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /target_relation/deleteTargetRelationByIds [delete]
func (target_relationApi *TargetRelationApi) DeleteTargetRelationByIds(c *gin.Context) {
	var IDS request.MIdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := target_relationService.DeleteTargetRelationByIds(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTargetRelation 更新TargetRelation
// @Tags TargetRelation
// @Summary 更新TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.TargetRelation true "更新TargetRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /target_relation/updateTargetRelation [put]
func (target_relationApi *TargetRelationApi) UpdateTargetRelation(c *gin.Context) {
	var target_relation project.TargetRelation
	_ = c.ShouldBindJSON(&target_relation)
	if err := target_relationService.UpdateTargetRelation(target_relation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTargetRelation 用id查询TargetRelation
// @Tags TargetRelation
// @Summary 用id查询TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询TargetRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /target_relation/findTargetRelation [get]
func (target_relationApi *TargetRelationApi) FindTargetRelation(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if retarget_relation, err := target_relationService.GetTargetRelationById(_id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retarget_relation": retarget_relation}, c)
	}
}

// GetTargetRelationList 分页获取TargetRelation列表
// @Tags TargetRelation
// @Summary 分页获取TargetRelation列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.TargetRelationSearch true "分页获取TargetRelation列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /target_relation/getTargetRelationList [post]
func (target_relationApi *TargetRelationApi) GetTargetRelationList(c *gin.Context) {
	var pageInfo projectReq.TargetRelationSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if list, total, err := target_relationService.GetTargetRelationInfoList(pageInfo.TargetRelation, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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
