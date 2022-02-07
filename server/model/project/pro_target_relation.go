// 自动生成模板project
package project

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TargetRelation 结构体
type TargetRelation struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	TargetName        string          `json:"target_name" bson:"target_name"`
	RegRate        float64          `json:"reg_rate" bson:"reg_rate"`
	ParentTargetId        string          `json:"parent_target_id" bson:"parent_target_id"`
	TargetId        string          `json:"target_id" bson:"target_id"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}


