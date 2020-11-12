package coreserver

import (
	"log"
	"net/http"

	"github.com/water25234/golang-line-chatbot/config"
	"github.com/water25234/golang-line-chatbot/router"
)

func init() {
	config.SetAppConfig()
}

// StartServer mean start server
func StartServer() {
	log.Fatal(http.ListenAndServe(config.GetAppConfig().GoAddrPort, router.SetupRouter()))
}
