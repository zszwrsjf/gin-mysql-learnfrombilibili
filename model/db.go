package model

import (
	"database/sql"
	"fmt"
	"goback/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() {
	db, err := sql.Open(utils.Db, fmt.Sprint("%s:%s@(%s:%s)/%s",
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

	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println("连接失败")
		panic(err)
	}
	fmt.Println("连接成功")
}
