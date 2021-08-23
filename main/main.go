package main

import (
	"fmt"
	"hiveview"
	"hiveview/router"
	"hiveview/utils"
)

// @title 运维平台
// @version 1.0.0
// @description 运维平台
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
// @BasePath /api/
func main() {
	//初始化配置文件
	hiveview.CONFIG.InitConfig("../config/settings.yaml")
	//初始化应用logger
	hiveview.CONFIG.InitLogger()
	//初始化访问日志logger
	hiveview.CONFIG.InitAccessLogger()
	//初始化数据库连接
	hiveview.CONFIG.InitDB()
	//数据库建表
	hiveview.CONFIG.Migrate()
	//初始化GIN
	hiveview.CONFIG.InitGin()
	//初始化enforcer
	err := hiveview.CONFIG.InitEnforcer()
	if err != nil {
		fmt.Println(err)
		return
	}
	//初始化router
	router.InitRouter(hiveview.CONFIG.Gin)
	//初始化admin用户
	hiveview.InitUser()
	//初始化CRON
	cron := hiveview.InitCron()
	utils.InitCronJob(cron)
	//utils.AnsibleAdhoc()
	////运行gin
	err = hiveview.CONFIG.Gin.Run(fmt.Sprintf("%s:%s", hiveview.CONFIG.Settings.Application.Host, hiveview.CONFIG.Settings.Application.Port))
	if err != nil {
		utils.LogPrint("err", err)
	}

}
