package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PathInfoRouter struct {
}

// InitPathInfoRouter 初始化 PathInfo 路由信息
func (s *PathInfoRouter) InitPathInfoRouter(Router *gin.RouterGroup) {
	path_infoRouter := Router.Group("path_info").Use(middleware.OperationRecord())
	path_infoRouterWithoutRecord := Router.Group("path_info")
	var path_infoApi = v1.ApiGroupApp.ProjectApiGroup.PathInfoApi
	{
		path_infoRouter.POST("createPathInfo", path_infoApi.CreatePathInfo)   // 新建PathInfo
		path_infoRouter.DELETE("deletePathInfo", path_infoApi.DeletePathInfo) // 删除PathInfo
		path_infoRouter.DELETE("deletePathInfoByIds", path_infoApi.DeletePathInfoByIds) // 批量删除PathInfo
		path_infoRouter.PUT("updatePathInfo", path_infoApi.UpdatePathInfo)    // 更新PathInfo
	}
	{
		path_infoRouterWithoutRecord.GET("findPathInfo", path_infoApi.FindPathInfo)        // 根据ID获取PathInfo
		path_infoRouterWithoutRecord.POST("getPathInfoList", path_infoApi.GetPathInfoList)  // 获取PathInfo列表
	}
}
