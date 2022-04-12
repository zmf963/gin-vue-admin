/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-17 21:36:08
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 20:01:37
 * @FilePath: /api/v1/project/enter.go
 * @Description:
 */
package project

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ProjectInfoApi
	ProjectDocApi
	TaskApi
	TargetApi
	TargetRelationApi
	DomainApi
	PortInfoApi
	PathInfoApi
	EmailInfoApi
	DocInfoApi
	KeysApi
	WechatOfficialAccountApi
	AppInfoApi
}

var (
	project_infoService            = service.ServiceGroupApp.ProjectServiceGroup.ProjectInfoService
	project_docService             = service.ServiceGroupApp.ProjectServiceGroup.ProjectDocService
	taskService                    = service.ServiceGroupApp.ProjectServiceGroup.TaskService
	targetService                  = service.ServiceGroupApp.ProjectServiceGroup.TargetService
	target_relationService         = service.ServiceGroupApp.ProjectServiceGroup.TargetRelationService
	domainService                  = service.ServiceGroupApp.ProjectServiceGroup.DomainService
	port_infoService               = service.ServiceGroupApp.ProjectServiceGroup.PortInfoService
	path_infoService               = service.ServiceGroupApp.ProjectServiceGroup.PathInfoService
	email_infoService              = service.ServiceGroupApp.ProjectServiceGroup.EmailInfoService
	doc_infoService                = service.ServiceGroupApp.ProjectServiceGroup.DocInfoService
	keysService                    = service.ServiceGroupApp.ProjectServiceGroup.KeysService
	wechat_official_accountService = service.ServiceGroupApp.ProjectServiceGroup.WechatOfficialAccountService
	app_infoService                = service.ServiceGroupApp.ProjectServiceGroup.AppInfoService
)
