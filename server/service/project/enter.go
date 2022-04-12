/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-17 21:36:08
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 19:05:52
 * @FilePath: /service/project/enter.go
 * @Description:
 */
package project

type ServiceGroup struct {
	ProjectInfoService
	ProjectDocService
	TaskService
	TargetService
	TargetRelationService
	DomainService
	PortInfoService
	PathInfoService
	EmailInfoService
	DocInfoService
	KeysService
	WechatOfficialAccountService
	AppInfoService
}
