/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-17 13:36:08
 * @LastEditors: zmf96
 * @LastEditTime: 2022-03-08 06:14:37
 * @FilePath: /model/project/pro_app_info.go
 * @Description:
 */
// 自动生成模板project
package project

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AppInfo 结构体
type AppInfo struct {
	ID_ primitive.ObjectID `json:"_id" bson:"_id"`

	Name      string `json:"name" bson:"name"`
	Link      string `json:"link" bson:"link"`
	Desc      string `json:"desc" bson:"desc"`
	Developer string `json:"developer" bson:"developer"`
	Source    string `json:"source" bson:"source"`
	TargetId  string `json:"target_id" bson:"target_id"`

	Tags     []string  `json:"tags" bson:"tags"`
	Remarks  string    `json:"remarks" bson:"remarks"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
	DeleteAt time.Time `json:"delete_at" bson:"delete_at"`
}
