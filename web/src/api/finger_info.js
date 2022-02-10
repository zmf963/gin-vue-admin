import service from '@/utils/request'

// @Tags FingerInfo
// @Summary 创建FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FingerInfo true "创建FingerInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /finger_info/createFingerInfo [post]
export const createFingerInfo = (data) => {
  return service({
    url: '/finger_info/createFingerInfo',
    method: 'post',
    data
  })
}

// @Tags FingerInfo
// @Summary 删除FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FingerInfo true "删除FingerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /finger_info/deleteFingerInfo [delete]
export const deleteFingerInfo = (params) => {
  return service({
    url: '/finger_info/deleteFingerInfo',
    method: 'delete',
    params
  })
}

// @Tags FingerInfo
// @Summary 批量删除FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FingerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /finger_info/deleteFingerInfo [delete]
export const deleteFingerInfoByIds = (data) => {
  return service({
    url: '/finger_info/deleteFingerInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags FingerInfo
// @Summary 更新FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FingerInfo true "更新FingerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /finger_info/updateFingerInfo [put]
export const updateFingerInfo = (data) => {
  return service({
    url: '/finger_info/updateFingerInfo',
    method: 'put',
    data
  })
}

// @Tags FingerInfo
// @Summary 用id查询FingerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FingerInfo true "用id查询FingerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /finger_info/findFingerInfo [get]
export const findFingerInfo = (params) => {
  return service({
    url: '/finger_info/findFingerInfo',
    method: 'get',
    params
  })
}

// @Tags FingerInfo
// @Summary 分页获取FingerInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FingerInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /finger_info/getFingerInfoList [post]
export const getFingerInfoList = (data) => {
  return service({
    url: '/finger_info/getFingerInfoList',
    method: 'post',
    data
  })
}
