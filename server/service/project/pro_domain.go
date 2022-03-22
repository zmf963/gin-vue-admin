package project

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project/response"

	//projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type DomainService struct {
}

var DomainServiceApp = new(DomainService)

//@author:
//@function: CreateDomain
//@description: 新增Domain
//@param: domain model.Domain
//@return: err error
func (domainService *DomainService) CreateDomain(domain project.Domain) (err error) {
	global.GVA_LOG.Info("[ CreateDomain]", zap.Any("domain", domain))
	domain.ID_ = primitive.NewObjectID()
	domain.CreateAt = time.Now().Local()
	domain.UpdateAt = time.Now().Local()
	result, err := global.Mongo_DB.Collection("pro_domain").InsertOne(context.TODO(), domain)
	if err != nil {
		global.GVA_LOG.Error("[CreateProject]", zap.Error(err))
		global.GVA_LOG.Error("[result]", zap.Any("result", result))
		return err
	}
	return nil
}

//@author:
//@function: DeleteDomain
//@description: 删除Domain
//@param: id
func (domainService *DomainService) DeleteDomain(_id primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteDomain]", zap.Any("id", _id))
	result, err := global.Mongo_DB.Collection("pro_domain").DeleteOne(context.TODO(), bson.M{"_id": _id})
	global.GVA_LOG.Debug("[DeleteDomain]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: DeleteDomainByIds
//@description: 批量删除Domain
//@param: ids []string
//@return: err error
func (domainService *DomainService) DeleteDomainByIds(ids []primitive.ObjectID) (err error) {
	global.GVA_LOG.Debug("[DeleteDomainByIds]", zap.Any("ids", ids))
	result, err := global.Mongo_DB.Collection("pro_domain").DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	global.GVA_LOG.Debug("[DeleteDomainByIds]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: UpdateDomain
//@description: 根据id更新Domain
//@param: domain project.Domain
//@return: err error
func (domainService *DomainService) UpdateDomain(domain response.RespDomain) (err error) {
	global.GVA_LOG.Debug("[UpdateDomain]", zap.Any("domain", domain))
	domain.UpdateAt = time.Now().Local()
	if domain.TargetName != "" {
		var _target_name project.Target
		err = global.Mongo_DB.Collection("pro_target").FindOne(context.TODO(), bson.M{"target_name": domain.TargetName}).Decode(&_target_name)
		if err != nil {
			global.GVA_LOG.Error("[UpdateDomain]", zap.Error(err))
			return err
		}
		domain.TargetId = _target_name.ID_
	}
	result, err := global.Mongo_DB.Collection("pro_domain").UpdateOne(context.TODO(), bson.M{"_id": domain.ID_}, bson.M{"$set": domain})
	global.GVA_LOG.Debug("[UpdateDomain]", zap.Any("result", result), zap.Any("err", err))
	return err
}

//@author:
//@function: GetDomainById
//@description: 根据id获取Domain
//@param: _id primitive.ObjectID
//@return: err error,  domain project.Domain
func (domainService *DomainService) GetDomainById(_id primitive.ObjectID) (dom response.RespDomain, err error) {
	global.GVA_LOG.Debug("[GetDomainById]", zap.Any("id", _id))
	err = global.Mongo_DB.Collection("pro_domain").FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&dom)
	global.GVA_LOG.Debug("[GetDomainById]", zap.Any("dom", dom), zap.Any("err", err))
	var _target_name project.Target
	err = global.Mongo_DB.Collection("pro_target").FindOne(context.TODO(), bson.M{"_id": dom.TargetId}).Decode(&_target_name)
	dom.TargetName = _target_name.TargetName
	return dom, err
}

//@author:
//@function: GetDomainInfoList
//@description: 分页获取数据,
//@param: project model.SysApi, pageInfo request.PageInfo, order string, desc bool
//@return: err error
func (domainService *DomainService) GetDomainInfoList(dom project.Domain, pageInfo request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var retList []response.RespDomain

	filter := bson.M{}
	// TODO 处理搜索条件

	if dom.Domain != "" {
		filter["domain"] = bson.M{"$regex": dom.Domain}
	}
	if len(dom.Ips) > 0 {
		filter["ips"] = bson.M{"$in": dom.Ips}
	}
	if len(dom.Hostnames) > 0 {
		filter["hostnames"] = bson.M{"$in": dom.Hostnames}
	}
	if dom.Os != "" {
		filter["os"] = bson.M{"$regex": dom.Os}
	}
	if dom.Cname != "" {
		filter["cname"] = bson.M{"$regex": dom.Cname}
	}
	if dom.Cidr != "" {
		filter["cidr"] = bson.M{"$regex": dom.Cidr}
	}
	if dom.Asn != "" {
		filter["asn"] = bson.M{"$regex": dom.Asn}
	}
	if dom.Org != "" {
		filter["org"] = bson.M{"$regex": dom.Org}
	}
	if dom.Addr != "" {
		filter["addr"] = bson.M{"$regex": dom.Addr}
	}
	if dom.Isp != "" {
		filter["isp"] = bson.M{"$regex": dom.Isp}
	}
	if dom.Source != "" {
		filter["source"] = bson.M{"$regex": dom.Source}
	}
	if !dom.TargetId.IsZero() {
		filter["target_id"] = bson.M{"$regex": dom.TargetId}
	}
	if len(dom.PortIds) > 0 {
		filter["port_ids"] = bson.M{"$in": dom.PortIds}
	}
	if len(dom.Tags) > 0 {
		filter["tags"] = bson.M{"$in": dom.Tags}
	}
	if dom.Remarks != "" {
		filter["remarks"] = bson.M{"$regex": dom.Remarks}
	}
	global.GVA_LOG.Info("[GetDomainInfoList]", zap.Any("TargetIdIsVerify", dom.TargetIdIsVerify))
	if dom.TargetIdIsVerify {
		filter["target_id_is_verify"] = dom.TargetIdIsVerify
	}
	total, err = global.Mongo_DB.Collection("pro_domain").CountDocuments(context.TODO(), filter)
	if err != nil {
		global.GVA_LOG.Error("[GetDomainInfoList]", zap.Error(err))
		return retList, total, err
	}

	pipline := []interface{}{
		bson.M{"$match": filter},
		bson.M{"$skip": offset},
		bson.M{"$limit": limit},
		bson.M{"$lookup": bson.M{"from": "pro_target", "localField": "target_id", "foreignField": "_id", "as": "pro_target"}},
		bson.M{"$project": bson.M{"domain": 1, "ips": 1, "hostnames": 1, "os": 1, "whois": 1, "alive": 1, "cname": 1, "cdn": 1, "cird": 1, "asn": 1, "org": 1, "addr": 1, "isp": 1, "source": 1, "target_id": 1, "target_id_is_verify": 1, "port_ids": 1, "tags": 1, "create_at": 1, "update_at": 1, "delete_at": 1, "target_name": "$pro_target.target_name"}},
		bson.M{"$unwind": bson.M{"path": "$target_name", "preserveNullAndEmptyArrays": true}},
	}
	if order != "" {
		_desc := 1
		if !desc {
			_desc = -1
		}
		pipline = append(pipline, bson.M{"$sort": bson.M{order: _desc}})
	} else {
		pipline = append(pipline, bson.M{"$sort": bson.M{"update_at": -1}})
	}

	cur, err := global.Mongo_DB.Collection("pro_domain").Aggregate(context.TODO(), pipline)
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
