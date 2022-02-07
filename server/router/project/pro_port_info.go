package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PortInfoRouter struct {
}

// InitPortInfoRouter 初始化 PortInfo 路由信息
func (s *PortInfoRouter) InitPortInfoRouter(Router *gin.RouterGroup) {
	port_infoRouter := Router.Group("port_info").Use(middleware.OperationRecord())
	port_infoRouterWithoutRecord := Router.Group("port_info")
	var port_infoApi = v1.ApiGroupApp.ProjectApiGroup.PortInfoApi
	{
		port_infoRouter.POST("createPortInfo", port_infoApi.CreatePortInfo)   // 新建PortInfo
		port_infoRouter.DELETE("deletePortInfo", port_infoApi.DeletePortInfo) // 删除PortInfo
		port_infoRouter.DELETE("deletePortInfoByIds", port_infoApi.DeletePortInfoByIds) // 批量删除PortInfo
		port_infoRouter.PUT("updatePortInfo", port_infoApi.UpdatePortInfo)    // 更新PortInfo
	}
	{
		port_infoRouterWithoutRecord.GET("findPortInfo", port_infoApi.FindPortInfo)        // 根据ID获取PortInfo
		port_infoRouterWithoutRecord.POST("getPortInfoList", port_infoApi.GetPortInfoList)  // 获取PortInfo列表
	}
}
