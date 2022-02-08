package project

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	projectReq "github.com/flipped-aurora/gin-vue-admin/server/model/project/request"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type DocInfoApi struct {
}

// CreateDocInfo 创建DocInfo
// @Tags DocInfo
// @Summary 创建DocInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.DocInfo true "创建DocInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /doc_info/createDocInfo [post]
func (doc_infoApi *DocInfoApi) CreateDocInfo(c *gin.Context) {
	var _doc_info project.DocInfo
	_ = c.ShouldBindJSON(&_doc_info)
	if err := doc_infoService.CreateDocInfo(_doc_info); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDocInfo 删除DocInfo
// @Tags DocInfo
// @Summary 删除DocInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除DocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /doc_info/deleteDocInfo [delete]
func (doc_infoApi *DocInfoApi) DeleteDocInfo(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := doc_infoService.DeleteDocInfo(_id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDocInfoByIds 批量删除DocInfo
// @Tags DocInfo
// @Summary 批量删除DocInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除DocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /doc_info/deleteDocInfoByIds [delete]
func (doc_infoApi *DocInfoApi) DeleteDocInfoByIds(c *gin.Context) {
	var IDS request.MIdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := doc_infoService.DeleteDocInfoByIds(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDocInfo 更新DocInfo
// @Tags DocInfo
// @Summary 更新DocInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.DocInfo true "更新DocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /doc_info/updateDocInfo [put]
func (doc_infoApi *DocInfoApi) UpdateDocInfo(c *gin.Context) {
	var doc_info project.DocInfo
	_ = c.ShouldBindJSON(&doc_info)
	if err := doc_infoService.UpdateDocInfo(doc_info); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDocInfo 用id查询DocInfo
// @Tags DocInfo
// @Summary 用id查询DocInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询DocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /doc_info/findDocInfo [get]
func (doc_infoApi *DocInfoApi) FindDocInfo(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if redoc_info, err := doc_infoService.GetDocInfoById(_id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redoc_info": redoc_info}, c)
	}
}

// GetDocInfoList 分页获取DocInfo列表
// @Tags DocInfo
// @Summary 分页获取DocInfo列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.DocInfoSearch true "分页获取DocInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /doc_info/getDocInfoList [post]
func (doc_infoApi *DocInfoApi) GetDocInfoList(c *gin.Context) {
	var pageInfo projectReq.DocInfoSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if list, total, err := doc_infoService.GetDocInfoInfoList(pageInfo.DocInfo, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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
