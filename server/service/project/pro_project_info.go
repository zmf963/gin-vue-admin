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

type ProjectInfoService struct {
}

var ProjectInfoServiceApp = new(ProjectInfoService)


//@author:
//@function: CreateProjectInfo
//@description: 新增ProjectInfo
//@param: project_info model.ProjectInfo
//@return: err error
func (project_infoService *ProjectInfoService) CreateProjectInfo(project_info project.ProjectInfo) (err error) {
	global.GVA_LOG.Info("[ CreateProjectInfo]", zap.Any("project_info", project_info))
	project_info.ID_ = primitive.NewObjectID()
	project_info.CreateAt = time.Now().Local()
	project_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_project_info").InsertOne(context.TODO(), project_info)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteProjectInfo
//@description: 删除ProjectInfo
//@param: id
func (project_infoService *ProjectInfoService)DeleteProjectInfo(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteProjectInfo]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_project_info").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteProjectInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteProjectInfoByIds
//@description: 批量删除ProjectInfo
//@param: ids []string
//@return: err error
func (project_infoService *ProjectInfoService)DeleteProjectInfoByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteProjectInfoByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_project_info").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteProjectInfoByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateProjectInfo
//@description: 根据id更新ProjectInfo
//@param: project_info project.ProjectInfo
//@return: err error
func (project_infoService *ProjectInfoService)UpdateProjectInfo(project_info project.ProjectInfo) (err error) {
	global.GVA_LOG.Debug("[UpdateProjectInfo]", zap.Any("project_info", project_info))
	project_info.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_project_info").UpdateOne(context.TODO(), bson.M{"_id": project_info.ID_}, bson.M{"$set": project_info})
	global.GVA_LOG.Debug("[UpdateProjectInfo]", zap.Any("result", result), zap.Any("err", err))
	return err
}


//@author:
//@function: GetProjectInfoById
//@description: 根据id获取ProjectInfo
//@param: _id primitive.ObjectID
//@return: err error,  project_info project.ProjectInfo
func (project_infoService *ProjectInfoService) GetProjectInfoById(_id primitive.ObjectID) (pro project.ProjectInfo, err error) {
	global.GVA_LOG.Debug("[GetProjectInfoById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_project_info").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&pro)
	global.GVA_LOG.Debug("[GetProjectInfoById]", zap.Any("pro", pro), zap.Any("err", err))
	return pro, err
}


//@author:
//@function: GetProjectInfoInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (project_infoService *ProjectInfoService) GetProjectInfoInfoList(pro project.ProjectInfo, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []project.ProjectInfo
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
	
	if pro.ProjectName != "" {
		filter["project_name"] = bson.M{"$regex": pro.ProjectName }
	}
	if pro.ProjectDesc != "" {
		filter["project_desc"] = bson.M{"$regex": pro.ProjectDesc }
	}
	if len(pro.TargetIds) > 0 {
		filter["target_ids"] = bson.M{"$in": pro.TargetIds }
	}
	if len(pro.TaskIds) > 0 {
		filter["task_ids"] = bson.M{"$in": pro.TaskIds }
	}

	if len(pro.Tags) > 0 {
		filter["tags"] = bson.M{"$in":pro.Tags }
	}
	if pro.Remarks != "" {
		filter["remarks"] = bson.M{"$regex":pro.Remarks}
	}

	total, err = global.Mongo_DB.Collection("pro_project_info").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetProjectInfoInfoList]", zap.Error(err))
		return retList, total, err
	}
	cur, err := global.Mongo_DB.Collection("pro_project_info").Find(context.TODO(), filter, findOptions)
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