package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomod/middleware"
	"gomod/middleware/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	// _ "net/http/pprof"
	"gomod/controller"
	"gomod/controller/api"
	"gomod/utils"
	"gomod/utils/uuid"
)

func init() {
	fmt.Println("router init")
}

// Foo GO语言中要提供给外面访问的方法或是结构体必须是首字母大写, 在模板中遍历的时候，field字段也要首字母大写
type Foo struct {
	I int
	J string
}

var Tmpl *template.Template

// SetupRouter func
func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// gin.Default 默认使用了两个中间件 engine.Use(Logger(), Recovery())
	router := gin.Default()

	router.NoMethod(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "请求方式不允许",
			"data":    map[string]string{},
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    http.StatusNotFound,
			"message": "接口不存在",
			"data":    map[string]string{},
		})
	})

	// 使用中间件
	// router.Use(Logger())
	// router.Use(middleware.Recovery())

	// router.LoadHTMLGlob("templates/*")
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
			"title":     "Hello,world",
			"valuesArr": valuesArr,
			"values":    values,
			"valuesMap": valuesMap,
		})
	})

	router.POST("/debug", func(c *gin.Context) {

		// Context对象中常用的属性

		// Request *http.Request
		// Writer  ResponseWriter
		// Params Params

		//log.Print("handle log")
		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("---body/--- \r\n " + string(body))
		fmt.Println("---header/---")
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		//fmt.Println("header \r\n",c.Request.Header)

		c.JSON(200, gin.H{
			"receive": "1024",
		})
	})

	// 首页
	router.GET("/", controller.Index)

	test := router.Group("/test")
	test.GET("yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey"})
	})
	test.GET("json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey"})
	})

	// API 接口
	apiGroup := router.Group("/api")
	apiGroup.Use(middleware.Response)
	apiGroup.Use(jwt.JWTAuth())

	apiGroup.Any("/", func(c *gin.Context) {
		c.Header("server", "lake/1.0.0")
		c.JSON(200, gin.H{
			"message": "lake service " + time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	apiGroup.POST("/login", api.AccountLogin)
	apiGroup.GET("/user/info", api.UserInfo)

	// 管理员路由
	// Authorized group (uses gin.BasicAuth() middleware)
	admin := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	admin.GET("/login", controller.AccountLogin)
	admin.POST("/login", controller.AccountLogin)

	admin.GET("/", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			// Value string `json:"value" binding:"required"`
			Value string `json:"value"`
		}

		if err := c.Bind(&json); err == nil {
			c.JSON(http.StatusOK, gin.H{"status": user + json.Value})
		} else {
			log.Println(err)
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
