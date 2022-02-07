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

type AppInfoService struct {
}

var AppInfoServiceApp = new(AppInfoService)


//@author:
//@function: CreateAppInfo
//@description: 新增AppInfo
//@param: app_info model.AppInfo
//@return: err error
func (app_infoService *AppInfoService) CreateAppInfo(app_info project.AppInfo) (err error) {
	global.GVA_LOG.Info("[ CreateAppInfo]", zap.Any("app_info", app_info))
	app_info.ID_ = primitive.NewObjectID()
	app_info.CreateAt = time.Now().Local()
	app_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_app_info").InsertOne(context.TODO(), app_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteAppInfo
//@description: 删除AppInfo
//@param: id
func (app_infoService *AppInfoService)DeleteAppInfo(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteAppInfo]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_app_info").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteAppInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteAppInfoByIds
//@description: 批量删除AppInfo
//@param: ids []string
//@return: err error
func (app_infoService *AppInfoService)DeleteAppInfoByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteAppInfoByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_app_info").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteAppInfoByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateAppInfo
//@description: 根据id更新AppInfo
//@param: app_info project.AppInfo
//@return: err error
func (app_infoService *AppInfoService)UpdateAppInfo(app_info project.AppInfo) (err error) {
	global.GVA_LOG.Debug("[UpdateAppInfo]", zap.Any("app_info", app_info))
	app_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_app_info").UpdateOne(context.TODO(), bson.M{"_id": app_info.ID_}, bson.M{"$set": app_info})
	global.GVA_LOG.Debug("[UpdateAppInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetAppInfoById
//@description: 根据id获取AppInfo
//@param: _id primitive.ObjectID
//@return: err error,  app_info project.AppInfo
func (app_infoService *AppInfoService) GetAppInfoById(_id primitive.ObjectID) (app project.AppInfo, err error) {
	global.GVA_LOG.Debug("[GetAppInfoById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_app_info").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&app)
	global.GVA_LOG.Debug("[GetAppInfoById]", zap.Any("app", app), zap.Any("err", err))
	return app, err
}


//@author:
//@function: GetAppInfoInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (app_infoService *AppInfoService) GetAppInfoInfoList(app project.AppInfo, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.AppInfo
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
	
	if app.Name != "" {
		filter["name"] = bson.M{"$regex": app.Name }
	}
	if app.Link != "" {
		filter["link"] = bson.M{"$regex": app.Link }
	}
	if app.Desc != "" {
		filter["desc"] = bson.M{"$regex": app.Desc }
	}
	if app.Developer != "" {
		filter["developer"] = bson.M{"$regex": app.Developer }
	}
	if app.Source != "" {
		filter["source"] = bson.M{"$regex": app.Source }
	}
	if app.TargetId != "" {
		filter["target_id"] = bson.M{"$regex": app.TargetId }
	}

	if len(app.Tags) > 0 {
		filter["tags"] = bson.M{"$in":app.Tags }
	}
	if app.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":app.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_app_info").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetAppInfoInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_app_info").Find(context.TODO(), filter, findOptions)
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