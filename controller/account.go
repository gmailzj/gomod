package controller

import (
	"encoding/json"
	"fmt"
	"gomod/models"
	"net/http"
	"gomod/utils"

	"github.com/gin-gonic/gin"
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

			var account models.Account
			err := Db.Get(&account, "select * from account where username=?", username)
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
			var jsonStr []byte
			var jsonData interface{}
			// jsonStr, err = json.Marshal(account)
			jsonStr, err = json.Marshal(struct {
				*models.Account
				// Salt   string `db:"salt" json:"salt"`
				// Status string `db:"status" json:"-"`
			}{
				Account: &account,
			})
			if err != nil {

			} else {
				err = json.Unmarshal(jsonStr, jsonData)
				fmt.Println("err:", string(jsonStr))
			}
			fmt.Println(jsonStr)
			c.String(http.StatusOK, string(jsonStr))
			return
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  utils.MD5(password),
				"data": account,
			})
		}

	} else {
		c.HTML(http.StatusOK, "admin/login.tmpl", gin.H{
			"title": "Hello,world",
		})
	}

}

//func AccountLoginGet(c *gin.Context) {
//	c.HTML(http.StatusOK, "account/login.tmpl", gin.H{
//		"title":  "Hello,world",
//	})
//}
