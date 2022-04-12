/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-04-12 18:48:53
 * @LastEditors: zmf96
 * @LastEditTime: 2022-04-12 20:29:47
 * @FilePath: /src/api/project_doc.js
 * @Description: 
 */
import service from '@/utils/request'

// @Tags ProjectDoc
// @Summary 创建ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectDoc true "创建ProjectDoc"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /project_doc/createProjectDoc [post]
export const createProjectDoc = (data) => {
  return service({
    url: '/project_doc/createProjectDoc',
    method: 'post',
    data
  })
}

// @Tags ProjectDoc
// @Summary 删除ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProjectDoc true "删除ProjectDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project_doc/deleteProjectDoc [delete]
export const deleteProjectDoc = (params) => {
  return service({
    url: '/project_doc/deleteProjectDoc',
    method: 'delete',
    params
  })
}

// @Tags ProjectDoc
// @Summary 批量删除ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project_doc/deleteProjectDoc [delete]
export const deleteProjectDocByIds = (data) => {
  return service({
    url: '/project_doc/deleteProjectDocByIds',
    method: 'delete',
    data
  })
}

// @Tags ProjectDoc
// @Summary 更新ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectDoc true "更新ProjectDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project_doc/updateProjectDoc [put]
export const updateProjectDoc = (data) => {
  return service({
    url: '/project_doc/updateProjectDoc',
    method: 'put',
    data
  })
}

// @Tags ProjectDoc
// @Summary 用id查询ProjectDoc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProjectDoc true "用id查询ProjectDoc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project_doc/findProjectDoc [get]
export const findProjectDoc = (params) => {
  return service({
    url: '/project_doc/findProjectDoc',
    method: 'get',
    params
  })
}

// @Tags ProjectDoc
// @Summary 分页获取ProjectDoc列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ProjectDoc列表"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"获取成功"}"
// @Router /project_doc/getProjectDocList [post]
export const getProjectDocList = (data) => {
  return service({
    url: '/project_doc/getProjectDocList',
    method: 'post',
    data
  })
}
