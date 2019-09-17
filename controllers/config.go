package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Unknwon/goconfig"
)

var config *goconfig.ConfigFile

func loadConfig() {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	fmt.Println(appConfigPath)
	config, err = goconfig.LoadConfigFile(appConfigPath)
	RunMode, err = config.GetValue(goconfig.DEFAULT_SECTION, "runmode")
	if err != nil {
		RunMode = goconfig.DEFAULT_SECTION
	}
}

//GetConfigValue 读取配置文件
func GetConfigValue(key string) string {

	value, _ := config.GetValue(RunMode, key)
	return value
}
