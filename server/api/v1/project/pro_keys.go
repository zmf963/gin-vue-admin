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

type KeysApi struct {
}

// CreateKeys 创建Keys
// @Tags Keys
// @Summary 创建Keys
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.Keys true "创建Keys"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /keys/createKeys [post]
func (keysApi *KeysApi) CreateKeys(c *gin.Context) {
	var _keys project.Keys
	_ = c.ShouldBindJSON(&_keys)
	if err := keysService.CreateKeys(_keys); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteKeys 删除Keys
// @Tags Keys
// @Summary 删除Keys
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除Keys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /keys/deleteKeys [delete]
func (keysApi *KeysApi) DeleteKeys(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := keysService.DeleteKeys(_id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteKeysByIds 批量删除Keys
// @Tags Keys
// @Summary 批量删除Keys
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除Keys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /keys/deleteKeysByIds [delete]
func (keysApi *KeysApi) DeleteKeysByIds(c *gin.Context) {
	var IDS request.MIdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := keysService.DeleteKeysByIds(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateKeys 更新Keys
// @Tags Keys
// @Summary 更新Keys
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.Keys true "更新Keys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /keys/updateKeys [put]
func (keysApi *KeysApi) UpdateKeys(c *gin.Context) {
	var keys project.Keys
	_ = c.ShouldBindJSON(&keys)
	if err := keysService.UpdateKeys(keys); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindKeys 用id查询Keys
// @Tags Keys
// @Summary 用id查询Keys
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询Keys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /keys/findKeys [get]
func (keysApi *KeysApi) FindKeys(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if rekeys, err := keysService.GetKeysById(_id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rekeys": rekeys}, c)
	}
}

// GetKeysList 分页获取Keys列表
// @Tags Keys
// @Summary 分页获取Keys列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.KeysSearch true "分页获取Keys列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /keys/getKeysList [post]
func (keysApi *KeysApi) GetKeysList(c *gin.Context) {
	var pageInfo projectReq.KeysSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if list, total, err := keysService.GetKeysInfoList(pageInfo.Keys, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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
