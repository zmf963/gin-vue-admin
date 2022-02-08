import service from '@/utils/request'

// @Tags TargetRelation
// @Summary 创建TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TargetRelation true "创建TargetRelation"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /target_relation/createTargetRelation [post]
export const createTargetRelation = (data) => {
  return service({
    url: '/target_relation/createTargetRelation',
    method: 'post',
    data
  })
}

// @Tags TargetRelation
// @Summary 删除TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TargetRelation true "删除TargetRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target_relation/deleteTargetRelation [delete]
export const deleteTargetRelation = (params) => {
  return service({
    url: '/target_relation/deleteTargetRelation',
    method: 'delete',
    params
  })
}

// @Tags TargetRelation
// @Summary 批量删除TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TargetRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /target_relation/deleteTargetRelation [delete]
export const deleteTargetRelationByIds = (data) => {
  return service({
    url: '/target_relation/deleteTargetRelationByIds',
    method: 'delete',
    data
  })
}

// @Tags TargetRelation
// @Summary 更新TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TargetRelation true "更新TargetRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /target_relation/updateTargetRelation [put]
export const updateTargetRelation = (data) => {
  return service({
    url: '/target_relation/updateTargetRelation',
    method: 'put',
    data
  })
}

// @Tags TargetRelation
// @Summary 用id查询TargetRelation
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TargetRelation true "用id查询TargetRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /target_relation/findTargetRelation [get]
export const findTargetRelation = (params) => {
  return service({
    url: '/target_relation/findTargetRelation',
    method: 'get',
    params
  })
}

// @Tags TargetRelation
// @Summary 分页获取TargetRelation列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TargetRelation列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /target_relation/getTargetRelationList [post]
export const getTargetRelationList = (data) => {
  return service({
    url: '/target_relation/getTargetRelationList',
    method: 'post',
    data
  })
}
