package project

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/project"
    projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TargetApi struct {
}


// CreateTarget 创建Target
// @Tags Target
// @Summary 创建Target
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.Target true "创建Target"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /target/createTarget [post]
func (targetApi *TargetApi) CreateTarget(c *gin.Context) {
    var _target project.Target
    _ = c.ShouldBindJSON(&_target)
    if err := targetService.CreateTarget(_target); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeleteTarget 删除Target
// @Tags Target
// @Summary 删除Target
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除Target"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target/deleteTarget [delete]
func (targetApi *TargetApi) DeleteTarget(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if err := targetService.DeleteTarget(_id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteTargetByIds 批量删除Target
// @Tags Target
// @Summary 批量删除Target
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除Target"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /target/deleteTargetByIds [delete]
func (targetApi *TargetApi) DeleteTargetByIds(c *gin.Context) {
    var IDS request.MIdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := targetService.DeleteTargetByIds(IDS.Ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdateTarget 更新Target
// @Tags Target
// @Summary 更新Target
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.Target true "更新Target"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /target/updateTarget [put]
func (targetApi *TargetApi) UpdateTarget(c *gin.Context) {
    var target project.Target
    _ = c.ShouldBindJSON(&target)
    if err := targetService.UpdateTarget(target); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindTarget 用id查询Target
// @Tags Target
// @Summary 用id查询Target
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询Target"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /target/findTarget [get]
func (targetApi *TargetApi) FindTarget(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if retarget,err := targetService.GetTargetById(_id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"retarget": retarget}, c)
    }
}

// GetTargetList 分页获取Target列表
// @Tags Target
// @Summary 分页获取Target列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.TargetSearch true "分页获取Target列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /target/getTargetList [post]
func (targetApi *TargetApi) GetTargetList(c *gin.Context) {
    var pageInfo projectReq.TargetSearch
    _ = c.ShouldBindJSON(&pageInfo)
    if list, total, err := targetService.GetTargetInfoList(pageInfo.Target,pageInfo.PageInfo,pageInfo.OrderKey,pageInfo.Desc); err != nil {
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
