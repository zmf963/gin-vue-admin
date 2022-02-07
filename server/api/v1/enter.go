package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/project"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/poc_manager"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	
	ProjectApiGroup project.ApiGroup
	Poc_managerApiGroup poc_manager.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
