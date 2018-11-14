package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	//加载模板
	router.LoadHTMLGlob("templates/*")
	// router.LoadHTMLFiles("templates/index.tmpl", "templates/index.html", "templates/index2.html")

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello,world",
		})
	})

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	// 带参数的路由
	// 获取路由匹配的参数
	router.GET("/user/:name", func(c *gin.Context) {
		// 下面两种方式都可以
		user := c.Params.ByName("name")
		user2 := c.Param("name")
		db["foo"] = "aaa" + user2
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// 获取querystring参数
	router.GET("/querystring", func(c *gin.Context) {

		// 获取参数?name=abc,如果没有取默认值
		name := c.DefaultQuery("name", "Guest") //可设置默认值
		// 是 c.Request.URL.Query().Get("lastname") 的简写

		lastname := c.Query("lastname")
		// ct := c.Header("Content-Type","text/html; charset=utf-8")

		// 获取请求头
		headVersion := c.GetHeader("version")
		fmt.Printf("%T", headVersion)
		// 设置响应头
		c.Header("lastname", lastname)
		c.String(http.StatusOK, name)
		// fmt.Println("Hello %s", name)
	})
	// POST 传递参数
	router.POST("post", func(c *gin.Context) {
		// 获取url里面的参数
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		message := c.PostForm("message")
		name := c.PostForm("name")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))

	// 路由群组
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK, "user-v1")
		})
		v1.GET("/list", func(c *gin.Context) {
			c.String(http.StatusOK, "list-v1")
		})

	}

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
		/*
			curl -X "POST" "http://127.0.0.1:8080/admin" \
					-H 'Content-Type: application/json; charset=utf-8' \
					-u 'foo:bar' \
					-d $'{
					"value": "1"
			}'
		*/
		if ok := c.Bind(&json); ok == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			fmt.Println(ok)
		}
	})

	// 设置静态文件目录
	router.Static("/assets", "./assets")
	// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead
	router.StaticFS("/static", http.Dir("./assets"))

	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	return router
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
