package project

import (
	"context"
	"encoding/json"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/streadway/amqp"

	//projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type TaskService struct {
}

var TaskServiceApp = new(TaskService)

//@author:
//@function: CreateTask
//@description: 新增Task
//@param: task model.Task
//@return: err error
func (taskService *TaskService) CreateTask(task project.Task) (err error) {
	global.GVA_LOG.Info("[ CreateTask]", zap.Any("task", task))
	task.ID_ = primitive.NewObjectID()
	task.CreateAt = time.Now().Local()
	task.UpdateAt = time.Now().Local()
	body, err := json.Marshal(task)
	if err != nil {
		global.GVA_LOG.Warn("[CreateTask] json.Marshal err: ", zap.Any("err", err))
	}
	err = global.RabbitMQ.Publish(
		"",
		"server:default",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Type:        "task",
			Body:        body,
		})
	if err != nil {
		global.GVA_LOG.Error("[CreateTask] rabbiqmq push err: ", zap.Any("err", err))
		return err
	}
	result, err := global.Mongo_DB.Collection("pro_task").InsertOne(context.TODO(), task)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}

	return nil
}

//@author:
//@function: DeleteTask
//@description: 删除Task
//@param: id
func (taskService *TaskService) DeleteTask(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteTask]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_task").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteTask]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteTaskByIds
//@description: 批量删除Task
//@param: ids []string
//@return: err error
func (taskService *TaskService) DeleteTaskByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteTaskByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_task").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteTaskByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateTask
//@description: 根据id更新Task
//@param: task project.Task
//@return: err error
func (taskService *TaskService) UpdateTask(task project.Task) (err error) {
	global.GVA_LOG.Debug("[UpdateTask]", zap.Any("task", task))
	task.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_task").UpdateOne(context.TODO(), bson.M{"_id": task.ID_}, bson.M{"$set": task})
	global.GVA_LOG.Debug("[UpdateTask]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: GetTaskById
//@description: 根据id获取Task
//@param: _id primitive.ObjectID
//@return: err error,  task project.Task
func (taskService *TaskService) GetTaskById(_id primitive.ObjectID) (tas project.Task, err error) {
	global.GVA_LOG.Debug("[GetTaskById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_task").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&tas)
	global.GVA_LOG.Debug("[GetTaskById]", zap.Any("tas", tas), zap.Any("err", err))
	return tas, err
}

//@author:
//@function: GetTaskInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (taskService *TaskService) GetTaskInfoList(tas project.Task, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.Task
	var findOptions = options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))
	if order != "" {
		_desc := 1
		if !desc {
			_desc = -1
		}
		findOptions.SetSort(map[string]int{order: _desc})
	}
	filter := bson.M{}
	// TODO 处理搜索条件

	if tas.TaskName != "" {
		filter["task_name"] = bson.M{"$regex": tas.TaskName}
	}
	if tas.Hosts != "" {
		filter["hosts"] = bson.M{"$regex": tas.Hosts}
	}
	if tas.Ports != "" {
		filter["ports"] = bson.M{"$regex": tas.Ports}
	}
	if tas.Keyword != "" {
		filter["keyword"] = bson.M{"$regex": tas.Keyword}
	}
	if len(tas.Tools) > 0 {
		filter["tools"] = bson.M{"$in": tas.Tools}
	}
	if tas.Status != "" {
		filter["status"] = bson.M{"$regex": tas.Status}
	}
	if tas.ProjectId != "" {
		filter["project_id"] = bson.M{"$regex": tas.ProjectId}
	}
	if tas.TargetId != "" {
		filter["target_id"] = bson.M{"$regex": tas.TargetId}
	}

	if len(tas.Tags) > 0 {
		filter["tags"] = bson.M{"$in": tas.Tags}
	}
	if tas.Remarks != "" {
		filter["remarks"] = bson.M{"$regex": tas.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_task").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetTaskInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_task").Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(context.TODO())
	err = cur.All(context.TODO(), &retList)
	if err != nil {
		return nil, 0, err
	}
	return retList, total, nil
}
