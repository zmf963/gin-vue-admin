/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-04-12 19:04:32
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 19:04:57
 * @FilePath: /model/project/request/pro_project_doc.go
 * @Description:
 */
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
)

type ProjectDocSearch struct {
	project.ProjectDoc
	request.PageInfo
	OrderKey string
	Desc     bool
}
