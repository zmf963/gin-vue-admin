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

type TaskApi struct {
}

// CreateTask 创建Task
// @Tags Task
// @Summary 创建Task
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.Task true "创建Task"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /task/createTask [post]
func (taskApi *TaskApi) CreateTask(c *gin.Context) {
	var _task project.Task
	_ = c.ShouldBindJSON(&_task)
	if err := taskService.CreateTask(_task); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTask 删除Task
// @Tags Task
// @Summary 删除Task
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTask [delete]
func (taskApi *TaskApi) DeleteTask(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := taskService.DeleteTask(_id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskByIds 批量删除Task
// @Tags Task
// @Summary 批量删除Task
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "批量删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /task/deleteTaskByIds [delete]
func (taskApi *TaskApi) DeleteTaskByIds(c *gin.Context) {
	var IDS request.MIdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := taskService.DeleteTaskByIds(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTask 更新Task
// @Tags Task
// @Summary 更新Task
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body project.Task true "更新Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/updateTask [put]
func (taskApi *TaskApi) UpdateTask(c *gin.Context) {
	var task project.Task
	_ = c.ShouldBindJSON(&task)
	if err := taskService.UpdateTask(task); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTask 用id查询Task
// @Tags Task
// @Summary 用id查询Task
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query string true "用id查询Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findTask [get]
func (taskApi *TaskApi) FindTask(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if retask, err := taskService.GetTaskById(_id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retask": retask}, c)
	}
}

// GetTaskList 分页获取Task列表
// @Tags Task
// @Summary 分页获取Task列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body projectReq.TaskSearch true "分页获取Task列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /task/getTaskList [post]
func (taskApi *TaskApi) GetTaskList(c *gin.Context) {
	var pageInfo projectReq.TaskSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if list, total, err := taskService.GetTaskInfoList(pageInfo.Task, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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
