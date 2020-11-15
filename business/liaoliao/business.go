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
						im.WebHookHandleMessage(job.event)
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

// WebHookHandleMessage mean start handle message
func (im *imple) WebHookHandleMessage(event *linebot.Event) {
	var sentence Sentence
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		if message.Text == "liaoliao --help" {
			sentence = &SendCommandLineContent{
				Command: "line chatbot, commands: liaoliao --help, translate-tw, translate-en",
			}
		} else if strings.Contains(message.Text, "translate-tw") {
			sentence = &SendTranslateContent{
				Sentence:   message.Text,
				SourceLang: "en",
				TargetLang: "zh-tw",
				Region:     "us-east-1",
			}
		} else if strings.Contains(message.Text, "translate-en") {
			sentence = &SendTranslateContent{
				Sentence:   message.Text,
				SourceLang: "zh-tw",
				TargetLang: "en",
				Region:     "us-east-1",
			}
		}
		sentence.HandleSentence(event.ReplyToken, im.bot)
	}
	sentence = nil
}

// HandleSentence mean send command line content logic
func (sc *SendCommandLineContent) HandleSentence(replyToken string, bot *linebot.Client) {
	_, err := bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(sc.Command),
	).Do()

	if err != nil {
		log.Printf("error: %v", err)
	}
}

// HandleSentence mean send translate content logic
func (sc *SendTranslateContent) HandleSentence(replyToken string, bot *linebot.Client) {
	cmd := strings.Fields(sc.Sentence)
	trText := strings.Join(append(cmd[1:]), " ")

	_, err := bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(
			tl.Translate(trText, sc.SourceLang, sc.TargetLang, sc.Region),
		),
	).Do()

	if err != nil {
		log.Print(err)
	}
}
