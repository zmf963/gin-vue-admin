package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WechatOfficialAccountSearch struct{
    project.WechatOfficialAccount
    request.PageInfo
    OrderKey string
    Desc bool
}