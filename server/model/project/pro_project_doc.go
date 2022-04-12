/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-04-12 19:02:48
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 19:12:44
 * @FilePath: /model/project/pro_project_doc.go
 * @Description:
 */
// 自动生成模板project
package project

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProjectDoc 结构体
type ProjectDoc struct {
	ID_ primitive.ObjectID `json:"_id" bson:"_id"`

	DocName   string `json:"doc_name" bson:"doc_name"`
	Link      string `json:"link" bson:"link"`
	Source    string `json:"source" bson:"source"`
	ProjectId primitive.ObjectID    `json:"project_id" bson:"project_id"`

	Tags     []string  `json:"tags" bson:"tags"`
	Remarks  string    `json:"remarks" bson:"remarks"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
	DeleteAt time.Time `json:"delete_at" bson:"delete_at"`
}
