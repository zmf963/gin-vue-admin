import service from '@/utils/request'

// @Tags DocInfo
// @Summary 创建DocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocInfo true "创建DocInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /doc_info/createDocInfo [post]
export const createDocInfo = (data) => {
  return service({
    url: '/doc_info/createDocInfo',
    method: 'post',
    data
  })
}

// @Tags DocInfo
// @Summary 删除DocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocInfo true "删除DocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /doc_info/deleteDocInfo [delete]
export const deleteDocInfo = (params) => {
  return service({
    url: '/doc_info/deleteDocInfo',
    method: 'delete',
    params
  })
}

// @Tags DocInfo
// @Summary 批量删除DocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /doc_info/deleteDocInfo [delete]
export const deleteDocInfoByIds = (data) => {
  return service({
    url: '/doc_info/deleteDocInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags DocInfo
// @Summary 更新DocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DocInfo true "更新DocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /doc_info/updateDocInfo [put]
export const updateDocInfo = (data) => {
  return service({
    url: '/doc_info/updateDocInfo',
    method: 'put',
    data
  })
}

// @Tags DocInfo
// @Summary 用id查询DocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DocInfo true "用id查询DocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /doc_info/findDocInfo [get]
export const findDocInfo = (params) => {
  return service({
    url: '/doc_info/findDocInfo',
    method: 'get',
    params
  })
}

// @Tags DocInfo
// @Summary 分页获取DocInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DocInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /doc_info/getDocInfoList [post]
export const getDocInfoList = (data) => {
  return service({
    url: '/doc_info/getDocInfoList',
    method: 'post',
    data
  })
}
