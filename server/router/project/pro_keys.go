package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KeysRouter struct {
}

// InitKeysRouter 初始化 Keys 路由信息
func (s *KeysRouter) InitKeysRouter(Router *gin.RouterGroup) {
	keysRouter := Router.Group("keys").Use(middleware.OperationRecord())
	keysRouterWithoutRecord := Router.Group("keys")
	var keysApi = v1.ApiGroupApp.ProjectApiGroup.KeysApi
	{
		keysRouter.POST("createKeys", keysApi.CreateKeys)   // 新建Keys
		keysRouter.DELETE("deleteKeys", keysApi.DeleteKeys) // 删除Keys
		keysRouter.DELETE("deleteKeysByIds", keysApi.DeleteKeysByIds) // 批量删除Keys
		keysRouter.PUT("updateKeys", keysApi.UpdateKeys)    // 更新Keys
	}
	{
		keysRouterWithoutRecord.GET("findKeys", keysApi.FindKeys)        // 根据ID获取Keys
		keysRouterWithoutRecord.POST("getKeysList", keysApi.GetKeysList)  // 获取Keys列表
	}
}
