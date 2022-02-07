package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmailInfoRouter struct {
}

// InitEmailInfoRouter 初始化 EmailInfo 路由信息
func (s *EmailInfoRouter) InitEmailInfoRouter(Router *gin.RouterGroup) {
	email_infoRouter := Router.Group("email_info").Use(middleware.OperationRecord())
	email_infoRouterWithoutRecord := Router.Group("email_info")
	var email_infoApi = v1.ApiGroupApp.ProjectApiGroup.EmailInfoApi
	{
		email_infoRouter.POST("createEmailInfo", email_infoApi.CreateEmailInfo)   // 新建EmailInfo
		email_infoRouter.DELETE("deleteEmailInfo", email_infoApi.DeleteEmailInfo) // 删除EmailInfo
		email_infoRouter.DELETE("deleteEmailInfoByIds", email_infoApi.DeleteEmailInfoByIds) // 批量删除EmailInfo
		email_infoRouter.PUT("updateEmailInfo", email_infoApi.UpdateEmailInfo)    // 更新EmailInfo
	}
	{
		email_infoRouterWithoutRecord.GET("findEmailInfo", email_infoApi.FindEmailInfo)        // 根据ID获取EmailInfo
		email_infoRouterWithoutRecord.POST("getEmailInfoList", email_infoApi.GetEmailInfoList)  // 获取EmailInfo列表
	}
}
