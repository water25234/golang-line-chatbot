package linebot

import "github.com/line/line-bot-sdk-go/linebot"

// Events mean Line bot entity by struct
type Events struct {
	Events []*linebot.Event `json:"events"`
}
