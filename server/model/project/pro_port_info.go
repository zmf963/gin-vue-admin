/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-17 21:36:08
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-01 21:05:58
 * @FilePath: /server/model/project/pro_port_info.go
 * @Description:
 */
// 自动生成模板project
package project

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PortInfo 结构体
type PortInfo struct {
	ID_ primitive.ObjectID `json:"_id" bson:"_id"`

	Port       string   `json:"port" bson:"port"`
	Host       string   `json:"host" bson:"host"`
	Hostinfo   string   `json:"hostinfo" bson:"hostinfo"`
	Url        string   `json:"url" bson:"url"`
	Title      string   `json:"title" bson:"title"`
	Favicons   string   `json:"favicons" bson:"favicons"`
	Screenshot string   `json:"screenshot" bson:"screenshot"`
	Products   []string `json:"products" bson:"products"`
	Protocols  []string `json:"protocols" bson:"protocols"`
	Alive      int      `json:"alive" bson:"alive"`
	Banner     string   `json:"banner" bson:"banner"`
	Status     string   `json:"status" bson:"status"`
	DomainId   string   `json:"domain_id" bson:"domain_id"`
	PathIds    []string `json:"path_ids" bson:"path_ids"`

	Tags     []string  `json:"tags" bson:"tags"`
	Remarks  string    `json:"remarks" bson:"remarks"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
	DeleteAt time.Time `json:"delete_at" bson:"delete_at"`
}
