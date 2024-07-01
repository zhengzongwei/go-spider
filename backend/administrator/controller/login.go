package controller

import (
	"backend/administrator"
	mongo_helper "backend/mongo-helper"
	"backend/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var jwt = utils.JwtUtil{}

type LoginController struct{}

type Users struct {
	Uuid      string `json:"uuid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LoginTime string `json:"login_time"`
}

func (u Users) String() string {
	return fmt.Sprintf("Uuid: %s, Username: %s, Password: %s, LoginTime: %s", u.Uuid, u.Username, u.Password, u.LoginTime)
}

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
	username := c.PostForm("username")
	password := c.PostForm("password")
	mongodb := mongo_helper.GetMongoHelper().GetConnection().Client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := &Users{}
	if err := mongodb.Database(administrator.DB).Collection("sys_user").FindOne(ctx, bson.M{"username": username}).Decode(user); err != nil {
		log.Println(err.Error())
		administrator.ErrorRep(c, "用户名或密码错误!")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		administrator.ErrorRep(c, "用户名或密码错误.")
		return
	}
	// 生成JWT凭证
	administrator.DataRep(c, jwt.GetToken(user.Uuid, user.Username, 120*time.Minute))
}

func (l *LoginController) logout(c *gin.Context) {
	// 执行登出方法
}
