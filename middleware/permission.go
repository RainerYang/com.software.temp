package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"com.motion.wallet/models"
	"github.com/gin-gonic/gin"
)

const (
	//ADMINPASSWORDSALT 密码盐
	ADMINPASSWORDSALT = "MOTION_ADMIN"
	//ADMINPERMISSIONSUPERADMIN 超级管理员
	ADMINPERMISSIONSUPERADMIN = "superadmin"
	//ADMINPERMISSIONASSET 资产模块
	ADMINPERMISSIONASSET = "asset"
	//ADMINPERMISSIONWITHDRAW 提现模块
	ADMINPERMISSIONWITHDRAW = "withdraw"
	//ADMINPERMISSIONUSER 用户模块
	ADMINPERMISSIONUSER = "user"
)

//CheckAdminSuperAdminPermission 超级管理员
func CheckAdminSuperAdminPermission(c *gin.Context) {
	checkAdminPermission(c, ADMINPERMISSIONSUPERADMIN)
}

//CheckAdminAssetPermission 资产模块
func CheckAdminAssetPermission(c *gin.Context) {
	checkAdminPermission(c, ADMINPERMISSIONASSET)
}

//CheckAdminWithdrawPermission 提现模块
func CheckAdminWithdrawPermission(c *gin.Context) {
	checkAdminPermission(c, ADMINPERMISSIONWITHDRAW)
}

//CheckAdminUserPermission 用户模块
func CheckAdminUserPermission(c *gin.Context) {
	checkAdminPermission(c, ADMINPERMISSIONUSER)
}

//CheckAdminPermission 查admin权限
func checkAdminPermission(c *gin.Context, permission string) {
	uid := c.GetInt64("uid")
	admin := models.GetAdmin(uid)
	if admin != nil && (strings.Contains(admin.Permission, permission) || strings.Contains(admin.Permission, ADMINPERMISSIONSUPERADMIN)) {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusForbidden)
		fmt.Println()
	}
}
