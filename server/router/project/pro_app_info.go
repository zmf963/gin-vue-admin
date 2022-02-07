package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppInfoRouter struct {
}

// InitAppInfoRouter 初始化 AppInfo 路由信息
func (s *AppInfoRouter) InitAppInfoRouter(Router *gin.RouterGroup) {
	app_infoRouter := Router.Group("app_info").Use(middleware.OperationRecord())
	app_infoRouterWithoutRecord := Router.Group("app_info")
	var app_infoApi = v1.ApiGroupApp.ProjectApiGroup.AppInfoApi
	{
		app_infoRouter.POST("createAppInfo", app_infoApi.CreateAppInfo)   // 新建AppInfo
		app_infoRouter.DELETE("deleteAppInfo", app_infoApi.DeleteAppInfo) // 删除AppInfo
		app_infoRouter.DELETE("deleteAppInfoByIds", app_infoApi.DeleteAppInfoByIds) // 批量删除AppInfo
		app_infoRouter.PUT("updateAppInfo", app_infoApi.UpdateAppInfo)    // 更新AppInfo
	}
	{
		app_infoRouterWithoutRecord.GET("findAppInfo", app_infoApi.FindAppInfo)        // 根据ID获取AppInfo
		app_infoRouterWithoutRecord.POST("getAppInfoList", app_infoApi.GetAppInfoList)  // 获取AppInfo列表
	}
}
