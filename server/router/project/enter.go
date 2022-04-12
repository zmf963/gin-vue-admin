/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-17 21:36:08
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 19:58:48
 * @FilePath: /router/project/enter.go
 * @Description:
 */
package project

type RouterGroup struct {
	ProjectInfoRouter
	ProjectDocRouter
	TaskRouter
	TargetRouter
	TargetRelationRouter
	DomainRouter
	PortInfoRouter
	PathInfoRouter
	EmailInfoRouter
	DocInfoRouter
	KeysRouter
	WechatOfficialAccountRouter
	AppInfoRouter
}
