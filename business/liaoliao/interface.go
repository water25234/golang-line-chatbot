package liaoliao

import (
	"github.com/line/line-bot-sdk-go/linebot"
	linebotE "github.com/water25234/golang-line-chatbot/entity/linebot"
)

// Business describe liaoliao business service function
type Business interface {

	// Message mean liaoliao business logic
	Message(linebotEvents *linebotE.Events)
}

// SendCommandLineContent describe send command line content
type SendCommandLineContent struct {
	Command string
}

// SendTranslateContent describe send translate content
type SendTranslateContent struct {
	Sentence   string
	SourceLang string
	TargetLang string
	Region     string
}

// Sentence describe sentence of interface
type Sentence interface {

	// HandleSentence mean handle sentence logic
	HandleSentence(replyToken string, bot *linebot.Client)
}
