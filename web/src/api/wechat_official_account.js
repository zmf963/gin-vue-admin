import service from '@/utils/request'

// @Tags WechatOfficialAccount
// @Summary 创建WechatOfficialAccount
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WechatOfficialAccount true "创建WechatOfficialAccount"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /wechat_official_account/createWechatOfficialAccount [post]
export const createWechatOfficialAccount = (data) => {
  return service({
    url: '/wechat_official_account/createWechatOfficialAccount',
    method: 'post',
    data
  })
}

// @Tags WechatOfficialAccount
// @Summary 删除WechatOfficialAccount
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WechatOfficialAccount true "删除WechatOfficialAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wechat_official_account/deleteWechatOfficialAccount [delete]
export const deleteWechatOfficialAccount = (params) => {
  return service({
    url: '/wechat_official_account/deleteWechatOfficialAccount',
    method: 'delete',
    params
  })
}

// @Tags WechatOfficialAccount
// @Summary 批量删除WechatOfficialAccount
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WechatOfficialAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wechat_official_account/deleteWechatOfficialAccount [delete]
export const deleteWechatOfficialAccountByIds = (data) => {
  return service({
    url: '/wechat_official_account/deleteWechatOfficialAccountByIds',
    method: 'delete',
    data
  })
}

// @Tags WechatOfficialAccount
// @Summary 更新WechatOfficialAccount
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WechatOfficialAccount true "更新WechatOfficialAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wechat_official_account/updateWechatOfficialAccount [put]
export const updateWechatOfficialAccount = (data) => {
  return service({
    url: '/wechat_official_account/updateWechatOfficialAccount',
    method: 'put',
    data
  })
}

// @Tags WechatOfficialAccount
// @Summary 用id查询WechatOfficialAccount
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WechatOfficialAccount true "用id查询WechatOfficialAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wechat_official_account/findWechatOfficialAccount [get]
export const findWechatOfficialAccount = (params) => {
  return service({
    url: '/wechat_official_account/findWechatOfficialAccount',
    method: 'get',
    params
  })
}

// @Tags WechatOfficialAccount
// @Summary 分页获取WechatOfficialAccount列表
// @Security ApiKeyAuth
// @Security BasicAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WechatOfficialAccount列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /wechat_official_account/getWechatOfficialAccountList [post]
export const getWechatOfficialAccountList = (data) => {
  return service({
    url: '/wechat_official_account/getWechatOfficialAccountList',
    method: 'post',
    data
  })
}
