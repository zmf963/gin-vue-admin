package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	
	"github.com/flipped-aurora/gin-vue-admin/server/service/project"
	"github.com/flipped-aurora/gin-vue-admin/server/service/poc_manager"
	"github.com/flipped-aurora/gin-vue-admin/server/service/finger_manager"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	
	ProjectServiceGroup project.ServiceGroup
	Poc_managerServiceGroup poc_manager.ServiceGroup
	Finger_managerServiceGroup finger_manager.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
