package controller

import (
	"backend/administrator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
}

func (this TestController) Claim() *administrator.Controller {
	return &administrator.Controller{
		Prefix: "/debug-test",
		Handles: []administrator.Handle{
			{Url: "/", Method: http.MethodGet, Handle: this.test},
		},
	}
}

func (this TestController) test(c *gin.Context) {
	//client := mongo_helper.GetMongoHelper().GetClientHelper().Client
	administrator.DefaultRep(c)
}
