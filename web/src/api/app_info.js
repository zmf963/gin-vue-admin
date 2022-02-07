import service from '@/utils/request'

// @Tags AppInfo
// @Summary 创建AppInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppInfo true "创建AppInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /app_info/createAppInfo [post]
export const createAppInfo = (data) => {
  return service({
    url: '/app_info/createAppInfo',
    method: 'post',
    data
  })
}

// @Tags AppInfo
// @Summary 删除AppInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppInfo true "删除AppInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app_info/deleteAppInfo [delete]
export const deleteAppInfo = (params) => {
  return service({
    url: '/app_info/deleteAppInfo',
    method: 'delete',
    params
  })
}

// @Tags AppInfo
// @Summary 批量删除AppInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app_info/deleteAppInfo [delete]
export const deleteAppInfoByIds = (data) => {
  return service({
    url: '/app_info/deleteAppInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags AppInfo
// @Summary 更新AppInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppInfo true "更新AppInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app_info/updateAppInfo [put]
export const updateAppInfo = (data) => {
  return service({
    url: '/app_info/updateAppInfo',
    method: 'put',
    data
  })
}

// @Tags AppInfo
// @Summary 用id查询AppInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppInfo true "用id查询AppInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app_info/findAppInfo [get]
export const findAppInfo = (params) => {
  return service({
    url: '/app_info/findAppInfo',
    method: 'get',
    params
  })
}

// @Tags AppInfo
// @Summary 分页获取AppInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /app_info/getAppInfoList [post]
export const getAppInfoList = (data) => {
  return service({
    url: '/app_info/getAppInfoList',
    method: 'post',
    data
  })
}
