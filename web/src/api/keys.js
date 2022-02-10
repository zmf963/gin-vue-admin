import service from '@/utils/request'

// @Tags Keys
// @Summary 创建Keys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Keys true "创建Keys"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /keys/createKeys [post]
export const createKeys = (data) => {
  return service({
    url: '/keys/createKeys',
    method: 'post',
    data
  })
}

// @Tags Keys
// @Summary 删除Keys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Keys true "删除Keys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /keys/deleteKeys [delete]
export const deleteKeys = (params) => {
  return service({
    url: '/keys/deleteKeys',
    method: 'delete',
    params
  })
}

// @Tags Keys
// @Summary 批量删除Keys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Keys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /keys/deleteKeys [delete]
export const deleteKeysByIds = (data) => {
  return service({
    url: '/keys/deleteKeysByIds',
    method: 'delete',
    data
  })
}

// @Tags Keys
// @Summary 更新Keys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Keys true "更新Keys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /keys/updateKeys [put]
export const updateKeys = (data) => {
  return service({
    url: '/keys/updateKeys',
    method: 'put',
    data
  })
}

// @Tags Keys
// @Summary 用id查询Keys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Keys true "用id查询Keys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /keys/findKeys [get]
export const findKeys = (params) => {
  return service({
    url: '/keys/findKeys',
    method: 'get',
    params
  })
}

// @Tags Keys
// @Summary 分页获取Keys列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Keys列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /keys/getKeysList [post]
export const getKeysList = (data) => {
  return service({
    url: '/keys/getKeysList',
    method: 'post',
    data
  })
}
