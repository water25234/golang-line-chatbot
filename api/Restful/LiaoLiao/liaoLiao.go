package liaoliao

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	liaoliaoB "github.com/water25234/golang-line-chatbot/business/liaoliao"
	linebotB "github.com/water25234/golang-line-chatbot/business/linebot"
	"github.com/water25234/golang-line-chatbot/config"
)

type liaoLiaoAPI struct {
	liaoliaoB liaoliaoB.Business
}

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

	linebot := linebotB.New()
	linebotEvents, err := linebot.BindingLineBotJSON(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	liaoliao := liaoliaoB.New(bot)
	liaoliao.Message(linebotEvents)
}
