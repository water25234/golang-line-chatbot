package linebot

import (
	"github.com/gin-gonic/gin"

	linebotE "github.com/water25234/golang-line-chatbot/entity/linebot"
)

// Business describe linebot business service function
type Business interface {

	// BindingLineBotJson mean Binding linebot json format from gin context
	BindingLineBotJSON(c *gin.Context) (linebotEvents *linebotE.Events, err error)
}
