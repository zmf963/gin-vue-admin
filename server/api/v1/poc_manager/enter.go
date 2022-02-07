package poc_manager

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct { 
    PocInfoApi
}

var ( 
    poc_infoService = service.ServiceGroupApp.Poc_managerServiceGroup.PocInfoService
)