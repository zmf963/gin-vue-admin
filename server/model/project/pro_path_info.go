// 自动生成模板project
package project

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PathInfo 结构体
type PathInfo struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	Hostinfo        string          `json:"hostinfo" bson:"hostinfo"`
	Path        string          `json:"path" bson:"path"`
	Url        string          `json:"url" bson:"url"`
	Title        string          `json:"title" bson:"title"`
	Status        string          `json:"status" bson:"status"`
	ResponseSize        int64          `json:"response_size" bson:"response_size"`
	PortId        string          `json:"port_id" bson:"port_id"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}


