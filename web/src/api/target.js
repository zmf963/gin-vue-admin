import service from '@/utils/request'

// @Tags Target
// @Summary 创建Target
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Target true "创建Target"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /target/createTarget [post]
export const createTarget = (data) => {
  return service({
    url: '/target/createTarget',
    method: 'post',
    data
  })
}

// @Tags Target
// @Summary 删除Target
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Target true "删除Target"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target/deleteTarget [delete]
export const deleteTarget = (params) => {
  return service({
    url: '/target/deleteTarget',
    method: 'delete',
    params
  })
}

// @Tags Target
// @Summary 批量删除Target
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Target"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target/deleteTarget [delete]
export const deleteTargetByIds = (data) => {
  return service({
    url: '/target/deleteTargetByIds',
    method: 'delete',
    data
  })
}

// @Tags Target
// @Summary 更新Target
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Target true "更新Target"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /target/updateTarget [put]
export const updateTarget = (data) => {
  return service({
    url: '/target/updateTarget',
    method: 'put',
    data
  })
}

// @Tags Target
// @Summary 用id查询Target
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Target true "用id查询Target"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /target/findTarget [get]
export const findTarget = (params) => {
  return service({
    url: '/target/findTarget',
    method: 'get',
    params
  })
}

// @Tags Target
// @Summary 分页获取Target列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Target列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /target/getTargetList [post]
export const getTargetList = (data) => {
  return service({
    url: '/target/getTargetList',
    method: 'post',
    data
  })
}
