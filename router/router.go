package router

import (
	"github.com/gin-gonic/gin"
	apiRestfulLiao "github.com/water25234/golang-line-chatbot/api/Restful/LiaoLiao"
)

// SetupRouter mean setup router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// api
	LiaoLiao := router.Group("/LiaoLiao")
	{
		LiaoLiao.GET("", apiRestfulLiao.GetLiaoLiaoMessage)
		LiaoLiao.POST("", apiRestfulLiao.PostLiaoLiaoMessage)
	}

	return router
}
