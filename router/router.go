package router

import (
	"com.motion.wallet/middleware"
	"com.motion.wallet/utils"
	"com.software.temp/api"
	"github.com/gin-gonic/gin"
)

//StartRouter 配置路由
func StartRouter(e *gin.Engine, port int) {
	e.Use(middleware.CheckToken)
	new(api.AccountAPI).SetRouter(e)
	e.Run(":" + utils.ParseString(port))
}
