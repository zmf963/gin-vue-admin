package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WechatOfficialAccountRouter struct {
}

// InitWechatOfficialAccountRouter 初始化 WechatOfficialAccount 路由信息
func (s *WechatOfficialAccountRouter) InitWechatOfficialAccountRouter(Router *gin.RouterGroup) {
	wechat_official_accountRouter := Router.Group("wechat_official_account").Use(middleware.OperationRecord())
	wechat_official_accountRouterWithoutRecord := Router.Group("wechat_official_account")
	var wechat_official_accountApi = v1.ApiGroupApp.ProjectApiGroup.WechatOfficialAccountApi
	{
		wechat_official_accountRouter.POST("createWechatOfficialAccount", wechat_official_accountApi.CreateWechatOfficialAccount)   // 新建WechatOfficialAccount
		wechat_official_accountRouter.DELETE("deleteWechatOfficialAccount", wechat_official_accountApi.DeleteWechatOfficialAccount) // 删除WechatOfficialAccount
		wechat_official_accountRouter.DELETE("deleteWechatOfficialAccountByIds", wechat_official_accountApi.DeleteWechatOfficialAccountByIds) // 批量删除WechatOfficialAccount
		wechat_official_accountRouter.PUT("updateWechatOfficialAccount", wechat_official_accountApi.UpdateWechatOfficialAccount)    // 更新WechatOfficialAccount
	}
	{
		wechat_official_accountRouterWithoutRecord.GET("findWechatOfficialAccount", wechat_official_accountApi.FindWechatOfficialAccount)        // 根据ID获取WechatOfficialAccount
		wechat_official_accountRouterWithoutRecord.POST("getWechatOfficialAccountList", wechat_official_accountApi.GetWechatOfficialAccountList)  // 获取WechatOfficialAccount列表
	}
}
