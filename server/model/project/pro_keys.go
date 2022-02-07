// 自动生成模板project
package project

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Keys 结构体
type Keys struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	Username        string          `json:"username" bson:"username"`
	Password        string          `json:"password" bson:"password"`
	PasswordType        string          `json:"password_type" bson:"password_type"`
	Content        string          `json:"content" bson:"content"`
	Source        string          `json:"source" bson:"source"`
	TargetId        string          `json:"target_id" bson:"target_id"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}


