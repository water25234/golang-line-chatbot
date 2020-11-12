package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	goAddrPort = flag.String("goAddrPort", ":8080", "RESTful api port")
)

var appConfig *AppConfig

func init() {
	godotenv.Load()
}

// AppConfig mean app config struct
type AppConfig struct {
	AppLogPath             string
	LineChannelSecret      string
	LineChannelAccessToken string
	GoAddrPort             string
}

// SetAppConfig mean set app config
func SetAppConfig() {
	appConfig = &AppConfig{
		AppLogPath:             os.Getenv("APP_LOG_PATH"),
		LineChannelSecret:      os.Getenv("LINE_CHANNEL_SECRET"),
		LineChannelAccessToken: os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		GoAddrPort:             *goAddrPort,
	}
}

// GetAppConfig mean get app config
func GetAppConfig() *AppConfig {
	return appConfig
}
