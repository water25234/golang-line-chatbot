package liaoliao

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	tl "github.com/water25234/golang-line-chatbot/common/translate"
	"github.com/water25234/golang-line-chatbot/config"
)

var bot *linebot.Client

// Message mean liaoliao business logic
func Message(ctx *gin.Context) {
	bot, _ = linebot.New(
		config.GetAppConfig().LineChannelSecret,
		config.GetAppConfig().LineChannelAccessToken)

	events, err := getLineEvents(ctx)
	if err != nil {
		log.Printf("error: %v", err)
	}

	for _, event := range events {
		log.Printf("Event ReplyToken: %v", event.ReplyToken)
		if event.Type == linebot.EventTypeMessage {
			switch event.Type {
			case linebot.EventTypeMessage:
				handleMessage(event)
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

func handleMessage(event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		if message.Text == "liaoliao --help" {
			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("line chatbot, commands: liaoliao --help, translate-tw, translate-en")).Do()
			if err != nil {
				log.Printf("error: %v", err)
			}
		} else if strings.Contains(message.Text, "translate-tw") {
			cmd := strings.Fields(message.Text)
			trText := strings.Join(append(cmd[1:]), " ")

			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tl.Translate(trText, "en", "zh-tw", "us-east-1"))).Do()
			if err != nil {
				log.Print(err)
			}
		} else if strings.Contains(message.Text, "translate-en") {
			cmd := strings.Fields(message.Text)
			trText := strings.Join(append(cmd[1:]), " ")

			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tl.Translate(trText, "zh-tw", "en", "us-east-1"))).Do()
			if err != nil {
				log.Print(err)
			}
		}
	}
}
