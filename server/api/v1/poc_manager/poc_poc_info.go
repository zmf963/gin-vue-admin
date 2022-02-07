package poc_manager

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/poc_manager"
    poc_managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/poc_manager/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type PocInfoApi struct {
}


// CreatePocInfo 创建PocInfo
// @Tags PocInfo
// @Summary 创建PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body poc_manager.PocInfo true "创建PocInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /poc_info/createPocInfo [post]
func (poc_infoApi *PocInfoApi) CreatePocInfo(c *gin.Context) {
    var _poc_info poc_manager.PocInfo
    _ = c.ShouldBindJSON(&_poc_info)
    if err := poc_infoService.CreatePocInfo(_poc_info); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeletePocInfo 删除PocInfo
// @Tags PocInfo
// @Summary 删除PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除PocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /poc_info/deletePocInfo [delete]
func (poc_infoApi *PocInfoApi) DeletePocInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if err := poc_infoService.DeletePocInfo(_id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeletePocInfoByIds 批量删除PocInfo
// @Tags PocInfo
// @Summary 批量删除PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除PocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /poc_info/deletePocInfoByIds [delete]
func (poc_infoApi *PocInfoApi) DeletePocInfoByIds(c *gin.Context) {
    var IDS request.MIdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := poc_infoService.DeletePocInfoByIds(IDS.Ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdatePocInfo 更新PocInfo
// @Tags PocInfo
// @Summary 更新PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body poc_manager.PocInfo true "更新PocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /poc_info/updatePocInfo [put]
func (poc_infoApi *PocInfoApi) UpdatePocInfo(c *gin.Context) {
    var poc_info poc_manager.PocInfo
    _ = c.ShouldBindJSON(&poc_info)
    if err := poc_infoService.UpdatePocInfo(poc_info); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindPocInfo 用id查询PocInfo
// @Tags PocInfo
// @Summary 用id查询PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询PocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /poc_info/findPocInfo [get]
func (poc_infoApi *PocInfoApi) FindPocInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if repoc_info,err := poc_infoService.GetPocInfoById(_id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"repoc_info": repoc_info}, c)
    }
}

// GetPocInfoList 分页获取PocInfo列表
// @Tags PocInfo
// @Summary 分页获取PocInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body poc_managerReq.PocInfoSearch true "分页获取PocInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /poc_info/getPocInfoList [post]
func (poc_infoApi *PocInfoApi) GetPocInfoList(c *gin.Context) {
    var pageInfo poc_managerReq.PocInfoSearch
    _ = c.ShouldBindJSON(&pageInfo)
    if list, total, err := poc_infoService.GetPocInfoInfoList(pageInfo.PocInfo,pageInfo.PageInfo,pageInfo.OrderKey,pageInfo.Desc); err != nil {
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
