package middleware

import (
	"net/http"
	"time"

	"com.motion.wallet/models"
	"github.com/gin-gonic/gin"
)

//ClientAuthConfirm 解析并确权
func ClientAuthConfirm(c *gin.Context) {
	ParseToken(c)
	if !c.GetBool("auth") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}

//ClientAuthParse 解析并放行
func ClientAuthParse(c *gin.Context) {
	ParseToken(c)
	c.Next()
}

//ParseToken 解析用户信息
func ParseToken(c *gin.Context) {
	tokenStr := c.GetHeader("token")
	valid := false
	var uid int64
	var expired int64
	var tYpe string
	if tokenStr != "" {
		uid, expired, tYpe, valid = models.ParseToken(tokenStr)
		if expired < time.Now().Unix() {
			valid = false
		}
	}
	if valid {
		switch tYpe {
		case "user":
			c.Set("auth", true)
		}
		c.Set("uid", uid)
		if expired < time.Now().Add(time.Hour*(-24)).Unix() { //token续期
			tokenStr = models.CreateToken(uid, tYpe, time.Hour*24*30)
			c.Header("token", tokenStr)
		}
	}
}
