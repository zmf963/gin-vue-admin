package finger_manager

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct { 
    FingerInfoApi
}

var ( 
    finger_infoService = service.ServiceGroupApp.Finger_managerServiceGroup.FingerInfoService
)