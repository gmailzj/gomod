package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)


var dbMap = make(map[string]string)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	// H is a shortcut for map[string]interface{}
	// type H map[string]interface{}

	// Running in "debug" mode. Switch to "release" mode in production.
	// - using env:	export GIN_MODE=release
	// - using code:	gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("server start ...")

	initDb()
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

func initDb() {
	var err error
	db, err = sql.Open("sqlite3", "./db.db")
	checkErr(err)

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)


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


