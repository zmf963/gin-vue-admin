/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-01 21:23:50
 * @LastEditors: zmf96
 * @LastEditTime: 2022-02-08 17:14:05
 * @FilePath: /server/main.go
 * @Description:
 */
package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
// @securityDefinitions.basic BasicAuth
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()                    // 初始化Viper
	global.GVA_LOG = core.Zap()                     // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()               // gorm连接数据库
	global.Mongo_DB = initialize.MongoInitConnect() // mongo连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
