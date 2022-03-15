package finger_manager

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/finger_manager"
	"time"
	//finger_managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/finger_manager/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type FingerInfoService struct {
}

var FingerInfoServiceApp = new(FingerInfoService)

//@author:
//@function: CreateFingerInfo
//@description: 新增FingerInfo
//@param: finger_info model.FingerInfo
//@return: err error
func (finger_infoService *FingerInfoService) CreateFingerInfo(finger_info finger_manager.FingerInfo) (err error) {
	global.GVA_LOG.Info("[ CreateFingerInfo]", zap.Any("finger_info", finger_info))
	if finger_info.ID_.IsZero() {
		finger_info.ID_ = primitive.NewObjectID()
	}
	finger_info.CreateAt = time.Now().Local()
	finger_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("fin_finger_info").InsertOne(context.TODO(), finger_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteFingerInfo
//@description: 删除FingerInfo
//@param: id
func (finger_infoService *FingerInfoService) DeleteFingerInfo(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteFingerInfo]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("fin_finger_info").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteFingerInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteFingerInfoByIds
//@description: 批量删除FingerInfo
//@param: ids []string
//@return: err error
func (finger_infoService *FingerInfoService) DeleteFingerInfoByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteFingerInfoByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("fin_finger_info").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteFingerInfoByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateFingerInfo
//@description: 根据id更新FingerInfo
//@param: finger_info finger_manager.FingerInfo
//@return: err error
func (finger_infoService *FingerInfoService) UpdateFingerInfo(finger_info finger_manager.FingerInfo) (err error) {
	global.GVA_LOG.Debug("[UpdateFingerInfo]", zap.Any("finger_info", finger_info))
	finger_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("fin_finger_info").UpdateOne(context.TODO(), bson.M{"_id": finger_info.ID_}, bson.M{"$set": finger_info})
	global.GVA_LOG.Debug("[UpdateFingerInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: GetFingerInfoById
//@description: 根据id获取FingerInfo
//@param: _id primitive.ObjectID
//@return: err error,  finger_info finger_manager.FingerInfo
func (finger_infoService *FingerInfoService) GetFingerInfoById(_id primitive.ObjectID) (fin finger_manager.FingerInfo, err error) {
	global.GVA_LOG.Debug("[GetFingerInfoById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("fin_finger_info").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&fin)
	global.GVA_LOG.Debug("[GetFingerInfoById]", zap.Any("fin", fin), zap.Any("err", err))
	return fin, err
}

//@author:
//@function: GetFingerInfoInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (finger_infoService *FingerInfoService) GetFingerInfoInfoList(fin finger_manager.FingerInfo, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []finger_manager.FingerInfo
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
	if !fin.ID_.IsZero() {
		filter = bson.M{"_id": fin.ID_}
	}
	if fin.Name != "" {
		filter["name"] = bson.M{"$regex": fin.Name}
	}
	if len(fin.LinkVul) > 0 {
		filter["link_vul"] = bson.M{"$in": fin.LinkVul}
	}

	if len(fin.Tags) > 0 {
		filter["tags"] = bson.M{"$in": fin.Tags}
	}
	if fin.Remarks != "" {
		filter["remarks"] = bson.M{"$regex": fin.Remarks}
	}

	total, err = global.Mongo_DB.Collection("fin_finger_info").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetFingerInfoInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("fin_finger_info").Find(context.TODO(), filter, findOptions)
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
