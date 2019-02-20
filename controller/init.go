package controller

import (
	//"database/sql"

	"utils/db"

	"github.com/jmoiron/sqlx"
)

//var Db *sql.DB
var Db *sqlx.DB

//控制器启动的时候 数据库连接初始化
func init() {
	Db = db.Db
	//dbSqlx = db.Db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
