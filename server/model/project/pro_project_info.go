/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-17 13:36:08
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-08 23:25:41
 * @FilePath: /server/model/project/pro_project_info.go
 * @Description:
 */
// 自动生成模板project
package project

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProjectInfo 结构体
type ProjectInfo struct {
	ID_ primitive.ObjectID `json:"_id" bson:"_id"`

	ProjectName string               `json:"project_name" bson:"project_name"`
	ProjectDesc string               `json:"project_desc" bson:"project_desc"`
	StartTime   time.Time            `json:"start_time" bson:"start_time"`
	EndTime     time.Time            `json:"end_time" bson:"end_time"`
	TargetIds   []primitive.ObjectID `json:"target_ids" bson:"target_ids"`
	TaskIds     []string             `json:"task_ids" bson:"task_ids"`

	Tags     []string  `json:"tags" bson:"tags"`
	Remarks  string    `json:"remarks" bson:"remarks"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
	DeleteAt time.Time `json:"delete_at" bson:"delete_at"`
}
