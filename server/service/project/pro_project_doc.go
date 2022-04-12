/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-04-12 19:05:56
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 20:49:56
 * @FilePath: /service/project/pro_project_doc.go
 * @Description:
 */
package project

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"

	//projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type ProjectDocService struct {
}

var ProjectDocServiceApp = new(ProjectDocService)

//@author:
//@function: CreateProjectDoc
//@description: 新增ProjectDoc
//@param: project_info model.ProjectDoc
//@return: err error
func (project_infoService *ProjectDocService) CreateProjectDoc(project_info project.ProjectDoc) (err error) {
	global.GVA_LOG.Info("[ CreateProjectDoc]", zap.Any("project_info", project_info))
	project_info.ID_ = primitive.NewObjectID()
	project_info.CreateAt = time.Now().Local()
	project_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_project_doc").InsertOne(context.TODO(), project_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteProjectDoc
//@description: 删除ProjectDoc
//@param: id
func (project_infoService *ProjectDocService) DeleteProjectDoc(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteProjectDoc]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_project_doc").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteProjectDoc]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteProjectDocByIds
//@description: 批量删除ProjectDoc
//@param: ids []string
//@return: err error
func (project_infoService *ProjectDocService) DeleteProjectDocByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteProjectDocByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_project_doc").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteProjectDocByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateProjectDoc
//@description: 根据id更新ProjectDoc
//@param: project_info project.ProjectDoc
//@return: err error
func (project_infoService *ProjectDocService) UpdateProjectDoc(project_info project.ProjectDoc) (err error) {
	global.GVA_LOG.Debug("[UpdateProjectDoc]", zap.Any("project_info", project_info))
	project_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_project_doc").UpdateOne(context.TODO(), bson.M{"_id": project_info.ID_}, bson.M{"$set": project_info})
	global.GVA_LOG.Debug("[UpdateProjectDoc]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: GetProjectDocById
//@description: 根据id获取ProjectDoc
//@param: _id primitive.ObjectID
//@return: err error,  project_info project.ProjectDoc
func (project_infoService *ProjectDocService) GetProjectDocById(_id primitive.ObjectID) (doc project.ProjectDoc, err error) {
	global.GVA_LOG.Debug("[GetProjectDocById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_project_doc").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&doc)
	global.GVA_LOG.Debug("[GetProjectDocById]", zap.Any("doc", doc), zap.Any("err", err))
	return doc, err
}

//@author:
//@function: GetProjectDocInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (project_infoService *ProjectDocService) GetProjectDocInfoList(doc project.ProjectDoc, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.ProjectDoc
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
		filter["doc_name"] = bson.M{"$regex": doc.DocName}
	}
	if doc.Link != "" {
		filter["link"] = bson.M{"$regex": doc.Link}
	}
	if doc.Source != "" {
		filter["source"] = bson.M{"$regex": doc.Source}
	}
	if !doc.ProjectId.IsZero() {
		filter["project_id"] = doc.ProjectId
	}

	if len(doc.Tags) > 0 {
		filter["tags"] = bson.M{"$in": doc.Tags}
	}
	if doc.Remarks != "" {
		filter["remarks"] = bson.M{"$regex": doc.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_project_doc").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetProjectDocInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_project_doc").Find(context.TODO(), filter, findOptions)
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
