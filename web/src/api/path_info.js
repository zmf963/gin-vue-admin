import service from '@/utils/request'

// @Tags PathInfo
// @Summary 创建PathInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PathInfo true "创建PathInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /path_info/createPathInfo [post]
export const createPathInfo = (data) => {
  return service({
    url: '/path_info/createPathInfo',
    method: 'post',
    data
  })
}

// @Tags PathInfo
// @Summary 删除PathInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PathInfo true "删除PathInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /path_info/deletePathInfo [delete]
export const deletePathInfo = (params) => {
  return service({
    url: '/path_info/deletePathInfo',
    method: 'delete',
    params
  })
}

// @Tags PathInfo
// @Summary 批量删除PathInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PathInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /path_info/deletePathInfo [delete]
export const deletePathInfoByIds = (data) => {
  return service({
    url: '/path_info/deletePathInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags PathInfo
// @Summary 更新PathInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PathInfo true "更新PathInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /path_info/updatePathInfo [put]
export const updatePathInfo = (data) => {
  return service({
    url: '/path_info/updatePathInfo',
    method: 'put',
    data
  })
}

// @Tags PathInfo
// @Summary 用id查询PathInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PathInfo true "用id查询PathInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /path_info/findPathInfo [get]
export const findPathInfo = (params) => {
  return service({
    url: '/path_info/findPathInfo',
    method: 'get',
    params
  })
}

// @Tags PathInfo
// @Summary 分页获取PathInfo列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PathInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /path_info/getPathInfoList [post]
export const getPathInfoList = (data) => {
  return service({
    url: '/path_info/getPathInfoList',
    method: 'post',
    data
  })
}
