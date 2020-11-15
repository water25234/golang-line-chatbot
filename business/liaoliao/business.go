package liaoliao

import (
	"log"
	"runtime"
	"strings"
	"sync"

	"github.com/line/line-bot-sdk-go/linebot"

	tl "github.com/water25234/golang-line-chatbot/common/translate"
	linebotE "github.com/water25234/golang-line-chatbot/entity/linebot"
)

type imple struct {
	bot *linebot.Client
}

// New mean liaoliao.Business by interface
func New(bot *linebot.Client) Business {
	return &imple{
		bot: bot,
	}
}

type jobChannel struct {
	event *linebot.Event
}

// Message mean liaoliao business logic
func (im *imple) Message(linebotEvents *linebotE.Events) {
	if linebotEvents == nil {
		log.Printf("linebotEvents is empty: %v", linebotEvents)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(linebotEvents.Events))

	jobChans := make(chan jobChannel, len(linebotEvents.Events))

	for i := 0; i <= runtime.NumCPU(); i++ {
		go func() {
			for job := range jobChans {
				log.Printf("Event ReplyToken: %v", job.event.ReplyToken)
				if job.event.Type == linebot.EventTypeMessage {
					switch job.event.Type {
					case linebot.EventTypeMessage:
						im.handleMessage(job.event)
					}
				}
				wg.Done()
			}
		}()
	}

	for _, event := range linebotEvents.Events {
		jobChans <- jobChannel{
			event: event,
		}
	}

	close(jobChans)

	wg.Wait()
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
