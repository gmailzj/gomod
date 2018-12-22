package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	//_ "net/http/pprof"
)

import "utils"
import "utils/uuid"
import (
	"controller"
)

func init() {
	fmt.Println("router init")
}

// GO语言中要提供给外面访问的方法或是结构体必须是首字母大写, 在模板中遍历的时候，field字段也要首字母大写
type Foo struct {
	I int
	J string
}

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	// 使用中间件
	//router.Use(Logger())



	//router.LoadHTMLGlob("templates/*")
	router.LoadHTMLGlob("templates/**/*")
	// 前台路由
	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		uid := uuid.NewV4()
		c.String(http.StatusOK, utils.GetEmptyString()+"index:"+uid.String())
	})

	router.GET("/index", func(c *gin.Context) {

		valuesArr := []int{1, 2, 3, 4, 5}
		var values []Foo
		for i := 0; i < 5; i++ {
			values = append(values, Foo{I: i, J: "value " + strconv.Itoa(i)})
		}

		valuesMap := make(map[string]string)
		valuesMap["language"] = "Go"
		valuesMap["version"] = "1.7.4"

		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "website/index.tmpl", gin.H{
			"title":  "Hello,world",
			"valuesArr":   valuesArr,
			"values": values,
			"valuesMap": valuesMap,
		})
	})

	router.GET("/debug", func(c *gin.Context){

		// Context对象中常用的属性

		// Request *http.Request
		// Writer  ResponseWriter
		// Params Params

		//log.Print("handle log")
		body,_ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("---body/--- \r\n "+string(body))

		fmt.Println("---header/--- \r\n")
		for k,v :=range c.Request.Header {
			fmt.Println(k,v)
		}
		//fmt.Println("header \r\n",c.Request.Header)

		c.JSON(200,gin.H{
			"receive":"1024",
		})
	})

	// 首页
	router.GET("/", controller.Index)

	router.GET("/account/login", controller.AccountLogin)
	router.POST("/account/login", controller.AccountLogin)

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
			c.JSON(http.StatusOK, gin.H{"status": user + json.Value})
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
