package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TaskSearch struct{
    project.Task
    request.PageInfo
    OrderKey string
    Desc bool
}