package main

import (
	"com.software.temp/controllers"
	"com.software.temp/router"
	"com.software.temp/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	controllers.StartUp()
	r := gin.New()
	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())
	httpport := controllers.GetConfigValue("httpport")
	router.StartRouter(r, utils.ParseInt(httpport))
}
