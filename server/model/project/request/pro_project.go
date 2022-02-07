package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProjectSearch struct{
    project.Project
    request.PageInfo
    OrderKey string
    Desc bool
}