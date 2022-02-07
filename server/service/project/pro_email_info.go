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

type EmailInfoService struct {
}

var EmailInfoServiceApp = new(EmailInfoService)


//@author:
//@function: CreateEmailInfo
//@description: 新增EmailInfo
//@param: email_info model.EmailInfo
//@return: err error
func (email_infoService *EmailInfoService) CreateEmailInfo(email_info project.EmailInfo) (err error) {
	global.GVA_LOG.Info("[ CreateEmailInfo]", zap.Any("email_info", email_info))
	email_info.ID_ = primitive.NewObjectID()
	email_info.CreateAt = time.Now().Local()
	email_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_email_info").InsertOne(context.TODO(), email_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteEmailInfo
//@description: 删除EmailInfo
//@param: id
func (email_infoService *EmailInfoService)DeleteEmailInfo(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteEmailInfo]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_email_info").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteEmailInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteEmailInfoByIds
//@description: 批量删除EmailInfo
//@param: ids []string
//@return: err error
func (email_infoService *EmailInfoService)DeleteEmailInfoByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteEmailInfoByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_email_info").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteEmailInfoByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateEmailInfo
//@description: 根据id更新EmailInfo
//@param: email_info project.EmailInfo
//@return: err error
func (email_infoService *EmailInfoService)UpdateEmailInfo(email_info project.EmailInfo) (err error) {
	global.GVA_LOG.Debug("[UpdateEmailInfo]", zap.Any("email_info", email_info))
	email_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_email_info").UpdateOne(context.TODO(), bson.M{"_id": email_info.ID_}, bson.M{"$set": email_info})
	global.GVA_LOG.Debug("[UpdateEmailInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetEmailInfoById
//@description: 根据id获取EmailInfo
//@param: _id primitive.ObjectID
//@return: err error,  email_info project.EmailInfo
func (email_infoService *EmailInfoService) GetEmailInfoById(_id primitive.ObjectID) (ema project.EmailInfo, err error) {
	global.GVA_LOG.Debug("[GetEmailInfoById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_email_info").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&ema)
	global.GVA_LOG.Debug("[GetEmailInfoById]", zap.Any("ema", ema), zap.Any("err", err))
	return ema, err
}


//@author:
//@function: GetEmailInfoInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (email_infoService *EmailInfoService) GetEmailInfoInfoList(ema project.EmailInfo, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.EmailInfo
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
	
	if ema.Email != "" {
		filter["email"] = bson.M{"$regex": ema.Email }
	}
	if ema.Source != "" {
		filter["source"] = bson.M{"$regex": ema.Source }
	}
	if ema.TargetId != "" {
		filter["target_id"] = bson.M{"$regex": ema.TargetId }
	}

	if len(ema.Tags) > 0 {
		filter["tags"] = bson.M{"$in":ema.Tags }
	}
	if ema.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":ema.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_email_info").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetEmailInfoInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_email_info").Find(context.TODO(), filter, findOptions)
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