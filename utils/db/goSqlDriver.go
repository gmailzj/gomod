package db

import (
	"database/sql"

	// mysql package
	_ "github.com/go-sql-driver/mysql"
)

//var dbMap = make(map[string]string)

// Db 对象
var DbNaitve *sql.DB

//func init() {
//	if DbNaitve == nil {
//		initDb()
//	}
//}

//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//		// log.Fatal(err)
//	}
//}
//
//func initDb() {
//	var err error
//	DbNaitve, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/goBlog?parseTime=true")
//	checkErr(err)
//
//	DbNaitve.SetMaxIdleConns(20)
//	DbNaitve.SetMaxOpenConns(20)
//
//	//defer DbNaitve.Close()
//
//}
