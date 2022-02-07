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

type WechatOfficialAccountService struct {
}

var WechatOfficialAccountServiceApp = new(WechatOfficialAccountService)


//@author:
//@function: CreateWechatOfficialAccount
//@description: 新增WechatOfficialAccount
//@param: wechat_official_account model.WechatOfficialAccount
//@return: err error
func (wechat_official_accountService *WechatOfficialAccountService) CreateWechatOfficialAccount(wechat_official_account project.WechatOfficialAccount) (err error) {
	global.GVA_LOG.Info("[ CreateWechatOfficialAccount]", zap.Any("wechat_official_account", wechat_official_account))
	wechat_official_account.ID_ = primitive.NewObjectID()
	wechat_official_account.CreateAt = time.Now().Local()
	wechat_official_account.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_wechat_official_account").InsertOne(context.TODO(), wechat_official_account)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteWechatOfficialAccount
//@description: 删除WechatOfficialAccount
//@param: id
func (wechat_official_accountService *WechatOfficialAccountService)DeleteWechatOfficialAccount(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteWechatOfficialAccount]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_wechat_official_account").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteWechatOfficialAccount]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteWechatOfficialAccountByIds
//@description: 批量删除WechatOfficialAccount
//@param: ids []string
//@return: err error
func (wechat_official_accountService *WechatOfficialAccountService)DeleteWechatOfficialAccountByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteWechatOfficialAccountByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_wechat_official_account").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteWechatOfficialAccountByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateWechatOfficialAccount
//@description: 根据id更新WechatOfficialAccount
//@param: wechat_official_account project.WechatOfficialAccount
//@return: err error
func (wechat_official_accountService *WechatOfficialAccountService)UpdateWechatOfficialAccount(wechat_official_account project.WechatOfficialAccount) (err error) {
	global.GVA_LOG.Debug("[UpdateWechatOfficialAccount]", zap.Any("wechat_official_account", wechat_official_account))
	wechat_official_account.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_wechat_official_account").UpdateOne(context.TODO(), bson.M{"_id": wechat_official_account.ID_}, bson.M{"$set": wechat_official_account})
	global.GVA_LOG.Debug("[UpdateWechatOfficialAccount]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetWechatOfficialAccountById
//@description: 根据id获取WechatOfficialAccount
//@param: _id primitive.ObjectID
//@return: err error,  wechat_official_account project.WechatOfficialAccount
func (wechat_official_accountService *WechatOfficialAccountService) GetWechatOfficialAccountById(_id primitive.ObjectID) (wec project.WechatOfficialAccount, err error) {
	global.GVA_LOG.Debug("[GetWechatOfficialAccountById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_wechat_official_account").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&wec)
	global.GVA_LOG.Debug("[GetWechatOfficialAccountById]", zap.Any("wec", wec), zap.Any("err", err))
	return wec, err
}


//@author:
//@function: GetWechatOfficialAccountInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (wechat_official_accountService *WechatOfficialAccountService) GetWechatOfficialAccountInfoList(wec project.WechatOfficialAccount, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.WechatOfficialAccount
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
	
	if wec.WechatName != "" {
		filter["wechat_name"] = bson.M{"$regex": wec.WechatName }
	}
	if wec.WechatNumber != "" {
		filter["wechat_number"] = bson.M{"$regex": wec.WechatNumber }
	}
	if wec.AccountPrincipal != "" {
		filter["account_principal"] = bson.M{"$regex": wec.AccountPrincipal }
	}
	if wec.Link != "" {
		filter["link"] = bson.M{"$regex": wec.Link }
	}
	if wec.Desc != "" {
		filter["desc"] = bson.M{"$regex": wec.Desc }
	}
	if wec.Source != "" {
		filter["source"] = bson.M{"$regex": wec.Source }
	}
	if wec.TargetId != "" {
		filter["target_id"] = bson.M{"$regex": wec.TargetId }
	}

	if len(wec.Tags) > 0 {
		filter["tags"] = bson.M{"$in":wec.Tags }
	}
	if wec.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":wec.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_wechat_official_account").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetWechatOfficialAccountInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_wechat_official_account").Find(context.TODO(), filter, findOptions)
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