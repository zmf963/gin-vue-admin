import service from '@/utils/request'

// @Tags EmailInfo
// @Summary 创建EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmailInfo true "创建EmailInfo"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /email_info/createEmailInfo [post]
export const createEmailInfo = (data) => {
  return service({
    url: '/email_info/createEmailInfo',
    method: 'post',
    data
  })
}

// @Tags EmailInfo
// @Summary 删除EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EmailInfo true "删除EmailInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /email_info/deleteEmailInfo [delete]
export const deleteEmailInfo = (params) => {
  return service({
    url: '/email_info/deleteEmailInfo',
    method: 'delete',
    params
  })
}

// @Tags EmailInfo
// @Summary 批量删除EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmailInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /email_info/deleteEmailInfo [delete]
export const deleteEmailInfoByIds = (data) => {
  return service({
    url: '/email_info/deleteEmailInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags EmailInfo
// @Summary 更新EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmailInfo true "更新EmailInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /email_info/updateEmailInfo [put]
export const updateEmailInfo = (data) => {
  return service({
    url: '/email_info/updateEmailInfo',
    method: 'put',
    data
  })
}

// @Tags EmailInfo
// @Summary 用id查询EmailInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EmailInfo true "用id查询EmailInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /email_info/findEmailInfo [get]
export const findEmailInfo = (params) => {
  return service({
    url: '/email_info/findEmailInfo',
    method: 'get',
    params
  })
}

// @Tags EmailInfo
// @Summary 分页获取EmailInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EmailInfo列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /email_info/getEmailInfoList [post]
export const getEmailInfoList = (data) => {
  return service({
    url: '/email_info/getEmailInfoList',
    method: 'post',
    data
  })
}
