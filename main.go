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

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
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

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	// 带参数的路由
	// 获取路由匹配的参数
	router.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		db["foo"] = "aaa"
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
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

	}

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
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

	return router
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
