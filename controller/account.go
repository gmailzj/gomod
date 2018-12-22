package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strings"
)

func AccountLogin(c *gin.Context) {
	//ctx := make(map[string]interface{})
	//c.String(http.StatusOK, strings.ToUpper("index"))
	if c.Request.Method == "POST" {

		// 处理 登陆
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != "" && password != "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
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
