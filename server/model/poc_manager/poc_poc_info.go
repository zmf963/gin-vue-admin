// 自动生成模板poc_manager
package poc_manager

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PocInfo 结构体
type PocInfo struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	PocName        string          `json:"poc_name" bson:"poc_name"`
	PocType        string          `json:"poc_type" bson:"poc_type"`
	PocLevel        string          `json:"poc_level" bson:"poc_level"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}


