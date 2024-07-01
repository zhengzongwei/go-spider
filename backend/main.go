/*
 * Copyright (c) 2023 zhengzongwei
 *
 * Licensed under the Mulan Permissive Software License，Version 2;
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://license.coscl.org.cn/MulanPSL2
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

package main

import (
	"backend/administrator"
	"backend/administrator/controller"
	mongo_helper "backend/mongo-helper"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 启动Mongodb服务
	client := mongo_helper.GetMongoHelper().GetConnection()
	defer client.Disconnect()
	// 启动Web服务
	engine := gin.Default()
	engine.Use(cors(), validateToken())
	// 注册Controller
	registry(engine)
	_ = engine.Run(os.Getenv("GO_SPIDER_SERVER"))

	// 启动爬虫服务
	// ...
}

func registry(engine *gin.Engine) {
	cons := []administrator.IController{
		&controller.LoginController{},
	}

	for _, con := range cons {
		con.Claim().Init(engine)
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "*")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		c.Next()
	}
}

var jwt = utils.JwtUtil{}
var white = [1]string{"/"}

func validateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		token := c.GetHeader("go-spider-token")
		check := true
		for _, w := range white {
			if strings.Contains(uri, w) {
				// 白名单，不校验Token
				check = false
				break
			}
		}

		if check && !jwt.ValidateToken(token) {
			c.AbortWithStatusJSON(http.StatusOK, administrator.Response{Code: http.StatusInternalServerError, Msg: "权限不足，请登录!"})
			return
		}

		c.Next()
	}
}
