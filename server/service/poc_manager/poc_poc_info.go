package poc_manager

import (
	"context"
	"time"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/poc_manager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	//poc_managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/poc_manager/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type PocInfoService struct {
}

var PocInfoServiceApp = new(PocInfoService)


//@author:
//@function: CreatePocInfo
//@description: 新增PocInfo
//@param: poc_info model.PocInfo
//@return: err error
func (poc_infoService *PocInfoService) CreatePocInfo(poc_info poc_manager.PocInfo) (err error) {
	global.GVA_LOG.Info("[ CreatePocInfo]", zap.Any("poc_info", poc_info))
	poc_info.ID_ = primitive.NewObjectID()
	poc_info.CreateAt = time.Now().Local()
	poc_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("poc_poc_info").InsertOne(context.TODO(), poc_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeletePocInfo
//@description: 删除PocInfo
//@param: id
func (poc_infoService *PocInfoService)DeletePocInfo(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeletePocInfo]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("poc_poc_info").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeletePocInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeletePocInfoByIds
//@description: 批量删除PocInfo
//@param: ids []string
//@return: err error
func (poc_infoService *PocInfoService)DeletePocInfoByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeletePocInfoByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("poc_poc_info").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeletePocInfoByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdatePocInfo
//@description: 根据id更新PocInfo
//@param: poc_info poc_manager.PocInfo
//@return: err error
func (poc_infoService *PocInfoService)UpdatePocInfo(poc_info poc_manager.PocInfo) (err error) {
	global.GVA_LOG.Debug("[UpdatePocInfo]", zap.Any("poc_info", poc_info))
	poc_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("poc_poc_info").UpdateOne(context.TODO(), bson.M{"_id": poc_info.ID_}, bson.M{"$set": poc_info})
	global.GVA_LOG.Debug("[UpdatePocInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetPocInfoById
//@description: 根据id获取PocInfo
//@param: _id primitive.ObjectID
//@return: err error,  poc_info poc_manager.PocInfo
func (poc_infoService *PocInfoService) GetPocInfoById(_id primitive.ObjectID) (poc poc_manager.PocInfo, err error) {
	global.GVA_LOG.Debug("[GetPocInfoById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("poc_poc_info").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&poc)
	global.GVA_LOG.Debug("[GetPocInfoById]", zap.Any("poc", poc), zap.Any("err", err))
	return poc, err
}


//@author:
//@function: GetPocInfoInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (poc_infoService *PocInfoService) GetPocInfoInfoList(poc poc_manager.PocInfo, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []poc_manager.PocInfo
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
	
	if poc.VulName != "" {
		filter["vul_name"] = bson.M{"$regex": poc.VulName }
	}
	if poc.VulType != "" {
		filter["vul_type"] = bson.M{"$regex": poc.VulType }
	}
	if poc.VulId != "" {
		filter["vul_id"] = bson.M{"$regex": poc.VulId }
	}
	if poc.VulManufacturer != "" {
		filter["vul_manufacturer"] = bson.M{"$regex": poc.VulManufacturer }
	}
	if poc.VulSystem != "" {
		filter["vul_system"] = bson.M{"$regex": poc.VulSystem }
	}
	if poc.Language != "" {
		filter["language"] = bson.M{"$regex": poc.Language }
	}
	if poc.Version != "" {
		filter["version"] = bson.M{"$regex": poc.Version }
	}
	if len(poc.LinkUrl) > 0 {
		filter["link_url"] = bson.M{"$in": poc.LinkUrl }
	}
	if poc.PocType != "" {
		filter["poc_type"] = bson.M{"$regex": poc.PocType }
	}
	if poc.PocContent != "" {
		filter["poc_content"] = bson.M{"$regex": poc.PocContent }
	}
	if len(poc.PocArgs) > 0 {
		filter["poc_args"] = bson.M{"$in": poc.PocArgs }
	}
	if poc.VulFingerId != "" {
		filter["vul_finger_id"] = bson.M{"$regex": poc.VulFingerId }
	}

	if len(poc.Tags) > 0 {
		filter["tags"] = bson.M{"$in":poc.Tags }
	}
	if poc.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":poc.Remarks}
	}

	total, err = global.Mongo_DB.Collection("poc_poc_info").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetPocInfoInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("poc_poc_info").Find(context.TODO(), filter, findOptions)
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