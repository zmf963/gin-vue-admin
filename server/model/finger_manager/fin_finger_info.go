/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-17 13:36:08
 * @LastEditors: zmf96
 * @LastEditTime: 2022-03-03 07:56:19
 * @FilePath: /model/finger_manager/fin_finger_info.go
 * @Description:
 */
// 自动生成模板finger_manager
package finger_manager

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FingerInfo 结构体
type FingerInfo struct {
	ID_ primitive.ObjectID `json:"_id" bson:"_id"`

	Name    string            `json:"name" bson:"name"`
	Content map[string]string `json:"content" bson:"content"`
	LinkVul []string          `json:"link_vul" bson:"link_vul"`

	Tags     []string  `json:"tags" bson:"tags"`
	Remarks  string    `json:"remarks" bson:"remarks"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
	DeleteAt time.Time `json:"delete_at" bson:"delete_at"`
}
