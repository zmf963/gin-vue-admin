package poc_manager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PocInfoRouter struct {
}

// InitPocInfoRouter 初始化 PocInfo 路由信息
func (s *PocInfoRouter) InitPocInfoRouter(Router *gin.RouterGroup) {
	poc_infoRouter := Router.Group("poc_info").Use(middleware.OperationRecord())
	poc_infoRouterWithoutRecord := Router.Group("poc_info")
	var poc_infoApi = v1.ApiGroupApp.Poc_managerApiGroup.PocInfoApi
	{
		poc_infoRouter.POST("createPocInfo", poc_infoApi.CreatePocInfo)   // 新建PocInfo
		poc_infoRouter.DELETE("deletePocInfo", poc_infoApi.DeletePocInfo) // 删除PocInfo
		poc_infoRouter.DELETE("deletePocInfoByIds", poc_infoApi.DeletePocInfoByIds) // 批量删除PocInfo
		poc_infoRouter.PUT("updatePocInfo", poc_infoApi.UpdatePocInfo)    // 更新PocInfo
	}
	{
		poc_infoRouterWithoutRecord.GET("findPocInfo", poc_infoApi.FindPocInfo)        // 根据ID获取PocInfo
		poc_infoRouterWithoutRecord.POST("getPocInfoList", poc_infoApi.GetPocInfoList)  // 获取PocInfo列表
	}
}
