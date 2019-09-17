package api

import (
	"com.software.temp/middleware"
	"github.com/gin-gonic/gin"
)

//AccountAPI account
type AccountAPI struct {
}

//SetRouter 路由
func (api *AccountAPI) SetRouter(e *gin.Engine) {
	e.GET("/:version/accounts/:id", api.GetAccount)
	e.POST("/:version/accounts", api.CreateAccount)
	e.PUT("/:version/accounts/:id", middleware.ClientAuthConfirm, api.UpdateAccount)
	e.DELETE("/:version/accounts/:id", middleware.ClientAuthConfirm, api.DeleteAccount)
}

//GetAccount ...
func (api *AccountAPI) GetAccount(c *gin.Context) {
	responseJSON(c, errAPINotDone)
}

//CreateAccount ...
func (api *AccountAPI) CreateAccount(c *gin.Context) {
	responseJSON(c, errAPINotDone)
}

//UpdateAccount ...
func (api *AccountAPI) UpdateAccount(c *gin.Context) {
	responseJSON(c, errAPINotDone)
}

//DeleteAccount ...
func (api *AccountAPI) DeleteAccount(c *gin.Context) {
	responseJSON(c, errAPINotDone)
}
