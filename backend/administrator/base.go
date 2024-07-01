package administrator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const DB = "administrator"

type Controller struct {
	Prefix  string
	Handles []Handle
}

type Handle struct {
	Url    string
	Method string
	Handle func(c *gin.Context)
}

type IController interface {
	Claim() (con *Controller)
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *Controller) Init(engine *gin.Engine) {
	for _, h := range c.Handles {
		switch h.Method {
		case http.MethodGet:
			engine.GET(fmt.Sprintf("%s%s", c.Prefix, h.Url), h.Handle)
		case http.MethodPost:
			engine.POST(fmt.Sprintf("%s%s", c.Prefix, h.Url), h.Handle)
		case http.MethodPut:
			engine.PUT(fmt.Sprintf("%s%s", c.Prefix, h.Url), h.Handle)
		case http.MethodDelete:
			engine.DELETE(fmt.Sprintf("%s%s", c.Prefix, h.Url), h.Handle)
		}
	}
}

func DefaultRep(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Msg: "", Data: "请求成功!"})
}

func DataRep(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Msg: "", Data: data})
}

func ErrorRep(c *gin.Context, err string) {
	c.JSON(http.StatusOK, Response{Code: http.StatusInternalServerError, Msg: err, Data: ""})
}
