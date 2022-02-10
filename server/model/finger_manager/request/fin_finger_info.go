package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/finger_manager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FingerInfoSearch struct{
    finger_manager.FingerInfo
    request.PageInfo
    OrderKey string
    Desc bool
}