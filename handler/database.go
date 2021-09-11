package handler

import (
	"database/sql"
	"sea_mod/conf"
	"time"
)

func Connect() *sql.DB {
	var conn string
	config := conf.GetConfiguration()
	conn = config.MysqlUsername + ":" + config.MysqlPassword + "@tcp(" + config.MysqlHost + ":" + config.MysqlPort + ")/" + config.MysqlDatabase
	db, err := sql.Open("mysql", conn)
	//defer db.Close()
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

