package controller

import (
	"backend/administrator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct{}

func (l *LoginController) Claim() (con *administrator.Controller) {
	con = &administrator.Controller{Prefix: "/sys",
		Handles: []administrator.Handle{
			{Url: "/login", Method: http.MethodPost, Handle: l.login},
			{Url: "/logout", Method: http.MethodGet, Handle: l.logout},
		}}
	return
}

func (l *LoginController) login(c *gin.Context) {
	// 执行登录方法
}

func (l *LoginController) logout(c *gin.Context) {
	// 执行登出方法
}
