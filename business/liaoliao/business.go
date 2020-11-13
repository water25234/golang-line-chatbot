package liaoliao

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	tl "github.com/water25234/golang-line-chatbot/common/translate"
)

type imple struct {
	ctx *gin.Context
	bot *linebot.Client
}

// New mean liaoliao.Business by interface
func New(ctx *gin.Context, bot *linebot.Client) Business {
	return &imple{
		ctx: ctx,
		bot: bot,
	}
}

// Message mean liaoliao business logic
func (im *imple) Message() {

	// result, _ := im.bot.ParseRequest(im.ctx.Request)

	events, err := getLineEvents(im.ctx)
	if err != nil {
		log.Printf("error: %v", err)
	}

	for _, event := range events {
		log.Printf("Event ReplyToken: %v", event.ReplyToken)
		if event.Type == linebot.EventTypeMessage {
			switch event.Type {
			case linebot.EventTypeMessage:
				im.handleMessage(event)
			}
		}
	}
}

func getLineEvents(c *gin.Context) ([]*linebot.Event, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	request := &struct {
		Events []*linebot.Event `json:"events"`
	}{}
	if err = json.Unmarshal(body, request); err != nil {
		return nil, err
	}
	return request.Events, nil
}

func (im *imple) handleMessage(event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		if message.Text == "liaoliao --help" {
			_, err := im.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("line chatbot, commands: liaoliao --help, translate-tw, translate-en")).Do()
			if err != nil {
				log.Printf("error: %v", err)
			}
		} else if strings.Contains(message.Text, "translate-tw") {
			cmd := strings.Fields(message.Text)
			trText := strings.Join(append(cmd[1:]), " ")

			_, err := im.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tl.Translate(trText, "en", "zh-tw", "us-east-1"))).Do()
			if err != nil {
				log.Print(err)
			}
		} else if strings.Contains(message.Text, "translate-en") {
			cmd := strings.Fields(message.Text)
			trText := strings.Join(append(cmd[1:]), " ")

			_, err := im.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tl.Translate(trText, "zh-tw", "en", "us-east-1"))).Do()
			if err != nil {
				log.Print(err)
			}
		}
	}
}
