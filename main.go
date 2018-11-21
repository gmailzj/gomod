package main

import (
	"database/sql"
	"fmt"
	log "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)
import "utils/uuid"

var dbMap = make(map[string]string)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func initDb() {
	var err error
	db, err = sql.Open("sqlite3", "./db.db")
	checkErr(err)

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	// //插入数据
	// stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	// checkErr(err)

	// res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)
	// id := int64(2)
	// fmt.Println(id)
	// //更新数据
	// stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err := stmt.Exec("astaxieupdate", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	//查询数据
	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)

	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var department string
	// 	var created time.Time
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Print(gin.H{
	// 		"uid":        uid,
	// 		"username":   username,
	// 		"department": department,
	// 		"created":    created,
	// 	})
	// }

	//删除数据
	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// checkErr(err)
	// defer stmt.Close() //关闭之

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/uuid", func(c *gin.Context) {
		// get a UUID instance
		uuidMy := guuid.New()
		str := uuidMy.String()

		uuids := uuid.NewV4()
		fmt.Print("version:", uuids.Version(), "\n")
		c.String(http.StatusOK, str)
	})

	// Get user value
	// 带参数的路由
	// 获取路由匹配的参数
	router.GET("/user/:name", func(c *gin.Context) {
		// 下面两种方式都可以
		user := c.Params.ByName("name")
		user2 := c.Param("name")
		dbMap["foo"] = "aaa" + user2
		value, ok := dbMap[user]
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
		fmt.Printf("%T", name)

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

	// 路由群组
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/user", func(c *gin.Context) {
			//查询数据
			rows, err := db.Query("SELECT * FROM userinfo")
			checkErr(err)

			list := []gin.H{}
			for rows.Next() {
				var uid int
				var username string
				var department string
				var created time.Time
				err = rows.Scan(&uid, &username, &department, &created)
				checkErr(err)
				list = append(list, gin.H{
					"uid":        uid,
					"username":   username,
					"department": department,
					"created":    created,
				})
			}
			c.JSON(http.StatusOK, list)
		})
		v1.GET("/list", func(c *gin.Context) {
			c.String(http.StatusOK, "list-v1")
		})

	}

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
		/*
			curl -X "POST" "http://127.0.0.1:8080/admin" \
					-H 'Content-Type: application/json; charset=utf-8' \
					-u 'foo:bar' \
					-d $'{
					"value": "1"
			}'
		*/
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

	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	router.GET("/json", func(c *gin.Context) {
		// data := []int{1, 2, 3}
		// c.JSON(http.StatusOK, gin.H{"errCode": 0, "msg": "abc", "data": data})

		var msg struct {
			// 右边的tag用来映射返回结果中的key
			Name    string `json:"user" xml:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/xml", func(c *gin.Context) {
		data := []int{1, 2, 3}
		c.XML(http.StatusOK, gin.H{"errCode": 0, "msg": "abc", "data": data})
	})

	// 设置静态文件目录
	router.Static("/assets", "./assets")
	// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead
	router.StaticFS("/static", http.Dir("./assets"))

	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	router.GET("/301", func(c *gin.Context) {
		//支持内部和外部的重定向
		// c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
		c.Redirect(http.StatusMovedPermanently, "/json")
	})

	return router
}

// Logger 中间件定义
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 在gin上下文中定义变量
		c.Set("example", "12345")

		// 请求前

		c.Next() //处理请求

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {

	// H is a shortcut for map[string]interface{}
	// type H map[string]interface{}

	// Running in "debug" mode. Switch to "release" mode in production.
	// - using env:	export GIN_MODE=release
	// - using code:	gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.ReleaseMode)

	initDb()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8081
	// r.Run(":8081")

	// 方法2
	// http.ListenAndServe(":8081", r)

	// 方法3
	s := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
