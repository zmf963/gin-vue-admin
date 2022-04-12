/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-04-12 19:59:02
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 20:15:09
 * @FilePath: /router/project/pro_project_doc.go
 * @Description:
 */
package project

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectDocRouter struct {
}

// InitProjectDocRouter 初始化 ProjectDoc 路由信息
func (s *ProjectDocRouter) InitProjectDocRouter(Router *gin.RouterGroup) {
	doc_infoRouter := Router.Group("project_doc").Use(middleware.OperationRecord())
	doc_infoRouterWithoutRecord := Router.Group("project_doc")
	var doc_infoApi = v1.ApiGroupApp.ProjectApiGroup.ProjectDocApi
	{
		doc_infoRouter.POST("createProjectDoc", doc_infoApi.CreateProjectDoc)             // 新建ProjectDoc
		doc_infoRouter.DELETE("deleteProjectDoc", doc_infoApi.DeleteProjectDoc)           // 删除ProjectDoc
		doc_infoRouter.DELETE("deleteProjectDocByIds", doc_infoApi.DeleteProjectDocByIds) // 批量删除ProjectDoc
		doc_infoRouter.PUT("updateProjectDoc", doc_infoApi.UpdateProjectDoc)              // 更新ProjectDoc
	}
	{
		doc_infoRouterWithoutRecord.GET("findProjectDoc", doc_infoApi.FindProjectDoc)        // 根据ID获取ProjectDoc
		doc_infoRouterWithoutRecord.POST("getProjectDocList", doc_infoApi.GetProjectDocList) // 获取ProjectDoc列表
	}
}
