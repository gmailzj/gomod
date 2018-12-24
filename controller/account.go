package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"utils"

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

			stmt, err := Db.Prepare("SELECT * FROM account WHERE username=? AND status = '1'")
			CheckErr(err)

			rows, err := stmt.Query(username)
			CheckErr(err)

			cloumns, err := rows.Columns()

			values := make([]sql.RawBytes, len(cloumns))
			scanArgs := make([]interface{}, len(values))
			for i := range values {
				scanArgs[i] = &values[i]
			}
			for rows.Next() {
				err = rows.Scan(scanArgs...)
				if err != nil {
					log.Fatal(err)
				}
				var value string
				for i, col := range values {
					if col == nil {
						value = "NULL"
					} else {
						value = string(col)
					}
					fmt.Println(cloumns[i], ": ", value)
				}
				fmt.Println("------------------")
			}

			//for rows.Next() {
			//	var password string
			//	var salt string
			//
			//	err = rows.Scan(&password, &salt)
			//	CheckErr(err)
			//	fmt.Println(password)
			//	fmt.Println(salt)
			//}

			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  utils.MD5(password),
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
