package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DocInfoRouter struct {
}

// InitDocInfoRouter 初始化 DocInfo 路由信息
func (s *DocInfoRouter) InitDocInfoRouter(Router *gin.RouterGroup) {
	doc_infoRouter := Router.Group("doc_info").Use(middleware.OperationRecord())
	doc_infoRouterWithoutRecord := Router.Group("doc_info")
	var doc_infoApi = v1.ApiGroupApp.ProjectApiGroup.DocInfoApi
	{
		doc_infoRouter.POST("createDocInfo", doc_infoApi.CreateDocInfo)   // 新建DocInfo
		doc_infoRouter.DELETE("deleteDocInfo", doc_infoApi.DeleteDocInfo) // 删除DocInfo
		doc_infoRouter.DELETE("deleteDocInfoByIds", doc_infoApi.DeleteDocInfoByIds) // 批量删除DocInfo
		doc_infoRouter.PUT("updateDocInfo", doc_infoApi.UpdateDocInfo)    // 更新DocInfo
	}
	{
		doc_infoRouterWithoutRecord.GET("findDocInfo", doc_infoApi.FindDocInfo)        // 根据ID获取DocInfo
		doc_infoRouterWithoutRecord.POST("getDocInfoList", doc_infoApi.GetDocInfoList)  // 获取DocInfo列表
	}
}
