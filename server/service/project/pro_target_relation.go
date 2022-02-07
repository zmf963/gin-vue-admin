package project

import (
	"context"
	"time"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	//projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type TargetRelationService struct {
}

var TargetRelationServiceApp = new(TargetRelationService)


//@author:
//@function: CreateTargetRelation
//@description: 新增TargetRelation
//@param: target_relation model.TargetRelation
//@return: err error
func (target_relationService *TargetRelationService) CreateTargetRelation(target_relation project.TargetRelation) (err error) {
	global.GVA_LOG.Info("[ CreateTargetRelation]", zap.Any("target_relation", target_relation))
	target_relation.ID_ = primitive.NewObjectID()
	target_relation.CreateAt = time.Now().Local()
	target_relation.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_target_relation").InsertOne(context.TODO(), target_relation)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteTargetRelation
//@description: 删除TargetRelation
//@param: id
func (target_relationService *TargetRelationService)DeleteTargetRelation(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteTargetRelation]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_target_relation").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteTargetRelation]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteTargetRelationByIds
//@description: 批量删除TargetRelation
//@param: ids []string
//@return: err error
func (target_relationService *TargetRelationService)DeleteTargetRelationByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteTargetRelationByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_target_relation").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteTargetRelationByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateTargetRelation
//@description: 根据id更新TargetRelation
//@param: target_relation project.TargetRelation
//@return: err error
func (target_relationService *TargetRelationService)UpdateTargetRelation(target_relation project.TargetRelation) (err error) {
	global.GVA_LOG.Debug("[UpdateTargetRelation]", zap.Any("target_relation", target_relation))
	target_relation.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_target_relation").UpdateOne(context.TODO(), bson.M{"_id": target_relation.ID_}, bson.M{"$set": target_relation})
	global.GVA_LOG.Debug("[UpdateTargetRelation]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetTargetRelationById
//@description: 根据id获取TargetRelation
//@param: _id primitive.ObjectID
//@return: err error,  target_relation project.TargetRelation
func (target_relationService *TargetRelationService) GetTargetRelationById(_id primitive.ObjectID) (tar project.TargetRelation, err error) {
	global.GVA_LOG.Debug("[GetTargetRelationById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_target_relation").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&tar)
	global.GVA_LOG.Debug("[GetTargetRelationById]", zap.Any("tar", tar), zap.Any("err", err))
	return tar, err
}


//@author:
//@function: GetTargetRelationInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (target_relationService *TargetRelationService) GetTargetRelationInfoList(tar project.TargetRelation, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.TargetRelation
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
	
	if tar.TargetName != "" {
		filter["target_name"] = bson.M{"$regex": tar.TargetName }
	}
	if tar.ParentTargetId != "" {
		filter["parent_target_id"] = bson.M{"$regex": tar.ParentTargetId }
	}
	if tar.TargetId != "" {
		filter["target_id"] = bson.M{"$regex": tar.TargetId }
	}

	if len(tar.Tags) > 0 {
		filter["tags"] = bson.M{"$in":tar.Tags }
	}
	if tar.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":tar.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_target_relation").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetTargetRelationInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_target_relation").Find(context.TODO(), filter, findOptions)
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