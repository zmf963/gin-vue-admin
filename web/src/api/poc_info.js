import service from '@/utils/request'

// @Tags PocInfo
// @Summary 创建PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PocInfo true "创建PocInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /poc_info/createPocInfo [post]
export const createPocInfo = (data) => {
  return service({
    url: '/poc_info/createPocInfo',
    method: 'post',
    data
  })
}

// @Tags PocInfo
// @Summary 删除PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PocInfo true "删除PocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /poc_info/deletePocInfo [delete]
export const deletePocInfo = (params) => {
  return service({
    url: '/poc_info/deletePocInfo',
    method: 'delete',
    params
  })
}

// @Tags PocInfo
// @Summary 批量删除PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /poc_info/deletePocInfo [delete]
export const deletePocInfoByIds = (data) => {
  return service({
    url: '/poc_info/deletePocInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags PocInfo
// @Summary 更新PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PocInfo true "更新PocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /poc_info/updatePocInfo [put]
export const updatePocInfo = (data) => {
  return service({
    url: '/poc_info/updatePocInfo',
    method: 'put',
    data
  })
}

// @Tags PocInfo
// @Summary 用id查询PocInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PocInfo true "用id查询PocInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /poc_info/findPocInfo [get]
export const findPocInfo = (params) => {
  return service({
    url: '/poc_info/findPocInfo',
    method: 'get',
    params
  })
}

// @Tags PocInfo
// @Summary 分页获取PocInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PocInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /poc_info/getPocInfoList [post]
export const getPocInfoList = (data) => {
  return service({
    url: '/poc_info/getPocInfoList',
    method: 'post',
    data
  })
}
