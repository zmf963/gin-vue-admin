// 自动生成模板project
package project

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Target 结构体
type Target struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	TargetName        string          `json:"target_name" bson:"target_name"`
	ProjectIds        string          `json:"project_ids" bson:"project_ids"`
	TaskIds        []string          `json:"task_ids" bson:"task_ids"`
	DomainIds        []string          `json:"domain_ids" bson:"domain_ids"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}

