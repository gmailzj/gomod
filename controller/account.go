package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"models"
	_ "models"
	"net/http"
	"utils"
	"encoding/json"
)

// AccountLogin 用户登录
func AccountLogin(c *gin.Context) {
	//ctx := make(map[string]interface{})
	//c.String(http.StatusOK, strings.ToUpper("index"))
	if c.Request.Method == "POST" {

		// 处理 登陆
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != "" && password != "" {

			var account []models.Account
			err := Db.Select(&account, "select * from account where username=?", username)
			if err != nil {
				fmt.Println("exec failed, ", err)
				return
			}

			fmt.Println("select succ:", account)

			//for rows.Next() {
			//	var password string
			//	var salt string
			//
			//	err = rows.Scan(&password, &salt)
			//	CheckErr(err)
			//	fmt.Println(password)
			//	fmt.Println(salt)
			//}
			var jsonStr []byte;
			jsonStr, _ = json.Marshal(account);
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  utils.MD5(password),
				"data": (jsonStr),
			})
		}

	} else {
		c.HTML(http.StatusOK, "account/login.tmpl", gin.H{
			"title": "Hello,world",
		})
	}

}

//func AccountLoginGet(c *gin.Context) {
//	c.HTML(http.StatusOK, "account/login.tmpl", gin.H{
//		"title":  "Hello,world",
//	})
//}
