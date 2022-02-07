// 自动生成模板project
package project

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Domain 结构体
type Domain struct {
	ID_             primitive.ObjectID  `json:"_id" bson:"_id"`
	
	Domain        string          `json:"domain" bson:"domain"`
	Ips        []string          `json:"ips" bson:"ips"`
	Hostnames        []string          `json:"hostnames" bson:"hostnames"`
	Os        string          `json:"os" bson:"os"`
	Whois        map[string]interface{}          `json:"whois" bson:"whois"`
	Alive        int          `json:"alive" bson:"alive"`
	Cname        string          `json:"cname" bson:"cname"`
	Cdn        int          `json:"cdn" bson:"cdn"`
	Cidr        string          `json:"cidr" bson:"cidr"`
	Asn        string          `json:"asn" bson:"asn"`
	Org        string          `json:"org" bson:"org"`
	Addr        string          `json:"addr" bson:"addr"`
	Isp        string          `json:"isp" bson:"isp"`
	Source        string          `json:"source" bson:"source"`
	TargetId        string          `json:"target_id" bson:"target_id"`
	TargetIdIsVerify        bool          `json:"target_id_is_verify" bson:"target_id_is_verify"`
	PortIds        []string          `json:"port_ids" bson:"port_ids"`

	Tags            []string               `json:"tags" bson:"tags"`
	Remarks         string              `json:"remarks" bson:"remarks"`
	CreateAt        time.Time              `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time              `json:"update_at" bson:"update_at"`
	DeleteAt        time.Time              `json:"delete_at" bson:"delete_at"`
}


