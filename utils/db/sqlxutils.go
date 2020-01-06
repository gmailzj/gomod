package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gomod/config"
	"log"
	"time"
)

var (
	Db *sqlx.DB
)

func init() {
	if Db == nil {
		//settings := conf.Settings
		//db_drivers, _ := settings.String("database", "db_pg_drivers")
		//db_contection, _ := settings.String("database", "db_pg_contection")
		var err error
		//Db, err = sqlx.Connect(db_drivers, db_contection)

		dataSourceName, _ := config.Config.GetValue("mysql", "dsn")
		Db, err = sqlx.Connect("mysql", dataSourceName)

		if err != nil || Db == nil {
			log.Fatalf("sqlx 初始化数据库出错：\n %#v", err)
		}

		Db.SetMaxIdleConns(2)                   //数据库最大闲置数
		Db.SetMaxOpenConns(12)                  //数据库最大连接数
		Db.SetConnMaxLifetime(20 * time.Second) //数据库最大生命周期
	}
}
