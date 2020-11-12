package router

import (
	"github.com/gin-gonic/gin"
	cntrLiaoLiao "github.com/water25234/golang-line-chatbot/app/controller/LiaoLiao"
)

// SetupRouter mean setup router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// api
	LiaoLiao := router.Group("/LiaoLiao")
	{
		LiaoLiao.GET("", cntrLiaoLiao.GetLiaoLiaoMessage)
		LiaoLiao.POST("", cntrLiaoLiao.PostLiaoLiaoMessage)
	}

	return router
}
