package liaoliao

import "github.com/gin-gonic/gin"

// Server describe liaoliao business service function
type Server interface {

	// Server mean liaoliao business logic
	Message(ctx *gin.Context)
}
