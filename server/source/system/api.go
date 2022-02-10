package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

func (a *api) TableName() string {
	return "sys_apis"
}

func (a *api) Initialize() error {
	entities := []system.SysApi{
		{ApiGroup: "base", Method: "POST", Path: "/base/login", Description: "用户登录(必选)"},

		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "jwt加入黑名单(退出，必选)"},

		{ApiGroup: "系统用户", Method: "DELETE", Path: "/user/deleteUser", Description: "删除用户"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/register", Description: "用户注册(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/getUserList", Description: "获取用户列表"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/user/setUserInfo", Description: "设置用户信息"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/user/setSelfInfo", Description: "设置自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "GET", Path: "/user/getUserInfo", Description: "获取自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserAuthorities", Description: "设置权限组"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/changePassword", Description: "修改密码（建(选择)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserAuthority", Description: "修改用户角色(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/resetPassword", Description: "重置用户密码"},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: "创建api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: "删除Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: "更新Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: "获取api列表"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: "获取所有api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: "获取api详细信息"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "批量删除api"},

		{ApiGroup: "角色", Method: "POST", Path: "/authority/copyAuthority", Description: "拷贝角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/createAuthority", Description: "创建角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/deleteAuthority", Description: "删除角色"},
		{ApiGroup: "角色", Method: "PUT", Path: "/authority/updateAuthority", Description: "更新角色信息"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/getAuthorityList", Description: "获取角色列表"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/setDataAuthority", Description: "设置角色资源权限"},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "更改角色api权限"},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表"},

		{ApiGroup: "菜单", Method: "POST", Path: "/menu/addBaseMenu", Description: "新增菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenu", Description: "获取菜单树(必选)"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/deleteBaseMenu", Description: "删除菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/updateBaseMenu", Description: "更新菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getBaseMenuById", Description: "根据id获取菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenuList", Description: "分页获取基础menu列表"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getBaseMenuTree", Description: "获取用户动态路由"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenuAuthority", Description: "获取指定角色menu"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/addMenuAuthority", Description: "增加menu和角色关联关系"},

		// {ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/findFile", Description: "寻找目标文件（秒传）"},
		// {ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "断点续传"},
		// {ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "断点续传完成"},
		// {ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "上传完成移除文件"},

		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "文件上传示例"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "删除文件"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "获取上传文件列表"},

		{ApiGroup: "系统服务", Method: "POST", Path: "/system/getServerInfo", Description: "获取服务器信息"},
		{ApiGroup: "系统服务", Method: "POST", Path: "/system/getSystemConfig", Description: "获取配置文件内容"},
		{ApiGroup: "系统服务", Method: "POST", Path: "/system/setSystemConfig", Description: "设置配置文件内容"},

		// {ApiGroup: "客户", Method: "PUT", Path: "/customer/customer", Description: "更新客户"},
		// {ApiGroup: "客户", Method: "POST", Path: "/customer/customer", Description: "创建客户"},
		// {ApiGroup: "客户", Method: "DELETE", Path: "/customer/customer", Description: "删除客户"},
		// {ApiGroup: "客户", Method: "GET", Path: "/customer/customer", Description: "获取单一客户"},
		// {ApiGroup: "客户", Method: "GET", Path: "/customer/customerList", Description: "获取客户列表"},

		// {ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getDB", Description: "获取所有数据库"},
		// {ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getTables", Description: "获取数据库表"},
		// {ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/createTemp", Description: "自动化代码"},
		// {ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/preview", Description: "预览自动化代码"},
		// {ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getColumn", Description: "获取所选table的所有字段"},

		// {ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/getMeta", Description: "获取meta信息"},
		// {ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/rollback", Description: "回滚自动生成代码"},
		// {ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/getSysHistory", Description: "查询回滚记录"},
		// {ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/delSysHistory", Description: "删除回滚记录"},

		// {ApiGroup: "系统字典详情", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容"},
		// {ApiGroup: "系统字典详情", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容"},
		// {ApiGroup: "系统字典详情", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容"},
		// {ApiGroup: "系统字典详情", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容"},
		// {ApiGroup: "系统字典详情", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表"},

		// {ApiGroup: "系统字典", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "新增字典"},
		// {ApiGroup: "系统字典", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "删除字典"},
		// {ApiGroup: "系统字典", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "更新字典"},
		// {ApiGroup: "系统字典", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "根据ID获取字典"},
		// {ApiGroup: "系统字典", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "获取字典列表"},

		{ApiGroup: "操作记录", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史"},

		// {ApiGroup: "断点续传(插件版)", Method: "POST", Path: "/simpleUploader/upload", Description: "插件版分片上传"},
		// {ApiGroup: "断点续传(插件版)", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "文件完整度验证"},
		// {ApiGroup: "断点续传(插件版)", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "上传完成合并文件"},

		{ApiGroup: "email", Method: "POST", Path: "/email/emailTest", Description: "发送测试邮件"},
		{ApiGroup: "email", Method: "POST", Path: "/email/emailSend", Description: "发送邮件示例"},

		{ApiGroup: "excel", Method: "POST", Path: "/excel/importExcel", Description: "导入excel"},
		{ApiGroup: "excel", Method: "GET", Path: "/excel/loadExcel", Description: "下载excel"},
		{ApiGroup: "excel", Method: "POST", Path: "/excel/exportExcel", Description: "导出excel"},
		{ApiGroup: "excel", Method: "GET", Path: "/excel/downloadTemplate", Description: "下载excel模板"},

		{ApiGroup: "ProjectInfo", Method: "POST", Path: "/project_info/createProjectInfo", Description: "新增项目信息"},
		{ApiGroup: "ProjectInfo", Method: "DELETE", Path: "/project_info/deleteProjectInfo", Description: "删除项目信息"},
		{ApiGroup: "ProjectInfo", Method: "DELETE", Path: "/project_info/delteProjectInfoByIds", Description: "批量删除项目信息"},
		{ApiGroup: "ProjectInfo", Method: "PUT", Path: "/project_info/updateProjectInfo", Description: "更新项目信息"},
		{ApiGroup: "ProjectInfo", Method: "GET", Path: "/project_info/findProjectInfo", Description: "根据ID获取项目信息"},
		{ApiGroup: "ProjectInfo", Method: "POST", Path: "/project_info/getProjectInfoList", Description: "获取项目信息列表"},

		
		
		{ApiGroup: "项目信息", Method: "POST", Path: "/project_info/createProjectInfo", Description: "新增项目信息"},
		{ApiGroup: "项目信息", Method: "DELETE", Path: "/project_info/deleteProjectInfo", Description: "删除项目信息"},
		{ApiGroup: "项目信息", Method: "DELETE", Path: "/project_info/delteProjectInfoByIds", Description: "批量删除项目信息"},
		{ApiGroup: "项目信息", Method: "PUT", Path: "/project_info/updateProjectInfo", Description: "更新项目信息"},
		{ApiGroup: "项目信息", Method: "GET", Path: "/project_info/findProjectInfo", Description: "根据ID获取项目信息"},
		{ApiGroup: "项目信息", Method: "POST", Path: "/project_info/getProjectInfoList", Description: "获取项目信息列表"},
		
		{ApiGroup: "任务", Method: "POST", Path: "/task/createTask", Description: "新增任务"},
		{ApiGroup: "任务", Method: "DELETE", Path: "/task/deleteTask", Description: "删除任务"},
		{ApiGroup: "任务", Method: "DELETE", Path: "/task/delteTaskByIds", Description: "批量删除任务"},
		{ApiGroup: "任务", Method: "PUT", Path: "/task/updateTask", Description: "更新任务"},
		{ApiGroup: "任务", Method: "GET", Path: "/task/findTask", Description: "根据ID获取任务"},
		{ApiGroup: "任务", Method: "POST", Path: "/task/getTaskList", Description: "获取任务列表"},
		
		{ApiGroup: "目标管理", Method: "POST", Path: "/target/createTarget", Description: "新增目标管理"},
		{ApiGroup: "目标管理", Method: "DELETE", Path: "/target/deleteTarget", Description: "删除目标管理"},
		{ApiGroup: "目标管理", Method: "DELETE", Path: "/target/delteTargetByIds", Description: "批量删除目标管理"},
		{ApiGroup: "目标管理", Method: "PUT", Path: "/target/updateTarget", Description: "更新目标管理"},
		{ApiGroup: "目标管理", Method: "GET", Path: "/target/findTarget", Description: "根据ID获取目标管理"},
		{ApiGroup: "目标管理", Method: "POST", Path: "/target/getTargetList", Description: "获取目标管理列表"},
		
		{ApiGroup: "TargetRelation", Method: "POST", Path: "/target_relation/createTargetRelation", Description: "新增TargetRelation"},
		{ApiGroup: "TargetRelation", Method: "DELETE", Path: "/target_relation/deleteTargetRelation", Description: "删除TargetRelation"},
		{ApiGroup: "TargetRelation", Method: "DELETE", Path: "/target_relation/delteTargetRelationByIds", Description: "批量删除TargetRelation"},
		{ApiGroup: "TargetRelation", Method: "PUT", Path: "/target_relation/updateTargetRelation", Description: "更新TargetRelation"},
		{ApiGroup: "TargetRelation", Method: "GET", Path: "/target_relation/findTargetRelation", Description: "根据ID获取TargetRelation"},
		{ApiGroup: "TargetRelation", Method: "POST", Path: "/target_relation/getTargetRelationList", Description: "获取TargetRelation列表"},
		
		{ApiGroup: "Domain", Method: "POST", Path: "/domain/createDomain", Description: "新增Domain"},
		{ApiGroup: "Domain", Method: "DELETE", Path: "/domain/deleteDomain", Description: "删除Domain"},
		{ApiGroup: "Domain", Method: "DELETE", Path: "/domain/delteDomainByIds", Description: "批量删除Domain"},
		{ApiGroup: "Domain", Method: "PUT", Path: "/domain/updateDomain", Description: "更新Domain"},
		{ApiGroup: "Domain", Method: "GET", Path: "/domain/findDomain", Description: "根据ID获取Domain"},
		{ApiGroup: "Domain", Method: "POST", Path: "/domain/getDomainList", Description: "获取Domain列表"},
		
		{ApiGroup: "端口信息", Method: "POST", Path: "/port_info/createPortInfo", Description: "新增端口信息"},
		{ApiGroup: "端口信息", Method: "DELETE", Path: "/port_info/deletePortInfo", Description: "删除端口信息"},
		{ApiGroup: "端口信息", Method: "DELETE", Path: "/port_info/deltePortInfoByIds", Description: "批量删除端口信息"},
		{ApiGroup: "端口信息", Method: "PUT", Path: "/port_info/updatePortInfo", Description: "更新端口信息"},
		{ApiGroup: "端口信息", Method: "GET", Path: "/port_info/findPortInfo", Description: "根据ID获取端口信息"},
		{ApiGroup: "端口信息", Method: "POST", Path: "/port_info/getPortInfoList", Description: "获取端口信息列表"},
		
		{ApiGroup: "路径信息", Method: "POST", Path: "/path_info/createPathInfo", Description: "新增路径信息"},
		{ApiGroup: "路径信息", Method: "DELETE", Path: "/path_info/deletePathInfo", Description: "删除路径信息"},
		{ApiGroup: "路径信息", Method: "DELETE", Path: "/path_info/deltePathInfoByIds", Description: "批量删除路径信息"},
		{ApiGroup: "路径信息", Method: "PUT", Path: "/path_info/updatePathInfo", Description: "更新路径信息"},
		{ApiGroup: "路径信息", Method: "GET", Path: "/path_info/findPathInfo", Description: "根据ID获取路径信息"},
		{ApiGroup: "路径信息", Method: "POST", Path: "/path_info/getPathInfoList", Description: "获取路径信息列表"},
		
		{ApiGroup: "Email信息", Method: "POST", Path: "/email_info/createEmailInfo", Description: "新增Email信息"},
		{ApiGroup: "Email信息", Method: "DELETE", Path: "/email_info/deleteEmailInfo", Description: "删除Email信息"},
		{ApiGroup: "Email信息", Method: "DELETE", Path: "/email_info/delteEmailInfoByIds", Description: "批量删除Email信息"},
		{ApiGroup: "Email信息", Method: "PUT", Path: "/email_info/updateEmailInfo", Description: "更新Email信息"},
		{ApiGroup: "Email信息", Method: "GET", Path: "/email_info/findEmailInfo", Description: "根据ID获取Email信息"},
		{ApiGroup: "Email信息", Method: "POST", Path: "/email_info/getEmailInfoList", Description: "获取Email信息列表"},
		
		{ApiGroup: "文档信息", Method: "POST", Path: "/doc_info/createDocInfo", Description: "新增文档信息"},
		{ApiGroup: "文档信息", Method: "DELETE", Path: "/doc_info/deleteDocInfo", Description: "删除文档信息"},
		{ApiGroup: "文档信息", Method: "DELETE", Path: "/doc_info/delteDocInfoByIds", Description: "批量删除文档信息"},
		{ApiGroup: "文档信息", Method: "PUT", Path: "/doc_info/updateDocInfo", Description: "更新文档信息"},
		{ApiGroup: "文档信息", Method: "GET", Path: "/doc_info/findDocInfo", Description: "根据ID获取文档信息"},
		{ApiGroup: "文档信息", Method: "POST", Path: "/doc_info/getDocInfoList", Description: "获取文档信息列表"},
		
		{ApiGroup: "Keys", Method: "POST", Path: "/keys/createKeys", Description: "新增Keys"},
		{ApiGroup: "Keys", Method: "DELETE", Path: "/keys/deleteKeys", Description: "删除Keys"},
		{ApiGroup: "Keys", Method: "DELETE", Path: "/keys/delteKeysByIds", Description: "批量删除Keys"},
		{ApiGroup: "Keys", Method: "PUT", Path: "/keys/updateKeys", Description: "更新Keys"},
		{ApiGroup: "Keys", Method: "GET", Path: "/keys/findKeys", Description: "根据ID获取Keys"},
		{ApiGroup: "Keys", Method: "POST", Path: "/keys/getKeysList", Description: "获取Keys列表"},
		
		{ApiGroup: "公众号", Method: "POST", Path: "/wechat_official_account/createWechatOfficialAccount", Description: "新增公众号"},
		{ApiGroup: "公众号", Method: "DELETE", Path: "/wechat_official_account/deleteWechatOfficialAccount", Description: "删除公众号"},
		{ApiGroup: "公众号", Method: "DELETE", Path: "/wechat_official_account/delteWechatOfficialAccountByIds", Description: "批量删除公众号"},
		{ApiGroup: "公众号", Method: "PUT", Path: "/wechat_official_account/updateWechatOfficialAccount", Description: "更新公众号"},
		{ApiGroup: "公众号", Method: "GET", Path: "/wechat_official_account/findWechatOfficialAccount", Description: "根据ID获取公众号"},
		{ApiGroup: "公众号", Method: "POST", Path: "/wechat_official_account/getWechatOfficialAccountList", Description: "获取公众号列表"},
		
		{ApiGroup: "应用信息", Method: "POST", Path: "/app_info/createAppInfo", Description: "新增应用信息"},
		{ApiGroup: "应用信息", Method: "DELETE", Path: "/app_info/deleteAppInfo", Description: "删除应用信息"},
		{ApiGroup: "应用信息", Method: "DELETE", Path: "/app_info/delteAppInfoByIds", Description: "批量删除应用信息"},
		{ApiGroup: "应用信息", Method: "PUT", Path: "/app_info/updateAppInfo", Description: "更新应用信息"},
		{ApiGroup: "应用信息", Method: "GET", Path: "/app_info/findAppInfo", Description: "根据ID获取应用信息"},
		{ApiGroup: "应用信息", Method: "POST", Path: "/app_info/getAppInfoList", Description: "获取应用信息列表"},
		
		
		{ApiGroup: "POC", Method: "POST", Path: "/poc_info/createPocInfo", Description: "新增POC"},
		{ApiGroup: "POC", Method: "DELETE", Path: "/poc_info/deletePocInfo", Description: "删除POC"},
		{ApiGroup: "POC", Method: "DELETE", Path: "/poc_info/deltePocInfoByIds", Description: "批量删除POC"},
		{ApiGroup: "POC", Method: "PUT", Path: "/poc_info/updatePocInfo", Description: "更新POC"},
		{ApiGroup: "POC", Method: "GET", Path: "/poc_info/findPocInfo", Description: "根据ID获取POC"},
		{ApiGroup: "POC", Method: "POST", Path: "/poc_info/getPocInfoList", Description: "获取POC列表"},
		
		
		{ApiGroup: "指纹", Method: "POST", Path: "/finger_info/createFingerInfo", Description: "新增指纹"},
		{ApiGroup: "指纹", Method: "DELETE", Path: "/finger_info/deleteFingerInfo", Description: "删除指纹"},
		{ApiGroup: "指纹", Method: "DELETE", Path: "/finger_info/delteFingerInfoByIds", Description: "批量删除指纹"},
		{ApiGroup: "指纹", Method: "PUT", Path: "/finger_info/updateFingerInfo", Description: "更新指纹"},
		{ApiGroup: "指纹", Method: "GET", Path: "/finger_info/findFingerInfo", Description: "根据ID获取指纹"},
		{ApiGroup: "指纹", Method: "POST", Path: "/finger_info/getFingerInfoList", Description: "获取指纹列表"},
		


	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *api) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("path = ? AND method = ?", "/excel/downloadTemplate", "GET").First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
