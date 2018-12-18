package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
import "utils"
import "utils/uuid"
import (
	"controller"
)

func init(){
	fmt.Println("router init")
}

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	// 使用中间件
	//router.Use(Logger())


	// 前台路由
	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		uid := uuid.NewV4()
		c.String(http.StatusOK, utils.GetEmptyString()+"index:"+uid.String())
	})

	// 首页
	router.GET("/", controller.Index)


	// 管理员路由
	// Authorized group (uses gin.BasicAuth() middleware)
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("/", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if ok := c.Bind(&json); ok == nil {
			dbMap[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			log.Println(ok)
		}
	})

	//加载模板
	// router.LoadHTMLGlob("templates/*")
	// router.LoadHTMLFiles("templates/index.tmpl", "templates/index.html", "templates/index2.html")
	/* router.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello,world",
		})
	}) */

	//router.LoadHTMLGlob("templates/**/*")
	//router.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
	//		"title": "Posts",
	//	})
	//})




	// 设置静态文件目录
	router.Static("/assets", "./assets")
	// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead
	router.StaticFS("/static", http.Dir("./assets"))

	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	//router.GET("/301", func(c *gin.Context) {
	//	//支持内部和外部的重定向
	//	// c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	//	c.Redirect(http.StatusMovedPermanently, "/json")
	//})

	return router
}
