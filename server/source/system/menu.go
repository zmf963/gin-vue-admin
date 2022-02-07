package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

func (m *menu) TableName() string {
	return "sys_base_menus"
}

func (m *menu) Initialize() error {
	entities := []system.SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: system.Meta{Title: "关于我们", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: system.Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: system.Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "system", Name: "system", Component: "view/superAdmin/system/system.vue", Sort: 3, Meta: system.Meta{Title: "系统配置", Icon: "operation"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: system.Meta{Title: "操作历史", Icon: "pie-chart"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: system.Meta{Title: "官方网站", Icon: "home-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 6, Meta: system.Meta{Title: "服务器状态", Icon: "cloudy"}},
		
		
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "project", Name: "project", Component: "view/project/index.vue", Sort: 1, Meta: system.Meta{Title: "项目", Icon: "operation"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "poc_manager", Name: "poc_manager", Component: "view/poc_manager/index.vue", Sort: 1, Meta: system.Meta{Title: "POC", Icon: "operation"}},
		
		
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "project_info", Name: "project_info", Component: "view/project/project_info/project_info.vue", Sort: 1, Meta: system.Meta{Title: "项目管理", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "task", Name: "task", Component: "view/project/task/task.vue", Sort: 1, Meta: system.Meta{Title: "任务", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "target", Name: "target", Component: "view/project/target/target.vue", Sort: 1, Meta: system.Meta{Title: "目标管理", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "target_relation", Name: "target_relation", Component: "view/project/target_relation/target_relation.vue", Sort: 1, Meta: system.Meta{Title: "TargetRelation", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "domain", Name: "domain", Component: "view/project/domain/domain.vue", Sort: 1, Meta: system.Meta{Title: "Domain", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "port_info", Name: "port_info", Component: "view/project/port_info/port_info.vue", Sort: 1, Meta: system.Meta{Title: "端口信息", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "path_info", Name: "path_info", Component: "view/project/path_info/path_info.vue", Sort: 1, Meta: system.Meta{Title: "路径信息", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "email_info", Name: "email_info", Component: "view/project/email_info/email_info.vue", Sort: 1, Meta: system.Meta{Title: "Email信息", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "doc_info", Name: "doc_info", Component: "view/project/doc_info/doc_info.vue", Sort: 1, Meta: system.Meta{Title: "文档信息", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "keys", Name: "keys", Component: "view/project/keys/keys.vue", Sort: 1, Meta: system.Meta{Title: "Keys", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "wechat_official_account", Name: "wechat_official_account", Component: "view/project/wechat_official_account/wechat_official_account.vue", Sort: 1, Meta: system.Meta{Title: "公众号", Icon: "operation", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "13", Path: "app_info", Name: "app_info", Component: "view/project/app_info/app_info.vue", Sort: 1, Meta: system.Meta{Title: "应用信息", Icon: "operation", KeepAlive: true}},
		
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "poc_info", Name: "poc_info", Component: "view/poc_manager/poc_info/poc_info.vue", Sort: 1, Meta: system.Meta{Title: "POC", Icon: "operation", KeepAlive: true}},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.Wrap(err, m.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (m *menu) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("path = ?", "autoCodeEdit/:id").First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
