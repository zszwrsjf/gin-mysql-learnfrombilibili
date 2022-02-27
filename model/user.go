package model

import (
	_ "github.com/go-sql-driver/mysql"
	// "github.com/jmoiron/sqlx"
	"goback/utils/errmsg"
	"fmt"
)

type User struct {
	Username string `form:"username" json:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
	// TODO 其他属性
	// Role int
	// Avater
}

// find user exist
func CheckUser(name string) (code int) {
	var id []uint
	err = db.Select(&id, "select id from user where username=?", name)
	if len(id) > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}
// add user
func CreateUser(data *User)int {

	r, err := db.Exec("insert into user(username, password) values (?, ?)", data.Username, data.Password)
    if err != nil {
        fmt.Println("exec failed, ", err)
        return errmsg.ERROR
    }
    id, err := r.LastInsertId()
    if err != nil {
        fmt.Println("exec failed, ", err)
        return errmsg.ERROR
    }

    fmt.Println("insert succ:", id)

	// err := db.Create(&data).Error
	// if err != nil {
	// 	return errmsg.ERROR // 500
	// }
	return errmsg.SUCCESS
}
