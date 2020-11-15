package liaoliao

import (
	linebotE "github.com/water25234/golang-line-chatbot/entity/linebot"
)

// Business describe liaoliao business service function
type Business interface {

	// Server mean liaoliao business logic
	Message(linebotEvents *linebotE.Events)
}
