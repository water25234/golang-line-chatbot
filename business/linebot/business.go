package linebot

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	linebotE "github.com/water25234/golang-line-chatbot/entity/linebot"
)

// New mean linebot.Business by interface
func New() Business {
	return &imple{}
}

type imple struct {
}

// BindingLineBotJson mean Binding linebot json format from gin context
func (im *imple) BindingLineBotJSON(c *gin.Context) (linebotEvents *linebotE.Events, err error) {
	// result, _ := im.bot.ParseRequest(im.ctx.Request)
	if c.Request.Body == nil {
		log.Printf("c.Request.Body is empty: %v", c.Request.Body)
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	linebotEvents = &linebotE.Events{}
	if err = json.Unmarshal(body, linebotEvents); err != nil {
		return nil, err
	}
	return linebotEvents, nil
}
