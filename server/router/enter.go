package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	
	"github.com/flipped-aurora/gin-vue-admin/server/router/project"
	"github.com/flipped-aurora/gin-vue-admin/server/router/poc_manager"
)

type RouterGroup struct {
	System   system.RouterGroup
	
	Project project.RouterGroup
	Poc_manager poc_manager.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
