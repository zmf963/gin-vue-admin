package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TargetRelationRouter struct {
}

// InitTargetRelationRouter 初始化 TargetRelation 路由信息
func (s *TargetRelationRouter) InitTargetRelationRouter(Router *gin.RouterGroup) {
	target_relationRouter := Router.Group("target_relation").Use(middleware.OperationRecord())
	target_relationRouterWithoutRecord := Router.Group("target_relation")
	var target_relationApi = v1.ApiGroupApp.ProjectApiGroup.TargetRelationApi
	{
		target_relationRouter.POST("createTargetRelation", target_relationApi.CreateTargetRelation)   // 新建TargetRelation
		target_relationRouter.DELETE("deleteTargetRelation", target_relationApi.DeleteTargetRelation) // 删除TargetRelation
		target_relationRouter.DELETE("deleteTargetRelationByIds", target_relationApi.DeleteTargetRelationByIds) // 批量删除TargetRelation
		target_relationRouter.PUT("updateTargetRelation", target_relationApi.UpdateTargetRelation)    // 更新TargetRelation
	}
	{
		target_relationRouterWithoutRecord.GET("findTargetRelation", target_relationApi.FindTargetRelation)        // 根据ID获取TargetRelation
		target_relationRouterWithoutRecord.POST("getTargetRelationList", target_relationApi.GetTargetRelationList)  // 获取TargetRelation列表
	}
}
