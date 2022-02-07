// 自动生成模板project
package project

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task 结构体
type Task struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	TaskName        string          `json:"task_name" bson:"task_name"`
	Hosts        string          `json:"hosts" bson:"hosts"`
	Ports        string          `json:"ports" bson:"ports"`
	Keyword        string          `json:"keyword" bson:"keyword"`
	Tools        []string          `json:"tools" bson:"tools"`
	ToolExt        map[string]interface{}          `json:"tool_ext" bson:"tool_ext"`
	Status        string          `json:"status" bson:"status"`
	ProjectId        string          `json:"project_id" bson:"project_id"`
	TargetId        string          `json:"target_id" bson:"target_id"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}


