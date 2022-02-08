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

type AppInfoApi struct {
}

// CreateAppInfo 创建AppInfo
// @Tags AppInfo
// @Summary 创建AppInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.AppInfo true "创建AppInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /app_info/createAppInfo [post]
func (app_infoApi *AppInfoApi) CreateAppInfo(c *gin.Context) {
	var _app_info project.AppInfo
	_ = c.ShouldBindJSON(&_app_info)
	if err := app_infoService.CreateAppInfo(_app_info); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppInfo 删除AppInfo
// @Tags AppInfo
// @Summary 删除AppInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除AppInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app_info/deleteAppInfo [delete]
func (app_infoApi *AppInfoApi) DeleteAppInfo(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := app_infoService.DeleteAppInfo(_id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppInfoByIds 批量删除AppInfo
// @Tags AppInfo
// @Summary 批量删除AppInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除AppInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app_info/deleteAppInfoByIds [delete]
func (app_infoApi *AppInfoApi) DeleteAppInfoByIds(c *gin.Context) {
	var IDS request.MIdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := app_infoService.DeleteAppInfoByIds(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppInfo 更新AppInfo
// @Tags AppInfo
// @Summary 更新AppInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.AppInfo true "更新AppInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app_info/updateAppInfo [put]
func (app_infoApi *AppInfoApi) UpdateAppInfo(c *gin.Context) {
	var app_info project.AppInfo
	_ = c.ShouldBindJSON(&app_info)
	if err := app_infoService.UpdateAppInfo(app_info); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppInfo 用id查询AppInfo
// @Tags AppInfo
// @Summary 用id查询AppInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询AppInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app_info/findAppInfo [get]
func (app_infoApi *AppInfoApi) FindAppInfo(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if reapp_info, err := app_infoService.GetAppInfoById(_id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapp_info": reapp_info}, c)
	}
}

// GetAppInfoList 分页获取AppInfo列表
// @Tags AppInfo
// @Summary 分页获取AppInfo列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.AppInfoSearch true "分页获取AppInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /app_info/getAppInfoList [post]
func (app_infoApi *AppInfoApi) GetAppInfoList(c *gin.Context) {
	var pageInfo projectReq.AppInfoSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if list, total, err := app_infoService.GetAppInfoInfoList(pageInfo.AppInfo, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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
