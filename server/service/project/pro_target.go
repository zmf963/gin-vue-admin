package project

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"time"
	//projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type TargetService struct {
}

var TargetServiceApp = new(TargetService)

//@author:
//@function: CreateTarget
//@description: 新增Target
//@param: target model.Target
//@return: err error
func (targetService *TargetService) CreateTarget(target project.Target) (err error) {
	global.GVA_LOG.Info("[ CreateTarget]", zap.Any("target", target))
	target.ID_ = primitive.NewObjectID()
	target.CreateAt = time.Now().Local()
	target.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_target").InsertOne(context.TODO(), target)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteTarget
//@description: 删除Target
//@param: id
func (targetService *TargetService) DeleteTarget(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteTarget]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_target").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteTarget]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteTargetByIds
//@description: 批量删除Target
//@param: ids []string
//@return: err error
func (targetService *TargetService) DeleteTargetByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteTargetByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_target").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteTargetByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateTarget
//@description: 根据id更新Target
//@param: target project.Target
//@return: err error
func (targetService *TargetService) UpdateTarget(target project.Target) (err error) {
	global.GVA_LOG.Debug("[UpdateTarget]", zap.Any("target", target))
	target.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_target").UpdateOne(context.TODO(), bson.M{"_id": target.ID_}, bson.M{"$set": target})
	global.GVA_LOG.Debug("[UpdateTarget]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: GetTargetById
//@description: 根据id获取Target
//@param: _id primitive.ObjectID
//@return: err error,  target project.Target
func (targetService *TargetService) GetTargetById(_id primitive.ObjectID) (tar project.Target, err error) {
	global.GVA_LOG.Debug("[GetTargetById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_target").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&tar)
	global.GVA_LOG.Debug("[GetTargetById]", zap.Any("tar", tar), zap.Any("err", err))
	return tar, err
}

//@author:
//@function: GetTargetInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (targetService *TargetService) GetTargetInfoList(tar project.Target, project_id primitive.ObjectID, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.Target
	var findOptions = options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))
	if order != "" {
		_desc := 1
		if !desc {
			_desc = -1
		}
		findOptions.SetSort(map[string]int{order: _desc})
	} else {
		findOptions.SetSort(map[string]int{"update_at": -1})
	}
	filter := bson.M{}
	// TODO 处理搜索条件

	if tar.TargetName != "" {
		filter["target_name"] = bson.M{"$regex": tar.TargetName}
	}
	if len(tar.DomainList) > 0 {
		filter["domain_list"] = bson.M{"$in": tar.DomainList}
	}

	if len(tar.Tags) > 0 {
		filter["tags"] = bson.M{"$in": tar.Tags}
	}
	if tar.Remarks != "" {
		filter["remarks"] = bson.M{"$regex": tar.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_target").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetTargetInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_target").Find(context.TODO(), filter, findOptions)
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
