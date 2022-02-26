package routes

import (
	"goback/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() { // 大写是public 小写是private
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}
