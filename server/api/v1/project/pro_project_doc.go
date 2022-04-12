/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-04-12 18:58:54
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 19:14:25
 * @FilePath: /api/v1/project/pro_project_doc.go
 * @Description:
 */
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

type ProjectDocApi struct {
}

// CreateProjectDoc 创建ProjectDoc
// @Tags ProjectDoc
// @Summary 创建ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.ProjectDoc true "创建ProjectDoc"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /project_doc/createProjectDoc [post]
func (project_docApi *ProjectDocApi) CreateProjectDoc(c *gin.Context) {
	var _project_doc project.ProjectDoc
	_ = c.ShouldBindJSON(&_project_doc)
	if err := project_docService.CreateProjectDoc(_project_doc); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProjectDoc 删除ProjectDoc
// @Tags ProjectDoc
// @Summary 删除ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除ProjectDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project_doc/deleteProjectDoc [delete]
func (project_docApi *ProjectDocApi) DeleteProjectDoc(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := project_docService.DeleteProjectDoc(_id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProjectDocByIds 批量删除ProjectDoc
// @Tags ProjectDoc
// @Summary 批量删除ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除ProjectDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /project_doc/deleteProjectDocByIds [delete]
func (project_docApi *ProjectDocApi) DeleteProjectDocByIds(c *gin.Context) {
	var IDS request.MIdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := project_docService.DeleteProjectDocByIds(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProjectDoc 更新ProjectDoc
// @Tags ProjectDoc
// @Summary 更新ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.ProjectDoc true "更新ProjectDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project_doc/updateProjectDoc [put]
func (project_docApi *ProjectDocApi) UpdateProjectDoc(c *gin.Context) {
	var project_doc project.ProjectDoc
	_ = c.ShouldBindJSON(&project_doc)
	if err := project_docService.UpdateProjectDoc(project_doc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProjectDoc 用id查询ProjectDoc
// @Tags ProjectDoc
// @Summary 用id查询ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询ProjectDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project_doc/findProjectDoc [get]
func (project_docApi *ProjectDocApi) FindProjectDoc(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if reproject_doc, err := project_docService.GetProjectDocById(_id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproject_doc": reproject_doc}, c)
	}
}

// GetProjectDocList 分页获取ProjectDoc列表
// @Tags ProjectDoc
// @Summary 分页获取ProjectDoc列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.ProjectDocSearch true "分页获取ProjectDoc列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /project_doc/getProjectDocList [post]
func (project_docApi *ProjectDocApi) GetProjectDocList(c *gin.Context) {
	var pageInfo projectReq.ProjectDocSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if list, total, err := project_docService.GetProjectDocInfoList(pageInfo.ProjectDoc, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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
