package api

import (
	"fmt"
	"gomod/controller/api/params"
	"gomod/models"
	"gomod/utils"
	"gomod/utils/logger"
	"log"
	"net/http"
	"strconv"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	myjwt "gomod/middleware/jwt"
)

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	Uid   int    `json:"uid"`
}

// AccountLogin 用户登录
func AccountLogin(c *gin.Context) {
	logId := c.GetString(utils.LogIdParam)
	//ctx := make(map[string]interface{})
	//c.String(http.StatusOK, strings.ToUpper("index"))

	// c.Set("response", gin.H{
	// 	"code": 1,
	// 	"msg":  "参数错误",
	// 	"data": "",
	// })
	// return
	logger.Info(logId, "xxx")
	if c.Request.Method == "POST" {
		var p params.LoginParamsJson
		if err := c.ShouldBind(&p); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "参数错误:" + err.Error(),
				"data": "",
			})
			return
		}
		// 处理 登陆
		username := p.Username
		password := p.Password
		if username != "" && password != "" {
			var account models.Account
			err := Db.Get(&account, "select * from account where username=?", username)
			if err != nil {
				fmt.Println("exec failed, ", err)
				return
			}

			if account.Password != password {
				c.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg":  "用户名或密码错误",
					"data": "",
				})
				return
			}

			// 生成token
			generateToken(c, account)
			return
		}

	}

}

// 生成令牌
func generateToken(c *gin.Context, account models.Account) {
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := myjwt.CustomClaims{
		strconv.Itoa(account.Id),
		account.Username,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		Uid:   account.Id,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}
