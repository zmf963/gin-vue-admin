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

type WechatOfficialAccountApi struct {
}


// CreateWechatOfficialAccount 创建WechatOfficialAccount
// @Tags WechatOfficialAccount
// @Summary 创建WechatOfficialAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.WechatOfficialAccount true "创建WechatOfficialAccount"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /wechat_official_account/createWechatOfficialAccount [post]
func (wechat_official_accountApi *WechatOfficialAccountApi) CreateWechatOfficialAccount(c *gin.Context) {
    var _wechat_official_account project.WechatOfficialAccount
    _ = c.ShouldBindJSON(&_wechat_official_account)
    if err := wechat_official_accountService.CreateWechatOfficialAccount(_wechat_official_account); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeleteWechatOfficialAccount 删除WechatOfficialAccount
// @Tags WechatOfficialAccount
// @Summary 删除WechatOfficialAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除WechatOfficialAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wechat_official_account/deleteWechatOfficialAccount [delete]
func (wechat_official_accountApi *WechatOfficialAccountApi) DeleteWechatOfficialAccount(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if err := wechat_official_accountService.DeleteWechatOfficialAccount(_id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteWechatOfficialAccountByIds 批量删除WechatOfficialAccount
// @Tags WechatOfficialAccount
// @Summary 批量删除WechatOfficialAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除WechatOfficialAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wechat_official_account/deleteWechatOfficialAccountByIds [delete]
func (wechat_official_accountApi *WechatOfficialAccountApi) DeleteWechatOfficialAccountByIds(c *gin.Context) {
    var IDS request.MIdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := wechat_official_accountService.DeleteWechatOfficialAccountByIds(IDS.Ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdateWechatOfficialAccount 更新WechatOfficialAccount
// @Tags WechatOfficialAccount
// @Summary 更新WechatOfficialAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.WechatOfficialAccount true "更新WechatOfficialAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wechat_official_account/updateWechatOfficialAccount [put]
func (wechat_official_accountApi *WechatOfficialAccountApi) UpdateWechatOfficialAccount(c *gin.Context) {
    var wechat_official_account project.WechatOfficialAccount
    _ = c.ShouldBindJSON(&wechat_official_account)
    if err := wechat_official_accountService.UpdateWechatOfficialAccount(wechat_official_account); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindWechatOfficialAccount 用id查询WechatOfficialAccount
// @Tags WechatOfficialAccount
// @Summary 用id查询WechatOfficialAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询WechatOfficialAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wechat_official_account/findWechatOfficialAccount [get]
func (wechat_official_accountApi *WechatOfficialAccountApi) FindWechatOfficialAccount(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if rewechat_official_account,err := wechat_official_accountService.GetWechatOfficialAccountById(_id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"rewechat_official_account": rewechat_official_account}, c)
    }
}

// GetWechatOfficialAccountList 分页获取WechatOfficialAccount列表
// @Tags WechatOfficialAccount
// @Summary 分页获取WechatOfficialAccount列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.WechatOfficialAccountSearch true "分页获取WechatOfficialAccount列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /wechat_official_account/getWechatOfficialAccountList [post]
func (wechat_official_accountApi *WechatOfficialAccountApi) GetWechatOfficialAccountList(c *gin.Context) {
    var pageInfo projectReq.WechatOfficialAccountSearch
    _ = c.ShouldBindJSON(&pageInfo)
    if list, total, err := wechat_official_accountService.GetWechatOfficialAccountInfoList(pageInfo.WechatOfficialAccount,pageInfo.PageInfo,pageInfo.OrderKey,pageInfo.Desc); err != nil {
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
