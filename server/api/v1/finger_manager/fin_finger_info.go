package finger_manager

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/finger_manager"
    finger_managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/finger_manager/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type FingerInfoApi struct {
}


// CreateFingerInfo 创建FingerInfo
// @Tags FingerInfo
// @Summary 创建FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body finger_manager.FingerInfo true "创建FingerInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /finger_info/createFingerInfo [post]
func (finger_infoApi *FingerInfoApi) CreateFingerInfo(c *gin.Context) {
    var _finger_info finger_manager.FingerInfo
    _ = c.ShouldBindJSON(&_finger_info)
    if err := finger_infoService.CreateFingerInfo(_finger_info); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeleteFingerInfo 删除FingerInfo
// @Tags FingerInfo
// @Summary 删除FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除FingerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /finger_info/deleteFingerInfo [delete]
func (finger_infoApi *FingerInfoApi) DeleteFingerInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if err := finger_infoService.DeleteFingerInfo(_id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteFingerInfoByIds 批量删除FingerInfo
// @Tags FingerInfo
// @Summary 批量删除FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除FingerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /finger_info/deleteFingerInfoByIds [delete]
func (finger_infoApi *FingerInfoApi) DeleteFingerInfoByIds(c *gin.Context) {
    var IDS request.MIdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := finger_infoService.DeleteFingerInfoByIds(IDS.Ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdateFingerInfo 更新FingerInfo
// @Tags FingerInfo
// @Summary 更新FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body finger_manager.FingerInfo true "更新FingerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /finger_info/updateFingerInfo [put]
func (finger_infoApi *FingerInfoApi) UpdateFingerInfo(c *gin.Context) {
    var finger_info finger_manager.FingerInfo
    _ = c.ShouldBindJSON(&finger_info)
    if err := finger_infoService.UpdateFingerInfo(finger_info); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindFingerInfo 用id查询FingerInfo
// @Tags FingerInfo
// @Summary 用id查询FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询FingerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /finger_info/findFingerInfo [get]
func (finger_infoApi *FingerInfoApi) FindFingerInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if refinger_info,err := finger_infoService.GetFingerInfoById(_id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"refinger_info": refinger_info}, c)
    }
}

// GetFingerInfoList 分页获取FingerInfo列表
// @Tags FingerInfo
// @Summary 分页获取FingerInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body finger_managerReq.FingerInfoSearch true "分页获取FingerInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /finger_info/getFingerInfoList [post]
func (finger_infoApi *FingerInfoApi) GetFingerInfoList(c *gin.Context) {
    var pageInfo finger_managerReq.FingerInfoSearch
    _ = c.ShouldBindJSON(&pageInfo)
    if list, total, err := finger_infoService.GetFingerInfoInfoList(pageInfo.FingerInfo,pageInfo.PageInfo,pageInfo.OrderKey,pageInfo.Desc); err != nil {
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
