import service from '@/utils/request'

// @Tags ProjectInfo
// @Summary 创建ProjectInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectInfo true "创建ProjectInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /project_info/createProjectInfo [post]
export const createProjectInfo = (data) => {
  return service({
    url: '/project_info/createProjectInfo',
    method: 'post',
    data
  })
}

// @Tags ProjectInfo
// @Summary 删除ProjectInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProjectInfo true "删除ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project_info/deleteProjectInfo [delete]
export const deleteProjectInfo = (params) => {
  return service({
    url: '/project_info/deleteProjectInfo',
    method: 'delete',
    params
  })
}

// @Tags ProjectInfo
// @Summary 批量删除ProjectInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project_info/deleteProjectInfo [delete]
export const deleteProjectInfoByIds = (data) => {
  return service({
    url: '/project_info/deleteProjectInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags ProjectInfo
// @Summary 更新ProjectInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectInfo true "更新ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project_info/updateProjectInfo [put]
export const updateProjectInfo = (data) => {
  return service({
    url: '/project_info/updateProjectInfo',
    method: 'put',
    data
  })
}

// @Tags ProjectInfo
// @Summary 用id查询ProjectInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProjectInfo true "用id查询ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project_info/findProjectInfo [get]
export const findProjectInfo = (params) => {
  return service({
    url: '/project_info/findProjectInfo',
    method: 'get',
    params
  })
}

// @Tags ProjectInfo
// @Summary 分页获取ProjectInfo列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ProjectInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /project_info/getProjectInfoList [post]
export const getProjectInfoList = (data) => {
  return service({
    url: '/project_info/getProjectInfoList',
    method: 'post',
    data
  })
}
