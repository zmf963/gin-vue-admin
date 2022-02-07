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

type PathInfoService struct {
}

var PathInfoServiceApp = new(PathInfoService)


//@author:
//@function: CreatePathInfo
//@description: 新增PathInfo
//@param: path_info model.PathInfo
//@return: err error
func (path_infoService *PathInfoService) CreatePathInfo(path_info project.PathInfo) (err error) {
	global.GVA_LOG.Info("[ CreatePathInfo]", zap.Any("path_info", path_info))
	path_info.ID_ = primitive.NewObjectID()
	path_info.CreateAt = time.Now().Local()
	path_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_path_info").InsertOne(context.TODO(), path_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeletePathInfo
//@description: 删除PathInfo
//@param: id
func (path_infoService *PathInfoService)DeletePathInfo(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeletePathInfo]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_path_info").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeletePathInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeletePathInfoByIds
//@description: 批量删除PathInfo
//@param: ids []string
//@return: err error
func (path_infoService *PathInfoService)DeletePathInfoByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeletePathInfoByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_path_info").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeletePathInfoByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdatePathInfo
//@description: 根据id更新PathInfo
//@param: path_info project.PathInfo
//@return: err error
func (path_infoService *PathInfoService)UpdatePathInfo(path_info project.PathInfo) (err error) {
	global.GVA_LOG.Debug("[UpdatePathInfo]", zap.Any("path_info", path_info))
	path_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_path_info").UpdateOne(context.TODO(), bson.M{"_id": path_info.ID_}, bson.M{"$set": path_info})
	global.GVA_LOG.Debug("[UpdatePathInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetPathInfoById
//@description: 根据id获取PathInfo
//@param: _id primitive.ObjectID
//@return: err error,  path_info project.PathInfo
func (path_infoService *PathInfoService) GetPathInfoById(_id primitive.ObjectID) (pat project.PathInfo, err error) {
	global.GVA_LOG.Debug("[GetPathInfoById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_path_info").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&pat)
	global.GVA_LOG.Debug("[GetPathInfoById]", zap.Any("pat", pat), zap.Any("err", err))
	return pat, err
}


//@author:
//@function: GetPathInfoInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (path_infoService *PathInfoService) GetPathInfoInfoList(pat project.PathInfo, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.PathInfo
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
	
	if pat.Hostinfo != "" {
		filter["hostinfo"] = bson.M{"$regex": pat.Hostinfo }
	}
	if pat.Path != "" {
		filter["path"] = bson.M{"$regex": pat.Path }
	}
	if pat.Url != "" {
		filter["url"] = bson.M{"$regex": pat.Url }
	}
	if pat.Title != "" {
		filter["title"] = bson.M{"$regex": pat.Title }
	}
	if pat.Status != "" {
		filter["status"] = bson.M{"$regex": pat.Status }
	}
	if pat.PortId != "" {
		filter["port_id"] = bson.M{"$regex": pat.PortId }
	}

	if len(pat.Tags) > 0 {
		filter["tags"] = bson.M{"$in":pat.Tags }
	}
	if pat.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":pat.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_path_info").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetPathInfoInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_path_info").Find(context.TODO(), filter, findOptions)
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