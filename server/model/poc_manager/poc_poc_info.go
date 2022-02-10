// 自动生成模板poc_manager
package poc_manager

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PocInfo 结构体
type PocInfo struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	VulName        string          `json:"vul_name" bson:"vul_name"`
	VulType        string          `json:"vul_type" bson:"vul_type"`
	IsLogin        bool          `json:"is_login" bson:"is_login"`
	VulId        string          `json:"vul_id" bson:"vul_id"`
	VulManufacturer        string          `json:"vul_manufacturer" bson:"vul_manufacturer"`
	VulSystem        string          `json:"vul_system" bson:"vul_system"`
	Language        string          `json:"language" bson:"language"`
	Version        string          `json:"version" bson:"version"`
	LinkUrl        []string          `json:"link_url" bson:"link_url"`
	PocType        string          `json:"poc_type" bson:"poc_type"`
	PocContent        string          `json:"poc_content" bson:"poc_content"`
	PocArgs        []string          `json:"poc_args" bson:"poc_args"`
	VulFingerId        string          `json:"vul_finger_id" bson:"vul_finger_id"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}


