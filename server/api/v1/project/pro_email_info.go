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

type EmailInfoApi struct {
}


// CreateEmailInfo 创建EmailInfo
// @Tags EmailInfo
// @Summary 创建EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.EmailInfo true "创建EmailInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /email_info/createEmailInfo [post]
func (email_infoApi *EmailInfoApi) CreateEmailInfo(c *gin.Context) {
    var _email_info project.EmailInfo
    _ = c.ShouldBindJSON(&_email_info)
    if err := email_infoService.CreateEmailInfo(_email_info); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeleteEmailInfo 删除EmailInfo
// @Tags EmailInfo
// @Summary 删除EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除EmailInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /email_info/deleteEmailInfo [delete]
func (email_infoApi *EmailInfoApi) DeleteEmailInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if err := email_infoService.DeleteEmailInfo(_id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteEmailInfoByIds 批量删除EmailInfo
// @Tags EmailInfo
// @Summary 批量删除EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除EmailInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /email_info/deleteEmailInfoByIds [delete]
func (email_infoApi *EmailInfoApi) DeleteEmailInfoByIds(c *gin.Context) {
    var IDS request.MIdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := email_infoService.DeleteEmailInfoByIds(IDS.Ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdateEmailInfo 更新EmailInfo
// @Tags EmailInfo
// @Summary 更新EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.EmailInfo true "更新EmailInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /email_info/updateEmailInfo [put]
func (email_infoApi *EmailInfoApi) UpdateEmailInfo(c *gin.Context) {
    var email_info project.EmailInfo
    _ = c.ShouldBindJSON(&email_info)
    if err := email_infoService.UpdateEmailInfo(email_info); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindEmailInfo 用id查询EmailInfo
// @Tags EmailInfo
// @Summary 用id查询EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询EmailInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /email_info/findEmailInfo [get]
func (email_infoApi *EmailInfoApi) FindEmailInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if reemail_info,err := email_infoService.GetEmailInfoById(_id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"reemail_info": reemail_info}, c)
    }
}

// GetEmailInfoList 分页获取EmailInfo列表
// @Tags EmailInfo
// @Summary 分页获取EmailInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.EmailInfoSearch true "分页获取EmailInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /email_info/getEmailInfoList [post]
func (email_infoApi *EmailInfoApi) GetEmailInfoList(c *gin.Context) {
    var pageInfo projectReq.EmailInfoSearch
    _ = c.ShouldBindJSON(&pageInfo)
    if list, total, err := email_infoService.GetEmailInfoInfoList(pageInfo.EmailInfo,pageInfo.PageInfo,pageInfo.OrderKey,pageInfo.Desc); err != nil {
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
