package model

import (
	"fmt"
	"goback/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

func InitDb() {
	db, err = sqlx.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName))
	if err != nil {
		fmt.Println("错误", err)
		panic(err)
	}
	
	db.SetConnMaxIdleTime(10)
	db.SetConnMaxLifetime(10 * time.Second)
	db.SetMaxOpenConns(100)

	fmt.Println("连接成功")
	
}

func ExitDb() {
	db.Close()
}