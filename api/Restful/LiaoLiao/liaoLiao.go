package liaoliao

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/water25234/golang-line-chatbot/business/liaoliao"
	"github.com/water25234/golang-line-chatbot/config"
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

	bot, _ := linebot.New(
		config.GetAppConfig().LineChannelSecret,
		config.GetAppConfig().LineChannelAccessToken)

	var liaoliao = liaoliao.New(ctx, bot)

	liaoliao.Message()
}
