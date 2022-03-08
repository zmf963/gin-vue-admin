/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-01 21:23:50
 * @LastEditors: zmf96
 * @LastEditTime: 2022-03-07 09:52:12
 * @FilePath: /server/model/project/pro_task.go
 * @Description:
 */
// 自动生成模板project
package project

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task 结构体
type Task struct {
	ID_ primitive.ObjectID `json:"_id" bson:"_id"`

	TaskName      string                 `json:"task_name" bson:"task_name"`
	Hosts         string                 `json:"hosts" bson:"hosts"`
	Ports         string                 `json:"ports" bson:"ports"`
	Keyword       string                 `json:"keyword" bson:"keyword"`
	Tools         []string               `json:"tools" bson:"tools"`
	ToolExt       map[string]interface{} `json:"tool_ext" bson:"tool_ext"`
	Status        string                 `json:"status" bson:"status"`
	ProjectId     string                 `json:"project_id" bson:"project_id"`
	TargetId      string                 `json:"target_id" bson:"target_id"`
	CeleryTaskIds []string               `json:"celery_task_ids" bson:"celery_task_ids"`

	Tags     []string  `json:"tags" bson:"tags"`
	Remarks  string    `json:"remarks" bson:"remarks"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
	DeleteAt time.Time `json:"delete_at" bson:"delete_at"`
}
