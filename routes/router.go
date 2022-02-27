package routes

import (
	"goback/utils"
	// "net/http"
	"goback/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() { // 大写是public 小写是private
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerv1 := r.Group("api/v1")
	{
		// user
		routerv1.POST("user/add", v1.AddUser)
		routerv1.PUT("user/:id", v1.EditUser)
		routerv1.DELETE("user/:id", v1.DeleteUser)
		// tag

		// article

	}

	r.Run(utils.HttpPort)
}
