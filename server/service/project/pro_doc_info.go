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

type DocInfoService struct {
}

var DocInfoServiceApp = new(DocInfoService)


//@author:
//@function: CreateDocInfo
//@description: 新增DocInfo
//@param: doc_info model.DocInfo
//@return: err error
func (doc_infoService *DocInfoService) CreateDocInfo(doc_info project.DocInfo) (err error) {
	global.GVA_LOG.Info("[ CreateDocInfo]", zap.Any("doc_info", doc_info))
	doc_info.ID_ = primitive.NewObjectID()
	doc_info.CreateAt = time.Now().Local()
	doc_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_doc_info").InsertOne(context.TODO(), doc_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteDocInfo
//@description: 删除DocInfo
//@param: id
func (doc_infoService *DocInfoService)DeleteDocInfo(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteDocInfo]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_doc_info").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteDocInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteDocInfoByIds
//@description: 批量删除DocInfo
//@param: ids []string
//@return: err error
func (doc_infoService *DocInfoService)DeleteDocInfoByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteDocInfoByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_doc_info").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteDocInfoByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateDocInfo
//@description: 根据id更新DocInfo
//@param: doc_info project.DocInfo
//@return: err error
func (doc_infoService *DocInfoService)UpdateDocInfo(doc_info project.DocInfo) (err error) {
	global.GVA_LOG.Debug("[UpdateDocInfo]", zap.Any("doc_info", doc_info))
	doc_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_doc_info").UpdateOne(context.TODO(), bson.M{"_id": doc_info.ID_}, bson.M{"$set": doc_info})
	global.GVA_LOG.Debug("[UpdateDocInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetDocInfoById
//@description: 根据id获取DocInfo
//@param: _id primitive.ObjectID
//@return: err error,  doc_info project.DocInfo
func (doc_infoService *DocInfoService) GetDocInfoById(_id primitive.ObjectID) (doc project.DocInfo, err error) {
	global.GVA_LOG.Debug("[GetDocInfoById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_doc_info").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&doc)
	global.GVA_LOG.Debug("[GetDocInfoById]", zap.Any("doc", doc), zap.Any("err", err))
	return doc, err
}


//@author:
//@function: GetDocInfoInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (doc_infoService *DocInfoService) GetDocInfoInfoList(doc project.DocInfo, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.DocInfo
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
	
	if doc.DocName != "" {
		filter["doc_name"] = bson.M{"$regex": doc.DocName }
	}
	if doc.Link != "" {
		filter["link"] = bson.M{"$regex": doc.Link }
	}
	if doc.Source != "" {
		filter["source"] = bson.M{"$regex": doc.Source }
	}
	if doc.TargetId != "" {
		filter["target_id"] = bson.M{"$regex": doc.TargetId }
	}

	if len(doc.Tags) > 0 {
		filter["tags"] = bson.M{"$in":doc.Tags }
	}
	if doc.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":doc.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_doc_info").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetDocInfoInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_doc_info").Find(context.TODO(), filter, findOptions)
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