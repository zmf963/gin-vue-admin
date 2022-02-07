package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TargetRelationSearch struct{
    project.TargetRelation
    request.PageInfo
    OrderKey string
    Desc bool
}