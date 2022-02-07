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

type ProjectInfoApi struct {
}


// CreateProjectInfo 创建ProjectInfo
// @Tags ProjectInfo
// @Summary 创建ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.ProjectInfo true "创建ProjectInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /project_info/createProjectInfo [post]
func (project_infoApi *ProjectInfoApi) CreateProjectInfo(c *gin.Context) {
    var _project_info project.ProjectInfo
    _ = c.ShouldBindJSON(&_project_info)
    if err := project_infoService.CreateProjectInfo(_project_info); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeleteProjectInfo 删除ProjectInfo
// @Tags ProjectInfo
// @Summary 删除ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project_info/deleteProjectInfo [delete]
func (project_infoApi *ProjectInfoApi) DeleteProjectInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if err := project_infoService.DeleteProjectInfo(_id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteProjectInfoByIds 批量删除ProjectInfo
// @Tags ProjectInfo
// @Summary 批量删除ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /project_info/deleteProjectInfoByIds [delete]
func (project_infoApi *ProjectInfoApi) DeleteProjectInfoByIds(c *gin.Context) {
    var IDS request.MIdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := project_infoService.DeleteProjectInfoByIds(IDS.Ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdateProjectInfo 更新ProjectInfo
// @Tags ProjectInfo
// @Summary 更新ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.ProjectInfo true "更新ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project_info/updateProjectInfo [put]
func (project_infoApi *ProjectInfoApi) UpdateProjectInfo(c *gin.Context) {
    var project_info project.ProjectInfo
    _ = c.ShouldBindJSON(&project_info)
    if err := project_infoService.UpdateProjectInfo(project_info); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindProjectInfo 用id查询ProjectInfo
// @Tags ProjectInfo
// @Summary 用id查询ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project_info/findProjectInfo [get]
func (project_infoApi *ProjectInfoApi) FindProjectInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if reproject_info,err := project_infoService.GetProjectInfoById(_id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"reproject_info": reproject_info}, c)
    }
}

// GetProjectInfoList 分页获取ProjectInfo列表
// @Tags ProjectInfo
// @Summary 分页获取ProjectInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.ProjectInfoSearch true "分页获取ProjectInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /project_info/getProjectInfoList [post]
func (project_infoApi *ProjectInfoApi) GetProjectInfoList(c *gin.Context) {
    var pageInfo projectReq.ProjectInfoSearch
    _ = c.ShouldBindJSON(&pageInfo)
    if list, total, err := project_infoService.GetProjectInfoInfoList(pageInfo.ProjectInfo,pageInfo.PageInfo,pageInfo.OrderKey,pageInfo.Desc); err != nil {
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
