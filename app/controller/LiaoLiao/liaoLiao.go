package liaoliao

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	liaoliao "github.com/water25234/golang-line-chatbot/app/server/liaoLiao"
)

// GetLiaoLiaoMessage mean get liaoliao message
func GetLiaoLiaoMessage(ctx *gin.Context) {
	fmt.Println("12345")

	// liaoliao.Message(ctx)

	ctx.JSON(
		http.StatusOK,
		map[string]string{"hello": "world"},
	)
}

// PostLiaoLiaoMessage mean get send laiolaio message
func PostLiaoLiaoMessage(ctx *gin.Context) {
	fmt.Println("67890")

	liaoliao.Message(ctx)

}
