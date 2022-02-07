package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DomainSearch struct{
    project.Domain
    request.PageInfo
    OrderKey string
    Desc bool
}