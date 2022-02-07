package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProjectInfoSearch struct{
    project.ProjectInfo
    request.PageInfo
    OrderKey string
    Desc bool
}