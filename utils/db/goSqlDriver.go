package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//var dbMap = make(map[string]string)
var db *sql.DB

func init() {
	if db == nil {
		initDb()
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
		// log.Fatal(err)
	}
}

func initDb() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/goBlog?parseTime=true")
	checkErr(err)

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	//defer db.Close()

}
