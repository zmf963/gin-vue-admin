// 自动生成模板project
package project

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WechatOfficialAccount 结构体
type WechatOfficialAccount struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	WechatName        string          `json:"wechat_name" bson:"wechat_name"`
	WechatNumber        string          `json:"wechat_number" bson:"wechat_number"`
	AccountPrincipal        string          `json:"account_principal" bson:"account_principal"`
	Link        string          `json:"link" bson:"link"`
	Desc        string          `json:"desc" bson:"desc"`
	Source        string          `json:"source" bson:"source"`
	TargetId        string          `json:"target_id" bson:"target_id"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}


