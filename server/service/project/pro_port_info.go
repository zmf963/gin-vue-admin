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

type PortInfoService struct {
}

var PortInfoServiceApp = new(PortInfoService)


//@author:
//@function: CreatePortInfo
//@description: 新增PortInfo
//@param: port_info model.PortInfo
//@return: err error
func (port_infoService *PortInfoService) CreatePortInfo(port_info project.PortInfo) (err error) {
	global.GVA_LOG.Info("[ CreatePortInfo]", zap.Any("port_info", port_info))
	port_info.ID_ = primitive.NewObjectID()
	port_info.CreateAt = time.Now().Local()
	port_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_port_info").InsertOne(context.TODO(), port_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeletePortInfo
//@description: 删除PortInfo
//@param: id
func (port_infoService *PortInfoService)DeletePortInfo(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeletePortInfo]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_port_info").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeletePortInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeletePortInfoByIds
//@description: 批量删除PortInfo
//@param: ids []string
//@return: err error
func (port_infoService *PortInfoService)DeletePortInfoByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeletePortInfoByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_port_info").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeletePortInfoByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdatePortInfo
//@description: 根据id更新PortInfo
//@param: port_info project.PortInfo
//@return: err error
func (port_infoService *PortInfoService)UpdatePortInfo(port_info project.PortInfo) (err error) {
	global.GVA_LOG.Debug("[UpdatePortInfo]", zap.Any("port_info", port_info))
	port_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_port_info").UpdateOne(context.TODO(), bson.M{"_id": port_info.ID_}, bson.M{"$set": port_info})
	global.GVA_LOG.Debug("[UpdatePortInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetPortInfoById
//@description: 根据id获取PortInfo
//@param: _id primitive.ObjectID
//@return: err error,  port_info project.PortInfo
func (port_infoService *PortInfoService) GetPortInfoById(_id primitive.ObjectID) (por project.PortInfo, err error) {
	global.GVA_LOG.Debug("[GetPortInfoById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_port_info").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&por)
	global.GVA_LOG.Debug("[GetPortInfoById]", zap.Any("por", por), zap.Any("err", err))
	return por, err
}


//@author:
//@function: GetPortInfoInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (port_infoService *PortInfoService) GetPortInfoInfoList(por project.PortInfo, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.PortInfo
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
	
	if por.Port != "" {
		filter["port"] = bson.M{"$regex": por.Port }
	}
	if por.Host != "" {
		filter["host"] = bson.M{"$regex": por.Host }
	}
	if por.Hostinfo != "" {
		filter["hostinfo"] = bson.M{"$regex": por.Hostinfo }
	}
	if por.Url != "" {
		filter["url"] = bson.M{"$regex": por.Url }
	}
	if por.Title != "" {
		filter["title"] = bson.M{"$regex": por.Title }
	}
	if por.Favicons != "" {
		filter["favicons"] = bson.M{"$regex": por.Favicons }
	}
	if len(por.Products) > 0 {
		// filter["products"] = bson.M{"$in": por.Products }
		for _, v := range por.Products {
			
			filter["products"] = bson.M{"$regex": v }
		}
	}
	if len(por.Protocols) > 0 {
		filter["protocols"] = bson.M{"$in": por.Protocols }
	}
	if por.Banner != "" {
		filter["banner"] = bson.M{"$regex": por.Banner }
	}
	if por.Status != "" {
		filter["status"] = bson.M{"$regex": por.Status }
	}
	if por.DomainId != "" {
		filter["domain_id"] = bson.M{"$regex": por.DomainId }
	}
	if len(por.PathIds) > 0 {
		filter["path_ids"] = bson.M{"$in": por.PathIds }
	}

	if len(por.Tags) > 0 {
		filter["tags"] = bson.M{"$in":por.Tags }
	}
	if por.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":por.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_port_info").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetPortInfoInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_port_info").Find(context.TODO(), filter, findOptions)
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