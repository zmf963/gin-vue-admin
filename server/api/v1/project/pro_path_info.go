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

type PathInfoApi struct {
}


// CreatePathInfo 创建PathInfo
// @Tags PathInfo
// @Summary 创建PathInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.PathInfo true "创建PathInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /path_info/createPathInfo [post]
func (path_infoApi *PathInfoApi) CreatePathInfo(c *gin.Context) {
    var _path_info project.PathInfo
    _ = c.ShouldBindJSON(&_path_info)
    if err := path_infoService.CreatePathInfo(_path_info); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// DeletePathInfo 删除PathInfo
// @Tags PathInfo
// @Summary 删除PathInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除PathInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /path_info/deletePathInfo [delete]
func (path_infoApi *PathInfoApi) DeletePathInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if err := path_infoService.DeletePathInfo(_id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeletePathInfoByIds 批量删除PathInfo
// @Tags PathInfo
// @Summary 批量删除PathInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除PathInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /path_info/deletePathInfoByIds [delete]
func (path_infoApi *PathInfoApi) DeletePathInfoByIds(c *gin.Context) {
    var IDS request.MIdsReq
    _ = c.ShouldBindJSON(&IDS)
    if err := path_infoService.DeletePathInfoByIds(IDS.Ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// UpdatePathInfo 更新PathInfo
// @Tags PathInfo
// @Summary 更新PathInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body project.PathInfo true "更新PathInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /path_info/updatePathInfo [put]
func (path_infoApi *PathInfoApi) UpdatePathInfo(c *gin.Context) {
    var path_info project.PathInfo
    _ = c.ShouldBindJSON(&path_info)
    if err := path_infoService.UpdatePathInfo(path_info); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// FindPathInfo 用id查询PathInfo
// @Tags PathInfo
// @Summary 用id查询PathInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询PathInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /path_info/findPathInfo [get]
func (path_infoApi *PathInfoApi) FindPathInfo(c *gin.Context) {
    _id, err := primitive.ObjectIDFromHex(c.Query("_id"))
    if err != nil {
        response.FailWithMessage("参数错误", c)
        return
    }
    if repath_info,err := path_infoService.GetPathInfoById(_id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"repath_info": repath_info}, c)
    }
}

// GetPathInfoList 分页获取PathInfo列表
// @Tags PathInfo
// @Summary 分页获取PathInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.PathInfoSearch true "分页获取PathInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /path_info/getPathInfoList [post]
func (path_infoApi *PathInfoApi) GetPathInfoList(c *gin.Context) {
    var pageInfo projectReq.PathInfoSearch
    _ = c.ShouldBindJSON(&pageInfo)
    if list, total, err := path_infoService.GetPathInfoInfoList(pageInfo.PathInfo,pageInfo.PageInfo,pageInfo.OrderKey,pageInfo.Desc); err != nil {
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
