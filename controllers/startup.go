package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

//StartUp 启动
func StartUp() {
	//初始化系统配置
	loadConfig()
	fmt.Println("init config done!")
	log.AddHook(newLfsHook())
	fmt.Println("init log engine done!")
	gin.SetMode(GetConfigValue("mode"))
	fmt.Printf("set system mode to %s.", GetConfigValue("mode"))
	//初始化数据库
	initDatabase()
	fmt.Println("init database engine done!")
	//初始化业务模块
	InitModels()
	//订时任务
	InitCron()
	fmt.Println("init pay model done!")
}

//InitModels 初始化Models
func InitModels() {

}

//InitCron 订时任务
func InitCron() {
	crontab := cron.New()
	// crontab.AddFunc("0/1 * * * * *", func() { CallbackRetry(1) })

	crontab.Start()
}
