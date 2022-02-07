package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectInfoRouter struct {
}

// InitProjectInfoRouter 初始化 ProjectInfo 路由信息
func (s *ProjectInfoRouter) InitProjectInfoRouter(Router *gin.RouterGroup) {
	project_infoRouter := Router.Group("project_info").Use(middleware.OperationRecord())
	project_infoRouterWithoutRecord := Router.Group("project_info")
	var project_infoApi = v1.ApiGroupApp.ProjectApiGroup.ProjectInfoApi
	{
		project_infoRouter.POST("createProjectInfo", project_infoApi.CreateProjectInfo)   // 新建ProjectInfo
		project_infoRouter.DELETE("deleteProjectInfo", project_infoApi.DeleteProjectInfo) // 删除ProjectInfo
		project_infoRouter.DELETE("deleteProjectInfoByIds", project_infoApi.DeleteProjectInfoByIds) // 批量删除ProjectInfo
		project_infoRouter.PUT("updateProjectInfo", project_infoApi.UpdateProjectInfo)    // 更新ProjectInfo
	}
	{
		project_infoRouterWithoutRecord.GET("findProjectInfo", project_infoApi.FindProjectInfo)        // 根据ID获取ProjectInfo
		project_infoRouterWithoutRecord.POST("getProjectInfoList", project_infoApi.GetProjectInfoList)  // 获取ProjectInfo列表
	}
}
