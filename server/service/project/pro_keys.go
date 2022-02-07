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

type KeysService struct {
}

var KeysServiceApp = new(KeysService)


//@author:
//@function: CreateKeys
//@description: 新增Keys
//@param: keys model.Keys
//@return: err error
func (keysService *KeysService) CreateKeys(keys project.Keys) (err error) {
	global.GVA_LOG.Info("[ CreateKeys]", zap.Any("keys", keys))
	keys.ID_ = primitive.NewObjectID()
	keys.CreateAt = time.Now().Local()
	keys.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_keys").InsertOne(context.TODO(), keys)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteKeys
//@description: 删除Keys
//@param: id
func (keysService *KeysService)DeleteKeys(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteKeys]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_keys").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteKeys]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteKeysByIds
//@description: 批量删除Keys
//@param: ids []string
//@return: err error
func (keysService *KeysService)DeleteKeysByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteKeysByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_keys").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteKeysByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateKeys
//@description: 根据id更新Keys
//@param: keys project.Keys
//@return: err error
func (keysService *KeysService)UpdateKeys(keys project.Keys) (err error) {
	global.GVA_LOG.Debug("[UpdateKeys]", zap.Any("keys", keys))
	keys.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_keys").UpdateOne(context.TODO(), bson.M{"_id": keys.ID_}, bson.M{"$set": keys})
	global.GVA_LOG.Debug("[UpdateKeys]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetKeysById
//@description: 根据id获取Keys
//@param: _id primitive.ObjectID
//@return: err error,  keys project.Keys
func (keysService *KeysService) GetKeysById(_id primitive.ObjectID) (key project.Keys, err error) {
	global.GVA_LOG.Debug("[GetKeysById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_keys").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&key)
	global.GVA_LOG.Debug("[GetKeysById]", zap.Any("key", key), zap.Any("err", err))
	return key, err
}


//@author:
//@function: GetKeysInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (keysService *KeysService) GetKeysInfoList(key project.Keys, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.Keys
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
	
	if key.Username != "" {
		filter["username"] = bson.M{"$regex": key.Username }
	}
	if key.Password != "" {
		filter["password"] = bson.M{"$regex": key.Password }
	}
	if key.PasswordType != "" {
		filter["password_type"] = bson.M{"$regex": key.PasswordType }
	}
	if key.Content != "" {
		filter["content"] = bson.M{"$regex": key.Content }
	}
	if key.Source != "" {
		filter["source"] = bson.M{"$regex": key.Source }
	}
	if key.TargetId != "" {
		filter["target_id"] = bson.M{"$regex": key.TargetId }
	}

	if len(key.Tags) > 0 {
		filter["tags"] = bson.M{"$in":key.Tags }
	}
	if key.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":key.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_keys").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetKeysInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_keys").Find(context.TODO(), filter, findOptions)
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