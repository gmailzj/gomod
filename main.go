package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)


func main() {

	// H is a shortcut for map[string]interface{}
	// type H map[string]interface{}

	// gin 1.3以上新加了这个功能
	// 如果要以指的格式（例如JSON，Key Values或其他格式）记录信息，则可以使用gin.DebugPrintRouteFunc指定格式。
	// 在下面的示例中，我们使用标准日志包记录所有路由，但您可以使用其他满足需求的日志工具。

	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}

	// Running in "debug" mode. Switch to "release" mode in production.
	// - using env:	export GIN_MODE=release
	// - using code:	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.ReleaseMode)
	fmt.Println("server start ...")

	r := SetupRouter()
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
		// status := c.Writer.Status()
		// log.Println(status)
	}
}


