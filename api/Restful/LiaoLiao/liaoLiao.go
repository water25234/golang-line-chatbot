package liaoliao

import (
	"net/http"

	"github.com/gin-gonic/gin"
	liaoB "github.com/water25234/golang-line-chatbot/business/liaoliao"
)

// GetLiaoLiaoMessage mean get liaoliao message
func GetLiaoLiaoMessage(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		map[string]string{"hello": "world"},
	)
}

// PostLiaoLiaoMessage mean get send laiolaio message
func PostLiaoLiaoMessage(ctx *gin.Context) {
	liaoB.Message(ctx)
}
