package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/poc_manager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PocInfoSearch struct{
    poc_manager.PocInfo
    request.PageInfo
    OrderKey string
    Desc bool
}