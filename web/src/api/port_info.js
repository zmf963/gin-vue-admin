import service from '@/utils/request'

// @Tags PortInfo
// @Summary 创建PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PortInfo true "创建PortInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /port_info/createPortInfo [post]
export const createPortInfo = (data) => {
  return service({
    url: '/port_info/createPortInfo',
    method: 'post',
    data
  })
}

// @Tags PortInfo
// @Summary 删除PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PortInfo true "删除PortInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /port_info/deletePortInfo [delete]
export const deletePortInfo = (params) => {
  return service({
    url: '/port_info/deletePortInfo',
    method: 'delete',
    params
  })
}

// @Tags PortInfo
// @Summary 批量删除PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PortInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /port_info/deletePortInfo [delete]
export const deletePortInfoByIds = (data) => {
  return service({
    url: '/port_info/deletePortInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags PortInfo
// @Summary 更新PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PortInfo true "更新PortInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /port_info/updatePortInfo [put]
export const updatePortInfo = (data) => {
  return service({
    url: '/port_info/updatePortInfo',
    method: 'put',
    data
  })
}

// @Tags PortInfo
// @Summary 用id查询PortInfo
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PortInfo true "用id查询PortInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /port_info/findPortInfo [get]
export const findPortInfo = (params) => {
  return service({
    url: '/port_info/findPortInfo',
    method: 'get',
    params
  })
}

// @Tags PortInfo
// @Summary 分页获取PortInfo列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PortInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /port_info/getPortInfoList [post]
export const getPortInfoList = (data) => {
  return service({
    url: '/port_info/getPortInfoList',
    method: 'post',
    data
  })
}
