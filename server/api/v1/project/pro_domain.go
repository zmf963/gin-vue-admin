package project

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/project"
    projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type DomainApi struct {
}


// CreateDomain 创建Domain
// @Tags Domain
// @Summary 创建Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.Domain true "创建Domain"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /domain/createDomain [post]
func (domainApi *DomainApi) CreateDomain(c *gin.Context) {
    var _domain project.Domain
    _ = c.ShouldBindJSON(&_domain)
    if err := domainService.CreateDomain(_domain); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeleteDomain 删除Domain
// @Tags Domain
// @Summary 删除Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /domain/deleteDomain [delete]
func (domainApi *DomainApi) DeleteDomain(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if err := domainService.DeleteDomain(_id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteDomainByIds 批量删除Domain
// @Tags Domain
// @Summary 批量删除Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /domain/deleteDomainByIds [delete]
func (domainApi *DomainApi) DeleteDomainByIds(c *gin.Context) {
    var IDS request.MIdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := domainService.DeleteDomainByIds(IDS.Ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdateDomain 更新Domain
// @Tags Domain
// @Summary 更新Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.Domain true "更新Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /domain/updateDomain [put]
func (domainApi *DomainApi) UpdateDomain(c *gin.Context) {
    var domain project.Domain
    _ = c.ShouldBindJSON(&domain)
    if err := domainService.UpdateDomain(domain); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindDomain 用id查询Domain
// @Tags Domain
// @Summary 用id查询Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /domain/findDomain [get]
func (domainApi *DomainApi) FindDomain(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if redomain,err := domainService.GetDomainById(_id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"redomain": redomain}, c)
    }
}

// GetDomainList 分页获取Domain列表
// @Tags Domain
// @Summary 分页获取Domain列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.DomainSearch true "分页获取Domain列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /domain/getDomainList [post]
func (domainApi *DomainApi) GetDomainList(c *gin.Context) {
    var pageInfo projectReq.DomainSearch
    _ = c.ShouldBindJSON(&pageInfo)
    if list, total, err := domainService.GetDomainInfoList(pageInfo.Domain,pageInfo.PageInfo,pageInfo.OrderKey,pageInfo.Desc); err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}