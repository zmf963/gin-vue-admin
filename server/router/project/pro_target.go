package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TargetRouter struct {
}

// InitTargetRouter 初始化 Target 路由信息
func (s *TargetRouter) InitTargetRouter(Router *gin.RouterGroup) {
	targetRouter := Router.Group("target").Use(middleware.OperationRecord())
	targetRouterWithoutRecord := Router.Group("target")
	var targetApi = v1.ApiGroupApp.ProjectApiGroup.TargetApi
	{
		targetRouter.POST("createTarget", targetApi.CreateTarget)   // 新建Target
		targetRouter.DELETE("deleteTarget", targetApi.DeleteTarget) // 删除Target
		targetRouter.DELETE("deleteTargetByIds", targetApi.DeleteTargetByIds) // 批量删除Target
		targetRouter.PUT("updateTarget", targetApi.UpdateTarget)    // 更新Target
	}
	{
		targetRouterWithoutRecord.GET("findTarget", targetApi.FindTarget)        // 根据ID获取Target
		targetRouterWithoutRecord.POST("getTargetList", targetApi.GetTargetList)  // 获取Target列表
	}
}
