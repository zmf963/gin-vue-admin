package finger_manager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FingerInfoRouter struct {
}

// InitFingerInfoRouter 初始化 FingerInfo 路由信息
func (s *FingerInfoRouter) InitFingerInfoRouter(Router *gin.RouterGroup) {
	finger_infoRouter := Router.Group("finger_info").Use(middleware.OperationRecord())
	finger_infoRouterWithoutRecord := Router.Group("finger_info")
	var finger_infoApi = v1.ApiGroupApp.Finger_managerApiGroup.FingerInfoApi
	{
		finger_infoRouter.POST("createFingerInfo", finger_infoApi.CreateFingerInfo)   // 新建FingerInfo
		finger_infoRouter.DELETE("deleteFingerInfo", finger_infoApi.DeleteFingerInfo) // 删除FingerInfo
		finger_infoRouter.DELETE("deleteFingerInfoByIds", finger_infoApi.DeleteFingerInfoByIds) // 批量删除FingerInfo
		finger_infoRouter.PUT("updateFingerInfo", finger_infoApi.UpdateFingerInfo)    // 更新FingerInfo
	}
	{
		finger_infoRouterWithoutRecord.GET("findFingerInfo", finger_infoApi.FindFingerInfo)        // 根据ID获取FingerInfo
		finger_infoRouterWithoutRecord.POST("getFingerInfoList", finger_infoApi.GetFingerInfoList)  // 获取FingerInfo列表
	}
}
