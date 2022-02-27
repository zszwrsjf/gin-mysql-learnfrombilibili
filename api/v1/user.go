package v1

import (
	"github.com/gin-gonic/gin"
	"goback/model"
	"net/http"
	"goback/utils/errmsg"
)

// add user
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	} 
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
// search user(single or list)

// edit user
func EditUser(c *gin.Context) {

}
// delete user
func DeleteUser(c *gin.Context) {

}
// find whether user existed
func UserExist(c *gin.Context) {

}